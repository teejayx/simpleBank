--name: CareAccount :one
Insert into accounts (
    owner,
    balance,
    currency
) values (
    $1,
    $2,
    $3
) RETURNING *;
