// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: adoption.sql

package sqlc

import (
	"context"
	"time"
)

const createAdoption = `-- name: CreateAdoption :exec
insert into adoption (
    pet_id,
    ` + "`" + `uid` + "`" + `,
    is_reviewed,
    is_approved
) values (
    ?, ?, false, false
)
`

type CreateAdoptionParams struct {
	PetID int32 `json:"pet_id"`
	Uid   int32 `json:"uid"`
}

func (q *Queries) CreateAdoption(ctx context.Context, arg CreateAdoptionParams) error {
	_, err := q.db.ExecContext(ctx, createAdoption, arg.PetID, arg.Uid)
	return err
}

const listReviewerAdoption = `-- name: ListReviewerAdoption :many
select adoption_id, pet.pet_id, pet.nickname, category.species, pet.publish_date, adopter.uid, adopter.username
from adoption
join pet on adoption.pet_id = pet.pet_id
join category on pet.category_id = category.category_id
join user as adopter on adoption.uid = adopter.uid
join user as reviewer on pet.uid = reviewer.uid
where reviewer.uid = ? and adoption.is_reviewed = false
`

type ListReviewerAdoptionRow struct {
	AdoptionID  int32     `json:"adoption_id"`
	PetID       int32     `json:"pet_id"`
	Nickname    string    `json:"nickname"`
	Species     string    `json:"species"`
	PublishDate time.Time `json:"publish_date"`
	Uid         int32     `json:"uid"`
	Username    string    `json:"username"`
}

func (q *Queries) ListReviewerAdoption(ctx context.Context, uid int32) ([]ListReviewerAdoptionRow, error) {
	rows, err := q.db.QueryContext(ctx, listReviewerAdoption, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListReviewerAdoptionRow{}
	for rows.Next() {
		var i ListReviewerAdoptionRow
		if err := rows.Scan(
			&i.AdoptionID,
			&i.PetID,
			&i.Nickname,
			&i.Species,
			&i.PublishDate,
			&i.Uid,
			&i.Username,
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

const listUserAdoption = `-- name: ListUserAdoption :many
select pet.pet_id, pet.nickname, category.species, adoption.is_reviewed, adoption.is_approved
from adoption
join pet on adoption.pet_id = pet.pet_id
join category on pet.category_id = category.category_id
where adoption.uid = ?
`

type ListUserAdoptionRow struct {
	PetID      int32  `json:"pet_id"`
	Nickname   string `json:"nickname"`
	Species    string `json:"species"`
	IsReviewed bool   `json:"is_reviewed"`
	IsApproved bool   `json:"is_approved"`
}

func (q *Queries) ListUserAdoption(ctx context.Context, uid int32) ([]ListUserAdoptionRow, error) {
	rows, err := q.db.QueryContext(ctx, listUserAdoption, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUserAdoptionRow{}
	for rows.Next() {
		var i ListUserAdoptionRow
		if err := rows.Scan(
			&i.PetID,
			&i.Nickname,
			&i.Species,
			&i.IsReviewed,
			&i.IsApproved,
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

const updateAdoptionReview = `-- name: UpdateAdoptionReview :exec
update adoption
set
    is_reviewed = true,
    is_approved = ?
where adoption_id = ?
`

type UpdateAdoptionReviewParams struct {
	IsApproved bool  `json:"is_approved"`
	AdoptionID int32 `json:"adoption_id"`
}

func (q *Queries) UpdateAdoptionReview(ctx context.Context, arg UpdateAdoptionReviewParams) error {
	_, err := q.db.ExecContext(ctx, updateAdoptionReview, arg.IsApproved, arg.AdoptionID)
	return err
}
