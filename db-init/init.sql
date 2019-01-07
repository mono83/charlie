CREATE SCHEMA IF NOT EXISTS `charlie`;

USE `charlie`;

CREATE TABLE IF NOT EXISTS `issue` (
  `id`         BIGINT(20)  NOT NULL                                                                                         AUTO_INCREMENT,
  `release_id` BIGINT(20)  NOT NULL,
  `issue_id`   VARCHAR(45) NOT NULL,
  `type`       ENUM ('info', 'added', 'changed', 'deprecated', 'removed', 'fixed', 'security', 'performance', 'unreleased') DEFAULT NULL,
  `components` TEXT                                                                                                         DEFAULT NULL,
  `message`    VARCHAR(255)                                                                                                 DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY (`release_id`)
);

CREATE TABLE IF NOT EXISTS `release` (
  `id`         BIGINT(20) NOT NULL AUTO_INCREMENT,
  `project_id` BIGINT(20) NOT NULL,
  `major`      VARCHAR(10)         DEFAULT NULL,
  `minor`      VARCHAR(10)         DEFAULT NULL,
  `patch`      VARCHAR(10)         DEFAULT NULL,
  `label`      VARCHAR(45)         DEFAULT NULL,
  `build`      VARCHAR(45)         DEFAULT NULL,
  `date`       INT(11)             DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY (`project_id`)
);

CREATE TABLE IF NOT EXISTS `project` (
  `id`          BIGINT(20)   NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(255) NOT NULL,
  `description` TEXT                  DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY (`name`)
);