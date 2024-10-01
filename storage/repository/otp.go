package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yogamandayu/ohmytp/domain/entity"
	"github.com/yogamandayu/ohmytp/domain/repository"
)

type OTPRepository struct {
	db *pgxpool.Conn
}

func NewOTPRepository(db *pgxpool.Conn) *OTPRepository {
	return &OTPRepository{
		db,
	}
}

var _ repository.OTPRepositoryInterface = &OTPRepository{}

func (o *OTPRepository) Save(ctx context.Context, value *entity.OTP) (*entity.OTP, error) {

	return value, nil
}

func (o *OTPRepository) Find(ctx context.Context, condition *entity.OTP) (*entity.OTP, error) {
	var otp *entity.OTP

	return otp, nil
}

func (o *OTPRepository) Update(ctx context.Context, target, value *entity.OTP) error {

	return nil
}
