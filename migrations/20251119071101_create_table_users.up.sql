CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(100),
    email varchar(100) not null,
    password varchar(255),
    created_at timestamp,
    updated_at timestamp
)
