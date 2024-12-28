CREATE TABLE IF NOT EXISTS user_invitations(
    token bytea PRIMARY KEY NOT NULL,
    email varchar NOT NULL
)