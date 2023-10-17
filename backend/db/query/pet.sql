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
    ?, ?, ?, ?, ?, false, now()
);

-- name: UpdatePet :exec
update pet
set
    nickname = ?,
    category_id = ?,
    `description` = ?,
    birthday = ?
where pet_id = ?;

-- name: DeletePet :exec
delete from pet
where pet_id = ?;

-- name: ListPet :many
select pet.pet_id, pet.uid, pet.category_id, pet.nickname, pet.birthday, pet.is_adopted, pet.description, pet.publish_date, category.species, category.color, category.gender
from pet
join category on pet.category_id = category.category_id;

-- name: ListAvailablePet :many
select pet.pet_id, pet.uid, pet.category_id, pet.nickname, pet.birthday, pet.is_adopted, pet.description, pet.publish_date, category.species, category.color, category.gender
from pet
join category on pet.category_id = category.category_id
left join adoption on pet.pet_id = adoption.pet_id and adoption.uid = ?
where adoption.uid is null;