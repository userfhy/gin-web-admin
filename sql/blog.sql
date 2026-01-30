/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80407 (8.4.7)
 Source Host           : 192.168.1.128:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80407 (8.4.7)
 File Encoding         : 65001

 Date: 30/01/2026 10:18:18
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, '2021-10-27 16:52:33.000', '2024-05-15 13:02:07.100', NULL, 'p', 'editor', '/v1/api/user/change_password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (2, '2021-10-27 17:00:15.000', '2021-10-27 17:00:29.000', NULL, 'p', 'editor', '/v1/api/user/logged_in', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (4, '2024-05-14 11:57:08.044', '2024-05-14 11:57:08.044', NULL, 'p', 'editor', '/v1/api/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `created_at`, `updated_at`, `deleted_at`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (8, '2026-01-29 16:17:07.042', '2026-01-29 16:17:07.042', NULL, 'p', 'editor', '/v1/api/refresh_token', 'POST', '', '', '');
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
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (1, '2024-05-10 16:39:36.066', '2026-01-30 10:14:41.089', NULL, 1, 1, '2026-01-30 10:14:41.089', 'admin', 'fhy', '13839999999', 'aa@qq.com', 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzcwMzA3MjAwLCJpYXQiOjE3Njk3MzkyODF9.P_rWslJ9DAZrwuhWaNR_dATwACMuLAJE8EQK1gR9xJ8');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (2, '2024-05-10 16:39:36.066', '2026-01-30 10:12:56.283', NULL, 2, 1, '2026-01-30 10:12:56.282', 'editor', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImVkaXRvciIsInJvbGVfa2V5IjoiZWRpdG9yIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzcwMzA3MjAwLCJpYXQiOjE3Njk3MzkxNzZ9.ZHVP-Fu7JDYVM2gaf0CgSG7vPJuc3fDt10xNvU3M0D0');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (3, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor2', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '3');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (4, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor3', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '4');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (5, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor4', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '5');
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`) VALUES (6, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor5', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '6');
COMMIT;

-- ----------------------------
-- Table structure for gin_dept
-- ----------------------------
DROP TABLE IF EXISTS `gin_dept`;
CREATE TABLE `gin_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint DEFAULT NULL COMMENT '父部门ID',
  `dept_name` varchar(128) DEFAULT NULL COMMENT '部门名称',
  `order_num` bigint DEFAULT NULL COMMENT '排序',
  `leader` varchar(64) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(32) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态(1启用0停用)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_dept
-- ----------------------------
BEGIN;
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (1, 0, '总部', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (2, 1, '郑州分公司', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (3, 1, '开封分公司', 0, '', '', '', 1, '');
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
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (67, '2026-01-30 10:00:03.349', '2026-01-30 10:00:03.349', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzY5NzM5NzcxLCJpYXQiOjE3Njk3MzI1NzF9.rzRzoJ82RDR-pZUiLXB08lgAGIRUmxJPRLnInJZiCpE');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (68, '2026-01-30 10:01:25.828', '2026-01-30 10:01:25.828', NULL, 2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImVkaXRvciIsInJvbGVfa2V5IjoiZWRpdG9yIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzY5NzQ1NjA5LCJpYXQiOjE3Njk3Mzg0MDl9.7sjuc6jrh9qqdFcs2XAtNVemxkd9Ea0N_WG43mlm_ds');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (69, '2026-01-30 10:12:34.489', '2026-01-30 10:12:34.489', NULL, 2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImVkaXRvciIsInJvbGVfa2V5IjoiZWRpdG9yIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzY5NzQ2MzExLCJpYXQiOjE3Njk3MzkxMTF9.vCVaXFJaXh8PvyzLDI5POWXNpfnGP2XANp4TzroNwJg');
INSERT INTO `gin_jwt_blacklist` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `jwt`) VALUES (70, '2026-01-30 10:12:49.615', '2026-01-30 10:12:49.615', NULL, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzY5NzQ2MzYyLCJpYXQiOjE3Njk3MzkxNjJ9.0trEEWa1NhUyLnrFQJDU1urPZuLf5bgQV21zDMxB8y4');
COMMIT;

-- ----------------------------
-- Table structure for gin_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_menu`;
CREATE TABLE `gin_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint DEFAULT NULL COMMENT '父菜单ID',
  `menu_type` bigint DEFAULT NULL COMMENT '菜单类型（0:目录 1:菜单 2:按钮）',
  `title` varchar(100) DEFAULT NULL COMMENT '菜单标题',
  `name` varchar(100) DEFAULT NULL COMMENT '菜单名称（唯一值）',
  `path` varchar(255) DEFAULT NULL COMMENT '路由地址',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `rank` bigint DEFAULT NULL COMMENT '排序序号',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `extra_icon` varchar(100) DEFAULT NULL COMMENT '额外图标',
  `enter_transition` varchar(100) DEFAULT NULL COMMENT '进入动画',
  `leave_transition` varchar(100) DEFAULT NULL COMMENT '离开动画',
  `active_path` varchar(255) DEFAULT NULL COMMENT '激活路径',
  `auths` varchar(255) DEFAULT NULL COMMENT '权限标识,逗号分隔',
  `frame_src` varchar(255) DEFAULT NULL COMMENT '内嵌 iframe 地址',
  `frame_loading` tinyint(1) DEFAULT NULL COMMENT '是否显示 iframe 加载动画',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存组件',
  `hidden_tag` tinyint(1) DEFAULT NULL COMMENT '是否隐藏标签',
  `fixed_tag` tinyint(1) DEFAULT NULL COMMENT '是否固定标签',
  `show_link` tinyint(1) DEFAULT NULL COMMENT '是否显示链接',
  `show_parent` tinyint(1) DEFAULT NULL COMMENT '是否显示父级菜单',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_menu
-- ----------------------------
BEGIN;
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (1, 0, 0, 'menus.pureSysManagement', '', '/system', NULL, 1, NULL, 'ri:settings-3-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (2, 1, 1, 'menus.pureUser', 'SystemUser', '/system/user/index', 'system/user/index', 0, NULL, 'ri:admin-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (3, 1, 1, 'menus.pureRole', 'SystemRole', '/system/role/index', 'system/role/index', 0, NULL, 'ri:admin-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (4, 1, 1, 'menus.pureSystemMenu', 'SystemMenu', '/system/menu/index', 'system/menu/index', 0, NULL, 'ep:menu', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (5, 1, 1, 'menus.pureSystemApi', 'SystemAPI', '/system/api/index', 'system/api/index', 0, NULL, 'ep:list', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (6, 1, 1, 'menus.pureDept', 'SystemDept', '/system/dept/index', 'system/dept/index', 0, NULL, 'ri:git-branch-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (7, 0, 0, 'menus.pureSysMonitor', '', '/monitor', NULL, 2, NULL, 'ep:monitor', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (8, 7, 1, 'menus.pureOnlineUser', 'OnlineUser', '/monitor/online-user', 'monitor/online/index', 0, NULL, 'ri:user-voice-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (9, 7, 1, 'menus.pureLoginLog', 'LoginLog', '/monitor/login-logs', 'monitor/logs/login/index', 0, NULL, 'ri:window-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (10, 7, 1, 'menus.pureOperationLog', 'OperationLog', '/monitor/operation-logs', 'monitor/logs/operation/index', 0, NULL, 'ri:history-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (11, 7, 1, 'menus.pureSystemLog', 'SystemLog', '/monitor/system-logs', 'monitor/logs/system/index', 0, NULL, 'ri:file-search-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (12, 0, 0, 'menus.purePermission', '', '/permission', NULL, 3, NULL, 'ep:lollipop', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (13, 12, 1, 'menus.purePermissionPage', 'PermissionPage', '/permission/page/index', 'permission/page/index', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (14, 12, 0, 'menus.purePermissionButton', '', '/permission/button', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (15, 14, 1, 'menus.purePermissionButtonRouter', 'PermissionButtonRouter', '/permission/button/router', 'permission/button/index', 0, NULL, '', NULL, NULL, NULL, '', 'permission:btn:add,permission:btn:edit,permission:btn:delete', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (16, 14, 1, 'menus.purePermissionButtonLogin', 'PermissionButtonLogin', '/permission/button/login', 'permission/button/perms', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (17, 0, 0, 'menus.pureExternalPage', '', '/iframe', NULL, 4, NULL, 'ri:links-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (18, 17, 0, 'menus.pureEmbeddedDoc', '', '/iframe/embedded', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (19, 18, 1, 'menus.pureColorHuntDoc', 'FrameColorHunt', '/iframe/colorhunt', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://colorhunt.co/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (20, 18, 1, 'menus.pureUiGradients', 'FrameUiGradients', '/iframe/uigradients', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://uigradients.com/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (21, 18, 1, 'menus.pureEpDoc', 'FrameEp', '/iframe/ep', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://element-plus.org/zh-CN/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (22, 18, 1, 'menus.pureTailwindcssDoc', 'FrameTailwindcss', '/iframe/tailwindcss', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://tailwindcss.com/docs/installation', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (23, 18, 1, 'menus.pureVueDoc', 'FrameVue', '/iframe/vue3', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://cn.vuejs.org/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (24, 18, 1, 'menus.pureViteDoc', 'FrameVite', '/iframe/vite', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://cn.vitejs.dev/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (25, 18, 1, 'menus.purePiniaDoc', 'FramePinia', '/iframe/pinia', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://pinia.vuejs.org/zh/index.html', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (26, 18, 1, 'menus.pureRouterDoc', 'FrameRouter', '/iframe/vue-router', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://router.vuejs.org/zh/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (27, 17, 0, 'menus.pureExternalDoc', '', '/iframe/external', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (28, 27, 1, 'menus.pureExternalLink', 'https://pure-admin.cn/', '/external', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (29, 27, 1, 'menus.pureUtilsLink', 'https://pure-admin-utils.netlify.app/', '/pureUtilsLink', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (30, 0, 0, 'menus.pureTabs', '', '/tabs', NULL, 5, NULL, 'ri:bookmark-2-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (31, 30, 1, 'menus.pureTabs', 'Tabs', '/tabs/index', 'tabs/index', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (32, 30, 1, '', 'TabQueryDetail', '/tabs/query-detail', NULL, 0, NULL, '', NULL, NULL, NULL, '/tabs/index', '', '', NULL, 0, NULL, NULL, 0, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (33, 30, 1, '', 'TabParamsDetail', '/tabs/params-detail/:id', 'params-detail', 0, NULL, '', NULL, NULL, NULL, '/tabs/index', '', '', NULL, 0, NULL, NULL, 0, NULL);
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
-- Table structure for gin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_role_menu`;
CREATE TABLE `gin_role_menu` (
  `role_id` bigint unsigned NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 1);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 2);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 3);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 4);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 5);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 6);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 7);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 8);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 9);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 10);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 11);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 12);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 13);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 14);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 15);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 19);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 20);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 21);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 22);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 23);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 24);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 25);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 26);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 27);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 28);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 29);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 30);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 31);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 32);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 33);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 12);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 13);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 14);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 15);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 16);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 17);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 18);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 19);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 20);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 21);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 22);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 23);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 24);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 25);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 26);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 27);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 28);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 29);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 30);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 31);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 32);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (2, 33);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
