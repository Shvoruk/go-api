CREATE DATABASE IF NOT EXISTS beastybonds;

CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    INDEX idx_users_email (email)
);

CREATE TABLE IF NOT EXISTS animals (
    id INT UNSIGNED AUTO_INCREMENT,
    user_id INT UNSIGNED NOT NULL,
    status ENUM('available','adopted','pending') NOT NULL DEFAULT 'available',
    name VARCHAR(100) NOT NULL,
    sex ENUM('male','female','unknown') NOT NULL,
    breed VARCHAR(100) NOT NULL,
    size ENUM('small','medium','large','giant') NOT NULL,
    age_in_month INT UNSIGNED NOT NULL,
    category VARCHAR(50)  NOT NULL,
    image_url VARCHAR(255),
    description TEXT,
    contact_info VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_animals_user
        FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE,

    PRIMARY KEY (id),
    INDEX idx_animals_category (category),
    INDEX idx_animals_user (user_id)
);