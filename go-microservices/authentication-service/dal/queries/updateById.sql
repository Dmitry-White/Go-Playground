UPDATE
    users
SET
    email = $1,
    first_name = $2,
    last_name = $3,
    user_active = $4,
    updated_at = $5
WHERE
    id = $6