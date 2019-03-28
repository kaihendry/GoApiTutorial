CREATE DATABASE rest_api_example;
USE rest_api_example;
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	age INT NOT NULL
);
