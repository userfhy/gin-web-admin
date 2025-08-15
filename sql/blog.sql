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

 Date: 14/08/2025 09:27:30
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
  `ptype` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1-女 2-男',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `refresh_token` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_gin_auth_username` (`username`),
  UNIQUE KEY `uni_gin_auth_refresh_token` (`refresh_token`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_auth
-- ----------------------------
BEGIN;
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (1, '2024-05-10 16:39:36.066', '2025-04-14 15:26:15.446', NULL, 1, 1, '2025-04-14 15:26:15.446', 'admin', 'fhy', '13839999999', 'aa@qq.com', 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQ1MTY0ODAwLCJpYXQiOjE3NDQ2MTU1NzV9.LLLDbntgYrMow0oWnkTtrUleA1WS9Y6Vr90zFdyWR_s');
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
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  KEY `idx_blog_jwt_blacklist_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_jwt_blacklist
-- ----------------------------
BEGIN;
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (53, '2025-03-09 16:20:06.457', '2025-03-09 16:20:06.457', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNTEyMzM5LCJpYXQiOjE3NDE1MDUxMzl9.ZQ3GcKcFTViabKxEE4UN8MDyJ_P3zk6ANpV8FOJQQ3s');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (54, '2025-03-10 15:49:37.684', '2025-03-10 15:49:37.684', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNTk1NzQ4LCJpYXQiOjE3NDE1ODg1NDh9.mBnz1WfSU3RFwRci0QTKuB_zpcX3J2MMqJbp1bMt9lk');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (55, '2025-03-10 15:55:15.811', '2025-03-10 15:55:15.811', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAwMTgxLCJpYXQiOjE3NDE1OTI5ODF9.8wN8UeYPxOlaorjEnLxHm0xxlbH9hRhDlNTDBw_-wt0');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (56, '2025-03-10 15:56:35.592', '2025-03-10 15:56:35.592', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAwNTE4LCJpYXQiOjE3NDE1OTMzMTh9.dKNrdibPC98tjgadSZ6F4ua9_MHDaoyroU4GjZREgFY');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (57, '2025-03-10 15:57:49.024', '2025-03-10 15:57:49.024', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAwNTk5LCJpYXQiOjE3NDE1OTMzOTl9.5IEJMUG6WpJPRwYMcHVSL-iYfOejUloRHvV-mhaOfsY');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (58, '2025-03-10 15:58:33.418', '2025-03-10 15:58:33.418', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAwNjczLCJpYXQiOjE3NDE1OTM0NzN9.E2b9gPfkrebnV5kwrvXkd4cok2X1oD-ihpxVnXN4UaU');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (59, '2025-03-10 15:59:14.634', '2025-03-10 15:59:14.634', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAwNzE3LCJpYXQiOjE3NDE1OTM1MTd9.Jivy3Z8E7vNCUPXCljef6GqFn2UhgX2n4Y00ChmtceM');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (60, '2025-03-10 16:15:16.614', '2025-03-10 16:15:16.614', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAxNTU2LCJpYXQiOjE3NDE1OTQzNTZ9.8LjrKDC_Nob9Q1vJZRITUj8Dyjw3BaccgogDwiFCYoI');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (61, '2025-03-10 16:16:29.314', '2025-03-10 16:16:29.314', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAxNzI0LCJpYXQiOjE3NDE1OTQ1MjR9.ZWBsS6bXvqR0efZFxu6uWUIv4LjPeQaDP7XayXX-CNo');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (62, '2025-03-10 16:17:36.074', '2025-03-10 16:17:36.074', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAxODI0LCJpYXQiOjE3NDE1OTQ2MjR9.Af17NI-8xhGfKapPSm_C6VtJEL1vgiYL3b8gHFMSVd8');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (63, '2025-03-10 16:18:03.198', '2025-03-10 16:18:03.198', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAxODYwLCJpYXQiOjE3NDE1OTQ2NjB9.0XSpn4xt6bLJhFkQ-C8r-BJe-MD2XaalQbpaIUsG7FU');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (64, '2025-03-10 16:22:34.045', '2025-03-10 16:22:34.045', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAxODg2LCJpYXQiOjE3NDE1OTQ2ODZ9.6cg898mZBtQ1aXP96atXK2w7FfRFhUhEAeSL283jnXs');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (65, '2025-03-10 16:47:26.300', '2025-03-10 16:47:26.300', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQxNjAyNTQ5LCJpYXQiOjE3NDE1OTUzNDl9.L4l5Na16nm61MkXNFpJSQnpS14Ntl6GT427vhXoMVMo');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (66, '2025-04-14 15:26:10.615', '2025-04-14 15:26:10.615', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzQ0NjIyNzYyLCJpYXQiOjE3NDQ2MTU1NjJ9.aNt7hq5hDr2MT0p1K9gcOWuRoQDZ87FKUn9ZdkOf6zc');
COMMIT;

-- ----------------------------
-- Table structure for gin_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_menu`;
CREATE TABLE `gin_menu` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `menu_name` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''路由名称''',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''路由路径''',
  `paths` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''组件路径''',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''菜单标题''',
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `menu_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `permission` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `visible` int DEFAULT '0',
  `is_frame` int DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_menu
-- ----------------------------
BEGIN;
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
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ip` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_admin` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '0',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_sort` int DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uni_gin_role_role_key` (`role_key`),
  CONSTRAINT `fk_gin_auth_role` FOREIGN KEY (`role_id`) REFERENCES `gin_auth` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_role
-- ----------------------------
BEGIN;
INSERT INTO `gin_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (1, NULL, '2024-05-21 11:03:04.239', NULL, '超管2233', 1, 0, 'admin', NULL, '超级管理员2');
INSERT INTO `gin_role` (`role_id`, `created_at`, `updated_at`, `deleted_at`, `role_name`, `is_admin`, `status`, `role_key`, `role_sort`, `remark`) VALUES (2, '2021-10-27 16:49:28.000', '2024-05-21 11:03:06.172', NULL, '编辑角色', 0, 0, 'editor', 0, '1111111111');
COMMIT;

-- ----------------------------
-- Table structure for gin_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_sys_menu`;
CREATE TABLE `gin_sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由名称',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单图标',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单标题',
  `sort` bigint DEFAULT NULL COMMENT '排序',
  `show_link` tinyint(1) DEFAULT '1' COMMENT '是否显示',
  `frame_src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'iframe外链地址',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `auths` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '权限标识（逗号分隔）',
  `roles` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '允许角色（逗号分隔）',
  `active_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '激活路径',
  `rank` bigint DEFAULT NULL COMMENT '菜单分类等级',
  PRIMARY KEY (`id`),
  KEY `idx_gin_sys_menu_deleted_at` (`deleted_at`),
  KEY `idx_gin_sys_menu_parent_id` (`parent_id`),
  KEY `idx_gin_sys_menu_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_sys_menu
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
