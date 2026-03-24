CREATE TABLE account (
    id uuid NOT NULL,
    user_name varchar NOT NULL,
    "password" varchar NOT NULL,
    email varchar NOT NULL,
    "name" varchar NOT NULL,
    active bool NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    CONSTRAINT account_pk UNIQUE (email),
    CONSTRAINT account_pk_id PRIMARY KEY (id)
);
