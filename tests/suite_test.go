package tests

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

// Define a test suite struct
type DatabaseTestSuite struct {
	suite.Suite
	DB *sqlx.DB
}

// SetupSuite is run once before all tests in the suite
func (suite *DatabaseTestSuite) SetupSuite() {
	// Open the database connection
	godotenv.Load("../.env")
	trueDsn := os.Getenv("MOCK_DATABASE_URL")
	db, err := sqlx.Open("postgres", trueDsn)
	if err != nil {
		suite.T().Fatal(err)
	}

	// Store the database connection in the suite
	suite.DB = db
}

// TearDownSuite is run once after all tests in the suite
func (suite *DatabaseTestSuite) TearDownSuite() {
	// Close the database connection
	suite.DB.Exec(gamesTruncate)
	suite.DB.Close()
}

// Entry point for the test suite
func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
