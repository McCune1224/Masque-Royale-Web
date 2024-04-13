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

// An acction that is associated with a player
type PlayerAction struct {
	ID          int    `db:"id"`
	ActionID    int    `db:"action_id"`
	PlayerID    int    `db:"player_id"`
	Target      string `db:"target"`
	Description string `db:"description"`
}

type ActionList struct {
	ID        int           `db:"id"`
	GameID    string        `db:"game_id"`
	ActionIds pq.Int64Array `db:"action_ids"`
}

type ActionModel struct {
	*sqlx.DB
}

func (a *ActionModel) Insert(action *Action) error {
	_, err := a.NamedExec(`INSERT INTO actions (ability_name, description, shorthand, rarity, role_association, categories) 
    VALUES (:ability_name, :description, :shorthand, :rarity, :role_association, :categories)`, action)
	return err
}

func (a *ActionModel) GetByID(id int) (*Action, error) {
	var action Action
	err := a.Get(&action, "SELECT * from actions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (a *ActionModel) GetByName(name string) (*Action, error) {
	var action Action
	err := a.Get(&action, "SELECT * from actions WHERE ability_name ILIKE $1", name)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

func (a *ActionModel) GetAll() ([]Action, error) {
	var actions []Action
	err := a.Select(&actions, "SELECT * from actions")
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (a *ActionModel) InsertActionList(gameID string) error {
	_, err := a.Exec(`INSERT INTO action_lists (game_id, action_ids) VALUES ($1, '{}')`, gameID)
	return err
}

func (a *ActionModel) GetActionList(gameId string) (*ActionList, error) {
	var actionList ActionList
	err := a.Get(&actionList, `SELECT * FROM action_lists WHERE game_id = $1`, gameId)
	if err != nil {
		return nil, err
	}
	return &actionList, nil
}

func (a *ActionModel) UpdateActionList(actionList *ActionList) error {
	// update just the action ids field
	_, err := a.Exec(`UPDATE action_lists SET action_ids = $2 WHERE id = $1`, actionList.ID, actionList.ActionIds)
	return err
}

func (a *ActionModel) RemoveActionList(gameID string, action Action) error {
	_, err := a.Exec(`DELETE FROM action_lists WHERE game_id = $1 AND action_ids = $2`, gameID)
	return err
}

func (a *ActionModel) ClearActionListIDs(gameId string) error {
	_, err := a.Exec(`UPDATE action_lists SET action_ids = '{}' WHERE game_id = $1`, gameId)
	return err
}

func (a *ActionModel) GetPlayerAction(id int) (*PlayerAction, error) {
	var playerAction PlayerAction
	err := a.Get(&playerAction, "SELECT * FROM player_actions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}

func (a *ActionModel) GetPlayerActionByPlayerID(id int) (*PlayerAction, error) {
	var playerAction PlayerAction
	err := a.Get(&playerAction, "SELECT * FROM player_actions WHERE player_id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}

func (a *ActionModel) GetPlayerActionByActionID(id int) (*PlayerAction, error) {
	var playerAction PlayerAction
	err := a.Get(&playerAction, "SELECT * FROM player_actions WHERE action_id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}
