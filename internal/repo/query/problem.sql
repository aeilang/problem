-- name: CreateProblem :exec
insert into problems (title, description, is_done, created_by)
values ($1, $2, $3, $4);


-- name: GetAllProblems :many
select id, title, description, is_done, created_by from problems;

-- name: GetProblemsByCreated_by :many
select id, title, description, is_done, created_by from problems
where created_by = $1;

-- name: GetProblemsById :one
select id, title, description, is_done, created_by from problems
where id = $1;

-- name: UpdateProblem :one
update problems set title = $1, description = $2
where created_by = $3 and id = $4 returning id, title, description, is_done;

-- name: DeleteProblem :exec
delete from problems where created_by = $1 and id = $2;

