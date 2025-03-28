CREATE TABLE IF NOT EXISTS animals (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO animals (name, category) VALUES
    ('Lion', 'Mammal'),
    ('Elephant', 'Mammal'),
    ('Eagle', 'Bird'),
    ('Shark', 'Fish'),
    ('Frog', 'Amphibian'),
    ('Cobra', 'Reptile'),
    ('Kangaroo', 'Mammal'),
    ('Panda', 'Mammal'),
    ('Penguin', 'Bird'),
    ('Tuna', 'Fish');