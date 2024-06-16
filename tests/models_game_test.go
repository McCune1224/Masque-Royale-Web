package tests

import (
	"context"

	"github.com/mccune1224/betrayal-widget/models"
)

func (suite *DatabaseTestSuite) TestGetGameByID() {
	ctx := context.Background()
	// Reset the games table before each test
	type test struct {
		targetId    int32
		expectedId  int32
		expectError bool
	}

	tests := []test{
		{targetId: 1, expectedId: 1, expectError: false},
		{targetId: 2, expectedId: 2, expectError: false},
		{targetId: 1 << 8, expectedId: 0, expectError: true}, // Expected to fail
	}

	q := models.New(suite.DB)
	for _, tt := range tests {
		game, err := q.GetGame(ctx, int32(tt.targetId))
		if tt.expectError {
			suite.Error(err)
		} else {
			suite.NoError(err)
			suite.NotNil(game)
			suite.Equal(tt.expectedId, game.ID)
		}
	}
}
