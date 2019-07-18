CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'name',
  `password` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `salt` varchar(16) DEFAULT NULL COMMENT '',
  `email` varchar(64) DEFAULT NULL,
  `mobile` varchar(15) DEFAULT NULL ,
  `reg_time` int(11) NOT NULL COMMENT '注册时间',
  `role` varchar(100) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;