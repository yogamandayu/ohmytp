package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yogamandayu/ohmytp/domain/repository"
)

type RepositoryRegistry struct {
	OTP repository.OTPRepositoryInterface
}

func NewRepositoryRegistry(db *pgxpool.Conn) *RepositoryRegistry {
	return &RepositoryRegistry{
		OTP: NewOTPRepository(db),
	}
}
