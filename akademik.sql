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
INSERT INTO `cpl` VALUES ('0916a26d-7141-4d1c-802c-e3aab221bc8a','CPL003','VPPPPPP','TESTESTESTES','published',2,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:28:56','2025-12-05 13:51:08',NULL),('2a78ae74-c36f-4499-9a98-3dbc03211d72','CPL-TES01','TES01','TESTESTESTES','published',2,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 11:09:06','2025-12-05 11:44:45',NULL),('eab004ce-0752-4447-9700-c71b25c4e050','CPL-TES02','TES-02','TESTESTESTESTES','published',7,'d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 11:22:12','2025-12-05 12:00:56',NULL);
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
INSERT INTO `cpl_assignments` VALUES ('8e8f0220-c9ac-486a-bdeb-6a4913f988ca','eab004ce-0752-4447-9700-c71b25c4e050','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','Pemrograman Web','33f3ed4c-1751-41c1-a70d-1b311e6bd20a','2025-12-05 00:00:00','cancelled','sasa','Penugasan dibatalkan oleh Kaprodi','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:04:03','2025-12-05 13:04:06',NULL,'2025-12-05 13:04:03','2025-12-05 13:04:06',NULL);
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
INSERT INTO `cpl_mk_mappings` VALUES ('263c12bc-9904-4178-a258-a7bbebb32696','2a78ae74-c36f-4499-9a98-3dbc03211d72','7f01fd62-3fdc-455c-8987-9fad9d4debc4','tinggi','2025-12-05 13:54:30','2025-12-05 13:54:30'),('294354f1-5f97-4e4c-99e8-5f67a48435a2','eab004ce-0752-4447-9700-c71b25c4e050','7f01fd62-3fdc-455c-8987-9fad9d4debc4','tinggi','2025-12-05 13:54:32','2025-12-05 13:54:32'),('2c14824f-e753-4000-abb3-6e219ac2aa69','0916a26d-7141-4d1c-802c-e3aab221bc8a','7f01fd62-3fdc-455c-8987-9fad9d4debc4','tinggi','2025-12-05 13:54:34','2025-12-05 13:54:34'),('35406972-7cf5-47c3-9390-d04b691396f8','0916a26d-7141-4d1c-802c-e3aab221bc8a','33f3ed4c-1751-41c1-a70d-1b311e6bd20a','tinggi','2025-12-05 13:54:25','2025-12-05 13:54:25'),('7cf376cf-2712-4a5b-8abd-9f55bae23551','eab004ce-0752-4447-9700-c71b25c4e050','d095e618-6a84-4949-bbf0-fdf687ed98f8','tinggi','2025-12-05 13:54:32','2025-12-05 13:54:32'),('aa3a8ad8-cf1f-49d9-8ee6-30b056f30278','2a78ae74-c36f-4499-9a98-3dbc03211d72','574f8a65-fe10-4890-abc7-5f1876533926','rendah','2025-12-05 13:54:27','2025-12-05 14:39:50'),('e864f9fb-4160-45e4-b9d3-39d64a69a983','2a78ae74-c36f-4499-9a98-3dbc03211d72','33f3ed4c-1751-41c1-a70d-1b311e6bd20a','rendah','2025-12-05 13:43:54','2025-12-05 14:39:49'),('fe4237e7-657b-46d7-a304-2b95d6f51501','0916a26d-7141-4d1c-802c-e3aab221bc8a','d095e618-6a84-4949-bbf0-fdf687ed98f8','tinggi','2025-12-05 13:54:34','2025-12-05 13:54:34');
/*!40000 ALTER TABLE `cpl_mk_mappings` ENABLE KEYS */;
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
INSERT INTO `mata_kuliah` VALUES ('33f3ed4c-1751-41c1-a70d-1b311e6bd20a','MK-001','Pemrograman Web',3,3,'wajib',NULL,'null',NULL,NULL,1,'aktif','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','2025-12-04 21:06:15','2025-12-04 21:06:15',NULL),('574f8a65-fe10-4890-abc7-5f1876533926','SASASASASA','aksan',1,1,'wajib','sasa','null',NULL,NULL,0,'aktif','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:38:42','2025-12-05 14:34:32',NULL),('7f01fd62-3fdc-455c-8987-9fad9d4debc4','MK011111','Rekayasa Perangkat Lunak',3,8,'wajib','MATAKULIAH WAJIB','null',NULL,NULL,0,'dihapus','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 13:05:14','2025-12-05 14:39:15',NULL),('b8ecc5e3-f924-419f-8fba-9084da126114','TIF501','Machine Learning',3,5,'pilihan','Pengantar machine learning','null',NULL,NULL,0,'dihapus','d713ddfb-1a39-4b33-9826-8be0814a3f57','2025-12-05 14:35:48','2025-12-05 14:39:42',NULL),('d095e618-6a84-4949-bbf0-fdf687ed98f8','MK-TEST-001','Pemrograman Web Test',3,3,'wajib',NULL,'null',NULL,NULL,1,'aktif','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','2025-12-04 21:08:39','2025-12-04 21:08:39',NULL);
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
INSERT INTO `notifications` VALUES ('023faf5d-608e-430d-b1c3-65036362b5f7','526e30af-f07e-4272-a63e-753ad48ab6c7','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web','assignment','cpl_assignment','52cc09b6-57e6-4d0a-acc6-49d301f33bbe',0,NULL,'2025-12-05 12:41:44',NULL),('14cd1d8f-ab58-4ec7-b752-bc086c82cbe9','526e30af-f07e-4272-a63e-753ad48ab6c7','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web','assignment','cpl_assignment','5a268d53-8c51-4c50-a8ee-57d09ed1e846',0,NULL,'2025-12-05 12:43:40',NULL),('23ba9039-92aa-41d8-9de3-e247269ab9ac','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web','assignment','cpl_assignment','8e8f0220-c9ac-486a-bdeb-6a4913f988ca',0,NULL,'2025-12-05 13:04:03',NULL),('4a28e9d0-15f9-450c-8440-de40b91579ef','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','Penugasan CPL Baru','Anda mendapat penugasan CPL baru untuk mata kuliah Pemrograman Web Test','assignment','cpl_assignment','5a405f46-157b-4ee4-b587-218ad1b83351',0,NULL,'2025-12-05 12:52:17',NULL);
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
INSERT INTO `refresh_tokens` VALUES ('05646387-ce87-4e93-9f5a-751c8126e9a0','2b704681-a615-43cb-8793-9e9e2d754ecf','235933db1acab8fd28992815822af90be7cfc03fbb2e4453f69d86fc0cce802e','2025-12-11 20:55:46',0,'2025-12-04 20:55:46'),('06cbd02b-b97e-4a94-abaa-cafb1d45d220','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','d88b9b5d4a8219c290011d402dccf4a1fcf48527c8984fa74344a8f4dd9942b3','2025-12-12 10:38:03',1,'2025-12-05 10:38:03'),('0b1e2976-c09d-4a5e-9a97-5ded2c898355','d713ddfb-1a39-4b33-9826-8be0814a3f57','44adde05725aa2644d8c25c9235f35f503ee9e4af59b7bdc17a9b43ab904d2dc','2025-12-12 14:35:05',1,'2025-12-05 14:35:05'),('0f3adefe-1d57-4ab6-8c72-0608ec0537d5','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','9941335e83f784333a6fc5c2b74ef0b772731a97a4d35a2c7a38ecd36a84072f','2025-12-11 21:06:15',1,'2025-12-04 21:06:15'),('1dba79bc-7d28-4b38-b452-0c8de888b150','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','e43039ec18678291968cb248cb19eb80f2982901fe8d9454f1afc6c8cf27fbc6','2025-12-11 21:06:15',1,'2025-12-04 21:06:15'),('22d3e1de-81c8-4f9b-96a6-42b4d901d313','d713ddfb-1a39-4b33-9826-8be0814a3f57','d378f80528df623dc0680fcc3002d5ff9851a47a0f52d960640b2507a51df7cf','2025-12-12 13:58:43',1,'2025-12-05 13:58:43'),('25de28c1-ad73-46ab-98cd-821987cf931f','526e30af-f07e-4272-a63e-753ad48ab6c7','7d45066b3c9ced1a6a3573d1d83060f135ed5556585769bb254e830cb5c07be8','2025-12-12 09:29:32',1,'2025-12-05 09:29:32'),('320cbd3f-669b-4bdf-82b0-844b50eb83e1','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','c4bf20d83c81c16c107410ace9a0ed93b0054384fc9bf6f8f951d59da5562e2b','2025-12-12 10:12:59',1,'2025-12-05 10:12:59'),('3261b298-c881-4906-ad1a-4771ec3a3233','d713ddfb-1a39-4b33-9826-8be0814a3f57','b89b693a8fffa69882120ab20efaf5e8d4b2b3d6640db9b6d9ef2118de8c2f35','2025-12-12 14:18:00',1,'2025-12-05 14:18:00'),('34bf52dc-34d9-4d4c-aca9-6e3c5903fd54','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','cae3ed817a653fc69f2cdadbad3922cbaf192218f935851c3c4b20b7b0c64bd9','2025-12-11 21:08:39',1,'2025-12-04 21:08:39'),('3785815c-2787-4cbd-9604-bfd45c5412ff','526e30af-f07e-4272-a63e-753ad48ab6c7','f5bd8bff9205b0ffdd9af8b6437340c672a56cf872c6db6b46fc9771274acb29','2025-12-11 22:35:08',1,'2025-12-04 22:35:08'),('3fab19d4-684a-4068-ab08-a0183663e590','2b704681-a615-43cb-8793-9e9e2d754ecf','ad382034e5edf3e20ab940c3d6e6b72a2f49ab0399a74db1da60a7f0877f590a','2025-12-11 20:55:19',1,'2025-12-04 20:55:19'),('4131c532-136b-4068-a9d9-9ed767d294ad','d713ddfb-1a39-4b33-9826-8be0814a3f57','26a257da42b34b2d91bd684865d750d21e0e237c91d5ba4c9d45db46deece81b','2025-12-12 12:41:06',1,'2025-12-05 12:41:06'),('440637fa-d6aa-4891-a959-099a730c20dc','d713ddfb-1a39-4b33-9826-8be0814a3f57','c94ecf426468f9084a334721772312419210130d575ab21c66daa3ed46b665bc','2025-12-12 10:21:33',1,'2025-12-05 10:21:33'),('467fa467-085c-4398-9286-d57126ffb1da','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','0261685d6aa967a71ca82f1d62783b4de0ced660f3c937313e767bc401d0c394','2025-12-11 20:56:29',1,'2025-12-04 20:56:29'),('4d934c65-8008-4a63-9681-7ece7f85fa37','d713ddfb-1a39-4b33-9826-8be0814a3f57','cdc56db10a05db73b4c779daad89350e9e61290b87a335c6fb01bc61082a293a','2025-12-12 12:09:19',1,'2025-12-05 12:09:19'),('5036c665-2017-4dc7-8b8c-37f226020b3f','d713ddfb-1a39-4b33-9826-8be0814a3f57','ca197aebd58eee2e98ff396eba6eaa23a0ced6a47ad34c1b4ea19d1c3fffc7bf','2025-12-12 12:53:28',1,'2025-12-05 12:53:28'),('5afd7058-8f5f-47ea-b0dc-396334a48c96','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','db1c6d0a256ece526f67c4fb9208d00001012d984dde64512e8a0510785dc1fa','2025-12-12 10:18:12',1,'2025-12-05 10:18:12'),('5bb67988-72a8-4b59-892e-212fc3c79504','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','ce6ca55cdb3dff50e52c5ade68bdcc1715d35afac7277a11455747d5d1bbbe54','2025-12-12 15:06:30',1,'2025-12-05 15:06:30'),('5f02de90-20af-4afd-87fa-344ac6011f83','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','653da350f86accc19c8c1132700549dc00c61c62ad841ac7d66dee09870ac25e','2025-12-12 12:11:58',1,'2025-12-05 12:11:58'),('5f109e79-4b82-4822-86d8-5aa515ee062b','526e30af-f07e-4272-a63e-753ad48ab6c7','d6eeb09b7b21697da0b958ceefefa86ee5670e5677ca5960e0e9f9f7b0c6bb3f','2025-12-11 21:10:41',1,'2025-12-04 21:10:41'),('73d83ac4-9580-43d1-a83c-bbb700cfb525','d713ddfb-1a39-4b33-9826-8be0814a3f57','8c83a1e3fdaea8ee72e584c2b22d5e9f0c1c8f95f92e8d08a894bb549ea8d834','2025-12-12 13:32:10',1,'2025-12-05 13:32:10'),('77f07539-15af-45af-9350-f65b456b29e4','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','7106c2191b91ce9cd3b1a27fb0b6d420beb5a5b13f8c2db4857b2656bc7c7334','2025-12-12 12:52:33',1,'2025-12-05 12:52:33'),('84537f8a-1bb0-4507-80ad-d36d5a6531aa','d713ddfb-1a39-4b33-9826-8be0814a3f57','73e6432621e3bbcf0a7657619d6ec2e5669d578ca60e571978c98b0aafafbe32','2025-12-12 14:35:09',1,'2025-12-05 14:35:09'),('86c2ff02-0b47-41a4-a459-7276c2a6aef7','d713ddfb-1a39-4b33-9826-8be0814a3f57','4dfcbec7eebf3dd8caa2a4eb3becca29a061f09a4a769d7ccfc03cba00d6f963','2025-12-12 12:38:54',1,'2025-12-05 12:38:54'),('884a2abd-4572-49c1-b239-c58244ef8655','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','552f8ed4d8c3f17ec1f3917969bafd21ebafcd893d8bbb84c900920195324b22','2025-12-12 12:11:16',1,'2025-12-05 12:11:16'),('8d7caaf8-e843-4ae3-abff-120c592fb673','526e30af-f07e-4272-a63e-753ad48ab6c7','4e40958d3d38c0fceb86dd976fea8d38164b4a818bea1afbbb1bfcd2902d6ef1','2025-12-12 09:29:19',1,'2025-12-05 09:29:19'),('a04293ed-bb6b-4f7f-807e-b5f74c352355','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','671338956d62bf0c7ee675558f597d114552cde952e864f21c5b87f21da1c773','2025-12-11 20:56:29',1,'2025-12-04 20:56:29'),('afaea6af-bfdb-45a2-9e32-3b209f459ce1','d713ddfb-1a39-4b33-9826-8be0814a3f57','e0edb3c40a21e9d8572cc0f26f053c61105a8da4fa278744de15e9a239dfe162','2025-12-12 13:27:13',1,'2025-12-05 13:27:13'),('b5351956-7768-4755-bee1-a1d5ec5b7701','d713ddfb-1a39-4b33-9826-8be0814a3f57','a504d845da3ec81b88cfd62d227f3027281d05c3c8f9123eb0c2700c8f442431','2025-12-12 13:27:21',1,'2025-12-05 13:27:21'),('c0ebb394-b6c1-41bd-bc8e-a4e6526b970c','526e30af-f07e-4272-a63e-753ad48ab6c7','20824114e8ee084331aba58fdda94fda547582a2ef090ef0a4fb98549b7e37f2','2025-12-11 21:42:17',1,'2025-12-04 21:42:17'),('c7700e79-b6c2-4b35-aa72-c1493fbaf7f1','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','6ca764c629604ea723ca40d26f5a06e02988113e1f2e5590eb88473c425af2f4','2025-12-11 21:06:37',1,'2025-12-04 21:06:37'),('c7cc93f3-3190-4ecb-8405-7537a0213cad','d713ddfb-1a39-4b33-9826-8be0814a3f57','7944d601851dfb0888079372b0af0331dbc5dcf2bebe0b8a8dff771a53ad2237','2025-12-12 14:56:32',1,'2025-12-05 14:56:32'),('cb34aa09-ddc1-43eb-896d-500766dea64e','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','b762dabd1296b21dfefc85d0041962f20b8c083696dee1c3a861f4891da15bb3','2025-12-12 10:11:16',1,'2025-12-05 10:11:16'),('db8d2876-8008-4e59-9f16-9a5779dbf67f','d713ddfb-1a39-4b33-9826-8be0814a3f57','53b94fe60ee35030815884699246dac9d428256c38f65d4b934e00643f2c9173','2025-12-12 15:10:07',1,'2025-12-05 15:10:07'),('df315c2f-596c-438c-bbea-f13ff522f61e','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','36acd783d7756d970958de9eb5de486e560a75cebe5b2ec701f6530df63b7f03','2025-12-11 21:06:37',1,'2025-12-04 21:06:37'),('e74697b5-606f-44fd-83ce-db70abc83051','d713ddfb-1a39-4b33-9826-8be0814a3f57','030dbd1ffad671f05368f672542ad7dcf9943e1fe585e75c9683ededa24f5f00','2025-12-12 13:01:26',1,'2025-12-05 13:01:26'),('ead49a0a-72f2-4916-8057-f6b6730788a9','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','6d1494d37953b85b2923d9084035706888b06d4a920e19bc23c2ee49622e1d5f','2025-12-11 21:08:40',1,'2025-12-04 21:08:40'),('f6993d50-7f06-435a-9d07-57b413ec9c2a','a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','3b84512348fad6adf8864c8c9f445dd8eaa2b5a77266c4434d21ddf9cbbdc59f','2025-12-12 12:38:28',1,'2025-12-05 12:38:28'),('fc368271-fdff-4666-aad9-4a6492707a33','2b704681-a615-43cb-8793-9e9e2d754ecf','2cf1533b702c5a05f6daefeb08b194e813e7f3972f8e4e67b3abc051ac6a6226','2025-12-11 20:54:29',1,'2025-12-04 20:54:29'),('fd1395a9-1387-43e6-9cfa-53105f82ce81','526e30af-f07e-4272-a63e-753ad48ab6c7','45a4d8df4eedf48cbb63cdc5e117f76767005adf4ec36a964524312a04023efb','2025-12-11 21:09:16',1,'2025-12-04 21:09:16'),('ff2e66e2-4a27-4be0-a37e-2aa100301575','6dfd7a45-dca1-46ba-a87a-9dfb8f870563','01e19ed6c9148f4ec6d4d7f1209ffbc2b71620502e09798f195fb067bbb2daf8','2025-12-12 12:25:24',0,'2025-12-05 12:25:24');
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
  `dosen_nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tahun_ajaran` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `semester_type` enum('ganjil','genap') COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('draft','submitted','approved','rejected','revision') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'draft',
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
/*!40000 ALTER TABLE `rps` ENABLE KEYS */;
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
  `jenis` enum('utama','pendukung') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'utama',
  `judul` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `penulis` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `penerbit` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tahun` int DEFAULT NULL,
  `isbn` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
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
  `kriteria_penilaian` text COLLATE utf8mb4_unicode_ci,
  `urutan` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
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
/*!40000 ALTER TABLE `rps_evaluasi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rps_rencana_pembelajaran`
--

DROP TABLE IF EXISTS `rps_rencana_pembelajaran`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rps_rencana_pembelajaran` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rps_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pertemuan` int NOT NULL,
  `kemampuan_akhir` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator` text COLLATE utf8mb4_unicode_ci,
  `materi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `metode_pembelajaran` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `waktu_menit` int NOT NULL DEFAULT '150',
  `pengalaman_belajar` text COLLATE utf8mb4_unicode_ci,
  `kriteria_penilaian` text COLLATE utf8mb4_unicode_ci,
  `bobot_nilai` decimal(5,2) DEFAULT NULL,
  `referensi` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_rps_rencana_rps` (`rps_id`),
  CONSTRAINT `fk_rps_rencana_rps` FOREIGN KEY (`rps_id`) REFERENCES `rps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rps_rencana_pembelajaran`
--

LOCK TABLES `rps_rencana_pembelajaran` WRITE;
/*!40000 ALTER TABLE `rps_rencana_pembelajaran` DISABLE KEYS */;
/*!40000 ALTER TABLE `rps_rencana_pembelajaran` ENABLE KEYS */;
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
INSERT INTO `users` VALUES ('2233d21b-c4fe-43dd-aba9-91f5d5df99a7','test@test.com','$2a$10$3ojSXiSrXcD/Q7VE.cuSA.S7jfAtmBsJgqFpDKzxPvJ91BY1Avefi','Test User','123456','dosen','active',NULL,NULL,NULL,'2025-12-04 20:39:37','2025-12-04 20:39:37',NULL),('2b704681-a615-43cb-8793-9e9e2d754ecf','testnew@test.com','$2a$10$W/McMduO3jwpHtM/GZ0HweXRi6yhFyhC1MZdwuwNXs5UMar196rpi','Test User New',NULL,'dosen','active',NULL,NULL,'2025-12-04 20:55:45','2025-12-04 20:50:46','2025-12-04 20:55:46',NULL),('526e30af-f07e-4272-a63e-753ad48ab6c7','ani.wijaya@university.ac.id','$2a$10$jpF6EGZIzZG2KofOPt1QceNVxu2T3JJzqd/bUYL0/YW4qkpsUoVQa','Dr. Ani Wijaya, M.T','197812102005011003','dosen','active','081234567899',NULL,'2025-12-05 09:29:32','2025-12-04 20:26:37','2025-12-05 09:29:32',NULL),('6dfd7a45-dca1-46ba-a87a-9dfb8f870563','kaprodi@test.com','$2a$10$5fV0Q49809ahb2XcVQ21oeYKhvSHqmzhU7tYaZi2nS7jnReIOUeVW','Kaprodi Updated',NULL,'kaprodi','active',NULL,NULL,'2025-12-05 12:25:23','2025-12-04 20:56:29','2025-12-05 12:25:24',NULL),('a6d2cfef-c36d-4dbb-97b7-d06152efd6ae','aksan@unismuh.ac.id','$2a$10$ejayq9j8X5qIPUhyVNOPFudkWY77ypGKYR90/sEPE7W2Z5RudLLy6','Dr. Muhammad Aksan S. T, M.T','197812102005011001','dosen','active','081234567899',NULL,'2025-12-05 15:06:30','2025-12-05 10:10:49','2025-12-05 15:06:30',NULL),('be4e85bd-d139-11f0-888e-ec750c0888b7','kaprodi@university.ac.id','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZRGdjGj/n3.QNqQ.0y.naPxhPj3Ey','Dr. John Doe, M.Kom','198501012010011001','kaprodi','active',NULL,NULL,NULL,'2025-12-04 17:50:38','2025-12-04 17:50:38',NULL),('d713ddfb-1a39-4b33-9826-8be0814a3f57','aksan1@unismuh.ac.id','$2a$10$oCEbTH57ZKYcUwiuLnsO7OElK3/f.8sId96/go66QovmVySUhwcA2','Dr. Muhammad Aksan S. T, M.T','197812102005011002','kaprodi','active','081234567899',NULL,'2025-12-05 15:10:07','2025-12-05 10:20:23','2025-12-05 15:10:07',NULL),('e27f2069-daa4-4f1f-9494-6218fdeb2f5e','test2@test.com','$2a$10$E88uRsx10bZewNCRHVGNpeR8tTmUKzXyWMCG.BHHaC71BGOWJZFeq','Test User2','1234567','dosen','active',NULL,NULL,NULL,'2025-12-04 20:40:00','2025-12-04 20:40:00',NULL);
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

-- Dump completed on 2025-12-05 23:21:47
