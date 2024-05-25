package tests

import (
	"context"

	"github.com/mccune1224/betrayal-widget/models"
)

const gamesTruncate = "TRUNCATE games RESTART IDENTITY;"

func (suite *DatabaseTestSuite) TestGetGameByID() {
	ctx := context.Background()
	// Reset the games table before each test
	suite.DB.Exec(ctx, gamesTruncate)
	suite.DB.Exec(ctx, `INSERT INTO games (name) VALUES ('Foo'), ('Bar'), ('Bababoey');`)

	type test struct {
		targetId    int
		expectedId  int
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
			suite.Nil(game)
		} else {
			suite.NoError(err)
			suite.NotNil(game)
			suite.Equal(tt.expectedId, game.ID)
		}
	}
}

func (suite *DatabaseTestSuite) TestGetAllGames() {
	ctx := context.Background()
	suite.DB.Exec(ctx, gamesTruncate)
	suite.DB.Exec(ctx, `INSERT INTO games (name) VALUES ('Foo'), ('Bar'), ('Bababoey');`)
	mockGameModel := models.New(suite.DB)

	got, err := mockGameModel.ListGames(ctx)
	suite.NoError(err)

	suite.Len(got, 3)
}

// func (suite *DatabaseTestSuite) TestInsertGame() {
// 	ctx := context.Background()
// 	mockGameModel := models.New(suite.DB)
//
// 	type test struct {
// 		targetName    string
// 		expectedError bool
// 	}
//
// 	tests := []test{
// 		{"unique", false},
// 		{"UniQuE", false},
// 		{"unique", true},
// 	}
//
// 	for _, tt := range tests {
// 		err := mockGameModel.InsertGame(tt.targetName)
// 		if tt.expectedError {
// 			suite.Error(err)
// 		} else {
// 			suite.NoError(err)
// 		}
//
// 	}
//
// }

// func (suite *DatabaseTestSuite) TestUpdateGame() {
// 	mockGameModel := data.GameModel{DB: suite.DB}
//
// 	suite.DB.Exec(gamesTruncate)
// 	suite.DB.Exec(`INSERT INTO games (name) VALUES ('Foo'), ('Bar');`)
//
// 	type test struct {
// 		game          *data.Game
// 		expectedError bool
// 	}
//
// 	updatedName := "A Brand New Name!"
// 	updatedIds := pq.Int32Array{20, 99}
//
// 	g1, err := mockGameModel.GetGameByID(1)
//
// 	suite.NoError(err)
//
// 	g1.Name = updatedName
// 	g1.Player_Ids = updatedIds
// 	err = mockGameModel.UpdateGame(g1)
// 	suite.NoError(err)
//
// 	t1Updated, err := mockGameModel.GetGameByID(1)
// 	suite.NoError(err)
//
// 	suite.Equal(updatedName, t1Updated.Name)
// 	suite.NotEqual(pq.Int32Array{}, t1Updated.Player_Ids)
//
// 	g1.Name = "Bar"
// 	err = mockGameModel.UpdateGame(g1)
// 	suite.Error(err)
// }
