DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`
(
    `id`         INT(11)   NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `first_name` CHAR(64)  NOT NULL,
    `last_name`  CHAR(64)  NOT NULL,
    `email`      CHAR(128) NOT NULL,
    `phone`      CHAR(16) DEFAULT '',
    UNIQUE KEY `employee_u_key` (`first_name`, `last_name`)
);