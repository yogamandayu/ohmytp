package repository_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/storage/repository"
	"github.com/yogamandayu/ohmytp/tests"
)

func TestSaveOtp(t *testing.T) {
	testSuite := tests.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests.FakeOtp().TransformToOtpRepository()
	otp, err := repo.SaveOtp(ctx, repository.SaveOtpParams{
		ID:            fakeOtp.ID,
		RequestID:     fakeOtp.RequestID,
		RouteType:     fakeOtp.RouteType,
		Code:          fakeOtp.Code,
		RequestedAt:   fakeOtp.RequestedAt,
		ConfirmedAt:   fakeOtp.ConfirmedAt,
		ExpiredAt:     fakeOtp.ExpiredAt,
		Attempt:       fakeOtp.Attempt,
		LastAttemptAt: fakeOtp.LastAttemptAt,
		ResendAttempt: fakeOtp.ResendAttempt,
		ResendAt:      fakeOtp.ResendAt,
		IpAddress:     fakeOtp.IpAddress,
		UserAgent:     fakeOtp.UserAgent,
		CreatedAt:     fakeOtp.CreatedAt,
		UpdatedAt:     fakeOtp.UpdatedAt,
	})
	require.NoError(t, err)
	assert.NotNil(t, otp)
}

func TestFindOtp(t *testing.T) {
	testSuite := tests.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests.FakeOtp().TransformToOtpRepository()
	otp, err := repo.SaveOtp(ctx, repository.SaveOtpParams{
		ID:            fakeOtp.ID,
		RequestID:     fakeOtp.RequestID,
		RouteType:     fakeOtp.RouteType,
		Code:          fakeOtp.Code,
		RequestedAt:   fakeOtp.RequestedAt,
		ConfirmedAt:   fakeOtp.ConfirmedAt,
		ExpiredAt:     fakeOtp.ExpiredAt,
		Attempt:       fakeOtp.Attempt,
		LastAttemptAt: fakeOtp.LastAttemptAt,
		ResendAttempt: fakeOtp.ResendAttempt,
		ResendAt:      fakeOtp.ResendAt,
		IpAddress:     fakeOtp.IpAddress,
		UserAgent:     fakeOtp.UserAgent,
		CreatedAt:     fakeOtp.CreatedAt,
		UpdatedAt:     fakeOtp.UpdatedAt,
	})
	require.NoError(t, err)
	assert.NotNil(t, otp)

	otp, err = repo.FindOtp(ctx, otp.ID)
	require.NoError(t, err)
	assert.Equal(t, fakeOtp.Code.String, otp.Code.String)
}

func TestUpdateOtp(t *testing.T) {
	testSuite := tests.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests.FakeOtp().TransformToOtpRepository()
	otp, err := repo.SaveOtp(ctx, repository.SaveOtpParams{
		ID:            fakeOtp.ID,
		RequestID:     fakeOtp.RequestID,
		RouteType:     fakeOtp.RouteType,
		Code:          fakeOtp.Code,
		RequestedAt:   fakeOtp.RequestedAt,
		ConfirmedAt:   fakeOtp.ConfirmedAt,
		ExpiredAt:     fakeOtp.ExpiredAt,
		Attempt:       fakeOtp.Attempt,
		LastAttemptAt: fakeOtp.LastAttemptAt,
		ResendAttempt: fakeOtp.ResendAttempt,
		ResendAt:      fakeOtp.ResendAt,
		IpAddress:     fakeOtp.IpAddress,
		UserAgent:     fakeOtp.UserAgent,
		CreatedAt:     fakeOtp.CreatedAt,
		UpdatedAt:     fakeOtp.UpdatedAt,
	})
	require.NoError(t, err)
	assert.NotNil(t, otp)

	otp1, err := repo.FindOtp(ctx, otp.ID)
	require.NoError(t, err)
	assert.Equal(t, otp.Code.String, otp1.Code.String)

	otp3, err := repo.UpdateOtp(ctx, repository.UpdateOtpParams{
		ID:        fakeOtp.ID,
		RequestID: fakeOtp.RequestID,
		RouteType: fakeOtp.RouteType,
		Code: pgtype.Text{
			String: "UPDATED",
			Valid:  true,
		},
		RequestedAt:   fakeOtp.RequestedAt,
		ConfirmedAt:   fakeOtp.ConfirmedAt,
		ExpiredAt:     fakeOtp.ExpiredAt,
		Attempt:       fakeOtp.Attempt,
		LastAttemptAt: fakeOtp.LastAttemptAt,
		ResendAttempt: fakeOtp.ResendAttempt,
		ResendAt:      fakeOtp.ResendAt,
		IpAddress:     fakeOtp.IpAddress,
		UserAgent:     fakeOtp.UserAgent,
	})
	require.NoError(t, err)
	assert.NotNil(t, otp)
	assert.NotEqual(t, otp.Code.String, otp3.Code.String)
	assert.Equal(t, otp.ID, otp3.ID)
	assert.Equal(t, otp.RequestID, otp3.RequestID)
	assert.Equal(t, "UPDATED", otp3.Code.String)
}
