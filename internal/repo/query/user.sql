-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: CreateUser :exec
insert into users(id, username, email, password_hash)
values ($1, $2, $3, $4);
