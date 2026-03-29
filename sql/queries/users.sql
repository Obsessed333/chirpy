-- name: CreateUser :one
insert into users(id, created_at, updated_at, email, hashed_password, is_chirpy_red)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
returning *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: UpdateUser :one
update users set email = $1, hashed_password = $2, updated_at = $3 where id = $4 returning *;

-- name: UpgradeToChirpyRed :one
update users set is_chirpy_red = true where id = $1 returning *;

-- name: GetUserByID :one
select * from users where id = $1;