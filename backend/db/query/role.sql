-- name: ApplyAllRoles :exec
insert into user_role (
    `uid`,
    role_id
) values
    (sqlc.arg(uid), 1),
    (sqlc.arg(uid), 2),
    (sqlc.arg(uid), 3);

-- name: ApplyUserRole :exec
insert into user_role (
    `uid`,
    role_id
) values (
    ?, 3
);
