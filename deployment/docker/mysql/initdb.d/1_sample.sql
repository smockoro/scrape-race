DROP DATABASE IF EXISTS raceservice;
CREATE DATABASE raceservice;
USE raceservice;

CREATE TABLE `races` (
              `id` bigint(20) NOT NULL,
              `distance` bigint(20) DEFAULT NULL,
              `racecource` varchar(200) DEFAULT NULL,
              `open_date` varchar(200) DEFAULT NULL,
              PRIMARY KEY (`ID`),
              UNIQUE KEY `ID_UNIQUE` (`ID`)
);

CREATE TABLE `horses` (
              `id` bigint(20) NOT NULL,
              `name` varchar(200) DEFAULT NULL,
              `link_url` varchar(200) DEFAULT NULL,
              PRIMARY KEY (`ID`),
              UNIQUE KEY `ID_UNIQUE` (`ID`)
);

CREATE TABLE `jockeys` (
              `id` bigint(20) NOT NULL,
              `name` varchar(200) DEFAULT NULL,
              `link_url` varchar(200) DEFAULT NULL,
              PRIMARY KEY (`ID`),
              UNIQUE KEY `ID_UNIQUE` (`ID`)
);

CREATE TABLE `relation_horse_race` (
              `race_id` bigint(20) NOT NULL,
              `horse_id` bigint(20) NOT NULL,
              `race_rank` bigint(20) DEFAULT NULL,
              `frame_number` bigint(20) DEFAULT NULL,
              `horse_number` bigint(20) DEFAULT NULL,
              `sex` varchar(200) DEFAULT NULL,
              `age` bigint(20) DEFAULT NULL,
              `handicap` bigint(20) DEFAULT NULL,
              `jockey_id` bigint(20) DEFAULT NULL,
              `goal_time` varchar(200) DEFAULT NULL,
              `final_3f` decimal(6,3) DEFAULT NULL,
              `odds` decimal(6,3) DEFAULT NULL,
              `choice` bigint(20) DEFAULT NULL,
              `horse_weight` bigint(20) DEFAULT NULL,
              `weight_diff` bigint(20) DEFAULT NULL,
              PRIMARY KEY (`RACE_ID`, `HORSE_ID`),
              UNIQUE KEY `ID_UNIQUE` (`RACE_ID`, `HORSE_ID`)
);

CREATE USER `conn-user`@`%` IDENTIFIED BY 'password';
GRANT SELECT,INSERT,UPDATE,DELETE ON raceservice.* TO `conn-user`@`%`;
