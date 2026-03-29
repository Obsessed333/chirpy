-- name: CreateRefreshToken :one
insert into refresh_tokens(token, created_at, updated_at, user_id, expires_at)
values(
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
returning *;

-- name: GetUserFromRefreshToken :one
select * from refresh_tokens where token = $1 and revoked_at is null and expires_at > NOW();

-- name: RevokeToken :exec
update refresh_tokens set revoked_at = NOW(), updated_at = NOW() where token = $1;