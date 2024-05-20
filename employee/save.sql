INSERT INTO `employee`
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE `first_name`=?,
                        `last_name`=?,
                        `email`=?,
                        `phone`=?