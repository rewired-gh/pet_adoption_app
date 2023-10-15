-- name: GetCategoryID :one
select category_id
from category
where species = ?
and color = ?
and gender = ?
limit 1;

-- name: CreateCategory :exec
insert into category (
    species,
    color,
    gender
) values (
    ?, ?, ?
);