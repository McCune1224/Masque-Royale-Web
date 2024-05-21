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

func (suite *DatabaseTestSuite) TestInsertGame() {
	mockGameModel := data.GameModel{DB: suite.DB}

	type test struct {
		targetName    string
		expectedError bool
	}

	tests := []test{
		{"unique", false},
		{"UniQuE", false},
		{"unique", true},
	}

	for _, tt := range tests {
		err := mockGameModel.InsertGame(tt.targetName)
		if tt.expectedError {
			suite.Error(err)
		} else {
			suite.NoError(err)
		}

	}

}
func (suite *DatabaseTestSuite) TestUpdateGame() {
	mockGameModel := data.GameModel{DB: suite.DB}

	suite.DB.Exec(gamesTruncate)
	suite.DB.Exec(`INSERT INTO games (name) VALUES ('Foo'), ('Bar');`)

	type test struct {
		game          *data.Game
		expectedError bool
	}

	updatedName := "A Brand New Name!"
	updatedPlayerCount := 1924

	g1, err := mockGameModel.GetGameByID(1)

	suite.NoError(err)

	g1.Name = updatedName
	g1.PlayerCount = updatedPlayerCount
	err = mockGameModel.UpdateGame(g1)
	suite.NoError(err)

	t1Updated, err := mockGameModel.GetGameByID(1)
	suite.NoError(err)

	suite.Equal(updatedName, t1Updated.Name)
	suite.NotEqual(0, t1Updated.PlayerCount)

	g1.Name = "Bar"
	err = mockGameModel.UpdateGame(g1)
	suite.Error(err)
}
