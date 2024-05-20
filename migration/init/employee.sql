DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`
(
    `id`         CHAR(36)  NOT NULL PRIMARY KEY,
    `first_name` CHAR(64)  NOT NULL,
    `last_name`  CHAR(64)  NOT NULL,
    `email`      CHAR(128) NOT NULL,
    `phone`      CHAR(16) DEFAULT ''
);