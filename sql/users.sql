SET NAMES utf8mb4; /* https://stackoverflow.com/a/35189650/4534 */
/* Above doesn't appear to work https://s.natalian.org/2019-03-31/1553993788_2560x1440.png */
CREATE DATABASE rest_api_example;
USE rest_api_example;
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	age INT NOT NULL
);
