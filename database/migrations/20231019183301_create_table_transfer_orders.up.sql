CREATE TABLE transfer_orders
(
    id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255),
    status ENUM('Pending', 'Proccessed', 'Finished') NOT NULL,
    fulfilled_date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

