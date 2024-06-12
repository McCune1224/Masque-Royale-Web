package tests

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

// Define a test suite struct
type DatabaseTestSuite struct {
	suite.Suite
	DB *pgx.Conn
}

// SetupSuite is run once before all tests in the suite
func (suite *DatabaseTestSuite) SetupSuite() {
	// Open the database connection
	godotenv.Load("../.env")
	trueDsn := os.Getenv("DATABASE_URL")
	db, err := pgx.Connect(context.Background(), trueDsn)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Store the database connection in the suite
	suite.DB = db
}

// TearDownSuite is run once after all tests in the suite
func (suite *DatabaseTestSuite) TearDownSuite() {
	ctx := context.Background()
	suite.DB.Close(ctx)
}

// Entry point for the test suite
func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
