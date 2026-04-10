/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80408 (8.4.8)
 Source Host           : 192.168.1.128:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80408 (8.4.8)
 File Encoding         : 65001

 Date: 10/04/2026 15:12:37
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
  `failed_login_count` int NOT NULL DEFAULT '0',
  `locked_until` datetime(3) DEFAULT NULL,
  `last_login_ip` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_gin_auth_username` (`username`),
  UNIQUE KEY `uni_gin_auth_refresh_token` (`refresh_token`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of gin_auth
-- ----------------------------
BEGIN;
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (1, '2024-05-10 16:39:36.066', '2026-04-10 14:05:54.454', NULL, 1, 1, '2026-04-10 14:05:54.454', 'admin', 'fhy', '13839999999', 'aa@qq.com', 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZV9rZXkiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzc2MzU1MjAwLCJpYXQiOjE3NzU4MDExNTR9.j2yFZ2xDfDFIePy6hUS6hNdKxV0zMD4k3dajERfDNW8', 0, NULL, NULL);
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (2, '2024-05-10 16:39:36.066', '2026-01-30 11:09:31.753', NULL, 2, 1, '2026-01-30 11:09:31.753', 'editor', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImVkaXRvciIsInJvbGVfa2V5IjoiZWRpdG9yIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzcwMzA3MjAwLCJpYXQiOjE3Njk3NDI1NzF9.YsRC6J1VXClDI_h2_-u3T1bCfh-duJasVkIE7aBiEcc', 0, NULL, NULL);
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (3, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor2', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '3', 0, NULL, NULL);
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (4, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor3', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '4', 0, NULL, NULL);
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (5, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor4', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '5', 0, NULL, NULL);
INSERT INTO `gin_auth` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_id`, `status`, `logged_in_at`, `username`, `nickname`, `phone`, `email`, `sex`, `password`, `refresh_token`, `failed_login_count`, `locked_until`, `last_login_ip`) VALUES (6, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor5', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '6', 0, NULL, NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_dept
-- ----------------------------
BEGIN;
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (1, 0, '总部', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (2, 1, '郑州分公司', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (3, 1, '开封分公司', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (4, 0, '顶级分类', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` (`id`, `parent_id`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `remark`) VALUES (5, 4, '北京公司', 0, '', '', '', 1, '');
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
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (34, 0, 0, '官网管理', '', '/site', NULL, 6, NULL, 'ri:global-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (35, 34, 1, '内容管理', 'SiteContent', '/site/content/index', 'site/content/index', 0, NULL, 'ri:article-line', NULL, NULL, NULL, '', '', '', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (36, 34, 1, '类别管理', 'SiteCategory', '/site/category/index', 'site/category/index', 0, NULL, 'ri:price-tag-3-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`) VALUES (37, 34, 1, '标签管理', 'SiteTag', '/site/tag/index', 'site/tag/index', 0, NULL, 'ri:price-tag-2-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
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
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 34);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 35);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 36);
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`) VALUES (1, 37);
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

-- ----------------------------
-- Table structure for gin_site_category
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_category`;
CREATE TABLE `gin_site_category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `slug` varchar(120) NOT NULL COMMENT '分类标识',
  `description` varchar(500) DEFAULT NULL COMMENT '分类描述',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT '0' COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_site_category_slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_site_category
-- ----------------------------
BEGIN;
INSERT INTO `gin_site_category` (`id`, `name`, `slug`, `description`, `status`, `sort`, `created_at`, `updated_at`) VALUES (1, '分类news', 'news', '', 1, 0, '2026-04-05 18:03:47.701', '2026-04-10 14:59:29.769');
COMMIT;

-- ----------------------------
-- Table structure for gin_site_content
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content`;
CREATE TABLE `gin_site_content` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(150) NOT NULL COMMENT '内容标题',
  `slug` varchar(150) NOT NULL COMMENT '页面标识',
  `summary` varchar(500) DEFAULT NULL COMMENT '摘要',
  `cover` varchar(500) DEFAULT NULL COMMENT '封面图',
  `content` longtext COMMENT '正文内容',
  `seo_keywords` varchar(255) DEFAULT NULL COMMENT 'SEO关键词',
  `seo_description` varchar(500) DEFAULT NULL COMMENT 'SEO描述',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态(1发布0草稿)',
  `sort` bigint NOT NULL DEFAULT '0' COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `published_at` datetime(3) DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_site_content_slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_site_content
-- ----------------------------
BEGIN;
INSERT INTO `gin_site_content` (`id`, `title`, `slug`, `summary`, `cover`, `content`, `seo_keywords`, `seo_description`, `status`, `sort`, `created_at`, `updated_at`, `published_at`) VALUES (1, '测试标题111', '1', '**d是否🎉️ 第三方分多少 &lt;script&gt;alert(222333);&lt;/script&gt;', '', '**d是否🎉️ 第三方分多少234**\n\ndsfsdf第三方士大夫水电费水电费水电费水电费水电费水电费水电费胜多负少的给对方回复过几年更好看❤️ 🎉️ 😕\n\n都是1111\n\n```javascript\n&lt;script&gt;alert(222333);&lt;/script&gt;\n```', '', '', 1, 1, '2026-04-05 17:41:45.080', '2026-04-10 15:01:40.695', '2026-04-05 18:22:55.089');
INSERT INTO `gin_site_content` (`id`, `title`, `slug`, `summary`, `cover`, `content`, `seo_keywords`, `seo_description`, `status`, `sort`, `created_at`, `updated_at`, `published_at`) VALUES (2, '测试22', '2', '第三方1', '', '第三方222', '', '', 1, 0, '2026-04-10 10:49:30.702', '2026-04-10 15:01:39.508', '2026-04-10 10:49:30.702');
COMMIT;

-- ----------------------------
-- Table structure for gin_site_content_category
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content_category`;
CREATE TABLE `gin_site_content_category` (
  `content_id` bigint NOT NULL,
  `category_id` bigint NOT NULL,
  PRIMARY KEY (`content_id`,`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_site_content_category
-- ----------------------------
BEGIN;
INSERT INTO `gin_site_content_category` (`content_id`, `category_id`) VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for gin_site_content_tag
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content_tag`;
CREATE TABLE `gin_site_content_tag` (
  `content_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  PRIMARY KEY (`content_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_site_content_tag
-- ----------------------------
BEGIN;
INSERT INTO `gin_site_content_tag` (`content_id`, `tag_id`) VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for gin_site_tag
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_tag`;
CREATE TABLE `gin_site_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(80) NOT NULL COMMENT '标签名',
  `slug` varchar(120) NOT NULL COMMENT '标签标识',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT '0' COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_site_tag_slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of gin_site_tag
-- ----------------------------
BEGIN;
INSERT INTO `gin_site_tag` (`id`, `name`, `slug`, `status`, `sort`, `created_at`, `updated_at`) VALUES (1, '标签', 'biaoqian', 1, 0, '2026-04-05 19:00:12.457', '2026-04-05 19:00:12.457');
COMMIT;

-- ----------------------------
-- Table structure for gin_audit_log
-- ----------------------------
DROP TABLE IF EXISTS `gin_audit_log`;
CREATE TABLE `gin_audit_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category` varchar(20) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `username` varchar(60) DEFAULT NULL,
  `ip` varchar(64) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `method` varchar(10) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `action` varchar(120) DEFAULT NULL,
  `message` text,
  PRIMARY KEY (`id`),
  KEY `idx_audit_category` (`category`),
  KEY `idx_audit_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
