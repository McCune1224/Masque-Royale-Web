package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
)

func Sanatize(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]string
	group := []string{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			if scanner.Text()[0] == '+' || scanner.Text()[0] == '-' || scanner.Text()[0] == '[' {
				data = append(data, group)
				group = []string{}
			}
			group = append(group, strings.TrimSpace(scanner.Text()))
		}
	}
	data = append(data, group)
	return data[1:], nil
}

func DataEntry(db *sqlx.DB) error {
	filepath := "./roles.txt"
	foo, err := Sanatize(filepath)
	if err != nil {
		return err
	}

	type chunk struct {
		r data.Role
		a []data.Ability
		p []data.Passive
	}

	entries := []chunk{}
	for _, v := range foo {
		role := RoleParse(v)
		abilities, err := AbilityParse(role, v)
		if err != nil {
			return err
		}
		fmt.Println("----")
		passives, err := PassiveParse(role, v)
		if err != nil {
			return err
		}
		entries = append(entries, chunk{role, abilities, passives})
	}

	for _, v := range entries {
		fmt.Println("Inserting", v.r.Name)
		role := v.r
		abilities := v.a
		passives := v.p

		// Insert abilities into psql database
		abResults := pq.Int32Array{}
		for _, a := range abilities {
			// insert the ability in psql
			_, err := db.NamedExec("INSERT INTO abilities (name, charges, description, role_specific, any_ability, categories) VALUES (:name, :charges, :description, :role_specific, :any_ability, :categories)", a)
			if err != nil {
				return err
			}
			// get all the ids of the abilities
			var id int32
			err = db.Get(&id, "SELECT id FROM abilities WHERE name = $1", a.Name)
			if err != nil {
				return err
			}
			abResults = append(abResults, id)
		}

		// Insert passives into psql database
		pResults := pq.Int32Array{}
		for _, p := range passives {
			_, err := db.NamedExec("INSERT INTO passives (name, description) VALUES (:name, :description)", p)
			if err != nil {
				return err
			}
			var id int32
			err = db.Get(&id, "SELECT id FROM passives WHERE name = $1", p.Name)
			if err != nil {
				return err
			}
			pResults = append(pResults, id)
		}
		role.AbilityIDs = append(role.AbilityIDs, abResults...)
		role.PassiveIDs = append(role.PassiveIDs, pResults...)

		// Insert role into psql database
		_, err := db.NamedExec("INSERT INTO roles (name, alignment, ability_ids, passive_ids) VALUES (:name, :alignment, :ability_ids, :passive_ids)", role)
		if err != nil {
			return err
		}
	}
	return nil
}

func RoleParse(entryLines []string) data.Role {
	alignment := entryLines[0]
	if alignment[0] == '[' {
		alignment = alignment[1 : len(alignment)-1]
	} else {
		alignment = alignment[1:]
	}

	name, _, _ := strings.Cut(entryLines[1], " ")


	return data.Role{
		Name:        name,
		Alignment:   alignment,
	}
}

func AbilityParse(r data.Role, entryLines []string) ([]data.Ability, error) {
	start := slices.Index(entryLines, "Abilities:")
	end := slices.Index(entryLines, "Passives:")
	if start == -1 {
		return nil, errors.New("Abilities not found: " + r.Name)
	}
	if end == -1 {
		return nil, errors.New("Passives not found: " + r.Name)
	}

	var abilities []data.Ability
	entries := entryLines[start+1 : end]
	for _, v := range entries {
		nameEnd := strings.Index(v, "[")
		name := v[:nameEnd]

		chargeOpen := strings.Index(v, "[")
		chargeClose := strings.Index(v, "]")
		charge := ""
		if v[chargeOpen+1] != 'x' && v[chargeOpen+1] != 'X' {
			charge = "-1"
		} else {
			charge = v[chargeOpen+2 : chargeClose]
		}

		iCharge, err := strconv.Atoi(strings.TrimSpace(charge))
		if err != nil {
			return nil, err
		}

		abilityType := v[chargeClose+1]

		roleSpecific := ""
		anyAbility := false
		switch abilityType {
		case '*':
			roleSpecific = ""
			anyAbility = true
		case ' ':
			anyAbility = false
			roleSpecific = r.Name

		case '^':
			roleSpecific = r.Name
			anyAbility = true
		default:
			return nil, errors.New("Unknown ability type: " + string(abilityType))
		}

		categoriesOpen := strings.Index(v, "(")
		categoriesClose := strings.Index(v, ")")
		categories := v[categoriesOpen+1 : categoriesClose]
		listedCategories := strings.Split(categories, "/")
		dbListedCategories := pq.StringArray{}
		for _, v := range listedCategories {
			dbListedCategories = append(dbListedCategories, strings.TrimSpace(v))
		}

		description := strings.TrimSpace(strings.Split(v, "-")[1])

		ability := data.Ability{
			Name:         name,
			Charges:      iCharge,
			Description:  description,
			RoleSpecific: roleSpecific,
			AnyAbility:   anyAbility,
			Categories:   dbListedCategories,
		}

		abilities = append(abilities, ability)
	}

	return abilities, nil
}

func PassiveParse(r data.Role, entryLines []string) ([]data.Passive, error) {
	var passives []data.Passive
	start := slices.Index(entryLines, "Passives:")
	if start == -1 {
		return nil, errors.New("Failed to find passives for role: " + r.Name)
	}

	entries := entryLines[start+1:]
	for _, v := range entries {
		entry := strings.Split(v, "-")

		passive := data.Passive{
			Name:        entry[0],
			Description: entry[1],
		}
		passives = append(passives, passive)
	}
	return passives, nil
}
