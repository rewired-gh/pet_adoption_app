-- name: CreateUser :exec
insert into `user` (
    location_id,
    username,
    pword,
    gender,
    birthday
) values (
    ?, ?, ?, ?, ?
);

-- name: AuthUser :one
select `uid` from `user`
where username = ? and pword = ?
limit 1;

-- name: GetUser :one
select username, gender, birthday, province, city, district
from `user`
join `location` on `user`.location_id = `location`.location_id
where `user`.`uid` = ?
limit 1;

-- name: AddContact :exec
insert into `contact` (
    `uid`,
    kind,
    content
) values (
    ?, ?, ?
);