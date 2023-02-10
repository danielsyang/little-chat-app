create table
  users (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID (),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    PRIMARY KEY (id)
  );

CREATE TABLE
  messages (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID (),
    user_id UUID NOT NULL,
    content VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    PRIMARY KEY (ID),
    CONSTRAINT fk_author FOREIGN KEY (user_id) REFERENCES users (id)
  );