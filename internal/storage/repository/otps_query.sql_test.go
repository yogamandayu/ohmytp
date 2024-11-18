package repository_test

import (
	"context"
	"testing"

	tests2 "github.com/yogamandayu/ohmytp/tests"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/internal/storage/repository"
)

func TestSaveOtp(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests2.FakeOtp().TransformToOtpRepository()
	otp, err := repo.SaveOtp(ctx, repository.SaveOtpParams{
		ID:            fakeOtp.ID,
		RequestID:     fakeOtp.RequestID,
		Identifier:    fakeOtp.Identifier,
		RouteType:     fakeOtp.RouteType,
		Code:          fakeOtp.Code,
		Purpose:       fakeOtp.Purpose,
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
}

func TestFindOtp(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests2.FakeOtp().TransformToOtpRepository()
	resSave, err := repo.SaveOtp(ctx, repository.SaveOtpParams{
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
	})
	require.NoError(t, err)
	assert.NotNil(t, resSave)

	resFind, err := repo.FindOtp(ctx, resSave.ID)
	require.NoError(t, err)
	assert.Equal(t, fakeOtp.Code.String, resFind.Code.String)
}

func TestUpdateOtp(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeOtp := tests2.FakeOtp().TransformToOtpRepository()
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
	})
	require.NoError(t, err)
	assert.NotNil(t, otp)

	otp1, err := repo.FindOtp(ctx, otp.ID)
	require.NoError(t, err)
	assert.Equal(t, otp.Code.String, otp1.Code.String)

	otp2, err := repo.UpdateOtp(ctx, repository.UpdateOtpParams{
		ID:        fakeOtp.ID,
		RequestID: fakeOtp.RequestID,
		RouteType: fakeOtp.RouteType,
		Code: pgtype.Text{
			String: "54321",
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
	assert.NotEqual(t, otp.Code.String, otp2.Code.String)
	assert.Equal(t, otp.ID, otp2.ID)
	assert.Equal(t, otp.RequestID, otp2.RequestID)
	assert.Equal(t, "54321", otp2.Code.String)
}
