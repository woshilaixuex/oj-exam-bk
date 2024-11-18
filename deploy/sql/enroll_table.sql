CREATE TABLE `enroll_table` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `student_number` char(10) NOT NULL,
  `name` varchar(20) NOT NULL,
  `qq_number` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `reason` varchar(100) NOT NULL,
  `grade` tinyint(4) NOT NULL,
  `had_experience` tinyint(1) NOT NULL,
  `orientation` char(10) DEFAULT NULL,
  `experience` varchar(100),
  PRIMARY KEY (`id`),
  KEY `idx_enroll_table_deleted_at` (`deleted_at`),
  INDEX `idx_student_created` (`student_number`, `created_at` DESC)
) ENGINE=InnoDB AUTO_INCREMENT=168 DEFAULT CHARSET=utf8mb4;
