USE `charlie`;

ALTER TABLE `project` DROP INDEX `name`;

ALTER TABLE `project` ADD UNIQUE INDEX (`name`);

ALTER TABLE `release` CHANGE COLUMN `date` `date` INT(11) NOT NULL ;