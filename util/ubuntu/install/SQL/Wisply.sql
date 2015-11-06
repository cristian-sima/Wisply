-- MySQL dump 10.13  Distrib 5.6.24, for Win32 (x86)
--
-- Host: localhost    Database: wisply
-- ------------------------------------------------------
-- Server version	5.6.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL COMMENT 'The name has this pattern [A-Za-z0-9\\s\\.]{3,25}',
  `password` binary(60) NOT NULL,
  `email` varchar(40) NOT NULL,
  `administrator` enum('false','true') NOT NULL DEFAULT 'false',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_2` (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `account_search`
--

DROP TABLE IF EXISTS `account_search`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_search` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` int(11) DEFAULT NULL,
  `query` varchar(200) NOT NULL,
  `timestamp` int(11) NOT NULL,
  `accessed` tinyint(1) NOT NULL COMMENT '0 if the search has not been accessed or 1 if the account accessed the search',
  PRIMARY KEY (`id`),
  KEY `account` (`account`),
  CONSTRAINT `account_search_ibfk_1` FOREIGN KEY (`account`) REFERENCES `account` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=420 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `account_token`
--

DROP TABLE IF EXISTS `account_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` int(11) NOT NULL,
  `value` varchar(200) NOT NULL,
  `timestamp` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `account` (`account`),
  CONSTRAINT `account_token_ibfk_1` FOREIGN KEY (`account`) REFERENCES `account` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=314 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `analyse`
--

DROP TABLE IF EXISTS `analyse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `analyse` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `start` int(11) NOT NULL,
  `end` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `digest_module`
--

DROP TABLE IF EXISTS `digest_module`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `digest_module` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `analyse` int(11) NOT NULL,
  `module` int(11) NOT NULL,
  `description` text NOT NULL,
  `formats` text NOT NULL,
  `keywords` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `analyse` (`analyse`),
  KEY `module` (`module`),
  CONSTRAINT `digest_module_ibfk_1` FOREIGN KEY (`analyse`) REFERENCES `analyse` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `digest_module_ibfk_2` FOREIGN KEY (`module`) REFERENCES `institution_module` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2478 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `digest_program`
--

DROP TABLE IF EXISTS `digest_program`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `digest_program` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `program` int(11) NOT NULL,
  `analyse` int(11) NOT NULL,
  `description` mediumtext NOT NULL,
  `formats` mediumtext NOT NULL,
  `keywords` mediumtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `program` (`program`),
  KEY `analyse` (`analyse`),
  CONSTRAINT `digest_program_ibfk_1` FOREIGN KEY (`analyse`) REFERENCES `analyse` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `digest_program_ibfk_2` FOREIGN KEY (`program`) REFERENCES `institution_program` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `download_table`
--

DROP TABLE IF EXISTS `download_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `download_table` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `description` varchar(1000) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `identifier`
--

DROP TABLE IF EXISTS `identifier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `identifier` (
  `identifier` varchar(1000) NOT NULL,
  `value` varchar(200) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repository` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `repository` (`repository`),
  KEY `identifier` (`identifier`(767)),
  CONSTRAINT `identifier_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=174604 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `identifier_set`
--

DROP TABLE IF EXISTS `identifier_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `identifier_set` (
  `identifier` varchar(1000) NOT NULL,
  `set_spec` varchar(3000) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repository` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `repository` (`repository`),
  CONSTRAINT `identifier_set_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=156930 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `institution`
--

DROP TABLE IF EXISTS `institution`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `institution` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(1000) NOT NULL,
  `url` varchar(2083) NOT NULL,
  `logo_URL` varchar(2083) NOT NULL,
  `wiki_URL` varchar(2083) NOT NULL,
  `wiki_ID` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `institution_module`
--

DROP TABLE IF EXISTS `institution_module`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `institution_module` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `content` text NOT NULL,
  `code` varchar(10) NOT NULL,
  `credits` varchar(5) NOT NULL COMMENT 'The value is representing the CATS credits',
  `year` varchar(2) NOT NULL,
  `institution` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `institution` (`institution`),
  CONSTRAINT `institution_module_ibfk_1` FOREIGN KEY (`institution`) REFERENCES `institution` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `institution_program`
--

DROP TABLE IF EXISTS `institution_program`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `institution_program` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `institution` int(11) NOT NULL,
  `title` varchar(200) NOT NULL,
  `code` varchar(10) NOT NULL,
  `year` varchar(4) NOT NULL,
  `ucas_code` varchar(20) NOT NULL,
  `level` enum('undergraduate','postgraduate') NOT NULL,
  `content` text NOT NULL,
  `subject` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `institution` (`institution`),
  KEY `subject` (`subject`),
  CONSTRAINT `institution_program_ibfk_1` FOREIGN KEY (`institution`) REFERENCES `institution` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `institution_program_ibfk_2` FOREIGN KEY (`subject`) REFERENCES `subject_area` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `institution_program_session`
--

DROP TABLE IF EXISTS `institution_program_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `institution_program_session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `program` int(11) NOT NULL,
  `module` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `module` (`module`),
  KEY `program` (`program`),
  CONSTRAINT `institution_program_session_ibfk_1` FOREIGN KEY (`program`) REFERENCES `institution_program` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `institution_program_session_ibfk_2` FOREIGN KEY (`module`) REFERENCES `institution_module` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `operation`
--

DROP TABLE IF EXISTS `operation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `process` int(11) NOT NULL,
  `content` varchar(1000) NOT NULL,
  `start` int(200) NOT NULL DEFAULT '0',
  `end` int(200) NOT NULL DEFAULT '0',
  `current_task` int(11) NOT NULL,
  `is_running` enum('true','false') NOT NULL DEFAULT 'true',
  `result` enum('danger','warning','success','normal') NOT NULL DEFAULT 'normal',
  PRIMARY KEY (`id`),
  KEY `process` (`process`),
  CONSTRAINT `operation_ibfk_1` FOREIGN KEY (`process`) REFERENCES `process` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=545 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `process`
--

DROP TABLE IF EXISTS `process`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `process` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` enum('Harvest','Analyse') NOT NULL,
  `start` int(11) NOT NULL DEFAULT '0',
  `end` int(11) NOT NULL DEFAULT '0',
  `is_running` enum('true','false') NOT NULL DEFAULT 'true',
  `current_operation` int(11) DEFAULT NULL,
  `result` enum('danger','warning','success','normal') NOT NULL DEFAULT 'normal',
  `is_suspended` enum('true','false') NOT NULL DEFAULT 'false',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `process_harvest`
--

DROP TABLE IF EXISTS `process_harvest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `process_harvest` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `process` int(11) NOT NULL,
  `repository` int(11) NOT NULL,
  `token_records` varchar(500) NOT NULL,
  `token_identifiers` varchar(500) NOT NULL,
  `token_collections` varchar(500) NOT NULL,
  `records` int(11) NOT NULL,
  `formats` int(11) NOT NULL,
  `collections` int(11) NOT NULL,
  `identifiers` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `process` (`process`),
  KEY `repository` (`repository`),
  CONSTRAINT `process_harvest_ibfk_1` FOREIGN KEY (`process`) REFERENCES `process` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `process_harvest_ibfk_2` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository`
--

DROP TABLE IF EXISTS `repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(500) NOT NULL,
  `url` varchar(2083) NOT NULL,
  `status` enum('unverified','verification-failed','ok','problems','verifying','updating','initializing','verified') NOT NULL,
  `institution` int(11) NOT NULL,
  `category` varchar(100) NOT NULL,
  `public_url` varchar(2083) NOT NULL,
  `filter` varchar(500) NOT NULL,
  `last_process` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `institution` (`institution`),
  CONSTRAINT `repository_ibfk_1` FOREIGN KEY (`institution`) REFERENCES `institution` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_collection`
--

DROP TABLE IF EXISTS `repository_collection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_collection` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spec` text NOT NULL,
  `path` varchar(500) NOT NULL,
  `description` text NOT NULL,
  `repository` int(11) NOT NULL,
  `number_of_records` int(11) NOT NULL,
  `name` varchar(300) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_2` (`id`),
  KEY `id` (`id`),
  KEY `repository` (`repository`),
  CONSTRAINT `repository_collection_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2997 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_email`
--

DROP TABLE IF EXISTS `repository_email`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_email` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(25) NOT NULL,
  `repository` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `id_2` (`id`),
  KEY `repository` (`repository`),
  CONSTRAINT `repository_email_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_format`
--

DROP TABLE IF EXISTS `repository_format`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_format` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `md_schema` varchar(1000) NOT NULL,
  `namespace` varchar(1000) NOT NULL,
  `prefix` varchar(1000) NOT NULL,
  `repository` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_2` (`id`),
  KEY `id` (`id`),
  KEY `repository` (`repository`),
  CONSTRAINT `repository_format_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=620 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_identification`
--

DROP TABLE IF EXISTS `repository_identification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_identification` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repository` int(11) NOT NULL,
  `protocol_version` varchar(10) NOT NULL,
  `earliest_datestamp` varchar(30) NOT NULL,
  `delete_policy` enum('persistent','transient','no') NOT NULL,
  `granularity` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `repository` (`repository`),
  CONSTRAINT `repository_identification_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_resource`
--

DROP TABLE IF EXISTS `repository_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repository` int(11) NOT NULL,
  `identifier` varchar(30) NOT NULL,
  `datestamp` varchar(30) NOT NULL,
  `is_visible` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'This field states if the repository can be accessed by Wisply or not. 1 if the resource can seen, 0 otherwise',
  PRIMARY KEY (`id`),
  KEY `repository` (`repository`),
  KEY `identifier` (`identifier`),
  CONSTRAINT `repository_resource_ibfk_1` FOREIGN KEY (`repository`) REFERENCES `repository` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=139505 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `resource_key`
--

DROP TABLE IF EXISTS `resource_key`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resource_key` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_key` varchar(200) NOT NULL,
  `repository` int(11) NOT NULL,
  `value` text NOT NULL,
  `resource` varchar(500) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `repository` (`repository`),
  KEY `resource` (`resource`)
) ENGINE=InnoDB AUTO_INCREMENT=1465119 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subject_area`
--

DROP TABLE IF EXISTS `subject_area`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(300) NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subject_area_definition`
--

DROP TABLE IF EXISTS `subject_area_definition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_area_definition` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subject` int(11) NOT NULL,
  `content` varchar(1000) NOT NULL,
  `source` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `subject` (`subject`),
  CONSTRAINT `subject_area_definition_ibfk_1` FOREIGN KEY (`subject`) REFERENCES `subject_area` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subject_area_ka`
--

DROP TABLE IF EXISTS `subject_area_ka`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subject_area_ka` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subject` int(11) NOT NULL,
  `code` varchar(10) NOT NULL,
  `content` text NOT NULL,
  `source` varchar(200) NOT NULL,
  `title` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `subject` (`subject`),
  CONSTRAINT `subject_area_ka_ibfk_1` FOREIGN KEY (`subject`) REFERENCES `subject_area` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `suggestion_resource`
--

DROP TABLE IF EXISTS `suggestion_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `suggestion_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `analyse` int(11) NOT NULL,
  `resource` varchar(500) NOT NULL,
  `module` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `analyse` (`analyse`),
  KEY `module` (`module`),
  CONSTRAINT `suggestion_resource_ibfk_1` FOREIGN KEY (`analyse`) REFERENCES `analyse` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `suggestion_resource_ibfk_3` FOREIGN KEY (`module`) REFERENCES `institution_module` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=650 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `start` int(11) NOT NULL DEFAULT '0',
  `end` int(11) NOT NULL DEFAULT '0',
  `content` text NOT NULL,
  `result` enum('danger','warning','success','normal') NOT NULL DEFAULT 'normal',
  `operation` int(11) NOT NULL,
  `is_running` enum('true','false') NOT NULL DEFAULT 'true',
  `process` int(11) NOT NULL,
  `explication` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `operation` (`operation`),
  KEY `process` (`process`),
  CONSTRAINT `task_ibfk_1` FOREIGN KEY (`operation`) REFERENCES `operation` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `task_ibfk_2` FOREIGN KEY (`process`) REFERENCES `process` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=9056 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-11-06  3:28:00
