INSERT INTO `employee`
    (`first_name`, `last_name`, `email`, `phone`)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE `email`=?,
                        `phone`=?