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

-- name: ListReviewerAdoption :many
select adoption_id, pet.pet_id, pet.nickname, category.species, pet.publish_date, adopter.uid, adopter.username
from adoption
join pet on adoption.pet_id = pet.pet_id
join category on pet.category_id = category.category_id
join user as adopter on adoption.uid = adopter.uid
join user as reviewer on pet.uid = reviewer.uid
where reviewer.uid = ? and adoption.is_reviewed = false;

-- name: UpdateAdoptionReview :exec
update adoption
set
    is_reviewed = true,
    is_approved = ?
where adoption_id = ?;