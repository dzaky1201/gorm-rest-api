CREATE TABLE users(
    user_id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE addresses(
    address_id serial PRIMARY KEY,
    city VARCHAR(40) NOT NULL,
    province VARCHAR(40) NOT NULL,
    postal_code INTEGER NOT NULL,
    user_id_fk INTEGER not NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id_fk) REFERENCES users(user_id),
    CONSTRAINT unique_user_id UNIQUE(user_id_fk)
);