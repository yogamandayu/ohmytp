package worker

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type ConsumerConfig struct {
	Concurrency uint16
	Priority    Priority
}

type Priority struct {
	Low      uint8
	Default  uint8
	Critical uint8
}

type ProducerConfig struct {
	Priority Priority
	MaxRetry uint8
}

type Worker struct {
	Redis *redis.Client

	Producer       *asynq.Client
	ProducerConfig *ProducerConfig

	Consumer       *asynq.Server
	ConsumerConfig *ConsumerConfig
}

func NewWorker(redis *redis.Client) *Worker {
	return &Worker{
		Redis: redis,
	}
}

func (w *Worker) AsProducer(config *ProducerConfig) {
	w.ProducerConfig = config
	w.Producer = asynq.NewClientFromRedisClient(w.Redis)
}

func (w *Worker) AsConsumer(config *ConsumerConfig) {
	w.ConsumerConfig = config
	w.Consumer = asynq.NewServerFromRedisClient(w.Redis, asynq.Config{
		Concurrency: int(config.Concurrency),
		Queues: map[string]int{
			"low":      int(config.Priority.Low),
			"default":  int(config.Priority.Default),
			"critical": int(config.Priority.Critical),
		},
	})
}

func (w *Worker) Produce(taskName string, payload []byte) error {
	var opts []asynq.Option

	if w.ProducerConfig.MaxRetry != 0 {
		opts = append(opts, asynq.MaxRetry(int(w.ProducerConfig.MaxRetry)))
	}

	task := asynq.NewTask(taskName, payload, opts...)

	_, err := w.Producer.Enqueue(task)
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) Consume(taskName string, handler func(ctx context.Context, task *asynq.Task) error) error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(taskName, handler)
	err := w.Consumer.Run(mux)
	if err != nil {
		return err
	}

	return nil
}
