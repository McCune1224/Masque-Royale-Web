package tests

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func setupMockDB(seedStatement string, dsn ...string) (*sqlx.DB, error) {
	trueDsn := ""
	if len(dsn) == 0 {
		godotenv.Load("../.env")
		trueDsn = os.Getenv("MOCK_DATABASE_URL")
	} else {
		trueDsn = dsn[0]
	}
	if trueDsn == "" {
		return nil, errors.New("no database url")
	}
	db, err := sqlx.Open("postgres", trueDsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(seedStatement)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func teardownMockDB(db *sqlx.DB, destroyStatement string) error {
	// Clean up the database
	_, err := db.Exec(destroyStatement)
	db.Close()
	return err
}
