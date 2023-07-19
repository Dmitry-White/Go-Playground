SELECT
    id,
    email,
    first_name,
    last_name,
    password,
    user_active,
    created_at,
    updated_at
FROM
    users
WHERE
    id = $1