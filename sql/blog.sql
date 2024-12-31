/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : 192.168.1.128:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 22/05/2024 09:49:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ptype` varchar(300) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, '2021-10-27 16:52:33.000', '2024-05-15 13:02:07.100', NULL, 'p', 'editor', '/v1/api/user/change_password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (2, '2021-10-27 17:00:15.000', '2021-10-27 17:00:29.000', NULL, 'p', 'editor', '/v1/api/user/logged_in', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (4, '2024-05-14 11:57:08.044', '2024-05-14 11:57:08.044', NULL, 'p', 'editor', '/v1/api/role', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for gin_auth
-- ----------------------------
DROP TABLE IF EXISTS `gin_auth`;
CREATE TABLE `gin_auth` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_id` bigint unsigned NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `logged_in_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(50) NOT NULL,
  `refresh_token` varchar(600) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_gin_auth_username` (`username`),
  UNIQUE KEY `uni_gin_auth_refresh_token` (`refresh_token`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_auth
-- ----------------------------
BEGIN;
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `password`, `refresh_token`) VALUES (1, '2024-05-10 16:39:36.066', '2024-05-21 16:20:21.101', NULL, 1, 0, '2024-05-21 16:19:12.097', 'admin', 'a203793c127cf17027b2cadbbff95355', '1');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `password`, `refresh_token`) VALUES (2, '2024-05-10 16:39:36.066', '2024-05-21 16:20:21.101', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor', 'a203793c127cf17027b2cadbbff95355', '2');
COMMIT;

-- ----------------------------
-- Table structure for gin_jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `gin_jwt_blacklist`;
CREATE TABLE `gin_jwt_blacklist` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `jwt` text,
  PRIMARY KEY (`id`),
  KEY `idx_blog_jwt_blacklist_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_jwt_blacklist
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gin_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_menu`;
CREATE TABLE `gin_menu` (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `menu_name` varchar(11) DEFAULT NULL COMMENT '''路由名称''',
  `path` varchar(128) DEFAULT NULL COMMENT '''路由路径''',
  `paths` varchar(128) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL COMMENT '''组件路径''',
  `title` varchar(64) DEFAULT NULL COMMENT '''菜单标题''',
  `icon` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `permission` varchar(32) DEFAULT NULL,
  `visible` int DEFAULT '0',
  `is_frame` int DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `idx_blog_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_menu
-- ----------------------------
BEGIN;
INSERT INTO `gin_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, 1, 'Upms', '/upms', NULL, 'Layout', '系统管理', 'example', 'M', NULL, 0, 0, '2021-11-05 15:49:22.000', '2021-11-05 15:49:30.000', NULL);
INSERT INTO `gin_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, 1, 'SysRole', '/permission/role', NULL, '/permission/role', '角色管理', NULL, 'M', NULL, 0, 0, '2021-11-05 15:51:53.000', '2021-11-05 15:51:57.000', NULL);
INSERT INTO `gin_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, 2, 'SysUser', '/permission/user', NULL, '/permission/user', '用户管理', NULL, 'M', NULL, 0, 0, '2021-11-05 16:01:20.000', '2021-11-05 16:01:24.000', NULL);
INSERT INTO `gin_menu` (`menu_id`, `parent_id`, `sort`, `menu_name`, `path`, `paths`, `component`, `title`, `icon`, `menu_type`, `permission`, `visible`, `is_frame`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 0, 1, 'dict', NULL, NULL, NULL, '字典管理', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for gin_report
-- ----------------------------
DROP TABLE IF EXISTS `gin_report`;
CREATE TABLE `gin_report` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `activity_id` bigint DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `ip` varchar(80) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_report
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gin_role
-- ----------------------------
DROP TABLE IF EXISTS `gin_role`;
CREATE TABLE `gin_role` (
  `role_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_name` varchar(128) DEFAULT NULL,
  `is_admin` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uni_gin_role_role_key` (`role_key`),
  CONSTRAINT `fk_gin_auth_role` FOREIGN KEY (`role_id`) REFERENCES `gin_auth` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_role
-- ----------------------------
BEGIN;
INSERT INTO `gin_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (1, NULL, '2024-05-21 11:03:04.239', NULL, '超管2233', 1, 0, 'admin', NULL, '超级管理员2');
INSERT INTO `gin_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (2, '2021-10-27 16:49:28.000', '2024-05-21 11:03:06.172', NULL, '编辑角色', 0, 0, 'editor', 0, '1111111111');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
