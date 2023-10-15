// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: pet.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createPet = `-- name: CreatePet :exec
insert into pet (
    nickname,
    ` + "`" + `uid` + "`" + `,
    category_id,
    ` + "`" + `description` + "`" + `,
    birthday,
    is_adopted,
    publish_date
) values (
    ?, ?, ?, ?, ?, true, now()
)
`

type CreatePetParams struct {
	Nickname    string       `json:"nickname"`
	Uid         int32        `json:"uid"`
	CategoryID  int32        `json:"category_id"`
	Description string       `json:"description"`
	Birthday    sql.NullTime `json:"birthday"`
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) error {
	_, err := q.db.ExecContext(ctx, createPet,
		arg.Nickname,
		arg.Uid,
		arg.CategoryID,
		arg.Description,
		arg.Birthday,
	)
	return err
}

const listPet = `-- name: ListPet :many
select pet_id, uid, pet.category_id, nickname, description, birthday, is_adopted, publish_date, category.category_id, species, color, gender
from pet
join category on pet.category_id = category.category_id
`

type ListPetRow struct {
	PetID        int32          `json:"pet_id"`
	Uid          int32          `json:"uid"`
	CategoryID   int32          `json:"category_id"`
	Nickname     string         `json:"nickname"`
	Description  string         `json:"description"`
	Birthday     sql.NullTime   `json:"birthday"`
	IsAdopted    bool           `json:"is_adopted"`
	PublishDate  time.Time      `json:"publish_date"`
	CategoryID_2 int32          `json:"category_id_2"`
	Species      string         `json:"species"`
	Color        sql.NullString `json:"color"`
	Gender       sql.NullString `json:"gender"`
}

func (q *Queries) ListPet(ctx context.Context) ([]ListPetRow, error) {
	rows, err := q.db.QueryContext(ctx, listPet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPetRow{}
	for rows.Next() {
		var i ListPetRow
		if err := rows.Scan(
			&i.PetID,
			&i.Uid,
			&i.CategoryID,
			&i.Nickname,
			&i.Description,
			&i.Birthday,
			&i.IsAdopted,
			&i.PublishDate,
			&i.CategoryID_2,
			&i.Species,
			&i.Color,
			&i.Gender,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}