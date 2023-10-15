-- name: CreatePet :exec
insert into pet (
    nickname,
    `uid`,
    category_id,
    `description`,
    birthday,
    is_adopted,
    publish_date
) values (
    ?, ?, ?, ?, ?, true, now()
);

-- name: ListPet :many
select *
from pet
join category on pet.category_id = category.category_id;