CREATE DATABASE  IF NOT EXISTS `payment` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `payment`;
-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: payment
-- ------------------------------------------------------
-- Server version	5.7.18

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
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `organisation` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `payment_id` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment`
--

LOCK TABLES `payment` WRITE;
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
INSERT INTO `payment` VALUES (1,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345671','2017-05-18 13:50:19','2017-05-18 13:50:19'),
                             (2,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345672','2017-05-18 13:50:19','2017-05-18 13:50:19'),
                             (3,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345673','2017-05-18 13:50:19','2017-05-18 13:50:19'),
                             (4,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345674','2017-05-18 13:50:19','2017-05-18 13:50:19'),
                             (5,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345675','2017-05-18 13:50:19','2017-05-18 13:50:19'),
                             (6,'43d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb','123456789012345676','2017-05-18 13:50:19','2017-05-18 13:50:19');
UNLOCK TABLES;
