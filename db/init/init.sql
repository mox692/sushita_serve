-- MySQL Script generated by MySQL Workbench
-- Mon Sep 28 17:18:29 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema golang-test-database
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema golang-test-database
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `golang-test-database` DEFAULT CHARACTER SET utf8 ;
USE `golang-test-database` ;

-- -----------------------------------------------------
-- Table `golang-test-database`.`user_score`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `golang-test-database`.`user_score` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` VARCHAR(45) NOT NULL,
  `user_name` VARCHAR(45) NOT NULL,
  `score` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;