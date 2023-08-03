-- name: CreateAccount :one
Insert into accounts (
    owner,
    balance,
    currency
) values (
     $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts 
ORDER BY id 
LIMIT $1 
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts SET
    balance = $2
WHERE id = $1
RETURNING *;    

-- name: DeleteAccount :exec    
DELETE FROM accounts WHERE id = $1;

