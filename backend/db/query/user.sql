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

-- name: GetUser :one
select * from `user`
where username = ? and pword = ?
limit 1;