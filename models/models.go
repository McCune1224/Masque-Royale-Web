// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Alignment string

const (
	AlignmentLAWFUL    Alignment = "LAWFUL"
	AlignmentOUTLANDER Alignment = "OUTLANDER"
	AlignmentCHAOTIC   Alignment = "CHAOTIC"
)

func (e *Alignment) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Alignment(s)
	case string:
		*e = Alignment(s)
	default:
		return fmt.Errorf("unsupported scan type for Alignment: %T", src)
	}
	return nil
}

type NullAlignment struct {
	Alignment Alignment `json:"alignment"`
	Valid     bool      `json:"valid"` // Valid is true if Alignment is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAlignment) Scan(value interface{}) error {
	if value == nil {
		ns.Alignment, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Alignment.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAlignment) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Alignment), nil
}

type Rarity string

const (
	RarityCOMMON       Rarity = "COMMON"
	RarityUNCOMMON     Rarity = "UNCOMMON"
	RarityRARE         Rarity = "RARE"
	RarityEPIC         Rarity = "EPIC"
	RarityLEGENDARY    Rarity = "LEGENDARY"
	RarityMYTHICAL     Rarity = "MYTHICAL"
	RarityROLESPECIFIC Rarity = "ROLE_SPECIFIC"
)

func (e *Rarity) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Rarity(s)
	case string:
		*e = Rarity(s)
	default:
		return fmt.Errorf("unsupported scan type for Rarity: %T", src)
	}
	return nil
}

type NullRarity struct {
	Rarity Rarity `json:"rarity"`
	Valid  bool   `json:"valid"` // Valid is true if Rarity is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRarity) Scan(value interface{}) error {
	if value == nil {
		ns.Rarity, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Rarity.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRarity) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Rarity), nil
}

type Ability struct {
	ID                int32       `json:"id"`
	AbilityDetailsID  pgtype.Int4 `json:"ability_details_id"`
	PlayerInventoryID pgtype.Int4 `json:"player_inventory_id"`
}

type AbilityDetail struct {
	ID             int32       `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	DefaultCharges pgtype.Int4 `json:"default_charges"`
	CategoryIds    []int32     `json:"category_ids"`
	Rarity         Rarity      `json:"rarity"`
	AnyAbility     pgtype.Bool `json:"any_ability"`
}

type Category struct {
	ID       int32       `json:"id"`
	Name     pgtype.Text `json:"name"`
	Priority pgtype.Int4 `json:"priority"`
}

type Game struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Phase     string           `json:"phase"`
	Round     int32            `json:"round"`
	PlayerIds []int32          `json:"player_ids"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type PassiveDetail struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Player struct {
	ID                int32       `json:"id"`
	Name              string      `json:"name"`
	GameID            pgtype.Int4 `json:"game_id"`
	RoleID            pgtype.Int4 `json:"role_id"`
	Alive             bool        `json:"alive"`
	AlignmentOverride pgtype.Text `json:"alignment_override"`
	Notes             string      `json:"notes"`
	RoomID            pgtype.Int4 `json:"room_id"`
}

type PlayerInventory struct {
	PlayerID        int32       `json:"player_id"`
	AbilityName     string      `json:"ability_name"`
	AbilityQuantity pgtype.Int4 `json:"ability_quantity"`
}

type Role struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Alignment Alignment `json:"alignment"`
}

type RoleAbilitiesJoin struct {
	RoleID    int32 `json:"role_id"`
	AbilityID int32 `json:"ability_id"`
}

type RolePassivesJoin struct {
	RoleID    int32 `json:"role_id"`
	PassiveID int32 `json:"passive_id"`
}

type Room struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type StatusDetail struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
}
