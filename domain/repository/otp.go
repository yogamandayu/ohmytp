package repository

import (
	"context"

	"github.com/yogamandayu/ohmytp/domain/entity"
)

type OTPRepositoryInterface interface {
	Save(ctx context.Context, value *entity.OTP) (*entity.OTP, error)
	Find(ctx context.Context, condition *entity.OTP) (*entity.OTP, error)
	Update(ctx context.Context, target *entity.OTP, value *entity.OTP) error
}
