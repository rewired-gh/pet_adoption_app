-- name: CreateAdoption :exec
insert into adoption (
    pet_id,
    `uid`,
    is_reviewed,
    is_approved
) values (
    ?, ?, false, false
);

-- name: ListUserAdoption :many
select pet.pet_id, pet.nickname, category.species, adoption.is_reviewed, adoption.is_approved
from adoption
join pet on adoption.pet_id = pet.pet_id
join category on pet.category_id = category.category_id
where adoption.uid = ?;