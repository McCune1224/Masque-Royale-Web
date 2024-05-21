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

// action that is associated with a player within a game
type PlayerRequest struct {
	ID          int    `db:"id"`
	ActionID    int    `db:"action_id"`
	PlayerID    int    `db:"player_id"`
	GameID      string `db:"game_id"`
	Target      string `db:"target"`
	Description string `db:"description"`
	RoundPhase  string `db:"round_phase"`
	Approved    bool   `db:"approved"`
	Note        string `db:"note"`
}

type ComplexPlayerRequest struct {
	P ComplexPlayer
	R PlayerRequest
	A Action
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

func (a ActionModel) GetAllPlayerActionsForGame(gameID string) ([]*PlayerRequest, error) {
	playerActions := []*PlayerRequest{}
	err := a.Select(&playerActions, "SELECT * from player_requests WHERE game_id = $1", gameID)
	if err != nil {
		return nil, err
	}
	return playerActions, nil
}

func (a *ActionModel) GetPlayerAction(id int) (*PlayerRequest, error) {
	var playerAction PlayerRequest
	err := a.Get(&playerAction, "SELECT * FROM player_requests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}

func (a *ActionModel) GetPlayerActionByPlayerID(id int) (*PlayerRequest, error) {
	var playerAction PlayerRequest
	err := a.Get(&playerAction, "SELECT * FROM player_requests WHERE player_id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}

func (a *ActionModel) GetPlayerActionByActionID(id int) (*PlayerRequest, error) {
	var playerAction PlayerRequest
	err := a.Get(&playerAction, "SELECT * FROM player_requests WHERE action_id = $1", id)
	if err != nil {
		return nil, err
	}

	return &playerAction, err
}

func (a *ActionModel) GetAllActionsByPlayerActions(actionList *ActionList) ([]Action, error) {
	playerActions := []PlayerRequest{}
	fullActions := []Action{}
	err := a.Select(&playerActions, "SELECT * FROM player_requests WHERE game_id = $1", actionList.GameID)
	if err != nil {
		return nil, err
	}

	actionIDS := pq.Int64Array{}
	for _, pa := range playerActions {
		actionIDS = append(actionIDS, int64(pa.ActionID))
	}

	err = a.Select(&fullActions, "SELECT * FROM actions WHERE id = ANY($1)", actionIDS)
	if err != nil {
		return nil, err
	}

	return fullActions, nil
}

func (a *ActionModel) GetPlayerRequest(id int) (*PlayerRequest, error) {
	action := &PlayerRequest{}
	err := a.Get(action, "SELECT * FROM player_requests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *ActionModel) GetAllPlayerUnapprovedRequestsByPlayerID(playerID int) ([]*Action, error) {
	playerActions := []*PlayerRequest{}
	fullActions := []*Action{}
	err := a.Select(&playerActions, "SELECT * FROM player_requests WHERE player_id = $1 and approved = false", playerID)
	if err != nil {
		return nil, err
	}

	ids := pq.Int64Array{}
	for _, pa := range playerActions {
		ids = append(ids, int64(pa.ActionID))
	}
	err = a.Select(&fullActions, "SELECT * FROM actions WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, err
	}
	return fullActions, nil
}

func (a *AbilityModel) GetAllPlayerRequestsByGameID(gameID string) ([]*PlayerRequest, error) {
	playerActions := []*PlayerRequest{}
	err := a.Select(&playerActions, "SELECT * from player_requests WHERE game_id = $1", gameID)
	if err != nil {
		return nil, err
	}
	return playerActions, err
}

func (a *AbilityModel) GetAllUnapprovedPlayerRequestsByGameID(gameID string) ([]*PlayerRequest, error) {
	playerActions := []*PlayerRequest{}
	err := a.Select(&playerActions, "SELECT * from player_requests WHERE game_id = $1 and approved = false", gameID)
	if err != nil {
		return nil, err
	}
	return playerActions, err
}

func (a *AbilityModel) GetAllApprovedPlayerRequestsByGameID(gameID string) ([]*PlayerRequest, error) {
	playerActions := []*PlayerRequest{}
	err := a.Select(&playerActions, "SELECT * from player_requests WHERE game_id = $1 and approved = true", gameID)
	if err != nil {
		return nil, err
	}
	return playerActions, err
}

// func (a *ActionModel) InsertPlayerRequest(pa *PlayerRequest) error {
// 	query := `INSERT INTO player_requests ` + PSQLGeneratedInsert(pa)
// 	_, err := a.NamedExec(query, &pa)
// 	return err
// }

func (a *ActionModel) ApprovePlayerRequest(id int) error {
	_, err := a.Exec("UPDATE player_requests SET approved = true WHERE id = $1", id)
	return err
}

func (a *ActionModel) UpdatePlayerRequestNote(id int, note string) error {
	_, err := a.Exec("UPDATE player_requests SET note = $1 WHERE id = $2", note, id)
	return err
}

func (a *ActionModel) DeletePlayerRequest(id int64) error {
	_, err := a.Exec("DELETE FROM player_requests WHERE id = $1", id)
	return err
}
