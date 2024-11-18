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

func TestSaveOtpRouteTypeSMS(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeSMS().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeSMS(ctx, repository.SaveOtpRouteTypeSMSParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Phone:     fakeRouteType.Phone,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)
}

func TestFindOtpRouteTypeSMS(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeSMS().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeSMS(ctx, repository.SaveOtpRouteTypeSMSParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Phone:     fakeRouteType.Phone,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)

	routeType, err = repo.FindOtpRouteTypeSMS(ctx, routeType.ID)
	require.NoError(t, err)
	assert.Equal(t, fakeRouteType.Phone.String, routeType.Phone.String)
}

func TestUpdateOtpRouteTypeSMS(t *testing.T) {
	testSuite := tests2.NewTestSuite()
	testSuite.LoadApp()
	t.Cleanup(testSuite.Clean)

	ctx := context.Background()
	repo := repository.New(testSuite.App.DB)
	fakeRouteType := tests2.FakeOtpRouteTypeSMS().TransformToOtpRepository()
	routeType, err := repo.SaveOtpRouteTypeSMS(ctx, repository.SaveOtpRouteTypeSMSParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Phone:     fakeRouteType.Phone,
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)

	routeType1, err := repo.FindOtpRouteTypeSMS(ctx, routeType.ID)
	require.NoError(t, err)
	assert.Equal(t, routeType.Phone.String, routeType1.Phone.String)

	routeType2, err := repo.UpdateOtpRouteTypeSMS(ctx, repository.UpdateOtpRouteTypeSMSParams{
		ID:        fakeRouteType.ID,
		RequestID: fakeRouteType.RequestID,
		OtpID:     fakeRouteType.OtpID,
		Phone: pgtype.Text{
			String: "1234567890",
			Valid:  true,
		},
	})
	require.NoError(t, err)
	assert.NotNil(t, routeType)
	assert.NotEqual(t, routeType.Phone.String, routeType2.Phone.String)
	assert.Equal(t, routeType.ID, routeType2.ID)
	assert.Equal(t, routeType.RequestID, routeType2.RequestID)
	assert.Equal(t, "1234567890", routeType2.Phone.String)
}
