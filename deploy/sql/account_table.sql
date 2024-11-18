CREATE TABLE account_table(
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime DEFAULT NULL,
    `account` varchar(12) NOT NULL,
    `student_number` char(10) NOT NULL,
    `password` varchar(12) NOT NULL,
    `exam_id` int NOT NULL,
    `score` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_account_student` (`account`,`student_number`)
) ENGINE=InnoDB AUTO_INCREMENT=168 DEFAULT CHARSET=utf8mb4;
