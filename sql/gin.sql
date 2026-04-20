/*
 Navicat Premium Dump SQL

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 90600 (9.6.0)
 Source Host           : 127.0.0.1:3306
 Source Schema         : gin

 Target Server Type    : MySQL
 Target Server Version : 90600 (9.6.0)
 File Encoding         : 65001

 Date: 20/04/2026 12:49:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ptype` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_index`(`v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, '2021-10-27 16:52:33.000', '2024-05-15 13:02:07.100', NULL, 'p', 'editor', '/v1/api/user/change_password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, '2021-10-27 17:00:15.000', '2021-10-27 17:00:29.000', NULL, 'p', 'editor', '/v1/api/user/logged_in', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, '2024-05-14 11:57:08.044', '2024-05-14 11:57:08.044', NULL, 'p', 'editor', '/v1/api/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, '2026-01-29 16:17:07.042', '2026-01-29 16:17:07.042', NULL, 'p', 'editor', '/v1/api/refresh_token', 'POST', '', '', '');

-- ----------------------------
-- Table structure for gin_audit_log
-- ----------------------------
DROP TABLE IF EXISTS `gin_audit_log`;
CREATE TABLE `gin_audit_log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `category` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `username` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT NULL,
  `action` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_audit_category`(`category` ASC) USING BTREE,
  INDEX `idx_audit_user`(`user_id` ASC) USING BTREE,
  INDEX `idx_gin_audit_log_category`(`category` ASC) USING BTREE,
  INDEX `idx_gin_audit_log_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_audit_log
-- ----------------------------

-- ----------------------------
-- Table structure for gin_auth
-- ----------------------------
DROP TABLE IF EXISTS `gin_auth`;
CREATE TABLE `gin_auth`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_id` bigint UNSIGNED NOT NULL DEFAULT 0,
  `status` int NOT NULL DEFAULT 0,
  `logged_in_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '1-女 2-男',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `refresh_token` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `failed_login_count` int NOT NULL DEFAULT 0,
  `locked_until` datetime(3) NULL DEFAULT NULL,
  `last_login_ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_gin_auth_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `uni_gin_auth_refresh_token`(`refresh_token` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_auth
-- ----------------------------
INSERT INTO `gin_auth` VALUES (1, '2024-05-10 16:39:36.066', '2026-04-20 12:46:56.318', NULL, 1, 1, '2026-04-20 12:46:56.302', 'admin', 'fhy', '13839999999', 'aa@qq.com', 2, 'a203793c127cf17027b2cadbbff95355', '', 0, NULL, '127.0.0.1');
INSERT INTO `gin_auth` VALUES (2, '2024-05-10 16:39:36.066', '2026-01-30 11:09:31.753', NULL, 2, 1, '2026-01-30 11:09:31.753', 'editor', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImVkaXRvciIsInJvbGVfa2V5IjoiZWRpdG9yIiwiaXNfYWRtaW4iOmZhbHNlLCJpc3MiOiJnaW4td2ViLWFkbWluIiwiZXhwIjoxNzcwMzA3MjAwLCJpYXQiOjE3Njk3NDI1NzF9.YsRC6J1VXClDI_h2_-u3T1bCfh-duJasVkIE7aBiEcc', 0, NULL, NULL);
INSERT INTO `gin_auth` VALUES (3, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor2', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '3', 0, NULL, NULL);
INSERT INTO `gin_auth` VALUES (4, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor3', NULL, NULL, NULL, 1, 'a203793c127cf17027b2cadbbff95355', '4', 0, NULL, NULL);
INSERT INTO `gin_auth` VALUES (5, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor4', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '5', 0, NULL, NULL);
INSERT INTO `gin_auth` VALUES (6, '2024-05-27 10:36:22.000', '2024-05-27 10:36:22.000', NULL, 2, 0, '2024-05-21 16:19:12.097', 'editor5', NULL, NULL, NULL, 2, 'a203793c127cf17027b2cadbbff95355', '6', 0, NULL, NULL);

-- ----------------------------
-- Table structure for gin_dept
-- ----------------------------
DROP TABLE IF EXISTS `gin_dept`;
CREATE TABLE `gin_dept`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父部门ID',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门名称',
  `order_num` bigint NULL DEFAULT NULL COMMENT '排序',
  `leader` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1启用0停用)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_dept
-- ----------------------------
INSERT INTO `gin_dept` VALUES (1, 0, '总部', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` VALUES (2, 1, '郑州分公司', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` VALUES (3, 1, '开封分公司', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` VALUES (4, 0, '顶级分类', 0, '', '', '', 1, '');
INSERT INTO `gin_dept` VALUES (5, 4, '北京公司', 0, '', '', '', 1, '');

-- ----------------------------
-- Table structure for gin_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `gin_dict_data`;
CREATE TABLE `gin_dict_data`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字典类型',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字典标签',
  `value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字典键值',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序值(越小越靠前)',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'CSS类名',
  `list_class` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '回显样式',
  `is_default` int NOT NULL DEFAULT 0 COMMENT '是否默认(1是0否)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_dict_data_type`(`dict_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_dict_data
-- ----------------------------
INSERT INTO `gin_dict_data` VALUES (1, 'user', 'type', 'type1', 1, 0, '', 'primary', 0, '', '2026-04-16 10:40:17.696', '2026-04-17 09:06:40.663');
INSERT INTO `gin_dict_data` VALUES (2, 'user', 'type2', 'type2', 1, 0, '', 'danger', 0, '', '2026-04-16 10:40:27.629', '2026-04-17 09:06:46.849');
INSERT INTO `gin_dict_data` VALUES (3, 'test2', 'type', '1', 1, 0, '', 'default', 0, '', '2026-04-17 09:29:44.884', '2026-04-17 09:29:44.884');
INSERT INTO `gin_dict_data` VALUES (4, 'user', 'type', 't3', 1, 0, '', 'default', 0, 't3备注', '2026-04-17 11:34:40.886', '2026-04-17 11:37:43.358');

-- ----------------------------
-- Table structure for gin_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `gin_dict_type`;
CREATE TABLE `gin_dict_type`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字典名称',
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字典类型',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序值(越小越靠前)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_dict_type`(`type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_dict_type
-- ----------------------------
INSERT INTO `gin_dict_type` VALUES (1, '测试', 'user', 1, 0, 'user的字典', '2026-04-16 09:32:24.990', '2026-04-17 11:36:19.637');
INSERT INTO `gin_dict_type` VALUES (2, '测试2', 'test2', 1, 0, '', '2026-04-16 10:41:05.127', '2026-04-16 10:41:05.127');

-- ----------------------------
-- Table structure for gin_jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `gin_jwt_blacklist`;
CREATE TABLE `gin_jwt_blacklist`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blog_jwt_blacklist_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 105 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_jwt_blacklist
-- ----------------------------

-- ----------------------------
-- Table structure for gin_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_menu`;
CREATE TABLE `gin_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父菜单ID',
  `menu_type` bigint NULL DEFAULT NULL COMMENT '菜单类型（0:目录 1:菜单 2:按钮）',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单标题',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单名称（唯一值）',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '组件路径',
  `rank` bigint NULL DEFAULT NULL COMMENT '排序序号',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '图标',
  `extra_icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '额外图标',
  `enter_transition` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '进入动画',
  `leave_transition` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '离开动画',
  `active_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '激活路径',
  `auths` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限标识,逗号分隔',
  `frame_src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '内嵌 iframe 地址',
  `frame_loading` tinyint(1) NULL DEFAULT NULL COMMENT '是否显示 iframe 加载动画',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '是否缓存组件',
  `hidden_tag` tinyint(1) NULL DEFAULT NULL COMMENT '是否隐藏标签',
  `fixed_tag` tinyint(1) NULL DEFAULT NULL COMMENT '是否固定标签',
  `show_link` tinyint(1) NULL DEFAULT NULL COMMENT '是否显示链接',
  `show_parent` tinyint(1) NULL DEFAULT NULL COMMENT '是否显示父级菜单',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_menu
-- ----------------------------
INSERT INTO `gin_menu` VALUES (1, 0, 0, '系统管理', '', '/system', NULL, 1, NULL, 'ri:settings-3-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (2, 1, 1, '用户管理', 'SystemUser', '/system/user/index', 'system/user/index', 0, NULL, 'ri:admin-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (3, 1, 1, '角色管理', 'SystemRole', '/system/role/index', 'system/role/index', 0, NULL, 'ri:admin-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (4, 1, 1, '菜单管理', 'SystemMenu', '/system/menu/index', 'system/menu/index', 0, NULL, 'ep:menu', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (5, 1, 1, 'API接口管理', 'SystemAPI', '/system/api/index', 'system/api/index', 0, NULL, 'ep:list', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (6, 1, 1, '部门管理', 'SystemDept', '/system/dept/index', 'system/dept/index', 0, NULL, 'ri:git-branch-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (7, 0, 0, '系统监控', '', '/monitor', NULL, 2, NULL, 'ep:monitor', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (8, 7, 1, '在线用户', 'OnlineUser', '/monitor/online-user', 'monitor/online/index', 0, NULL, 'ri:user-voice-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (9, 7, 1, '登录日志', 'LoginLog', '/monitor/login-logs', 'monitor/logs/login/index', 0, NULL, 'ri:window-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (10, 7, 1, '操作日志', 'OperationLog', '/monitor/operation-logs', 'monitor/logs/operation/index', 0, NULL, 'ri:history-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (11, 7, 1, '系统日志', 'SystemLog', '/monitor/system-logs', 'monitor/logs/system/index', 0, NULL, 'ri:file-search-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (12, 0, 0, '权限管理', '', '/permission', NULL, 3, NULL, 'ep:lollipop', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (13, 12, 1, '页面权限', 'PermissionPage', '/permission/page/index', 'permission/page/index', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (14, 12, 0, '按钮权限', '', '/permission/button', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (15, 14, 1, '路由返回按钮权限', 'PermissionButtonRouter', '/permission/button/router', 'permission/button/index', 0, NULL, '', NULL, NULL, NULL, '', 'permission:btn:add,permission:btn:edit,permission:btn:delete', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (16, 14, 1, '登录接口返回按钮权限', 'PermissionButtonLogin', '/permission/button/login', 'permission/button/perms', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (17, 0, 0, '外部页面', '', '/iframe', NULL, 4, NULL, 'ri:links-fill', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (18, 17, 0, '文档内嵌', '', '/iframe/embedded', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (19, 18, 1, '调色板', 'FrameColorHunt', '/iframe/colorhunt', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://colorhunt.co/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (20, 18, 1, '渐变色', 'FrameUiGradients', '/iframe/uigradients', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://uigradients.com/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (21, 18, 1, 'element-plus', 'FrameEp', '/iframe/ep', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://element-plus.org/zh-CN/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (22, 18, 1, 'tailwindcss', 'FrameTailwindcss', '/iframe/tailwindcss', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://tailwindcss.com/docs/installation', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (23, 18, 1, 'vue3', 'FrameVue', '/iframe/vue3', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://cn.vuejs.org/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (24, 18, 1, 'vite', 'FrameVite', '/iframe/vite', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://cn.vitejs.dev/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (25, 18, 1, 'pinia', 'FramePinia', '/iframe/pinia', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://pinia.vuejs.org/zh/index.html', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (26, 18, 1, 'vue-router', 'FrameRouter', '/iframe/vue-router', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', 'https://router.vuejs.org/zh/', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (27, 17, 0, '文档外链', '', '/iframe/external', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (28, 27, 1, 'vue-pure-admin', 'https://pure-admin.cn/', '/external', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (29, 27, 1, 'pure-admin-utils', 'https://pure-admin-utils.netlify.app/', '/pureUtilsLink', NULL, 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (30, 0, 0, '标签页操作', '', '/tabs', NULL, 5, NULL, 'ri:bookmark-2-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (31, 30, 1, '标签页操作', 'Tabs', '/tabs/index', 'tabs/index', 0, NULL, '', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (32, 30, 1, '', 'TabQueryDetail', '/tabs/query-detail', NULL, 0, NULL, '', NULL, NULL, NULL, '/tabs/index', '', '', NULL, 0, NULL, NULL, 0, NULL);
INSERT INTO `gin_menu` VALUES (33, 30, 1, '', 'TabParamsDetail', '/tabs/params-detail/:id', 'params-detail', 0, NULL, '', NULL, NULL, NULL, '/tabs/index', '', '', NULL, 0, NULL, NULL, 0, NULL);
INSERT INTO `gin_menu` VALUES (34, 0, 0, '官网管理', '', '/site', NULL, 6, NULL, 'ri:global-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (35, 34, 1, '内容管理', 'SiteContent', '/site/content/index', 'site/content/index', 0, NULL, 'ri:article-line', NULL, NULL, NULL, '', '', '', NULL, 1, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (36, 34, 1, '类别管理', 'SiteCategory', '/site/category/index', 'site/category/index', 0, NULL, 'ri:price-tag-3-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (37, 34, 1, '标签管理', 'SiteTag', '/site/tag/index', 'site/tag/index', 0, NULL, 'ri:price-tag-2-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (38, 1, 1, '字典管理', 'SystemDict', '/system/dict/index', 'system/dict/index', 0, NULL, 'ep:files', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);
INSERT INTO `gin_menu` VALUES (39, 7, 1, '服务器监控', 'ServerMonitor', '/monitor/server', 'monitor/server/index', 0, NULL, 'ri:line-chart-line', NULL, NULL, NULL, '', '', '', NULL, 0, NULL, NULL, 1, NULL);

-- ----------------------------
-- Table structure for gin_report
-- ----------------------------
DROP TABLE IF EXISTS `gin_report`;
CREATE TABLE `gin_report`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `activity_id` bigint NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ip` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_report
-- ----------------------------

-- ----------------------------
-- Table structure for gin_role
-- ----------------------------
DROP TABLE IF EXISTS `gin_role`;
CREATE TABLE `gin_role`  (
  `role_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_admin` int NOT NULL DEFAULT 0,
  `status` int NOT NULL DEFAULT 0,
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_sort` int NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE,
  UNIQUE INDEX `uni_gin_role_role_key`(`role_key` ASC) USING BTREE,
  CONSTRAINT `fk_gin_auth_role` FOREIGN KEY (`role_id`) REFERENCES `gin_auth` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_role
-- ----------------------------
INSERT INTO `gin_role` VALUES (1, NULL, '2024-05-21 11:03:04.239', NULL, '超管2233', 1, 0, 'admin', NULL, '超级管理员2');
INSERT INTO `gin_role` VALUES (2, '2021-10-27 16:49:28.000', '2024-05-21 11:03:06.172', NULL, '编辑角色', 0, 0, 'editor', 0, '1111111111');

-- ----------------------------
-- Table structure for gin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `gin_role_menu`;
CREATE TABLE `gin_role_menu`  (
  `role_id` bigint UNSIGNED NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_role_menu
-- ----------------------------
INSERT INTO `gin_role_menu` VALUES (1, 1);
INSERT INTO `gin_role_menu` VALUES (1, 2);
INSERT INTO `gin_role_menu` VALUES (1, 3);
INSERT INTO `gin_role_menu` VALUES (1, 4);
INSERT INTO `gin_role_menu` VALUES (1, 5);
INSERT INTO `gin_role_menu` VALUES (1, 6);
INSERT INTO `gin_role_menu` VALUES (1, 7);
INSERT INTO `gin_role_menu` VALUES (1, 8);
INSERT INTO `gin_role_menu` VALUES (1, 9);
INSERT INTO `gin_role_menu` VALUES (1, 10);
INSERT INTO `gin_role_menu` VALUES (1, 11);
INSERT INTO `gin_role_menu` VALUES (1, 12);
INSERT INTO `gin_role_menu` VALUES (1, 13);
INSERT INTO `gin_role_menu` VALUES (1, 14);
INSERT INTO `gin_role_menu` VALUES (1, 15);
INSERT INTO `gin_role_menu` VALUES (1, 16);
INSERT INTO `gin_role_menu` VALUES (1, 17);
INSERT INTO `gin_role_menu` VALUES (1, 18);
INSERT INTO `gin_role_menu` VALUES (1, 19);
INSERT INTO `gin_role_menu` VALUES (1, 20);
INSERT INTO `gin_role_menu` VALUES (1, 21);
INSERT INTO `gin_role_menu` VALUES (1, 22);
INSERT INTO `gin_role_menu` VALUES (1, 23);
INSERT INTO `gin_role_menu` VALUES (1, 24);
INSERT INTO `gin_role_menu` VALUES (1, 25);
INSERT INTO `gin_role_menu` VALUES (1, 26);
INSERT INTO `gin_role_menu` VALUES (1, 27);
INSERT INTO `gin_role_menu` VALUES (1, 28);
INSERT INTO `gin_role_menu` VALUES (1, 29);
INSERT INTO `gin_role_menu` VALUES (1, 30);
INSERT INTO `gin_role_menu` VALUES (1, 31);
INSERT INTO `gin_role_menu` VALUES (1, 32);
INSERT INTO `gin_role_menu` VALUES (1, 33);
INSERT INTO `gin_role_menu` VALUES (1, 34);
INSERT INTO `gin_role_menu` VALUES (1, 35);
INSERT INTO `gin_role_menu` VALUES (1, 36);
INSERT INTO `gin_role_menu` VALUES (1, 37);
INSERT INTO `gin_role_menu` VALUES (1, 38);
INSERT INTO `gin_role_menu` VALUES (1, 39);
INSERT INTO `gin_role_menu` VALUES (2, 12);
INSERT INTO `gin_role_menu` VALUES (2, 13);
INSERT INTO `gin_role_menu` VALUES (2, 14);
INSERT INTO `gin_role_menu` VALUES (2, 15);
INSERT INTO `gin_role_menu` VALUES (2, 16);
INSERT INTO `gin_role_menu` VALUES (2, 17);
INSERT INTO `gin_role_menu` VALUES (2, 18);
INSERT INTO `gin_role_menu` VALUES (2, 19);
INSERT INTO `gin_role_menu` VALUES (2, 20);
INSERT INTO `gin_role_menu` VALUES (2, 21);
INSERT INTO `gin_role_menu` VALUES (2, 22);
INSERT INTO `gin_role_menu` VALUES (2, 23);
INSERT INTO `gin_role_menu` VALUES (2, 24);
INSERT INTO `gin_role_menu` VALUES (2, 25);
INSERT INTO `gin_role_menu` VALUES (2, 26);
INSERT INTO `gin_role_menu` VALUES (2, 27);
INSERT INTO `gin_role_menu` VALUES (2, 28);
INSERT INTO `gin_role_menu` VALUES (2, 29);
INSERT INTO `gin_role_menu` VALUES (2, 30);
INSERT INTO `gin_role_menu` VALUES (2, 31);
INSERT INTO `gin_role_menu` VALUES (2, 32);
INSERT INTO `gin_role_menu` VALUES (2, 33);

-- ----------------------------
-- Table structure for gin_site_category
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_category`;
CREATE TABLE `gin_site_category`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `slug` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类标识',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '分类描述',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_site_category_slug`(`slug` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_site_category
-- ----------------------------
INSERT INTO `gin_site_category` VALUES (1, '分类news', 'news', '', 1, 0, '2026-04-05 18:03:47.701', '2026-04-10 14:59:29.769');

-- ----------------------------
-- Table structure for gin_site_content
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content`;
CREATE TABLE `gin_site_content`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容标题',
  `slug` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '页面标识',
  `summary` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '摘要',
  `cover` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '封面图',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '正文内容',
  `seo_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SEO关键词',
  `seo_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SEO描述',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1发布0草稿)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `published_at` datetime(3) NULL DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_site_content_slug`(`slug` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_site_content
-- ----------------------------
INSERT INTO `gin_site_content` VALUES (1, '测试标题111', '1', '**d是否🎉️ 第三方分多少 &lt;script&gt;alert(222333);&lt;/script&gt;', '', '**d是否🎉️ 第三方分多少234**\n\ndsfsdf第三方士大夫水电费水电费水电费水电费水电费水电费水电费胜多负少的给对方回复过几年更好看❤️ 🎉️ 😕\n\n都是1111\n\n```javascript\n&lt;script&gt;alert(222333);&lt;/script&gt;\n```', '', '', 1, 1, '2026-04-05 17:41:45.080', '2026-04-10 15:01:40.695', '2026-04-05 18:22:55.089');
INSERT INTO `gin_site_content` VALUES (2, '测试22', '2', '第三方1', '', '第三方222', '', '', 1, 0, '2026-04-10 10:49:30.702', '2026-04-10 15:01:39.508', '2026-04-10 10:49:30.702');

-- ----------------------------
-- Table structure for gin_site_content_category
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content_category`;
CREATE TABLE `gin_site_content_category`  (
  `content_id` bigint NOT NULL,
  `category_id` bigint NOT NULL,
  PRIMARY KEY (`content_id`, `category_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_site_content_category
-- ----------------------------
INSERT INTO `gin_site_content_category` VALUES (1, 1);

-- ----------------------------
-- Table structure for gin_site_content_tag
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_content_tag`;
CREATE TABLE `gin_site_content_tag`  (
  `content_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  PRIMARY KEY (`content_id`, `tag_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_site_content_tag
-- ----------------------------
INSERT INTO `gin_site_content_tag` VALUES (1, 1);

-- ----------------------------
-- Table structure for gin_site_tag
-- ----------------------------
DROP TABLE IF EXISTS `gin_site_tag`;
CREATE TABLE `gin_site_tag`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签名',
  `slug` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签标识',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态(1启用0停用)',
  `sort` bigint NOT NULL DEFAULT 0 COMMENT '排序值(越小越靠前)',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_site_tag_slug`(`slug` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gin_site_tag
-- ----------------------------
INSERT INTO `gin_site_tag` VALUES (1, '标签', 'biaoqian', 1, 0, '2026-04-05 19:00:12.457', '2026-04-05 19:00:12.457');

SET FOREIGN_KEY_CHECKS = 1;
