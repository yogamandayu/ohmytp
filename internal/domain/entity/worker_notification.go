package entity

// WorkerNotification is a worker notification payload contract.
type WorkerNotification struct {
	Via  string      `json:"via"`
	Data interface{} `json:"data"`
}

// WorkerNotificationViaTelegramData is a contract payload for notification via telegram.
type WorkerNotificationViaTelegramData struct {
	Message string `json:"message"`
}
