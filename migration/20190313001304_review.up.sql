CREATE TABLE user_reviews(
    id int NOT NULL AUTO_INCREMENT,
    order_id int NOT NUll,
    product_id int NOT NULL,
    user_id int NOT NULL,
    rating float(2,1),
    review varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
