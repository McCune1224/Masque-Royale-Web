package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Action struct {
	ID              int            `db:"id"`
	AbilityName     string         `db:"ability_name"`
	Description     string         `db:"description"`
	Shorthand       string         `db:"shorthand"`
	Rarity          string         `db:"rarity"`
	RoleAssociation string         `db:"role_association"`
	Categories      pq.StringArray `db:"categories"`
}

type ActionModel struct {
	*sqlx.DB
}

func (a *ActionModel) Insert(action *Action) error {
	_, err := a.NamedExec(`INSERT INTO actions (ability_name, description, shorthand, rarity, role_association, categories) 
    VALUES (:ability_name, :description, :shorthand, :rarity, :role_association, :categories)`, action)
	return err
}

func (a *ActionModel) GetByName(name string) (*Action, error) {
	var action *Action
	err := a.Get(&action, "SELECT * from actions WHERE ability_name ILIKE $1", name)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *ActionModel) GetAll() ([]Action, error) {
	var actions []Action
	err := a.Select(&actions, "SELECT * from actions")
	if err != nil {
		return nil, err
	}

	return actions, nil
}
