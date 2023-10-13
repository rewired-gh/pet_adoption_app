// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package sqlc

import (
	"database/sql"
	"time"
)

type Adoption struct {
	AdoptionID int32 `json:"adoption_id"`
	PetID      int32 `json:"pet_id"`
	Uid        int32 `json:"uid"`
	IsReviewed bool  `json:"is_reviewed"`
	IsApproved bool  `json:"is_approved"`
}

type Blacklist struct {
	BlacklistID int32         `json:"blacklist_id"`
	Uid         sql.NullInt32 `json:"uid"`
	Reason      string        `json:"reason"`
}

type Category struct {
	CategoryID int32          `json:"category_id"`
	Species    string         `json:"species"`
	Color      sql.NullString `json:"color"`
	Gender     sql.NullString `json:"gender"`
}

type Contact struct {
	Uid     int32  `json:"uid"`
	Kind    string `json:"kind"`
	Content string `json:"content"`
}

type Location struct {
	LocationID int32          `json:"location_id"`
	Province   string         `json:"province"`
	City       sql.NullString `json:"city"`
	District   string         `json:"district"`
}

type Pet struct {
	PetID       int32        `json:"pet_id"`
	Uid         int32        `json:"uid"`
	CategoryID  int32        `json:"category_id"`
	Nickname    string       `json:"nickname"`
	Description string       `json:"description"`
	Birthday    sql.NullTime `json:"birthday"`
	IsAdopted   bool         `json:"is_adopted"`
	PublishDate time.Time    `json:"publish_date"`
}

type Role struct {
	RoleID   int32  `json:"role_id"`
	Rolename string `json:"rolename"`
}

type User struct {
	Uid        int32     `json:"uid"`
	LocationID int32     `json:"location_id"`
	Username   string    `json:"username"`
	Pword      string    `json:"pword"`
	Gender     string    `json:"gender"`
	Birthday   time.Time `json:"birthday"`
}

type UserRole struct {
	Uid    int32 `json:"uid"`
	RoleID int32 `json:"role_id"`
}
