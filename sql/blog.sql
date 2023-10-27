/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : 192.168.1.128:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 27/10/2023 14:35:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_id` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `logged_in_at` datetime DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_blog_auth_username` (`username`),
  KEY `idx_blog_auth_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `password`) VALUES (1, NULL, '2023-08-14 17:28:33', NULL, 1, 0, '2023-08-14 17:28:33', 'admin', '1B44F66F94C34FB0485462E883D81AB5');
INSERT INTO `blog_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `password`) VALUES (2, NULL, '2021-10-27 17:00:41', NULL, 2, 0, '2021-10-27 17:00:41', 'editor', '1B44F66F94C34FB0485462E883D81AB5');
COMMIT;

-- ----------------------------
-- Table structure for blog_jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `blog_jwt_blacklist`;
CREATE TABLE `blog_jwt_blacklist` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int unsigned DEFAULT NULL,
  `jwt` text,
  PRIMARY KEY (`id`),
  KEY `idx_blog_jwt_blacklist_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_jwt_blacklist
-- ----------------------------
BEGIN;
INSERT INTO `blog_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (1, '2021-10-27 11:36:40', '2021-10-27 11:36:40', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiZXhwIjoxNjM1MzUwNDAwLCJpYXQiOjE2MzUzMDQ0MTIsImlzcyI6Imdpbi10ZXN0In0.qsfy-DV4cfkrV3mimq_kIqWNYJJkSdiaWiFmW2K37v8');
INSERT INTO `blog_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (2, '2021-10-27 16:54:24', '2021-10-27 16:54:24', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiZXhwIjoxNjM1MzUwNDAwLCJpYXQiOjE2MzUzMjQ0MjQsImlzcyI6Imdpbi10ZXN0In0.PlFXCJS5gbCKFAkjun-vmo8SwPHHIJnWCeqWw_rWWGQ');
COMMIT;

-- ----------------------------
-- Table structure for blog_menu
-- ----------------------------
DROP TABLE IF EXISTS `blog_menu`;
CREATE TABLE `blog_menu` (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `menu_name` varchar(11) DEFAULT NULL COMMENT '路由名称',
  `path` varchar(128) DEFAULT NULL COMMENT '路由路径',
  `paths` varchar(128) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `title` varchar(64) DEFAULT NULL COMMENT '菜单标题',
  `icon` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `permission` varchar(32) DEFAULT NULL,
  `visible` int DEFAULT '0',
  `is_frame` int DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `idx_blog_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_menu
-- ----------------------------
BEGIN;
INSERT INTO `blog_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, 1, 'Upms', '/upms', NULL, 'Layout', '系统管理', 'example', 'M', NULL, 0, 0, '2021-11-05 15:49:22', '2021-11-05 15:49:30', NULL);
INSERT INTO `blog_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, 1, 'SysRole', '/permission/role', NULL, '/permission/role', '角色管理', NULL, 'M', NULL, 0, 0, '2021-11-05 15:51:53', '2021-11-05 15:51:57', NULL);
INSERT INTO `blog_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, 2, 'SysUser', '/permission/user', NULL, '/permission/user', '用户管理', NULL, 'M', NULL, 0, 0, '2021-11-05 16:01:20', '2021-11-05 16:01:24', NULL);
INSERT INTO `blog_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 0, 1, 'dict', NULL, NULL, NULL, '字典管理', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for blog_report
-- ----------------------------
DROP TABLE IF EXISTS `blog_report`;
CREATE TABLE `blog_report` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `activity_id` int DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `ip` varchar(80) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blog_report_deleted_at` (`deleted_at`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_report
-- ----------------------------
BEGIN;
INSERT INTO `blog_report` (`id`, `created_at`, `updated_at`, `deleted_at`, `activity_id`, `name`, `phone`, `ip`) VALUES (1, '2021-10-26 16:06:22', '2021-10-26 16:06:22', NULL, 1, '', '110000', '192.168.1.128');
COMMIT;

-- ----------------------------
-- Table structure for blog_role
-- ----------------------------
DROP TABLE IF EXISTS `blog_role`;
CREATE TABLE `blog_role` (
  `role_id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(128) DEFAULT NULL,
  `is_admin` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uix_blog_role_role_key` (`role_key`),
  KEY `idx_blog_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of blog_role
-- ----------------------------
BEGIN;
INSERT INTO `blog_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (1, NULL, '2021-11-05 10:39:00', NULL, '超管223331', 1, 0, 'admin', NULL, '超级管理员');
INSERT INTO `blog_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (2, '2021-10-27 16:49:28', '2021-10-27 16:49:28', NULL, '编辑角色', 0, 0, 'editor', 0, '1111111111');
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `ptype` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`),
  KEY `idx_casbin_rule_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `ptype`) VALUES (1, '2021-10-27 16:52:33', '2021-11-05 14:21:22', NULL, 'p', 'editor', '/v1/api/user/change_password', 'PUT', '', '', '', NULL);
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `ptype`) VALUES (2, '2021-10-27 17:00:15', '2021-10-27 17:00:29', NULL, 'p', 'editor', '/v1/api/user/logged_in', 'GET', '', '', '', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
