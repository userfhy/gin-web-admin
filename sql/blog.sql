/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80402 (8.4.2)
 Source Host           : 192.168.1.128:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80402 (8.4.2)
 File Encoding         : 65001

 Date: 03/01/2025 16:18:58
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
  `nickname` varchar(30) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `email` varchar(40) DEFAULT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1-女 2-男',
  `password` varchar(50) NOT NULL,
  `refresh_token` varchar(600) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_gin_auth_username` (`username`),
  UNIQUE KEY `uni_gin_auth_refresh_token` (`refresh_token`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_auth
-- ----------------------------
BEGIN;
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (1, '2024-05-10 16:39:36.066', '2024-12-31 11:41:42.832', NULL, 1, 1, '2024-12-31 11:41:42.831', 'admin', 'fhy', '13839999999', 'aa@qq.com', 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTczNjE3OTIwMCwiaWF0IjoxNzM1NjE2NTAyfQ.tu6YlMj33bKjH-IhDXTFQ1PPCAajxFVf7zeKTdzf0nI');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (2, '2024-05-10 16:39:36.066', '2024-08-27 11:36:07.524', NULL, 2, 1, '2024-08-27 11:34:54.024', 'editor', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '2');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (3, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor2', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '3');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (4, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor3', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '4');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (5, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor4', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '5');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (6, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor5', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '6');
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
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_jwt_blacklist
-- ----------------------------
BEGIN;
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (32, '2024-05-24 15:39:29.863', '2024-05-24 15:39:29.863', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzE2NTQxOTc2LCJpYXQiOjE3MTY1MzQ3NzZ9.TyZ7wLxGSKBLN9CHXjx5qhaxQQ_ErMEoo_ARbjRmwdY');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (33, '2024-05-24 16:49:32.059', '2024-05-24 16:49:32.059', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNjU0MzU3MywiaWF0IjoxNzE2NTM2MzczfQ.ibNNusn6Lv__AJQtK6Lp_8HuPomrarOdugCGInzsWlM');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (34, '2024-05-24 16:56:08.396', '2024-05-24 16:56:08.396', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNjU0ODEyMiwiaWF0IjoxNzE2NTQwOTIyfQ.kKnjAPMkVpgATyCzCIg3pF_hYkxoxTHtdHgIKwbqPEE');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (35, '2024-05-27 10:35:42.441', '2024-05-27 10:35:42.441', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzE2NzgxNzUwLCJpYXQiOjE3MTY3NzQ1NTB9.sdjry8xX2UNHlK4Ha05VZu0LVme_zcbjqp0yQqISwhQ');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (36, '2024-05-27 15:02:03.164', '2024-05-27 15:02:03.164', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzE2Nzk5NzI1LCJpYXQiOjE3MTY3OTI1MjV9.VovJoBLTgs13cJDAflfng3V3oJEOFrkApMHBZ9Q8LmY');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (37, '2024-05-27 16:30:20.148', '2024-05-27 16:30:20.148', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNjgwMDYzNywiaWF0IjoxNzE2NzkzNDM3fQ.3OYA_B4dHQcO5p6OOm3EcvwkBOhmcBYNXne2Sdvkw7c');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (38, '2024-06-08 08:59:33.010', '2024-06-08 08:59:33.010', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNzgxMzA4NSwiaWF0IjoxNzE3ODA1ODg1fQ.8VmyLhQZRBrUdn3_5_YiMfDuazg0i70GNhsK4P8VQ4M');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (39, '2024-06-08 09:01:22.238', '2024-06-08 09:01:22.238', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNzgxNTY4MiwiaWF0IjoxNzE3ODA4NDgyfQ.tYpbx7Je0MVTqpsEv4Zhvt7fpk4PkOGhv-4SPqVKqoU');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (40, '2024-06-08 09:07:33.350', '2024-06-08 09:07:33.350', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNzgxNjA1MywiaWF0IjoxNzE3ODA4ODUzfQ.dK0Yh8kwgBagXy_yL06XdcLE0TtkfN0lwAh-fWVwTHo');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (41, '2024-06-08 09:31:13.986', '2024-06-08 09:31:13.986', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcxNzgxNjA1NiwiaWF0IjoxNzE3ODA4ODU2fQ.iwsEhR2CxiOzmYiq7VSRSh8gUT-VlN3fbGGywiqCZGg');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (42, '2024-08-27 11:25:42.846', '2024-08-27 11:25:42.846', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcyNDczNTczNywiaWF0IjoxNzI0NzI4NTM3fQ.DSRSOMahNBeagYUr3cQwIt3tVkDYt172iozcJQ-o8ho');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (43, '2024-08-27 11:34:04.007', '2024-08-27 11:34:04.007', NULL, 2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJlZGl0b3IiLCJyb2xlX2tleSI6ImVkaXRvciIsImlzX2FkbWluIjpmYWxzZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcyNDczNjQ5MywiaWF0IjoxNzI0NzI5MjkzfQ.6GYn3-_yrrpJa5E1xD6E4lwZiLfgnh7j375lmGcfYOU');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (44, '2024-08-27 11:36:07.523', '2024-08-27 11:36:07.523', NULL, 2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJlZGl0b3IiLCJyb2xlX2tleSI6ImVkaXRvciIsImlzX2FkbWluIjpmYWxzZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcyNDczNjg5NCwiaWF0IjoxNzI0NzI5Njk0fQ.ALKb5uN-VOdGzBsJ067cArLhOEcHuA9Uuf04lkhhiSI');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (45, '2024-08-27 11:39:17.727', '2024-08-27 11:39:17.727', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcyNDczNjk3MiwiaWF0IjoxNzI0NzI5NzcyfQ.VntUeYfRWSMcXrUAlOYruXABsVadbOa2-q0cJjPzk20');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (46, '2024-09-23 11:15:16.421', '2024-09-23 11:15:16.421', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJhZG1pbiIsInJvbGVfa2V5IjoiYWRtaW4iLCJpc19hZG1pbiI6dHJ1ZSwiaXNzIjoiZ2luLXdlYi1hZG1pbiIsImV4cCI6MTcyNzA2ODAwMywiaWF0IjoxNzI3MDYwODAzfQ.e_JYdURCrBV-IN1u6Saqe6tf7v-epO2hfYGjQ6eMAKs');
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
