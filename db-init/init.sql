CREATE DATABASE IF NOT EXISTS `docentre`;

CREATE USER 'my_user'@'localhost' IDENTIFIED BY 'my_password';
GRANT ALL PRIVILEGES ON docentre.* TO 'my_user'@'localhost';
