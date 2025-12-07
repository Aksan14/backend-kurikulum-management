-- MySQL dump 10.13  Distrib 8.0.44, for Linux (x86_64)
--
-- Host: localhost    Database: akademik-management
-- ------------------------------------------------------
-- Server version	8.0.44-0ubuntu0.22.04.1

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
-- Table structure for table `audit_logs`
--

DROP TABLE IF EXISTS `audit_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `audit_logs` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `action` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `table_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `record_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `old_values` json DEFAULT NULL,
  `new_values` json DEFAULT NULL,
  `ip_address` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_audit_logs_user` (`user_id`),
  KEY `idx_audit_logs_table` (`table_name`),
  KEY `idx_audit_logs_record` (`record_id`),
  KEY `idx_audit_logs_created_at` (`created_at`),
  CONSTRAINT `fk_audit_logs_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `audit_logs`
--

LOCK TABLES `audit_logs` WRITE;
/*!40000 ALTER TABLE `audit_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `audit_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cpl`
--

DROP TABLE IF EXISTS `cpl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cpl` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deskripsi` text COLLATE utf8mb4_unicode_ci,
  `status` enum('draft','published','archived') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'draft',
  `version` int NOT NULL DEFAULT '1',
  `created_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_cpl_kode` (`kode`),
  KEY `idx_cpl_status` (`status`),
  KEY `idx_cpl_deleted_at` (`deleted_at`),
  KEY `fk_cpl_created_by` (`created_by`),
  CONSTRAINT `fk_cpl_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cpl`
--

LOCK TABLES `cpl` WRITE;
/*!40000 ALTER TABLE `cpl` DISABLE KEYS */;
INSERT INTO `cpl` VALUES ('0916a26d-7141-4d1c-802c-e3aab221bc8a','CPL003','VPPPPPP','TESTESTESTES','published',2,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:28:56','2025-12-05 13:51:08',NULL),('eab004ce-0752-4447-9700-c71b25c4e050','CPL-TES02','TES-02','TESTESTESTESTES','published',7,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 11:22:12','2025-12-05 12:00:56',NULL);
/*!40000 ALTER TABLE `cpl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cpl_assignments`
--

DROP TABLE IF EXISTS `cpl_assignments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cpl_assignments` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpl_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `dosen_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mata_kuliah` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mata_kuliah_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deadline` timestamp NULL DEFAULT NULL,
  `status` enum('assigned','accepted','rejected','completed','cancelled') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'assigned',
  `catatan` text COLLATE utf8mb4_unicode_ci,
  `rejection_reason` text COLLATE utf8mb4_unicode_ci,
  `assigned_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `assigned_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `response_at` timestamp NULL DEFAULT NULL,
  `completed_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_cpl_assignments_cpl` (`cpl_id`),
  KEY `idx_cpl_assignments_dosen` (`dosen_id`),
  KEY `idx_cpl_assignments_status` (`status`),
  KEY `idx_cpl_assignments_deleted_at` (`deleted_at`),
  KEY `fk_cpl_assignments_assigned_by` (`assigned_by`),
  CONSTRAINT `fk_cpl_assignments_assigned_by` FOREIGN KEY (`assigned_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_cpl_assignments_cpl` FOREIGN KEY (`cpl_id`) REFERENCES `cpl` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cpl_assignments_dosen` FOREIGN KEY (`dosen_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cpl_assignments`
--

LOCK TABLES `cpl_assignments` WRITE;
/*!40000 ALTER TABLE `cpl_assignments` DISABLE KEYS */;
INSERT INTO `cpl_assignments` VALUES ('97a4ae23-e9a4-49a0-99d5-082ea642c2ef','0916a26d-7141-4d1c-802c-e3aab221bc8a','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','aksan','574f8a65-fe10-4890-abc7-5f1876533926','2025-12-11 00:00:00','accepted','tidak ada',NULL,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-06 19:27:48','2025-12-06 19:28:29',NULL,'2025-12-06 19:27:48','2025-12-06 19:28:29',NULL);
/*!40000 ALTER TABLE `cpl_assignments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cpl_mk_mappings`
--

DROP TABLE IF EXISTS `cpl_mk_mappings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cpl_mk_mappings` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpl_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mata_kuliah_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `level` enum('tinggi','sedang','rendah') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_cpl_mk` (`cpl_id`,`mata_kuliah_id`),
  KEY `idx_cpl_mk_cpl_id` (`cpl_id`),
  KEY `idx_cpl_mk_mk_id` (`mata_kuliah_id`),
  KEY `idx_cpl_mk_level` (`level`),
  CONSTRAINT `fk_cpl_mk_mapping_cpl` FOREIGN KEY (`cpl_id`) REFERENCES `cpl` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cpl_mk_mapping_mk` FOREIGN KEY (`mata_kuliah_id`) REFERENCES `mata_kuliah` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cpl_mk_mappings`
--

LOCK TABLES `cpl_mk_mappings` WRITE;
/*!40000 ALTER TABLE `cpl_mk_mappings` DISABLE KEYS */;
INSERT INTO `cpl_mk_mappings` VALUES ('294354f1-5f97-4e4c-99e8-5f67a48435a2','eab004ce-0752-4447-9700-c71b25c4e050','7f01fd62-3fdc-455c-8987-9fad9d4debc4','tinggi','2025-12-05 13:54:32','2025-12-05 13:54:32'),('2c14824f-e753-4000-abb3-6e219ac2aa69','0916a26d-7141-4d1c-802c-e3aab221bc8a','7f01fd62-3fdc-455c-8987-9fad9d4debc4','tinggi','2025-12-05 13:54:34','2025-12-05 13:54:34'),('35406972-7cf5-47c3-9390-d04b691396f8','0916a26d-7141-4d1c-802c-e3aab221bc8a','33f3ed4c-1751-41c1-a70d-1b311e6bd20a','tinggi','2025-12-05 13:54:25','2025-12-05 13:54:25'),('7cf376cf-2712-4a5b-8abd-9f55bae23551','eab004ce-0752-4447-9700-c71b25c4e050','d095e618-6a84-4949-bbf0-fdf687ed98f8','tinggi','2025-12-05 13:54:32','2025-12-05 13:54:32'),('fe4237e7-657b-46d7-a304-2b95d6f51501','0916a26d-7141-4d1c-802c-e3aab221bc8a','d095e618-6a84-4949-bbf0-fdf687ed98f8','tinggi','2025-12-05 13:54:34','2025-12-05 13:54:34');
/*!40000 ALTER TABLE `cpl_mk_mappings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cpmk_cpl_mapping`
--

DROP TABLE IF EXISTS `cpmk_cpl_mapping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cpmk_cpl_mapping` (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpmk_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpl_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_cpmk_cpl` (`cpmk_id`,`cpl_id`),
  KEY `idx_cpmk_cpl_cpmk_id` (`cpmk_id`),
  KEY `idx_cpmk_cpl_cpl_id` (`cpl_id`),
  CONSTRAINT `cpmk_cpl_mapping_ibfk_1` FOREIGN KEY (`cpmk_id`) REFERENCES `rps_cpmk` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cpmk_cpl_mapping_ibfk_2` FOREIGN KEY (`cpl_id`) REFERENCES `cpl` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cpmk_cpl_mapping`
--

LOCK TABLES `cpmk_cpl_mapping` WRITE;
/*!40000 ALTER TABLE `cpmk_cpl_mapping` DISABLE KEYS */;
/*!40000 ALTER TABLE `cpmk_cpl_mapping` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `document_templates`
--

DROP TABLE IF EXISTS `document_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `document_templates` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deskripsi` text COLLATE utf8mb4_unicode_ci,
  `sections` json DEFAULT NULL,
  `file_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `version` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1.0',
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_document_templates_is_active` (`is_active`),
  KEY `fk_document_templates_created_by` (`created_by`),
  CONSTRAINT `fk_document_templates_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_templates`
--

LOCK TABLES `document_templates` WRITE;
/*!40000 ALTER TABLE `document_templates` DISABLE KEYS */;
/*!40000 ALTER TABLE `document_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `files`
--

DROP TABLE IF EXISTS `files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `files` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `original_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `stored_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_path` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_type` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mime_type` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` bigint NOT NULL,
  `uploaded_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `reference_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_files_uploaded_by` (`uploaded_by`),
  KEY `idx_files_reference` (`reference_type`,`reference_id`),
  CONSTRAINT `fk_files_uploaded_by` FOREIGN KEY (`uploaded_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `files`
--

LOCK TABLES `files` WRITE;
/*!40000 ALTER TABLE `files` DISABLE KEYS */;
/*!40000 ALTER TABLE `files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `generated_documents`
--

DROP TABLE IF EXISTS `generated_documents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `generated_documents` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `template_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `template_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tahun` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('pending','processing','ready','failed') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pending',
  `file_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `file_type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pdf',
  `file_size` bigint DEFAULT NULL,
  `sections` json DEFAULT NULL,
  `generation_data` json DEFAULT NULL,
  `progress` int NOT NULL DEFAULT '0',
  `error_message` text COLLATE utf8mb4_unicode_ci,
  `created_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `completed_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_generated_documents_template` (`template_id`),
  KEY `idx_generated_documents_status` (`status`),
  KEY `idx_generated_documents_tahun` (`tahun`),
  KEY `fk_generated_documents_created_by` (`created_by`),
  CONSTRAINT `fk_generated_documents_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_generated_documents_template` FOREIGN KEY (`template_id`) REFERENCES `document_templates` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `generated_documents`
--

LOCK TABLES `generated_documents` WRITE;
/*!40000 ALTER TABLE `generated_documents` DISABLE KEYS */;
/*!40000 ALTER TABLE `generated_documents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mata_kuliah`
--

DROP TABLE IF EXISTS `mata_kuliah`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mata_kuliah` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sks` int NOT NULL DEFAULT '3',
  `semester` int NOT NULL,
  `jenis` enum('wajib','pilihan') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'wajib',
  `deskripsi` text COLLATE utf8mb4_unicode_ci,
  `prasyarat` json DEFAULT NULL,
  `dosen_pengampu_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `koordinator_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `status` enum('aktif','nonaktif','dihapus') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'aktif',
  `created_by` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_mata_kuliah_kode` (`kode`),
  KEY `idx_mata_kuliah_semester` (`semester`),
  KEY `idx_mata_kuliah_dosen` (`dosen_pengampu_id`),
  KEY `idx_mata_kuliah_koordinator` (`koordinator_id`),
  KEY `idx_mata_kuliah_deleted_at` (`deleted_at`),
  KEY `fk_mata_kuliah_created_by` (`created_by`),
  CONSTRAINT `fk_mata_kuliah_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_mata_kuliah_dosen` FOREIGN KEY (`dosen_pengampu_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_mata_kuliah_koordinator` FOREIGN KEY (`koordinator_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mata_kuliah`
--

LOCK TABLES `mata_kuliah` WRITE;
/*!40000 ALTER TABLE `mata_kuliah` DISABLE KEYS */;
INSERT INTO `mata_kuliah` VALUES ('33f3ed4c-1751-41c1-a70d-1b311e6bd20a','MK-001','Pemrograman Web',3,3,'wajib',NULL,'null',NULL,NULL,1,'aktif','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','2025-12-04 21:06:15','2025-12-06 02:52:48',NULL),('574f8a65-fe10-4890-abc7-5f1876533926','SASASASASA','aksan',1,1,'wajib','sasa','null','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae',0,'aktif','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:38:42','2025-12-06 04:58:57',NULL),('7f01fd62-3fdc-455c-8987-9fad9d4debc4','MK011111','Rekayasa Perangkat Lunak',3,8,'wajib','MATAKULIAH WAJIB','null',NULL,NULL,0,'dihapus','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:05:14','2025-12-05 14:39:15',NULL),('b8ecc5e3-f924-419f-8fba-9084da126114','TIF501','Machine Learning',3,5,'pilihan','Pengantar machine learning','null',NULL,NULL,0,'dihapus','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 14:35:48','2025-12-05 14:39:42',NULL),('d095e618-6a84-4949-bbf0-fdf687ed98f8','MK-TEST-001','Pemrograman Web Test',3,3,'wajib',NULL,'null',NULL,NULL,1,'aktif','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','2025-12-04 21:08:39','2025-12-04 21:08:39',NULL);
/*!40000 ALTER TABLE `mata_kuliah` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notifications`
--

DROP TABLE IF EXISTS `notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notifications` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` enum('info','warning','success','error','assignment','approval','rejection') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'info',
  `reference_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_read` tinyint(1) NOT NULL DEFAULT '0',
  `read_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `action_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_notifications_user` (`user_id`),
  KEY `idx_notifications_is_read` (`is_read`),
  KEY `idx_notifications_type` (`type`),
  KEY `idx_notifications_reference` (`reference_type`,`reference_id`),
  CONSTRAINT `fk_notifications_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notifications`
--

LOCK TABLES `notifications` WRITE;
/*!40000 ALTER TABLE `notifications` DISABLE KEYS */;
INSERT INTO `notifications` VALUES ('023faf5d-608e-430d-b1c3-65036362b5f7','526e30af-f07e-4272-a63e-753ad48ab6c7','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web','assignment','cpl_assignment','52cc09b6-57e6-4d0a-acc6-49d301f33bbe',0,NULL,'2025-12-05 12:41:44',NULL),('14cd1d8f-ab58-4ec7-b752-bc086c82cbe9','526e30af-f07e-4272-a63e-753ad48ab6c7','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web','assignment','cpl_assignment','5a268d53-8c51-4c50-a8ee-57d09ed1e846',0,NULL,'2025-12-05 12:43:40',NULL),('e0bd0510-f892-42bf-9a01-77938cf7d83e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','Revisi RPS Diminta','Kaprodi meminta revisi RPS untuk mata kuliah undefined. Catatan: revisi','rejection','rps','99abcad5-d3f5-42a5-b201-902a606278c4',0,NULL,'2025-12-07 15:14:34',NULL);
/*!40000 ALTER TABLE `notifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refresh_tokens`
--

DROP TABLE IF EXISTS `refresh_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `refresh_tokens` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expires_at` timestamp NOT NULL,
  `revoked` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_refresh_tokens_user_id` (`user_id`),
  KEY `idx_refresh_tokens_token` (`token`(255)),
  CONSTRAINT `fk_refresh_tokens_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refresh_tokens`
--

LOCK TABLES `refresh_tokens` WRITE;
/*!40000 ALTER TABLE `refresh_tokens` DISABLE KEYS */;
INSERT INTO `refresh_tokens` VALUES ('01e80b51-54d4-4263-8b51-2a0ed22f2026','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','f19168e8e3fcfff2488624fc32cf63c87a299d904d60fef89e5d93ed7c6fa8e6','2025-12-14 02:39:09',1,'2025-12-07 02:39:09'),('03bb860a-8b45-4e44-9b04-7ce652248daa','d713ddfb-1a39-4b33-9826-8be0814a3f57','16c5a39ef4fa97d74b5a20f273e9c2d79daa765bef1b57520ddf4f76577f093e','2025-12-14 10:57:09',1,'2025-12-07 10:57:09'),('05646387-ce87-4e93-9f5a-751c8126e9a0','2b704681-a615-43cb-8793-9e9e2d754ecf','235933db1acab8fd28992815822af90be7cfc03fbb2e4453f69d86fc0cce802e','2025-12-11 20:55:46',0,'2025-12-04 20:55:46'),('0616a58e-0021-4157-8e6d-3d5b1752161d','d713ddfb-1a39-4b33-9826-8be0814a3f57','f1fcfe872f48263342e4b71b38e1c18e5a3a7fe62170c0322f0b02a406562bb7','2025-12-14 13:17:48',1,'2025-12-07 13:17:48'),('06cbd02b-b97e-4a94-abaa-cafb1d45d220','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','d88b9b5d4a8219c290011d402dccf4a1fcf48527c8984fa74344a8f4dd9942b3','2025-12-12 10:38:03',1,'2025-12-05 10:38:03'),('0ae33508-8d62-4ff0-93bc-13e4918a6bc1','d713ddfb-1a39-4b33-9826-8be0814a3f57','48bbd417aee90b73997b2b0accfd8159c5a22f287b95bd7e68ff4d1f33e92b20','2025-12-13 07:48:50',1,'2025-12-06 07:48:50'),('0b1e2976-c09d-4a5e-9a97-5ded2c898355','d713ddfb-1a39-4b33-9826-8be0814a3f57','44adde05725aa2644d8c25c9235f35f503ee9e4af59b7bdc17a9b43ab904d2dc','2025-12-12 14:35:05',1,'2025-12-05 14:35:05'),('0e1a8875-69c1-4844-a525-0e0b0e8fa42a','d713ddfb-1a39-4b33-9826-8be0814a3f57','ce644095db77084fb7ffb671f962e4181a51109e0c0e299c6be0a52560836d23','2025-12-14 15:05:50',1,'2025-12-07 15:05:50'),('0f3adefe-1d57-4ab6-8c72-0608ec0537d5','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','9941335e83f784333a6fc5c2b74ef0b772731a97a4d35a2c7a38ecd36a84072f','2025-12-11 21:06:15',1,'2025-12-04 21:06:15'),('0faf708e-dcc8-43c2-9fc0-75ca5d6875ae','d713ddfb-1a39-4b33-9826-8be0814a3f57','14394c7bafbfc53c15154b83e52f71b8f765de89e200315b7929314193239458','2025-12-14 13:38:39',1,'2025-12-07 13:38:39'),('148923d4-ed4b-47dc-a18e-34078cc7933f','d713ddfb-1a39-4b33-9826-8be0814a3f57','99aa74dc4886c1dec68ce3485a028830a20bc86b5d904cbf865394b4fefddf5a','2025-12-14 13:49:28',1,'2025-12-07 13:49:28'),('15bca504-4941-4912-b3e2-c912790466b2','d713ddfb-1a39-4b33-9826-8be0814a3f57','27e7078ea3826d0c18fbcdd71fc6d6e7416bfc677cf790b95ebb692c8b309acc','2025-12-14 02:09:18',1,'2025-12-07 02:09:18'),('1612f456-fe16-45d8-8d9d-924a736f168d','d713ddfb-1a39-4b33-9826-8be0814a3f57','b3e0051ae12415ea684d13c56eaf4ef360fa6c695b58d378ad450b35bd32c77e','2025-12-13 01:06:02',1,'2025-12-06 01:06:02'),('1dba79bc-7d28-4b38-b452-0c8de888b150','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','e43039ec18678291968cb248cb19eb80f2982901fe8d9454f1afc6c8cf27fbc6','2025-12-11 21:06:15',1,'2025-12-04 21:06:15'),('1f8d1240-d79d-444a-8503-7fb0228782f9','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','93ee46f11ca319413419ab4b39063f905a2be6d4085a6ad65f961edf67758b42','2025-12-13 10:03:37',1,'2025-12-06 10:03:37'),('22d3e1de-81c8-4f9b-96a6-42b4d901d313','d713ddfb-1a39-4b33-9826-8be0814a3f57','d378f80528df623dc0680fcc3002d5ff9851a47a0f52d960640b2507a51df7cf','2025-12-12 13:58:43',1,'2025-12-05 13:58:43'),('25b08c47-b821-4a60-b20f-29cdb6bda556','d713ddfb-1a39-4b33-9826-8be0814a3f57','a5319a87c9d44db69cf6216e68e3caf2a000255daffa62523acb5044a587bfef','2025-12-13 19:26:33',1,'2025-12-06 19:26:33'),('25de28c1-ad73-46ab-98cd-821987cf931f','526e30af-f07e-4272-a63e-753ad48ab6c7','7d45066b3c9ced1a6a3573d1d83060f135ed5556585769bb254e830cb5c07be8','2025-12-12 09:29:32',1,'2025-12-05 09:29:32'),('27bff371-d9e1-4af9-a214-7842f39b69f0','d713ddfb-1a39-4b33-9826-8be0814a3f57','5a4e1f54570a446f92fd65f8119349eef7bc5063e95e6b548a5c23e20bc97f6e','2025-12-14 13:45:20',1,'2025-12-07 13:45:20'),('30021d44-e48b-4bf6-a2e5-f697f662e3f1','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','4e8e6f813595838b8f9c7753ac6f83ea390bf205cf4c48a6b245a03f907d251b','2025-12-13 11:16:09',1,'2025-12-06 11:16:09'),('304d2b81-c309-4f38-badc-b11edbe01333','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','a33761bf195da14bcd974f73862a05a9974c44806d8f8fb88fdcc097ecd5bdd6','2025-12-14 15:05:08',1,'2025-12-07 15:05:08'),('320cbd3f-669b-4bdf-82b0-844b50eb83e1','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c4bf20d83c81c16c107410ace9a0ed93b0054384fc9bf6f8f951d59da5562e2b','2025-12-12 10:12:59',1,'2025-12-05 10:12:59'),('3261b298-c881-4906-ad1a-4771ec3a3233','d713ddfb-1a39-4b33-9826-8be0814a3f57','b89b693a8fffa69882120ab20efaf5e8d4b2b3d6640db9b6d9ef2118de8c2f35','2025-12-12 14:18:00',1,'2025-12-05 14:18:00'),('34b5928c-1fac-42eb-90d0-1c76554fcea7','d713ddfb-1a39-4b33-9826-8be0814a3f57','9bffb1811118d046455f3f82bc6b5e27461018f2f9d81aa90df5f796bb51be80','2025-12-13 01:06:13',1,'2025-12-06 01:06:13'),('34bf52dc-34d9-4d4c-aca9-6e3c5903fd54','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','cae3ed817a653fc69f2cdadbad3922cbaf192218f935851c3c4b20b7b0c64bd9','2025-12-11 21:08:39',1,'2025-12-04 21:08:39'),('34d11e22-f83d-43c2-abed-e52e20080873','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','05f09cc59da47454a4f83d759a025d06f8c62d328b70bb6eebd9702f44da331e','2025-12-13 11:16:14',1,'2025-12-06 11:16:14'),('3785815c-2787-4cbd-9604-bfd45c5412ff','526e30af-f07e-4272-a63e-753ad48ab6c7','f5bd8bff9205b0ffdd9af8b6437340c672a56cf872c6db6b46fc9771274acb29','2025-12-11 22:35:08',1,'2025-12-04 22:35:08'),('3fab19d4-684a-4068-ab08-a0183663e590','2b704681-a615-43cb-8793-9e9e2d754ecf','ad382034e5edf3e20ab940c3d6e6b72a2f49ab0399a74db1da60a7f0877f590a','2025-12-11 20:55:19',1,'2025-12-04 20:55:19'),('4131c532-136b-4068-a9d9-9ed767d294ad','d713ddfb-1a39-4b33-9826-8be0814a3f57','26a257da42b34b2d91bd684865d750d21e0e237c91d5ba4c9d45db46deece81b','2025-12-12 12:41:06',1,'2025-12-05 12:41:06'),('423dcb82-548e-4137-8d74-65e3e2441137','d713ddfb-1a39-4b33-9826-8be0814a3f57','df80562d5aa1eb3bf77f10cd0079a2060f54eed6b46c8a603d78e3d1bbf0517a','2025-12-14 12:24:51',1,'2025-12-07 12:24:51'),('440637fa-d6aa-4891-a959-099a730c20dc','d713ddfb-1a39-4b33-9826-8be0814a3f57','c94ecf426468f9084a334721772312419210130d575ab21c66daa3ed46b665bc','2025-12-12 10:21:33',1,'2025-12-05 10:21:33'),('467fa467-085c-4398-9286-d57126ffb1da','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','0261685d6aa967a71ca82f1d62783b4de0ced660f3c937313e767bc401d0c394','2025-12-11 20:56:29',1,'2025-12-04 20:56:29'),('49178b79-8309-4c44-bc1e-184879c350e1','d713ddfb-1a39-4b33-9826-8be0814a3f57','ef6e4382c5f148ea1f21e17159193eda173b723eef38e2eb39f43cb254141f93','2025-12-13 07:20:34',1,'2025-12-06 07:20:34'),('4d934c65-8008-4a63-9681-7ece7f85fa37','d713ddfb-1a39-4b33-9826-8be0814a3f57','cdc56db10a05db73b4c779daad89350e9e61290b87a335c6fb01bc61082a293a','2025-12-12 12:09:19',1,'2025-12-05 12:09:19'),('4dbf76d3-4c9f-49be-be9e-419474e84110','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','414bfd5cf55e02a3cb52e1834e615b3e2bfb5582848f510ae567c3a53fcc8fd5','2025-12-13 02:56:15',1,'2025-12-06 02:56:15'),('5036c665-2017-4dc7-8b8c-37f226020b3f','d713ddfb-1a39-4b33-9826-8be0814a3f57','ca197aebd58eee2e98ff396eba6eaa23a0ced6a47ad34c1b4ea19d1c3fffc7bf','2025-12-12 12:53:28',1,'2025-12-05 12:53:28'),('50d6e2bb-8d1f-4870-9c6b-2390e0de891f','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','8667cf9ee8358e3c13d3114120eef0cdf3071b7c2a73ce10e3a9da46f0870272','2025-12-13 10:03:32',1,'2025-12-06 10:03:32'),('516bb601-fcd0-4642-b036-b441f3c9d298','d713ddfb-1a39-4b33-9826-8be0814a3f57','dbb7f2f5ab92ccdb9860e77b7da19a61fcb5ee967016edd8fe0bc39837a57740','2025-12-13 07:58:26',1,'2025-12-06 07:58:26'),('54fd03c2-fc5f-48a9-8307-80337830e5dc','d713ddfb-1a39-4b33-9826-8be0814a3f57','591f66f95528ece2dfa54b7e3cd5be878a3fe8315040abce3aaf233ac6e0d511','2025-12-14 12:08:23',1,'2025-12-07 12:08:23'),('57847e5c-2ef2-4f19-a8fc-262e2ef90f80','d713ddfb-1a39-4b33-9826-8be0814a3f57','009d9d94867caf0404e171cd9f0829b2118d1bfbde8ae1e0614c864b3947dab4','2025-12-14 12:10:18',1,'2025-12-07 12:10:18'),('57b27084-0066-40e3-b61f-fa4c731cbbc1','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','b9b943f6e812c7a1edab9a39ca2c51baf6e09e3defcba691109983a6162beccf','2025-12-14 10:26:05',1,'2025-12-07 10:26:05'),('5afd7058-8f5f-47ea-b0dc-396334a48c96','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','db1c6d0a256ece526f67c4fb9208d00001012d984dde64512e8a0510785dc1fa','2025-12-12 10:18:12',1,'2025-12-05 10:18:12'),('5bb67988-72a8-4b59-892e-212fc3c79504','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','ce6ca55cdb3dff50e52c5ade68bdcc1715d35afac7277a11455747d5d1bbbe54','2025-12-12 15:06:30',1,'2025-12-05 15:06:30'),('5ec4bbaf-a2ec-4aaf-95bd-85b2da96f829','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','4e72356b68d361de97a7b8acd824577c410f0d8fe60ad4506312fd6a1122381f','2025-12-13 07:52:09',1,'2025-12-06 07:52:09'),('5f02de90-20af-4afd-87fa-344ac6011f83','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','653da350f86accc19c8c1132700549dc00c61c62ad841ac7d66dee09870ac25e','2025-12-12 12:11:58',1,'2025-12-05 12:11:58'),('5f109e79-4b82-4822-86d8-5aa515ee062b','526e30af-f07e-4272-a63e-753ad48ab6c7','d6eeb09b7b21697da0b958ceefefa86ee5670e5677ca5960e0e9f9f7b0c6bb3f','2025-12-11 21:10:41',1,'2025-12-04 21:10:41'),('5f9d233c-9295-4eb5-bbc6-cbcd0ec73d82','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','df6eed1b6f7eece23ead26fa0015d241fa8049b5a25e7f80f98e0388d4cf2901','2025-12-13 17:14:34',1,'2025-12-06 17:14:34'),('642fa5d1-480b-4bb3-ae1e-594da9d4c78f','d713ddfb-1a39-4b33-9826-8be0814a3f57','a439209c224eb2d3e5156ef3a881ceaa7e6dbda6d79f74c5c1781dda26b6e89b','2025-12-13 07:50:37',1,'2025-12-06 07:50:37'),('64510f8f-2117-446f-8810-d6c2705e8607','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','2b361829be58aee8fba73d307214aa3ffbc05b72af0131b19080259578c56cf1','2025-12-13 15:02:10',1,'2025-12-06 15:02:10'),('65fea108-67b4-4910-a681-98e4e6a6380f','d713ddfb-1a39-4b33-9826-8be0814a3f57','0cb5b2ad3277ed078995a6392d945deeab9fd68d7fe42ce5c5456bf0116f90b6','2025-12-14 13:53:47',1,'2025-12-07 13:53:47'),('6ebe7cd2-c396-49ad-9cab-3613431f1fc6','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','619ff86ee21ef4abf5bd130f8763249606f13b6481880f2fa9289f13ede7b763','2025-12-13 19:38:59',1,'2025-12-06 19:38:59'),('70d8ca61-ef96-426b-a394-893cf0974e24','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','2025c786177be4bc1d875c579613a75354bcd894eba70d8367f41798dd4f0e44','2025-12-13 19:28:07',1,'2025-12-06 19:28:07'),('7204d667-448d-4234-97a8-9358ccb6faba','2233d21b-c4fe-43dd-aba9-91f5d5df99a7','0758fae0f56f92b9c7bec4fe674c634d210069401921e996cf9f8481b7ab3eee','2025-12-13 18:59:08',0,'2025-12-06 18:59:08'),('73d83ac4-9580-43d1-a83c-bbb700cfb525','d713ddfb-1a39-4b33-9826-8be0814a3f57','8c83a1e3fdaea8ee72e584c2b22d5e9f0c1c8f95f92e8d08a894bb549ea8d834','2025-12-12 13:32:10',1,'2025-12-05 13:32:10'),('7785f4c5-8147-47fe-9a55-3feebe3cb8fb','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','eec5ac4e10084795e6d2c26acef8f79eff7598fcfad0b4389d3d1674210cce9a','2025-12-13 19:35:48',1,'2025-12-06 19:35:48'),('77f07539-15af-45af-9350-f65b456b29e4','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','7106c2191b91ce9cd3b1a27fb0b6d420beb5a5b13f8c2db4857b2656bc7c7334','2025-12-12 12:52:33',1,'2025-12-05 12:52:33'),('7946f84d-e517-4173-a307-ed7dc93f2553','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','fe06bba6607c344ae1e18521090588d7d4778ad779c8f9c070516fcbd30d56d3','2025-12-13 15:37:30',1,'2025-12-06 15:37:30'),('7df60dd0-c047-46a8-a64b-6cccd0848d3a','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','01d89d668d3ad588a6d5ea29c46e232e7ecdc27f2578a5e5974386d3854d3b96','2025-12-14 12:06:55',1,'2025-12-07 12:06:55'),('7fffca23-e2e1-4976-8d8e-3325180bbe36','d713ddfb-1a39-4b33-9826-8be0814a3f57','4f30b96f5d45047e36a09ccaf42e4a098a3b231d338c1033405b823f07fe36d1','2025-12-13 05:51:37',1,'2025-12-06 05:51:37'),('80ce3965-050e-4796-b5b7-1ad3a6ea6794','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','185171794a623cf2c0ecedac3098fbee1422c24fe1200442fce68cbb7d43b235','2025-12-13 16:05:26',1,'2025-12-06 16:05:26'),('81019270-c6ff-481b-938b-f070e444900f','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','eb75a27bbef54fda185c42f24fe112d19ecd6fc4aaf37fe9be7f0b89ac1793b5','2025-12-13 16:31:25',1,'2025-12-06 16:31:25'),('84537f8a-1bb0-4507-80ad-d36d5a6531aa','d713ddfb-1a39-4b33-9826-8be0814a3f57','73e6432621e3bbcf0a7657619d6ec2e5669d578ca60e571978c98b0aafafbe32','2025-12-12 14:35:09',1,'2025-12-05 14:35:09'),('86c2ff02-0b47-41a4-a459-7276c2a6aef7','d713ddfb-1a39-4b33-9826-8be0814a3f57','4dfcbec7eebf3dd8caa2a4eb3becca29a061f09a4a769d7ccfc03cba00d6f963','2025-12-12 12:38:54',1,'2025-12-05 12:38:54'),('884a2abd-4572-49c1-b239-c58244ef8655','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','552f8ed4d8c3f17ec1f3917969bafd21ebafcd893d8bbb84c900920195324b22','2025-12-12 12:11:16',1,'2025-12-05 12:11:16'),('8867397f-97ff-41ec-9781-2c5f5f94884a','2233d21b-c4fe-43dd-aba9-91f5d5df99a7','387ecfd0a661795b881091833b1e9bc04a8f8ef57c41239b44512517a5cee656','2025-12-13 18:13:27',1,'2025-12-06 18:13:27'),('8a64f047-3ece-4a2d-8638-0a428dddd92b','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','fadbb98cbee02b7d691e5dbfb8e321c37e7f51beb198a979a0eee9b869c1edd2','2025-12-13 15:01:45',1,'2025-12-06 15:01:45'),('8d7caaf8-e843-4ae3-abff-120c592fb673','526e30af-f07e-4272-a63e-753ad48ab6c7','4e40958d3d38c0fceb86dd976fea8d38164b4a818bea1afbbb1bfcd2902d6ef1','2025-12-12 09:29:19',1,'2025-12-05 09:29:19'),('90be8597-e41e-4a50-a732-f8e7db284fdf','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','e9f200f24fe515cd81c149ac66c1694e3500aa3b8e2b566c453cbb206c45ab63','2025-12-14 13:52:35',1,'2025-12-07 13:52:35'),('9253f7e8-2276-415c-af27-913c9e466ae2','d713ddfb-1a39-4b33-9826-8be0814a3f57','a3256124c140cebf12138901cd272bd27f1098a14d8ee8a187645e1ddba22269','2025-12-14 15:04:01',1,'2025-12-07 15:04:01'),('9318eab3-a785-4943-b996-6ced369d6594','d713ddfb-1a39-4b33-9826-8be0814a3f57','be5db3fd1b51899ca742de39c1310b060578cc39b84d7145dbc5c766b4d42917','2025-12-14 02:15:00',1,'2025-12-07 02:15:00'),('9c2a0560-c238-47fb-b259-b0a958840497','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','b69e80ca2e3acb21a642974ddba1d71ced8284b4c6595d3523c3a3b0a0852fcf','2025-12-14 14:36:53',1,'2025-12-07 14:36:53'),('9c9dc536-f9a9-4a92-b47c-24dd00d06442','d713ddfb-1a39-4b33-9826-8be0814a3f57','7144f053eba8ae34249f1ee497e302e46a5b3d72dd907733159d53f2899bde50','2025-12-14 13:08:10',1,'2025-12-07 13:08:10'),('9ffcef4b-91d1-4e1d-9fce-7c319d749576','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c70849a9152a50ea63955a42fa2df616e2abc579a433eecdb5f6ca2743f16a69','2025-12-13 15:31:34',1,'2025-12-06 15:31:34'),('a029a6d2-abd8-45fd-b465-eb86a8b06cc4','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','e0f09b0c2c54ad991e8338582dfe807b85e986ceb39628f7cf946b83c90330f9','2025-12-13 17:14:23',1,'2025-12-06 17:14:23'),('a04293ed-bb6b-4f7f-807e-b5f74c352355','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','671338956d62bf0c7ee675558f597d114552cde952e864f21c5b87f21da1c773','2025-12-11 20:56:29',1,'2025-12-04 20:56:29'),('a060f78b-1849-4ea2-8800-5223ba21951e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','cfcf7e913e1ab58f863206e05add4d89789f72be5c7cb6e1123293738ad85e5b','2025-12-13 16:31:17',1,'2025-12-06 16:31:17'),('a6ac80e6-c23a-49e6-add8-a1fe0b1f74e8','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','e0c761737a6bcb06be92326861172e2fb01768f6d8a90a3a94c6213ebc01d6b9','2025-12-13 19:43:05',1,'2025-12-06 19:43:05'),('a8792678-6578-4db0-951a-076846548da2','d713ddfb-1a39-4b33-9826-8be0814a3f57','8a0cc254e92ca75d069a2fa222fce39109d03305d8d4c4b7388e61019134ecf2','2025-12-13 07:50:45',1,'2025-12-06 07:50:45'),('afaea6af-bfdb-45a2-9e32-3b209f459ce1','d713ddfb-1a39-4b33-9826-8be0814a3f57','e0edb3c40a21e9d8572cc0f26f053c61105a8da4fa278744de15e9a239dfe162','2025-12-12 13:27:13',1,'2025-12-05 13:27:13'),('b3884aeb-d85c-4064-a9c1-748d2505d838','d713ddfb-1a39-4b33-9826-8be0814a3f57','25548f015cad48271402495da0f366b7cb2b878fe2de064bff86a04f8bd4fdde','2025-12-13 04:58:35',1,'2025-12-06 04:58:35'),('b5351956-7768-4755-bee1-a1d5ec5b7701','d713ddfb-1a39-4b33-9826-8be0814a3f57','a504d845da3ec81b88cfd62d227f3027281d05c3c8f9123eb0c2700c8f442431','2025-12-12 13:27:21',1,'2025-12-05 13:27:21'),('bfd3455c-5469-46cc-b85e-aefd3d812311','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','18f005fc75f0b3ccf78263dcbc041469f2422f96df0fb5a4d1a370cb3b8ad58f','2025-12-13 15:02:03',1,'2025-12-06 15:02:03'),('bffb9ebb-5a8d-4765-b766-e2e560b1574b','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','1f4e29b97de02e6d7d57b00b3a20c2224166bd7972c66b2b660cffcc9f7c9f25','2025-12-13 17:05:30',1,'2025-12-06 17:05:30'),('c01b8270-a959-421a-ad86-e372e70d0a55','d713ddfb-1a39-4b33-9826-8be0814a3f57','ca6c7c2708f794dd9eaf045f190964348727de1f20e37e52f6ae2fd894fea021','2025-12-14 14:16:54',1,'2025-12-07 14:16:54'),('c0ebb394-b6c1-41bd-bc8e-a4e6526b970c','526e30af-f07e-4272-a63e-753ad48ab6c7','20824114e8ee084331aba58fdda94fda547582a2ef090ef0a4fb98549b7e37f2','2025-12-11 21:42:17',1,'2025-12-04 21:42:17'),('c75f6d66-4fcc-4fd2-aaba-da366b65a60b','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','dc0d0926c7c2830d1b643695d5f34226adea8ac3f0e7ccde5c61822940f584a0','2025-12-13 19:01:56',1,'2025-12-06 19:01:56'),('c7700e79-b6c2-4b35-aa72-c1493fbaf7f1','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','6ca764c629604ea723ca40d26f5a06e02988113e1f2e5590eb88473c425af2f4','2025-12-11 21:06:37',1,'2025-12-04 21:06:37'),('c7cc93f3-3190-4ecb-8405-7537a0213cad','d713ddfb-1a39-4b33-9826-8be0814a3f57','7944d601851dfb0888079372b0af0331dbc5dcf2bebe0b8a8dff771a53ad2237','2025-12-12 14:56:32',1,'2025-12-05 14:56:32'),('c8113600-2690-4329-9ea7-38d00e429c80','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','93c243298785482bd7edd07d1e9f192721384e998d8d5984dd5b43fb69b16da5','2025-12-14 10:58:33',1,'2025-12-07 10:58:33'),('cb34aa09-ddc1-43eb-896d-500766dea64e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','b762dabd1296b21dfefc85d0041962f20b8c083696dee1c3a861f4891da15bb3','2025-12-12 10:11:16',1,'2025-12-05 10:11:16'),('cbb7d2fc-2b15-4492-9bdb-0b4fefce67bd','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','3d427287a472b2f0572e82de0974f4487e813db4729db97d74d7d8b7ebeea8bd','2025-12-13 15:40:37',1,'2025-12-06 15:40:37'),('cd43497a-be93-488d-8b90-6c5de8c8a715','d713ddfb-1a39-4b33-9826-8be0814a3f57','820cf8af7d3dacc30cc88436c4dd43a92a4be0103eeaa51a13586c6e49fb9995','2025-12-14 13:08:22',1,'2025-12-07 13:08:22'),('ce217584-20c8-414c-b946-27727789b7e4','d713ddfb-1a39-4b33-9826-8be0814a3f57','b3dfc9b9f6e1af2ec52ad3882e53ad20f4f70d2148140e7b27b16006cc39b8e9','2025-12-14 13:51:06',1,'2025-12-07 13:51:06'),('ce88cafd-e6ba-444f-a225-104ea34f277f','d713ddfb-1a39-4b33-9826-8be0814a3f57','a3c209d47fe2f38796eb140cbbfb249bd9e16779730b66d98cd4dce91b01d649','2025-12-14 13:38:51',1,'2025-12-07 13:38:51'),('ce9b0852-d255-478d-8d69-67b25617bf4c','d713ddfb-1a39-4b33-9826-8be0814a3f57','48c73c7aa44a95a88a3d6cde16a403dff8dfe9dc3e491812b3ff122d6e6c978e','2025-12-14 10:24:20',1,'2025-12-07 10:24:20'),('cec1bc57-812a-4fd8-a778-dcb34c823bcf','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','9e08e33d6ac6b5df21cf2f08dc8db81261bfc63f38862988a1495ddfb1bd4f37','2025-12-14 12:09:06',1,'2025-12-07 12:09:06'),('d0d990e5-9ae7-4c36-bc63-01087502b192','d713ddfb-1a39-4b33-9826-8be0814a3f57','c0b5e67215c5408eef271204cbb918d8af11bf56d2db192d2040f627cb04883c','2025-12-13 19:32:22',1,'2025-12-06 19:32:22'),('d3623569-6f63-483b-aab5-9135837af46a','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','ef9bfb7f629a6776eb92fb6967850c68d050c44f998b00fee3ad293f93de348a','2025-12-13 15:40:29',1,'2025-12-06 15:40:29'),('d5257981-c7c8-4406-96fe-18a0689affbb','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','5bb7d206bf9e9c7aaec617c3f01ba57700ac00a161c04e5b88357a2d79814f38','2025-12-13 16:05:34',1,'2025-12-06 16:05:34'),('d5440f64-b362-4e98-b327-c25871072eca','d713ddfb-1a39-4b33-9826-8be0814a3f57','2a4448d99c19a0059527b6ad5eaf7d6a950ad205fe61261cda3627d0e2d1f8b0','2025-12-14 14:18:07',1,'2025-12-07 14:18:07'),('dabee517-33cd-45f3-92cf-d214a0a6ce2b','d713ddfb-1a39-4b33-9826-8be0814a3f57','47bc784df4ee5dfa6bc8a88a928c526da58ac66e9c38dda2e19f069b9a7a31ae','2025-12-13 01:02:37',1,'2025-12-06 01:02:37'),('db8d2876-8008-4e59-9f16-9a5779dbf67f','d713ddfb-1a39-4b33-9826-8be0814a3f57','53b94fe60ee35030815884699246dac9d428256c38f65d4b934e00643f2c9173','2025-12-12 15:10:07',1,'2025-12-05 15:10:07'),('dc860027-877f-402c-a1d2-62ab919f8869','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','6343630370d05789138336fae15b8e941c37c127da1f2592226b584681be4680','2025-12-12 23:59:40',1,'2025-12-05 23:59:40'),('deab2206-bb09-4b31-b115-d513be1c9f1d','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','b6c7271eb32190a207f7bf334d7c73b26e635dfb9fa366f8789e00a3b1cabc41','2025-12-14 15:43:54',1,'2025-12-07 15:43:54'),('deee1577-c249-42f4-b85d-c0e5f8269bca','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c75d3ac2e1dcd3386bfe7e180ccd0bbf40e5eecf0d4c5c6492a3a2b214956d79','2025-12-14 02:05:32',1,'2025-12-07 02:05:32'),('df315c2f-596c-438c-bbea-f13ff522f61e','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','36acd783d7756d970958de9eb5de486e560a75cebe5b2ec701f6530df63b7f03','2025-12-11 21:06:37',1,'2025-12-04 21:06:37'),('e469c4d0-1ec5-4e67-90f6-cb3b19b72f40','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','2833a1c420448fafa7cf750881d165a69115c73e65a3ee80e20c814caf408aec','2025-12-13 19:39:06',1,'2025-12-06 19:39:06'),('e74697b5-606f-44fd-83ce-db70abc83051','d713ddfb-1a39-4b33-9826-8be0814a3f57','030dbd1ffad671f05368f672542ad7dcf9943e1fe585e75c9683ededa24f5f00','2025-12-12 13:01:26',1,'2025-12-05 13:01:26'),('ead49a0a-72f2-4916-8057-f6b6730788a9','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','6d1494d37953b85b2923d9084035706888b06d4a920e19bc23c2ee49622e1d5f','2025-12-11 21:08:40',1,'2025-12-04 21:08:40'),('f09d97ed-db7c-4f7b-be07-32f278d3bd55','d713ddfb-1a39-4b33-9826-8be0814a3f57','3b1a2856964fe50a2e815c8cd34398885b3e11c28d6208ac8e203af12fb1a49e','2025-12-13 00:00:43',1,'2025-12-06 00:00:43'),('f0ed3b90-ed8b-46e5-83b3-15bd8d792ed5','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','cfc980d7781e91fd17a00cbef3035ac114edf4422c3e7f90f99768879e2265fa','2025-12-13 04:59:34',1,'2025-12-06 04:59:34'),('f292837c-179e-47d1-b980-ab98069b6507','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','fdd7602929c7d3624937e3fdda348b790a760bc217c46e3f2b3f8f0ffdc04eb7','2025-12-13 06:13:25',1,'2025-12-06 06:13:25'),('f3d0a2e4-aa2f-41d0-994b-b96faa911a03','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','f28a27765e6d37cde106ac38d62aa640d6a6b4d7390852a157420e11d5280cbe','2025-12-13 19:32:52',1,'2025-12-06 19:32:52'),('f4e6b737-ac8d-456b-ba23-ecffa281f90e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','eda167edb5c5dbc2bb83b9bc490a506aaab128c1c6764a3a3c2d917503e63c37','2025-12-13 19:35:38',1,'2025-12-06 19:35:38'),('f67b7b20-1994-4bcb-890a-a535fc308b43','d713ddfb-1a39-4b33-9826-8be0814a3f57','23e0343656bd9c52a3a8203f164ff11dd4e27ecb547ee02c0928b389b1561f48','2025-12-13 19:40:58',1,'2025-12-06 19:40:58'),('f6993d50-7f06-435a-9d07-57b413ec9c2a','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','3b84512348fad6adf8864c8c9f445dd8eaa2b5a77266c4434d21ddf9cbbdc59f','2025-12-12 12:38:28',1,'2025-12-05 12:38:28'),('f82a97e6-2f48-447b-953f-93e07c8f807e','d713ddfb-1a39-4b33-9826-8be0814a3f57','087dadbf0a8543c58233742482480ac8dc0296fbfaaffacd102e7d525d4dd999','2025-12-14 14:59:03',1,'2025-12-07 14:59:03'),('fc368271-fdff-4666-aad9-4a6492707a33','2b704681-a615-43cb-8793-9e9e2d754ecf','2cf1533b702c5a05f6daefeb08b194e813e7f3972f8e4e67b3abc051ac6a6226','2025-12-11 20:54:29',1,'2025-12-04 20:54:29'),('fc600bb8-96d5-4c26-b092-7421d18d812e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c475929a887dc65e37889f3f971a46e2b9a51efe4e9ad874d5468d58b71e19c0','2025-12-14 15:15:09',1,'2025-12-07 15:15:09'),('fca971da-6dab-4365-aa1c-ac7b71dfbc05','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c94ad7183fdcdd2827f4c4144c6c012ea26d5252547aa0e79cff25ae2d8124dd','2025-12-14 15:57:32',1,'2025-12-07 15:57:32'),('fd1395a9-1387-43e6-9cfa-53105f82ce81','526e30af-f07e-4272-a63e-753ad48ab6c7','45a4d8df4eedf48cbb63cdc5e117f76767005adf4ec36a964524312a04023efb','2025-12-11 21:09:16',1,'2025-12-04 21:09:16'),('fd33ff85-3287-4b6b-8c1a-61602bad9b38','d713ddfb-1a39-4b33-9826-8be0814a3f57','d443974df4e102b750397ee7bfc614729919d4f53c9b5cf6c787f0bacbeb2156','2025-12-14 15:09:54',1,'2025-12-07 15:09:54'),('feeaac9d-3b9c-4f7a-98ae-30675b553c6c','d713ddfb-1a39-4b33-9826-8be0814a3f57','85e3441f7746f72af13fca4fa84877639ff8897e9fbb0f4a0bc0f06194ae4429','2025-12-14 02:03:23',1,'2025-12-07 02:03:23'),('ff2e66e2-4a27-4be0-a37e-2aa100301575','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','01e19ed6c9148f4ec6d4d7f1209ffbc2b71620502e09798f195fb067bbb2daf8','2025-12-12 12:25:24',0,'2025-12-05 12:25:24');
/*!40000 ALTER TABLE `refresh_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps`
--

DROP TABLE IF EXISTS `rps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mata_kuliah_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `dosen_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `dosen_nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `tahun_ajaran` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `semester_type` enum('ganjil','genap') COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('draft','submitted','approved','rejected','revision') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'draft',
  `submitted_at` timestamp NULL DEFAULT NULL,
  `deskripsi_mk` text COLLATE utf8mb4_unicode_ci,
  `capaian_pembelajaran` text COLLATE utf8mb4_unicode_ci,
  `metode_pembelajaran` json DEFAULT NULL,
  `media_pembelajaran` json DEFAULT NULL,
  `reviewer_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `review_catatan` text COLLATE utf8mb4_unicode_ci,
  `reviewed_at` timestamp NULL DEFAULT NULL,
  `approved_at` timestamp NULL DEFAULT NULL,
  `version` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `semester_tipe` enum('ganjil','genap') COLLATE utf8mb4_unicode_ci DEFAULT 'ganjil',
  `tanggal_penyusunan` date DEFAULT NULL,
  `penyusun_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `penyusun_nama` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `penyusun_nidn` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `koordinator_rmk_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `koordinator_rmk_nama` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `koordinator_rmk_nidn` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kaprodi_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kaprodi_nama` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kaprodi_nidn` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fakultas` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `program_studi` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_rps_mata_kuliah` (`mata_kuliah_id`),
  KEY `idx_rps_dosen` (`dosen_id`),
  KEY `idx_rps_status` (`status`),
  KEY `idx_rps_tahun_ajaran` (`tahun_ajaran`),
  KEY `idx_rps_deleted_at` (`deleted_at`),
  KEY `fk_rps_reviewer` (`reviewer_id`),
  CONSTRAINT `fk_rps_dosen` FOREIGN KEY (`dosen_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_rps_mata_kuliah` FOREIGN KEY (`mata_kuliah_id`) REFERENCES `mata_kuliah` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_rps_reviewer` FOREIGN KEY (`reviewer_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps`
--

LOCK TABLES `rps` WRITE;
/*!40000 ALTER TABLE `rps` DISABLE KEYS */;
INSERT INTO `rps` VALUES ('99abcad5-d3f5-42a5-b201-902a606278c4','574f8a65-fe10-4890-abc7-5f1876533926','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','','2024/2025','genap','draft','2025-12-07 15:15:33','aksan','aksan','[\"Ceramah\", \"Diskusi\", \"Praktikum\"]','[\"Laptop\", \"Projector\", \"LMS\"]','d713ddfb-1a39-4b33-9826-8be0814a3f57','revisi','2025-12-07 15:14:34','2025-12-07 13:49:16',1,'2025-12-07 10:55:46','2025-12-07 15:53:03',NULL,'ganjil','2025-12-07',NULL,'aksan','1101001011',NULL,'aksan','1101001011',NULL,'aksan','1101001011','aksan','1101001011');
/*!40000 ALTER TABLE `rps` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_analisis_ketercapaian_cpl`
--

DROP TABLE IF EXISTS `rps_analisis_ketercapaian_cpl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_analisis_ketercapaian_cpl` (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `minggu_mulai` int NOT NULL,
  `minggu_selesai` int DEFAULT NULL,
  `cpl_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpmk_ids` json DEFAULT NULL COMMENT 'Array of CPMK IDs terkait',
  `sub_cpmk_ids` json DEFAULT NULL COMMENT 'Array of Sub-CPMK IDs terkait',
  `topik_materi` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `jenis_assessment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'UTS, UAS, Kuis, Tugas, Praktikum, dll',
  `bobot_kontribusi` int NOT NULL DEFAULT '0' COMMENT 'Kontribusi ke CPL (%)',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_analisis_cpl_rps_id` (`rps_id`),
  KEY `idx_analisis_cpl_cpl_id` (`cpl_id`),
  CONSTRAINT `rps_analisis_ketercapaian_cpl_ibfk_1` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE,
  CONSTRAINT `rps_analisis_ketercapaian_cpl_ibfk_2` FOREIGN KEY (`cpl_id`) REFERENCES `cpl` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_analisis_ketercapaian_cpl`
--

LOCK TABLES `rps_analisis_ketercapaian_cpl` WRITE;
/*!40000 ALTER TABLE `rps_analisis_ketercapaian_cpl` DISABLE KEYS */;
INSERT INTO `rps_analisis_ketercapaian_cpl` VALUES ('06ea691b-5565-4fa5-b4b2-a66bbbaaa532','99abcad5-d3f5-42a5-b201-902a606278c4',1,8,'0916a26d-7141-4d1c-802c-e3aab221bc8a','[\"CPMK-01\"]','null','aksan','aksan',10,'2025-12-07 11:48:09','2025-12-07 11:48:09'),('ccda90dd-a96e-4f8f-8659-aa80aa5d5535','99abcad5-d3f5-42a5-b201-902a606278c4',1,8,'0916a26d-7141-4d1c-802c-e3aab221bc8a','[\"CPMK-01\"]','null','aksan','aksan',10,'2025-12-07 11:48:24','2025-12-07 11:48:24');
/*!40000 ALTER TABLE `rps_analisis_ketercapaian_cpl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_bahan_bacaan`
--

DROP TABLE IF EXISTS `rps_bahan_bacaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_bahan_bacaan` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jenis` enum('buku','jurnal','artikel','website','modul','utama','pendukung') COLLATE utf8mb4_unicode_ci DEFAULT 'buku',
  `judul` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `penulis` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `penerbit` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tahun` int DEFAULT NULL,
  `isbn` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `halaman` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `urutan` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rps_bahan_rps` (`rps_id`),
  CONSTRAINT `fk_rps_bahan_rps` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_bahan_bacaan`
--

LOCK TABLES `rps_bahan_bacaan` WRITE;
/*!40000 ALTER TABLE `rps_bahan_bacaan` DISABLE KEYS */;
INSERT INTO `rps_bahan_bacaan` VALUES ('1f72c2db-8f36-4760-8adc-a214ed91e35e','99abcad5-d3f5-42a5-b201-902a606278c4','buku','AKSAN','AKSAN',NULL,2025,'AKSAN','1-50','AKSAN',1,'2025-12-07 11:55:43','2025-12-07 11:55:43'),('8de10ab3-d6d1-48fb-b037-44b79e0f6190','99abcad5-d3f5-42a5-b201-902a606278c4','buku','AKSAN','AKSAN',NULL,2025,'AKSAN','1-50','AKSAN',1,'2025-12-07 11:56:11','2025-12-07 11:56:10'),('cb3b9835-d9cb-4d56-8cc5-46080735f4c8','99abcad5-d3f5-42a5-b201-902a606278c4','buku','AKSAN','AKSAN',NULL,2025,'AKSAN','1-50','AKSAN',2,'2025-12-07 11:56:00','2025-12-07 11:56:00'),('f28aad3c-eb75-41b4-a1a8-09c2dad4849e','99abcad5-d3f5-42a5-b201-902a606278c4','buku','AKSAN','AKSAN',NULL,2025,'AKSAN','1-50','AKSAN',1,'2025-12-07 11:56:00','2025-12-07 11:56:00');
/*!40000 ALTER TABLE `rps_bahan_bacaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_cpl_mappings`
--

DROP TABLE IF EXISTS `rps_cpl_mappings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_cpl_mappings` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpl_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpmk_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tingkat_kemampuan` enum('dasar','menengah','lanjut') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'menengah',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rps_cpl_mappings_rps` (`rps_id`),
  KEY `idx_rps_cpl_mappings_cpl` (`cpl_id`),
  KEY `fk_rps_cpl_mappings_cpmk` (`cpmk_id`),
  CONSTRAINT `fk_rps_cpl_mappings_cpl` FOREIGN KEY (`cpl_id`) REFERENCES `cpl` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_rps_cpl_mappings_cpmk` FOREIGN KEY (`cpmk_id`) REFERENCES `rps_cpmk` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_rps_cpl_mappings_rps` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_cpl_mappings`
--

LOCK TABLES `rps_cpl_mappings` WRITE;
/*!40000 ALTER TABLE `rps_cpl_mappings` DISABLE KEYS */;
/*!40000 ALTER TABLE `rps_cpl_mappings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_cpmk`
--

DROP TABLE IF EXISTS `rps_cpmk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_cpmk` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deskripsi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `bobot` decimal(5,2) DEFAULT NULL,
  `urutan` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rps_cpmk_rps` (`rps_id`),
  CONSTRAINT `fk_rps_cpmk_rps` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_cpmk`
--

LOCK TABLES `rps_cpmk` WRITE;
/*!40000 ALTER TABLE `rps_cpmk` DISABLE KEYS */;
INSERT INTO `rps_cpmk` VALUES ('1a7147dc-3067-4ea6-81ae-2322a3044b5e','99abcad5-d3f5-42a5-b201-902a606278c4','CPMK-01','CPMK1',NULL,1,'2025-12-07 11:03:49','2025-12-07 11:03:49');
/*!40000 ALTER TABLE `rps_cpmk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_evaluasi`
--

DROP TABLE IF EXISTS `rps_evaluasi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_evaluasi` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `komponen` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `teknik_penilaian` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `instrumen` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bobot` decimal(5,2) NOT NULL,
  `minggu_mulai` int DEFAULT NULL,
  `minggu_selesai` int DEFAULT NULL,
  `cpl_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kriteria_penilaian` text COLLATE utf8mb4_unicode_ci,
  `urutan` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `cpmk_ids` json DEFAULT NULL,
  `sub_cpmk_ids` json DEFAULT NULL,
  `topik_materi` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `jenis_assessment` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_rps_evaluasi_rps` (`rps_id`),
  CONSTRAINT `fk_rps_evaluasi_rps` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_evaluasi`
--

LOCK TABLES `rps_evaluasi` WRITE;
/*!40000 ALTER TABLE `rps_evaluasi` DISABLE KEYS */;
INSERT INTO `rps_evaluasi` VALUES ('33f0119d-9390-4e79-9756-334b22f99cf3','99abcad5-d3f5-42a5-b201-902a606278c4','Tugas','Penugasan','Rubrik Penilaian',20.00,1,16,NULL,NULL,3,'2025-12-07 11:56:38','2025-12-07 11:56:38','null','null',NULL,'Formatif'),('69e48026-4209-4416-ac8f-a309b245e640','99abcad5-d3f5-42a5-b201-902a606278c4','UTS','Tes Tertulis','Soal Essay',30.00,1,8,NULL,NULL,1,'2025-12-07 11:56:38','2025-12-07 11:56:38','null','null',NULL,'Sumatif'),('b24c4682-a596-469f-8fd0-0fac4a6f97ad','99abcad5-d3f5-42a5-b201-902a606278c4','Kehadiran','Observasi','Daftar Hadir',10.00,1,16,NULL,NULL,4,'2025-12-07 11:56:38','2025-12-07 11:56:38','null','null',NULL,'Formatif'),('e5f030cc-2135-41fe-a488-e602069286ca','99abcad5-d3f5-42a5-b201-902a606278c4','UAS','Tes Tertulis','Soal Essay',40.00,9,16,NULL,NULL,2,'2025-12-07 11:56:38','2025-12-07 11:56:38','null','null',NULL,'Sumatif');
/*!40000 ALTER TABLE `rps_evaluasi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_rencana_pembelajaran`
--

DROP TABLE IF EXISTS `rps_rencana_pembelajaran`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_rencana_pembelajaran` (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `pertemuan` int NOT NULL,
  `minggu_mulai` int NOT NULL DEFAULT '1',
  `minggu_selesai` int DEFAULT NULL,
  `topik` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `sub_topik` json DEFAULT NULL,
  `sub_cpmk_ids` json DEFAULT NULL,
  `cpmk_ids` json DEFAULT NULL,
  `indikator` json DEFAULT NULL,
  `metode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `media_lms` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `waktu` int DEFAULT NULL,
  `waktu_tm` int DEFAULT NULL,
  `waktu_bm` int DEFAULT NULL,
  `waktu_pt` int DEFAULT NULL,
  `materi` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `teknik_penilaian` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `kriteria_penilaian` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `bobot_penilaian` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rencana_rps_id` (`rps_id`),
  CONSTRAINT `rps_rencana_pembelajaran_ibfk_1` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_rencana_pembelajaran`
--

LOCK TABLES `rps_rencana_pembelajaran` WRITE;
/*!40000 ALTER TABLE `rps_rencana_pembelajaran` DISABLE KEYS */;
INSERT INTO `rps_rencana_pembelajaran` VALUES ('07c4e645-72c9-4960-84ef-244a5b07a3df','99abcad5-d3f5-42a5-b201-902a606278c4',2,1,NULL,'Teknik Penilaian Observasi dan Praktik, dll Kriteria Penilaian Kriteria penilaian... Bobot Penilaian (%)','null',NULL,'null',NULL,'Ceramah dan Praktik',NULL,150,NULL,NULL,NULL,'Teknik Penilaian\nObservasi dan Praktik, dll\nKriteria Penilaian\nKriteria penilaian...\nBobot Penilaian (%)',NULL,NULL,NULL,'2025-12-07 11:50:06'),('5516c97b-bb02-4701-ba0e-8954d6a9c1f4','99abcad5-d3f5-42a5-b201-902a606278c4',2,1,NULL,'Teknik Penilaian Observasi dan Praktik, dll Kriteria Penilaian Kriteria penilaian... Bobot Penilaian (%)','null',NULL,'null',NULL,'Ceramah dan Praktik',NULL,150,NULL,NULL,NULL,'Teknik Penilaian\nObservasi dan Praktik, dll\nKriteria Penilaian\nKriteria penilaian...\nBobot Penilaian (%)',NULL,NULL,NULL,'2025-12-07 11:50:33'),('7259fcaa-0528-4422-8c76-62451a7e4b00','99abcad5-d3f5-42a5-b201-902a606278c4',1,1,NULL,'aksan','null',NULL,'null',NULL,'Ceramah dan Praktik',NULL,150,NULL,NULL,NULL,'sasa',NULL,NULL,NULL,'2025-12-07 11:50:33'),('935a3989-5954-49f9-9d56-e01b0a00c675','99abcad5-d3f5-42a5-b201-902a606278c4',1,1,NULL,'aksan','null',NULL,'null',NULL,'Ceramah dan Praktik',NULL,150,NULL,NULL,NULL,'sasa',NULL,NULL,NULL,'2025-12-07 11:35:47'),('dfed22a6-afac-4a51-a0dc-654f3def3992','99abcad5-d3f5-42a5-b201-902a606278c4',1,1,NULL,'aksan','null',NULL,'null',NULL,'Ceramah dan Praktik',NULL,150,NULL,NULL,NULL,'sasa',NULL,NULL,NULL,'2025-12-07 11:50:05');
/*!40000 ALTER TABLE `rps_rencana_pembelajaran` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_rencana_tugas`
--

DROP TABLE IF EXISTS `rps_rencana_tugas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_rencana_tugas` (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nomor_tugas` int NOT NULL,
  `judul` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_keberhasilan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `batas_waktu_minggu` int DEFAULT NULL,
  `petunjuk_pengerjaan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `jenis_tugas` enum('individu','kelompok') COLLATE utf8mb4_unicode_ci DEFAULT 'individu',
  `luaran_tugas` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Misal: Kode program + laporan PDF',
  `kriteria_penilaian` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Rubrik inti penilaian',
  `teknik_penilaian` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Teknik penilaian yang digunakan',
  `bobot` int NOT NULL DEFAULT '0' COMMENT 'Bobot terhadap nilai akhir (%)',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rencana_tugas_rps_id` (`rps_id`),
  CONSTRAINT `rps_rencana_tugas_ibfk_1` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_rencana_tugas`
--

LOCK TABLES `rps_rencana_tugas` WRITE;
/*!40000 ALTER TABLE `rps_rencana_tugas` DISABLE KEYS */;
INSERT INTO `rps_rencana_tugas` VALUES ('3b405ea6-cf70-452b-9d15-d50ac3dea326','99abcad5-d3f5-42a5-b201-902a606278c4',1,'aksan','aksan',1,'aksan','individu','aksan','aksan','Checklist',10,'2025-12-07 11:46:52','2025-12-07 11:46:52'),('f36eb0c8-d926-47e5-a0d1-3114ee56692f','99abcad5-d3f5-42a5-b201-902a606278c4',2,'aksan','aksan',2,'aksan','individu','aksan','aksan','Presentasi',10,'2025-12-07 11:46:52','2025-12-07 11:46:52'),('fb92e6c9-3ed5-4098-92d9-d57a3c4504ae','99abcad5-d3f5-42a5-b201-902a606278c4',1,'aksan','aksan',1,'aksan','individu','aksan','aksan','',10,'2025-12-07 11:46:32','2025-12-07 11:46:32');
/*!40000 ALTER TABLE `rps_rencana_tugas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_skala_penilaian`
--

DROP TABLE IF EXISTS `rps_skala_penilaian`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_skala_penilaian` (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nilai_min` int NOT NULL COMMENT 'Batas bawah nilai',
  `nilai_max` int NOT NULL COMMENT 'Batas atas nilai',
  `huruf_mutu` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'A, B+, B, C+, C, D, E',
  `bobot_nilai` decimal(4,2) NOT NULL COMMENT '4.0, 3.5, 3.0, dst',
  `is_lulus` tinyint(1) DEFAULT '1' COMMENT 'Apakah dianggap lulus',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_skala_penilaian_rps_id` (`rps_id`),
  CONSTRAINT `rps_skala_penilaian_ibfk_1` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_skala_penilaian`
--

LOCK TABLES `rps_skala_penilaian` WRITE;
/*!40000 ALTER TABLE `rps_skala_penilaian` DISABLE KEYS */;
INSERT INTO `rps_skala_penilaian` VALUES ('1fc9cb34-e624-485f-91ad-2c6d57904ab8','99abcad5-d3f5-42a5-b201-902a606278c4',10,59,'A',4.00,1,'2025-12-07 16:03:10'),('509533db-0938-4ca8-9d29-133530dcd48f','99abcad5-d3f5-42a5-b201-902a606278c4',60,79,'E',1.70,0,'2025-12-07 16:18:27');
/*!40000 ALTER TABLE `rps_skala_penilaian` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sub_cpmk`
--

DROP TABLE IF EXISTS `sub_cpmk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sub_cpmk` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cpmk_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deskripsi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `urutan` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_sub_cpmk_kode_cpmk` (`cpmk_id`,`kode`),
  KEY `idx_sub_cpmk_cpmk_id` (`cpmk_id`),
  CONSTRAINT `sub_cpmk_ibfk_1` FOREIGN KEY (`cpmk_id`) REFERENCES `rps_cpmk` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sub_cpmk`
--

LOCK TABLES `sub_cpmk` WRITE;
/*!40000 ALTER TABLE `sub_cpmk` DISABLE KEYS */;
INSERT INTO `sub_cpmk` VALUES ('2335836c-e737-40ce-b216-008909248787','1a7147dc-3067-4ea6-81ae-2322a3044b5e','Sub-CPMK-03','SUBCPMK3',3,'2025-12-07 11:34:38','2025-12-07 11:34:38'),('ecd1ffcb-7508-430a-994a-a5ff51044231','1a7147dc-3067-4ea6-81ae-2322a3044b5e','Sub-CPMK-02','SUBCPMK2',2,'2025-12-07 11:23:32','2025-12-07 11:34:38'),('f8e314b8-6bbf-4440-839f-d7b47c272d2d','1a7147dc-3067-4ea6-81ae-2322a3044b5e','Sub-CPMK-01','SUBCPMK1',1,'2025-12-07 11:23:22','2025-12-07 11:34:38');
/*!40000 ALTER TABLE `sub_cpmk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_settings`
--

DROP TABLE IF EXISTS `system_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_settings` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `key_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` text COLLATE utf8mb4_unicode_ci,
  `description` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_system_settings_key` (`key_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_settings`
--

LOCK TABLES `system_settings` WRITE;
/*!40000 ALTER TABLE `system_settings` DISABLE KEYS */;
INSERT INTO `system_settings` VALUES ('be4fbb58-d139-11f0-888e-ec750c0888b7','app_name','Kurikulum Management System','Application name','2025-12-04 17:50:38','2025-12-04 17:50:38'),('be4fbfa5-d139-11f0-888e-ec750c0888b7','tahun_ajaran_aktif','2024/2025','Active academic year','2025-12-04 17:50:38','2025-12-04 17:50:38'),('be4fc0d2-d139-11f0-888e-ec750c0888b7','semester_aktif','ganjil','Active semester','2025-12-04 17:50:38','2025-12-04 17:50:38'),('be4fc1a9-d139-11f0-888e-ec750c0888b7','max_sks_per_semester','24','Maximum SKS per semester','2025-12-04 17:50:38','2025-12-04 17:50:38'),('be4fc270-d139-11f0-888e-ec750c0888b7','deadline_rps_submission','14','Days before class start for RPS submission','2025-12-04 17:50:38','2025-12-04 17:50:38');
/*!40000 ALTER TABLE `system_settings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password_hash` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nip` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` enum('kaprodi','dosen') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'dosen',
  `status` enum('active','inactive') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'active',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar_url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `last_login` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_users_email` (`email`),
  UNIQUE KEY `uk_users_nip` (`nip`),
  KEY `idx_users_role` (`role`),
  KEY `idx_users_status` (`status`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('2233d21b-c4fe-43dd-aba9-91f5d5df99a7','test@test.com','$2a$10$3ojSXiSrXcD/Q7VE.cuSA.S7jfAtmBsJgqFpDKzxPvJ91BY1Avefi','Test User','123456','dosen','active',NULL,NULL,'2025-12-06 18:59:07','2025-12-04 20:39:37','2025-12-06 18:59:08',NULL),('2b704681-a615-43cb-8793-9e9e2d754ecf','testnew@test.com','$2a$10$W/McMduO3jwpHtM/GZ0HweXRi6yhFyhC1MZdwuwNXs5UMar196rpi','Test User New',NULL,'dosen','active',NULL,NULL,'2025-12-04 20:55:45','2025-12-04 20:50:46','2025-12-04 20:55:46',NULL),('526e30af-f07e-4272-a63e-753ad48ab6c7','ani.wijaya@university.ac.id','$2a$10$jpF6EGZIzZG2KofOPt1QceNVxu2T3JJzqd/bUYL0/YW4qkpsUoVQa','Dr. Ani Wijaya, M.T','197812102005011003','dosen','active','081234567899',NULL,'2025-12-05 09:29:32','2025-12-04 20:26:37','2025-12-05 09:29:32',NULL),('6dfd7a45-dca1-46ba-a87a-9dfb8f870563','kaprodi@test.com','$2a$10$5fV0Q49809ahb2XcVQ21oeYKhvSHqmzhU7tYaZi2nS7jnReIOUeVW','Kaprodi Updated',NULL,'kaprodi','active',NULL,NULL,'2025-12-05 12:25:23','2025-12-04 20:56:29','2025-12-05 12:25:24',NULL),('a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','aksan@unismuh.ac.id','$2a$10$ejayq9j8X5qIPUhyVNOPFudkWY77ypGKYR90/sEPE7W2Z5RudLLy6','Dr. Muhammad Aksan S. T, M.T','197812102005011001','dosen','active','081234567899',NULL,'2025-12-07 15:57:31','2025-12-05 10:10:49','2025-12-07 15:57:32',NULL),('be4e85bd-d139-11f0-888e-ec750c0888b7','kaprodi@university.ac.id','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZRGdjGj/n3.QNqQ.0y.naPxhPj3Ey','Dr. John Doe, M.Kom','198501012010011001','kaprodi','active',NULL,NULL,NULL,'2025-12-04 17:50:38','2025-12-04 17:50:38',NULL),('d713ddfb-1a39-4b33-9826-8be0814a3f57','aksan1@unismuh.ac.id','$2a$10$oCEbTH57ZKYcUwiuLnsO7OElK3/f.8sId96/go66QovmVySUhwcA2','Dr. Muhammad Aksan S. T, M.T','197812102005011002','kaprodi','active','081234567899',NULL,'2025-12-07 15:09:54','2025-12-05 10:20:23','2025-12-07 15:09:54',NULL),('e27f2069-daa4-4f1f-9494-6218fdeb2f5e','test2@test.com','$2a$10$E88uRsx10bZewNCRHVGNpeR8tTmUKzXyWMCG.BHHaC71BGOWJZFeq','Test User2','1234567','dosen','active',NULL,NULL,NULL,'2025-12-04 20:40:00','2025-12-04 20:40:00',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `v_cpl_assignment_summary`
--

DROP TABLE IF EXISTS `v_cpl_assignment_summary`;
/*!50001 DROP VIEW IF EXISTS `v_cpl_assignment_summary`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `v_cpl_assignment_summary` AS SELECT 
 1 AS `id`,
 1 AS `status`,
 1 AS `cpl_kode`,
 1 AS `cpl_nama`,
 1 AS `dosen_nama`,
 1 AS `dosen_email`,
 1 AS `assigned_by_nama`,
 1 AS `assigned_at`,
 1 AS `response_at`,
 1 AS `completed_at`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `v_rps_summary`
--

DROP TABLE IF EXISTS `v_rps_summary`;
/*!50001 DROP VIEW IF EXISTS `v_rps_summary`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `v_rps_summary` AS SELECT 
 1 AS `id`,
 1 AS `status`,
 1 AS `tahun_ajaran`,
 1 AS `semester_type`,
 1 AS `mata_kuliah_kode`,
 1 AS `mata_kuliah_nama`,
 1 AS `sks`,
 1 AS `dosen_nama`,
 1 AS `dosen_email`,
 1 AS `reviewer_nama`,
 1 AS `created_at`,
 1 AS `updated_at`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `v_cpl_assignment_summary`
--

/*!50001 DROP VIEW IF EXISTS `v_cpl_assignment_summary`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_cpl_assignment_summary` AS select `ca`.`id` AS `id`,`ca`.`status` AS `status`,`c`.`kode` AS `cpl_kode`,`c`.`nama` AS `cpl_nama`,`d`.`nama` AS `dosen_nama`,`d`.`email` AS `dosen_email`,`ab`.`nama` AS `assigned_by_nama`,`ca`.`assigned_at` AS `assigned_at`,`ca`.`response_at` AS `response_at`,`ca`.`completed_at` AS `completed_at` from (((`cpl_assignments` `ca` join `cpl` `c` on((`ca`.`cpl_id` = `c`.`id`))) join `users` `d` on((`ca`.`dosen_id` = `d`.`id`))) join `users` `ab` on((`ca`.`assigned_by` = `ab`.`id`))) where (`ca`.`deleted_at` is null) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `v_rps_summary`
--

/*!50001 DROP VIEW IF EXISTS `v_rps_summary`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_rps_summary` AS select `r`.`id` AS `id`,`r`.`status` AS `status`,`r`.`tahun_ajaran` AS `tahun_ajaran`,`r`.`semester_type` AS `semester_type`,`mk`.`kode` AS `mata_kuliah_kode`,`mk`.`nama` AS `mata_kuliah_nama`,`mk`.`sks` AS `sks`,`u`.`nama` AS `dosen_nama`,`u`.`email` AS `dosen_email`,`rv`.`nama` AS `reviewer_nama`,`r`.`created_at` AS `created_at`,`r`.`updated_at` AS `updated_at` from (((`rps` `r` join `mata_kuliah` `mk` on((`r`.`mata_kuliah_id` = `mk`.`id`))) join `users` `u` on((`r`.`dosen_id` = `u`.`id`))) left join `users` `rv` on((`r`.`reviewer_id` = `rv`.`id`))) where (`r`.`deleted_at` is null) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-12-08  1:48:31
