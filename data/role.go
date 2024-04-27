package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Role struct {
	ID         int           `db:"id"`
	Name       string        `db:"name"`
	Alignment  string        `db:"alignment"`
	AbilityIDs pq.Int32Array `db:"ability_ids"`
	PassiveIDs pq.Int32Array `db:"passive_ids"`
}

type ComplexRole struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Alignment string     `db:"alignment"`
	Abilities []*Ability `db:"abilities"`
	Passives  []*Passive `db:"passives"`
}

// Used for when searching for roles on flashcard
type SearchComplexRoleResult struct {
	CR          *ComplexRole
	MatchedName string
}

type RoleAbilityPassiveJoin struct {
	RoleID         int            `db:"role_id"`
	RoleName       string         `db:"role_name"`
	RoleAlignment  string         `db:"role_alignment"`
	RoleAbilityIDs pq.Int32Array  `db:"role_ability_ids"`
	RolePassiveIDs pq.Int32Array  `db:"role_passive_ids"`
	AbilityID      int            `db:"ability_id"`
	AbilityName    string         `db:"ability_name"`
	AbilityDesc    string         `db:"ability_description"`
	AbilityCharges int            `db:"ability_charges"`
	AbilityRarity  string         `db:"ability_rarity"`
	AbilityAny     bool           `db:"ability_any_ability"`
	AbilityRole    string         `db:"ability_role_specific"`
	AbilityCats    pq.StringArray `db:"ability_categories"`
	PassiveID      int            `db:"passive_id"`
	PassiveName    string         `db:"passive_name"`
	PassiveDesc    string         `db:"passive_description"`
}

type RoleModel struct {
	*sqlx.DB
}

func (rm *RoleModel) Get(id int) (*Role, error) {
	query := `SELECT * FROM roles WHERE id = $1`
	var role Role
	err := rm.DB.Get(&role, query, id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rm *RoleModel) GetByName(name string) (*Role, error) {
	var role Role
	err := rm.DB.Get(&role, "SELECT * FROM roles WHERE name ILIKE $1", name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rm *RoleModel) GetByAbilityID(id int) (*Role, error) {
	var role *Role
	err := rm.DB.Select(&role, "SELECT * FROM roles WHERE $1 = ANY(ability_ids)", id)
	if err != nil {
		return nil, err
	}
	return role, err
}

func (rm *RoleModel) GetByPassiveID(id int) (*Role, error) {
	var role *Role
	err := rm.DB.Select(&role, "SELECT * FROM roles WHERE $1 = ANY(passive_ids)", id)
	if err != nil {
		return nil, err
	}
	return role, err
}

func (rm *RoleModel) GetAll() ([]*Role, error) {
	var roles []*Role
	err := rm.DB.Select(&roles, "SELECT * FROM roles")
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (rm *RoleModel) GetComplexByName(name string) (*ComplexRole, error) {
	roleQuery := &RoleAbilityPassiveJoin{}
	complexRole := &ComplexRole{}

	query := `SELECT r.id AS role_id, r.name AS role_name, r.alignment AS role_alignment, r.ability_ids AS role_ability_ids, r.passive_ids AS role_passive_ids, 
  a.id AS ability_id, a.name AS ability_name, a.description AS ability_description, a.charges AS ability_charges, a.rarity AS ability_rarity, a.any_ability AS ability_any_ability, a.role_specific AS ability_role_specific, a.categories AS ability_categories, 
  p.id AS passive_id, p.name AS passive_name, p.description AS passive_description FROM roles r LEFT JOIN abilities a ON a.id = ANY(r.ability_ids) LEFT JOIN passives p ON p.id = ANY(r.passive_ids) WHERE r.name ILIKE $1`
	err := rm.DB.Get(roleQuery, query, name)
	if err != nil {
		return nil, err
	}

	complexRole.ID = roleQuery.RoleID
	complexRole.Name = roleQuery.RoleName
	complexRole.Alignment = roleQuery.RoleAlignment
	var abilities []*Ability
	var passives []*Passive
	for _, id := range roleQuery.RoleAbilityIDs {
		ability := &Ability{
			ID:           int(id),
			Name:         roleQuery.AbilityName,
			Description:  roleQuery.AbilityDesc,
			Charges:      roleQuery.AbilityCharges,
			Rarity:       roleQuery.AbilityRarity,
			AnyAbility:   roleQuery.AbilityAny,
			RoleSpecific: roleQuery.AbilityRole,
			Categories:   roleQuery.AbilityCats,
		}
		abilities = append(abilities, ability)
	}
	for _, id := range roleQuery.RolePassiveIDs {
		passive := &Passive{
			ID:          int(id),
			Name:        roleQuery.PassiveName,
			Description: roleQuery.PassiveDesc,
		}
		passives = append(passives, passive)
	}
	complexRole.Abilities = abilities
	complexRole.Passives = passives

	return complexRole, nil
}

func (rm *RoleModel) GetComplexByID(id int) (*ComplexRole, error) {
	role := &Role{}
	abilities := []*Ability{}
	passives := []*Passive{}

	err := rm.DB.Get(role, "SELECT * FROM roles WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	err = rm.DB.Select(&abilities, "SELECT a.* FROM UNNEST($1::int[]) WITH ORDINALITY t(id, ord) JOIN abilities a ON a.id = t.id", role.AbilityIDs)
	if err != nil {
		return nil, err
	}
	// get all passives associated with provided role passive Ids
	err = rm.DB.Select(&passives, "SELECT p.* FROM UNNEST($1::int[]) WITH ORDINALITY t(id, ord) JOIN passives p ON p.id = t.id", role.PassiveIDs)
	if err != nil {
		return nil, err
	}

	return &ComplexRole{
		ID:        role.ID,
		Name:      role.Name,
		Alignment: role.Alignment,
		Abilities: abilities,
		Passives:  passives,
	}, nil
}

func (rm *RoleModel) GetAllComplex() ([]*ComplexRole, error) {
	cRoles := []*ComplexRole{}
	roles := []*Role{}
	abilities := []*Ability{}
	passives := []*Passive{}

	// Get basic role first
	err := rm.DB.Select(&roles, "SELECT * FROM roles")
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		err = rm.DB.Select(&abilities, "SELECT a.* FROM UNNEST($1::int[]) WITH ORDINALITY t(id, ord) JOIN abilities a ON a.id = t.id", role.AbilityIDs)
		if err != nil {
			return nil, err
		}
		// get all passives associated with provided role passive Ids
		err = rm.DB.Select(&passives, "SELECT p.* FROM UNNEST($1::int[]) WITH ORDINALITY t(id, ord) JOIN passives p ON p.id = t.id", role.PassiveIDs)
		if err != nil {
			return nil, err
		}

		cRole := &ComplexRole{
			ID:        role.ID,
			Name:      role.Name,
			Alignment: role.Alignment,
		}

		cRole.Abilities = append(cRole.Abilities, abilities...)
		cRole.Passives = append(cRole.Passives, passives...)

		cRoles = append(cRoles, cRole)
	}

	return cRoles, nil
}
