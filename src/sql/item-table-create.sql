CREATE TABLE IF NOT EXISTS `item` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NULL,
  `priority` INT NOT NULL,
  `completed` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `priority_UNIQUE` (`priority` ASC),
  INDEX `completed` (`completed` ASC),
  INDEX `completed_priority` (`completed` ASC, `priority` ASC));