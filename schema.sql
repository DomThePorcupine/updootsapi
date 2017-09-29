-- Adminer 4.3.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` varchar(250) DEFAULT NULL,
  `userid` varchar(20) DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `userid` varchar(20) NOT NULL,
  `admin` tinyint(1) NOT NULL DEFAULT '0',
  UNIQUE KEY `userid` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `votes`;
CREATE TABLE `votes` (
  `message` int(11) NOT NULL,
  `userid` varchar(20) DEFAULT NULL,
  `updoot` tinyint(1) NOT NULL DEFAULT '0',
  `downdoot` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- 2017-09-29 00:24:21