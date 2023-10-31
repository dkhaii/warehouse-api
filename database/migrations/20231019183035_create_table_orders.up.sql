CREATE TABLE orders
(
    id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    notes VARCHAR(255) NOT NULL,
    request_transfer_date DATE,
    created_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);