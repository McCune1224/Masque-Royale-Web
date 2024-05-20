package tests

import (
	"testing"

	"github.com/mccune1224/betrayal-widget/data"
)

func TestGetAllGames(t *testing.T) {
	db, err := setupMockDB(
		`INSERT INTO games (name) VALUES ('Foo'), ('Bar'), ('Bababoey');`,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer teardownMockDB(db,
		`TRUNCATE TABLE games RESTART IDENTITY CASCADE;`,
	)
	mockGameModel := data.GameModel{db}

	got, err := mockGameModel.GetAllGames()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 3 {
		t.Fatalf("got %d, expected 3", len(got))
	}

}
