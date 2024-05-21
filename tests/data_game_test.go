package tests

import (
	"github.com/mccune1224/betrayal-widget/data"
)

const gamesTruncate = "TRUNCATE games RESTART IDENTITY;"

func (suite *DatabaseTestSuite) TestGetGameByID() {
	// Reset the games table before each test
	suite.DB.Exec(gamesTruncate)
	suite.DB.Exec(`INSERT INTO games (name) VALUES ('Foo'), ('Bar'), ('Bababoey');`)

	mockGameModel := data.GameModel{DB: suite.DB}

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

	for _, tt := range tests {
		game, err := mockGameModel.GetGameByID(tt.targetId)

		if tt.expectError {
			suite.Error(err)
			suite.Nil(game)
		} else {
			suite.NoError(err)
			suite.NotNil(game)
			suite.Equal(tt.expectedId, game.Id)
		}
	}
}

func (suite *DatabaseTestSuite) TestGetAllGames() {

	suite.DB.Exec(gamesTruncate)
	suite.DB.Exec(`INSERT INTO games (name) VALUES ('Foo'), ('Bar'), ('Bababoey');`)
	mockGameModel := data.GameModel{DB: suite.DB}

	got, err := mockGameModel.GetAllGames()
	suite.NoError(err)

	suite.Len(got, 3)
}
