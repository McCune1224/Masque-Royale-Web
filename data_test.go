package main

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// func TestSanatize(t *testing.T) {
// 	filepath := "./roles.txt"
// 	foo, err := Sanatize(filepath)
// 	if err != nil {
// 		t.Errorf("Sanatize failed: %v", err)
// 	}
// 	if len(foo) == 0 {
// 		t.Errorf("Sanatize failed: %v", foo)
// 	}
// }

// func TestRoleParse(t *testing.T) {
// 	data, err := Sanatize("./roles.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, v := range data {
// 		RoleParse(v)
// 	}
// }

func TestDataEntry(t *testing.T) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal("error opening database,", err)
	}

	err = DataEntry(db)
	if err != nil {
		t.Fatal(err)
	}
}
