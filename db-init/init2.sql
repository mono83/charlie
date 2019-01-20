USE `charlie`;

ALTER TABLE `project` DROP INDEX `name`;

ALTER TABLE `project` ADD UNIQUE INDEX (`name`);
