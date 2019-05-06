CREATE DATABASE rest_api_example character set utf8mb4 collate utf8mb4_unicode_ci;
USE rest_api_example;
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	age INT
);
