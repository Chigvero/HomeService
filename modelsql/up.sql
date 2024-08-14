-- Включение расширения uuid-ossp для генерации UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создание таблицы Users
CREATE TABLE Users (
                       id        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email     VARCHAR(255) UNIQUE NOT NULL,
                       user_type VARCHAR(255) NOT NULL CHECK (user_type IN ('moderator', 'client')),
                       password  VARCHAR(255) NOT NULL
);

CREATE TABLE houses (
                        id SERIAL PRIMARY KEY,
                        address VARCHAR(255) NOT NULL,
                        year INT NOT NULL,
                        developer VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE flats (
                       id INT,
                       house_id INT NOT NULL,
                       price INT NOT NULL,
                       rooms INT NOT NULL,
                       status VARCHAR(50) CHECK (status IN ('created', 'approved','declined','on moderation')) default 'created' ,
                       moderator_id UUID REFERENCES Users(id),
                       FOREIGN KEY (house_id) REFERENCES houses(id),
                        PRIMARY KEY (house_id, id)
);
