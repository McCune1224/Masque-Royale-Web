package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Ability struct {
	ID           int            `db:"id"`
	Name         string         `db:"name"`
	Description  string         `db:"description"`
	Charges      int            `db:"charges"`
	Rarity       string         `db:"rarity"`
	AnyAbility   bool           `db:"any_ability"`
	RoleSpecific string         `db:"role_specific"`
	Categories   pq.StringArray `db:"categories"`
}

type AnyAbility struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Rarity      string         `db:"rarity"`
	Categories  pq.StringArray `db:"categories"`
}

type AbilityFlashcardSearch struct {
	Ability
	AssociatedRoles []string
}

// create table psql statement for any_ability table
type AbilityModel struct {
	*sqlx.DB
}

func (m *AbilityModel) GetAllAnyAbilities() ([]*AnyAbility, error) {
	abilities := []*AnyAbility{}
	err := m.Select(&abilities, "SELECT * FROM any_abilities")
	if err != nil {
		return nil, err
	}
	return abilities, nil
}
