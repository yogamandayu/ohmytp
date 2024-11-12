package repository_test

import (
	"context"
	"testing"

	repository2 "github.com/yogamandayu/ohmytp/internal/storage/repository"
	tests2 "github.com/yogamandayu/ohmytp/internal/tests"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveOtpRouteTypeEmail(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	defer t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository2.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeEmail().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeEmail(ctx, repository2.SaveOtpRouteTypeEmailParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Email:     fakeRouteType.Email,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)
}

func TestFindOtpRouteTypeEmail(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository2.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeEmail().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeEmail(ctx, repository2.SaveOtpRouteTypeEmailParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Email:     fakeRouteType.Email,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)

	routeType, err = repo.FindOtpRouteTypeEmail(ctx, routeType.ID)
	require.NoError(t, err)
	assert.Equal(t, fakeRouteType.Email.String, routeType.Email.String)
}

func TestUpdateOtpRouteTypeEmail(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository2.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeEmail().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeEmail(ctx, repository2.SaveOtpRouteTypeEmailParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Email:     fakeRouteType.Email,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)

	routeType1, err := repo.FindOtpRouteTypeEmail(ctx, routeType.ID)
	require.NoError(t, err)
	assert.Equal(t, routeType.Email.String, routeType1.Email.String)

	routeType2, err := repo.UpdateOtpRouteTypeEmail(ctx, repository2.UpdateOtpRouteTypeEmailParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Email: pgtype.Text{
			String: "update-email@example.com",
			Valid:  true,
		},
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)
	assert.NotEqual(t, routeType.Email.String, routeType2.Email.String)
	assert.Equal(t, routeType.ID, routeType2.ID)
	assert.Equal(t, routeType.RequestID, routeType2.RequestID)
	assert.Equal(t, "update-email@example.com", routeType2.Email.String)
}
