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
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (15,'Cristian Sima','$2a$10$Gd/sGXRIli/UrZ/ZTLkpJeUn0tOnpopLqsSuk9vDzQ0Xy8a6TiTJi','cristian.sima93@yahoo.com','true'),(16,'Dr Su White','$2a$10$MjzVUPF5hK3q7ERKi7Kio..ktdMo6eFn/5igfc6u2RopewDGDm1yy','saw@ecs.soton.ac.uk','true');
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `account_token`
--

LOCK TABLES `account_token` WRITE;
/*!40000 ALTER TABLE `account_token` DISABLE KEYS */;
INSERT INTO `account_token` VALUES (1,15,'To5bS4P8kZqJggppls1sClMRaJ8=','1444495660');
/*!40000 ALTER TABLE `account_token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `history_event`
--

LOCK TABLES `history_event` WRITE;
/*!40000 ALTER TABLE `history_event` DISABLE KEYS */;
/*!40000 ALTER TABLE `history_event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `identifier`
--

LOCK TABLES `identifier` WRITE;
/*!40000 ALTER TABLE `identifier` DISABLE KEYS */;
/*!40000 ALTER TABLE `identifier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `identifier_set`
--

LOCK TABLES `identifier_set` WRITE;
/*!40000 ALTER TABLE `identifier_set` DISABLE KEYS */;
/*!40000 ALTER TABLE `identifier_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `institution`
--

LOCK TABLES `institution` WRITE;
/*!40000 ALTER TABLE `institution` DISABLE KEYS */;
INSERT INTO `institution` VALUES (1,'University of Southampton','The University of Southampton (occasionally abbreviated as Soton) is a public research university located in Southampton, England. Southampton is a research intensive university and a founding member of the Russell Group of elite British universities.\nThe origins of the university date back to the founding of the Hartley Institution in 1862 following a legacy to the Corporation of Southampton by Henry Robinson Hartley. In 1902, the Institution developed into the Hartley University College, with degrees awarded by the University of London. On 29 April 1952, the institution was granted a Royal Charter to give the University of Southampton full university status. It is a member of the European University Association, the Association of Commonwealth Universities and is an accredited institution of the Worldwide Universities Network.','http://www.southampton.ac.uk/','https://upload.wikimedia.org/wikipedia/en/thumb/9/9e/Southampton_crest.png/100px-Southampton_crest.png','https://en.wikipedia.org/wiki/University_of_Southampton','NULL');
/*!40000 ALTER TABLE `institution` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `operation`
--

LOCK TABLES `operation` WRITE;
/*!40000 ALTER TABLE `operation` DISABLE KEYS */;
/*!40000 ALTER TABLE `operation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `process`
--

LOCK TABLES `process` WRITE;
/*!40000 ALTER TABLE `process` DISABLE KEYS */;
/*!40000 ALTER TABLE `process` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `process_harvest`
--

LOCK TABLES `process_harvest` WRITE;
/*!40000 ALTER TABLE `process_harvest` DISABLE KEYS */;
/*!40000 ALTER TABLE `process_harvest` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository`
--

LOCK TABLES `repository` WRITE;
/*!40000 ALTER TABLE `repository` DISABLE KEYS */;
INSERT INTO `repository` VALUES (16,'EdShare','EdShare is the teaching and learning repository for your University. It was created by the University\'s Learning Societies Lab, and is a University wide digital resource for all staff and students, and much of its content is also open to the rest of the w','http://edshare.soton.ac.uk/cgi/oai2','ok',1,'EPrints','http://edshare.soton.ac.uk',0);
/*!40000 ALTER TABLE `repository` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository_collection`
--

LOCK TABLES `repository_collection` WRITE;
/*!40000 ALTER TABLE `repository_collection` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_collection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository_email`
--

LOCK TABLES `repository_email` WRITE;
/*!40000 ALTER TABLE `repository_email` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_email` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository_format`
--

LOCK TABLES `repository_format` WRITE;
/*!40000 ALTER TABLE `repository_format` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_format` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository_identification`
--

LOCK TABLES `repository_identification` WRITE;
/*!40000 ALTER TABLE `repository_identification` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_identification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `repository_resource`
--

LOCK TABLES `repository_resource` WRITE;
/*!40000 ALTER TABLE `repository_resource` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `resource_collection`
--

LOCK TABLES `resource_collection` WRITE;
/*!40000 ALTER TABLE `resource_collection` DISABLE KEYS */;
/*!40000 ALTER TABLE `resource_collection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `resource_key`
--

LOCK TABLES `resource_key` WRITE;
/*!40000 ALTER TABLE `resource_key` DISABLE KEYS */;
/*!40000 ALTER TABLE `resource_key` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2015-10-13 12:47:46
