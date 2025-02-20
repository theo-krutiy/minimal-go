CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    count_available INT NOT NULL CHECK (count_available >= 0),
    price INT NOT NULL CHECK (price >= 0),
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
);

CREATE TABLE items_in_cart (
    item_id INT REFERENCES items(id) ON DELETE CASCADE,
    cart_id INT REFERENCES carts(id) ON DELETE CASCADE,
    count_in_cart INT NOT NULL CHECK (count_in_cart >= 0),
    CONSTRAINT unique_cart_item UNIQUE(item_id, cart_id)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    total_price INT NOT NULL CHECK (total_price >= 0),
    state VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP 
);

CREATE TABLE items_in_order (
    item_id INT REFERENCES items(id) ON DELETE CASCADE,    
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    count INT NOT NULL CHECK (count >= 0),
    price INT NOT NULL CHECK (price >= 0),
);
