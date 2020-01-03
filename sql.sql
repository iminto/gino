/*
SQLyog Ultimate v12.4.1 (64 bit)
MySQL - 5.5.64-MariaDB : Database - test
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`test` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `test`;

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL COMMENT 'name',
  `password` varchar(60) CHARACTER SET utf8 NOT NULL,
  `salt` varchar(16) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `mobile` varchar(15) DEFAULT NULL,
  `reg_time` int(11) NOT NULL,
  `role` varchar(100) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `users` */

insert  into `users`(`id`,`username`,`password`,`salt`,`email`,`mobile`,`reg_time`,`role`,`name`,`updated_at`,`created_at`) values 
(1,'wait','123456',NULL,'waitfox@qq.com',NULL,1577672367,'1','waitfox@qq.com','2020-01-03 09:52:44','2020-01-01 09:52:39'),
(2,'admin','123456',NULL,'admin@qq.com',NULL,1578014188,'0','admin@qq.com','2020-01-03 09:52:46','2020-01-01 09:52:42');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
