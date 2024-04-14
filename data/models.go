package data

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Models struct {
	Games     GameModel
	Players   PlayerModel
	Roles     RoleModel
	Abilities AbilityModel
	Passives  PassiveModel
	Alliances AllianceModel
	Actions   ActionModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		Games:     GameModel{DB: db},
		Players:   PlayerModel{DB: db},
		Roles:     RoleModel{DB: db},
		Abilities: AbilityModel{DB: db},
		Passives:  PassiveModel{DB: db},
		Alliances: AllianceModel{DB: db},
		Actions:   ActionModel{DB: db},
	}
}

// Helper to automatically generate the PSQL query for the keys of a struct
func SqlGenKeys(model interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(model))
	var query []string
	for i := 0; i < v.NumField(); i++ {
		columnName := v.Type().Field(i).Tag.Get("db")

		switch t := v.Field(i).Interface().(type) {
		case string:
			if t != "" {
				query = append(query, fmt.Sprintf("%s=$%s", columnName, columnName))
			}
		case int:
			if t != 0 {
				query = append(query, fmt.Sprintf("%s=$%s", columnName, columnName))
			}
		default:
			if reflect.ValueOf(t).Kind() == reflect.Ptr {
				if reflect.Indirect(reflect.ValueOf(t)) != reflect.ValueOf(nil) {
					query = append(query, fmt.Sprintf("%s=$%s", columnName, columnName))
				}
			} else {
				if reflect.ValueOf(t) != reflect.ValueOf(nil) {
					query = append(query, fmt.Sprintf("%s=$%s", columnName, columnName))
				}
			}
		}
	}
	return strings.Join(query, ", ")
}

// Function to generate an SQLX insert statement with placeholders
// i.e for query INSERT INTO table (column1, column2...) VALUES (:column1, :column2...)
// Will check if default field value is not nil/empty before adding to query
func PSQLGeneratedInsert(model interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(model))
	var columns []string
	var tally int

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("db")
		switch t := v.Field(i).Interface().(type) {
		case string:
			if t != "" {
				columns = append(columns, tag)
				tally += 1
			}
		case int:
			if t != 0 {
				columns = append(columns, tag)
				tally += 1
			}
		case int64:
			if t != 0 {
				columns = append(columns, tag)
				tally += 1
			}
		default:
			if reflect.ValueOf(t).Kind() == reflect.Ptr {
				if reflect.Indirect(reflect.ValueOf(t)) != reflect.ValueOf(nil) {
					columns = append(columns, tag)
					tally += 1
				}
			} else {
				if reflect.ValueOf(t) != reflect.ValueOf(nil) {
					columns = append(columns, tag)
					tally += 1
				}
			}
		}
	}
	left := ""
	right := ""
	for _, v := range columns {
		left = left + v + ", "
		right = right + ":" + v + ", "
	}
	left = strings.TrimSuffix(left, ", ")
	right = strings.TrimSuffix(right, ", ")

	left = "(" + left + ")"
	right = "VALUES (" + right + ")"

	return fmt.Sprintf("%s %s", left, right)
}

func PSQLGeneratedUpdate(model interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(model))
	var query []string
	for i := 0; i < v.NumField(); i++ {
		columnName := v.Type().Field(i).Tag.Get("db")

		switch t := v.Field(i).Interface().(type) {
		case string:
			if t != "" {
				query = append(query, fmt.Sprintf("%s = :%s", columnName, columnName))
			}
		case int:
			if t != 0 {
				query = append(query, fmt.Sprintf("%s = :%s", columnName, columnName))
			}
		case pq.StringArray:
			if len(t) != 0 {
				query = append(query, fmt.Sprintf("%s = :%s", columnName, columnName))
			}
		}
	}

	ret := strings.Join(query, ", ")
	return ret
}
