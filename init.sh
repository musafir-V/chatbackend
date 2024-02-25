#!/bin/bash

# Create database and user
docker exec -i chatbackend-db-1 mysql -uroot -ppassword <<EOF
CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;
CREATE TABLE IF NOT EXISTS user (
    username VARCHAR(255) PRIMARY KEY,
    logged_in BOOLEAN DEFAULT 0,
    password VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sender VARCHAR(255) NOT NULL,
    receiver VARCHAR(255) NOT NULL,
    is_read BOOLEAN DEFAULT 0,
    message TEXT NOT NULL
);
EOF

echo "Database and user created successfully."
