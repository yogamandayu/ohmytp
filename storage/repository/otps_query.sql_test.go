package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/storage/repository"
	"github.com/yogamandayu/ohmytp/tests"
	"testing"
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
