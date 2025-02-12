CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    count_available INT NOT NULL CHECK (count_available >= 0),
    price_integer INT NOT NULL CHECK (price_integer >= 0),
    price_decimal INT NOT NULL CHECK (price_decimal >= 0 AND price_decimal < 100)
);

CREATE TABLE items_in_cart (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id) ON DELETE CASCADE,
    count_in_cart INT NOT NULL CHECK (count_in_cart >= 0),
    CONSTRAINT unique_cart_item UNIQUE(user_id, item_id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP 
);

CREATE TABLE items_in_transaction (
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    item_id INT REFERENCES items(id) ON DELETE CASCADE,
    count_in_transaction INT NOT NULL CHECK (count_in_transaction >= 0),
    price_integer INT NOT NULL CHECK (price_integer >= 0),
    price_decimal INT NOT NULL CHECK (price_decimal >= 0 AND price_decimal < 100)
);
