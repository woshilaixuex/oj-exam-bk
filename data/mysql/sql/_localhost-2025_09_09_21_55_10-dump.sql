-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: ojexam
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account_table`
--

DROP TABLE IF EXISTS `account_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_table` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `account` varchar(12) NOT NULL,
  `student_number` char(10) NOT NULL,
  `password` varchar(12) NOT NULL,
  `exam_id` int NOT NULL,
  `score` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_account_student` (`account`,`student_number`)
) ENGINE=InnoDB AUTO_INCREMENT=168 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_table`
--

LOCK TABLES `account_table` WRITE;
/*!40000 ALTER TABLE `account_table` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `enroll_table`
--

DROP TABLE IF EXISTS `enroll_table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `enroll_table` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `student_number` char(10) NOT NULL,
  `name` varchar(20) NOT NULL,
  `qq_number` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `reason` varchar(100) NOT NULL,
  `grade` tinyint NOT NULL,
  `had_experience` tinyint(1) NOT NULL,
  `orientation` char(10) DEFAULT NULL,
  `experience` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_enroll_table_deleted_at` (`deleted_at`),
  KEY `idx_student_created` (`student_number`,`created_at` DESC)
) ENGINE=InnoDB AUTO_INCREMENT=169 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `enroll_table`
--

LOCK TABLES `enroll_table` WRITE;
/*!40000 ALTER TABLE `enroll_table` DISABLE KEYS */;
INSERT INTO `enroll_table` VALUES (1,'2025-08-20 17:18:00','2025-08-13 17:18:06','2025-08-13 17:18:07','2200720230','去','去','0','0',0,0,NULL,NULL),(168,'2025-08-13 19:08:50','2025-08-13 19:08:51','2025-08-13 19:08:51','2400320107','白文博','2',' 3','3',3,3,'3','3');
/*!40000 ALTER TABLE `enroll_table` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exam_account`
--

DROP TABLE IF EXISTS `exam_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exam_account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(30) NOT NULL,
  `password` varchar(30) NOT NULL,
  `student_id` varchar(30) DEFAULT NULL,
  `score` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=302 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exam_account`
--

LOCK TABLES `exam_account` WRITE;
/*!40000 ALTER TABLE `exam_account` DISABLE KEYS */;
INSERT INTO `exam_account` VALUES (201,'csd202501','gerN0KXvS/Y=','1',582),(202,'csd202502','VVPqj1v/Usk=','1',NULL),(203,'csd202503','1z9UfRDykz4=','1',NULL),(204,'csd202504','g7fsbUpB+yA=','1',NULL),(205,'csd202505','Eql8hIFsizA=','1',NULL),(206,'csd202506','FREIbP/wES8=','1',NULL),(207,'csd202507','bsiRnP4org4=','1',NULL),(208,'csd202508','TMzqZazhKiI=','1',582),(209,'csd202509','dF7pRmASsck=','1',NULL),(210,'csd2025010','BaiZTyu1tt0=','1',NULL),(211,'csd2025011','g+oZwrX74W8=','1',NULL),(212,'csd2025012','4p6psfup1vA=','1',NULL),(213,'csd2025013','Zk08ytnJcl0=','1',NULL),(214,'csd2025014','i2ZjIGG8FjI=','1',NULL),(215,'csd2025015','Ft3xP1xr9ns=','1',NULL),(216,'csd2025016','kfK+SKIYaVY=','1',NULL),(217,'csd2025017','V+FWKWb3cBU=','1',NULL),(218,'csd2025018','zVoQdtQlz/E=','1',NULL),(219,'csd2025019','SvgVX9hSNj4=','1',NULL),(220,'csd2025020','qopq6nhUfdA=','1',NULL),(221,'csd2025021','uCUBaYt3pUs=','1',NULL),(222,'csd2025022','fAnX+VK09Yk=','1',NULL),(223,'csd2025023','/DOfcvt0Fpc=','1',NULL),(224,'csd2025024','456nYkr+9ao=','1',NULL),(225,'csd2025025','uq6oxBPE+ic=','1',NULL),(226,'csd2025026','cbawC/t3TUs=','1',NULL),(227,'csd2025027','EbW/E+fDCec=','1',NULL),(228,'csd2025028','bOE6KQGrohI=','1',NULL),(229,'csd2025029','e6yoCpvios8=','1',NULL),(230,'csd2025030','U6Oh1bpHT8E=','1',NULL),(231,'csd2025031','Vpt0t59fNCk=','1',NULL),(232,'csd2025032','btc9hRRgxRk=','1',NULL),(233,'csd2025033','ueSQEeupnLc=','1',NULL),(234,'csd2025034','IYslr4ucXdY=','1',NULL),(235,'csd2025035','kvaKvP1bKZ8=','1',NULL),(236,'csd2025036','wTYaTE5P+E0=','1',NULL),(237,'csd2025037','hkySDoHouqc=','1',NULL),(238,'csd2025038','MgDgsEPAkyQ=','1',NULL),(239,'csd2025039','IcrtU5Iz8OQ=','1',NULL),(240,'csd2025040','CK/JezwpIaU=','1',NULL),(241,'csd2025041','rFpGYiPquq0=','1',NULL),(242,'csd2025042','Xkc1OJFrhiU=','1',NULL),(243,'csd2025043','aDeVR583eVY=','1',NULL),(244,'csd2025044','eDe0CyjGQ/A=','1',NULL),(245,'csd2025045','CSZOwXtW0Cw=','1',NULL),(246,'csd2025046','6lRC1M9k6ow=','1',NULL),(247,'csd2025047','yWzAm5bbV7E=','1',NULL),(248,'csd2025048','xPJ9vVUwt/8=','1',NULL),(249,'csd2025049','XmvqPdo836E=','1',NULL),(250,'csd2025050','LPF6MbV0LzI=','1',NULL),(251,'csd2025051','xlY1m8mQ8YI=','1',NULL),(252,'csd2025052','B++zIEwTTQo=','1',NULL),(253,'csd2025053','46anSY7Ed0Y=','1',NULL),(254,'csd2025054','5W8iui4wXSw=','1',NULL),(255,'csd2025055','8chm4NTs0O4=','1',NULL),(256,'csd2025056','Chg2lRKwrHk=','1',NULL),(257,'csd2025057','0Lx0vg94Tc4=','1',NULL),(258,'csd2025058','l0Ye0wIcOMA=','1',NULL),(259,'csd2025059','1MFgsI2exDk=','1',NULL),(260,'csd2025060','y2MhJDuqn5s=','1',NULL),(261,'csd2025061','L8F+iqLUeLE=','1',NULL),(262,'csd2025062','LoDUj04ki/E=','1',NULL),(263,'csd2025063','JfcwcFBC+Ag=','1',NULL),(264,'csd2025064','nDoFwB4yHho=','1',NULL),(265,'csd2025065','lPsU5ZEpit0=','1',NULL),(266,'csd2025066','3UxynGD0kAQ=','1',NULL),(267,'csd2025067','HKErbefUsvA=','1',NULL),(268,'csd2025068','oXamx939E78=','1',NULL),(269,'csd2025069','hCQjwtJJC5Q=','1',NULL),(270,'csd2025070','OHoYgGoZWBA=','1',NULL),(271,'csd2025071','8Yf98bPVPzU=','1',NULL),(272,'csd2025072','1lqQnO8vv/o=','1',NULL),(273,'csd2025073','QlHMCBkvhO8=','1',NULL),(274,'csd2025074','ReC5/O+KVFA=','1',NULL),(275,'csd2025075','XNEedQifW6o=','1',NULL),(276,'csd2025076','Yv/Ia2qao8g=','1',NULL),(277,'csd2025077','lbS8y4CyCEY=','1',NULL),(278,'csd2025078','caPtABLdKns=','1',NULL),(279,'csd2025079','PkjJK9rzt7U=','1',NULL),(280,'csd2025080','lK/Qw6w2MFQ=','1',NULL),(281,'csd2025081','XsWg5VJlZOs=','1',NULL),(282,'csd2025082','V0fE1s9k1aY=','1',NULL),(283,'csd2025083','yj8FSpPcWh0=','1',NULL),(284,'csd2025084','nexyA3K2DTo=','1',NULL),(285,'csd2025085','a0cxdbEI/B4=','1',NULL),(286,'csd2025086','oyl0xfGSFi8=','1',NULL),(287,'csd2025087','JwruabryAms=','1',NULL),(288,'csd2025088','nTv2PxYkVFQ=','1',NULL),(289,'csd2025089','pQbut//Rw10=','1',NULL),(290,'csd2025090','tzOq+FzNsY8=','1',NULL),(291,'csd2025091','EpcTSxOczlg=','1',NULL),(292,'csd2025092','fACnQNi5WPc=','1',NULL),(293,'csd2025093','5zDagR25P8Q=','1',NULL),(294,'csd2025094','tKHBF6WRLXc=','1',NULL),(295,'csd2025095','r7EOELehCgQ=','1',NULL),(296,'csd2025096','3kI18cFjI+k=','1',NULL),(297,'csd2025097','sJrSWQYUm1o=','1',NULL),(298,'csd2025098','IZA4KukSqko=','1',NULL),(299,'csd2025099','Qk4YFjzDHgA=','1',NULL),(300,'csd20250100','7ekphELGwgU=','1',NULL),(301,'csd24008','bai','2400320107',582);
/*!40000 ALTER TABLE `exam_account` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-09-09 21:55:11
