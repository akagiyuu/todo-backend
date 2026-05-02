-- name: CreateAccount :one
INSERT INTO accounts(email, password)
VALUES (@email, @password)
RETURNING id;

-- name: GetAccountByEmail :one
SELECT id, password
FROM accounts
WHERE email = @email;
