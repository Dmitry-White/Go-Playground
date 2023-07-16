--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY users
ADD
    CONSTRAINT users_pkey PRIMARY KEY (id);

INSERT INTO
    users(
        "email",
        "first_name",
        "last_name",
        "password",
        "user_active",
        "created_at",
        "updated_at"
    )
VALUES
    (
        'admin@example.com',
        'Admin',
        'User',
        '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',
        1,
        '2022-03-14 00:00:00',
        '2022-03-14 00:00:00'
    );