-- Mock data converted from /home/fhy/myWeb/web-admin-frontend/mock/asyncRoutes.ts
-- Execute this file in your database to import the menu structure.

-- Clear existing menu data to avoid conflicts
DELETE FROM `gin_menu`;
ALTER TABLE `gin_menu` AUTO_INCREMENT = 1;
-- Rebuild admin role menu relation (role_id=1) to avoid missing menus after import
DELETE FROM `gin_role_menu` WHERE `role_id` = 1;

-- Insert new menu data
INSERT INTO `gin_menu` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `icon`, `auths`, `frame_src`, `keep_alive`, `show_link`, `active_path`) VALUES
(1, 0, 0, 'menus.pureSysManagement', '', '/system', NULL, 1, 'ri:settings-3-line', '', '', 0, 1, ''),
(2, 1, 1, 'menus.pureUser', 'SystemUser', '/system/user/index', 'system/user/index', 0, 'ri:admin-line', '', '', 0, 1, ''),
(3, 1, 1, 'menus.pureRole', 'SystemRole', '/system/role/index', 'system/role/index', 0, 'ri:admin-fill', '', '', 0, 1, ''),
(4, 1, 1, 'menus.pureSystemMenu', 'SystemMenu', '/system/menu/index', 'system/menu/index', 0, 'ep:menu', '', '', 0, 1, ''),
(5, 1, 1, 'menus.pureSystemApi', 'SystemAPI', '/system/api/index', 'system/api/index', 0, 'ep:list', '', '', 0, 1, ''),
(6, 1, 1, 'menus.pureDept', 'SystemDept', '/system/dept/index', 'system/dept/index', 0, 'ri:git-branch-line', '', '', 0, 1, ''),
(7, 0, 0, 'menus.pureSysMonitor', '', '/monitor', NULL, 2, 'ep:monitor', '', '', 0, 1, ''),
(8, 7, 1, 'menus.pureOnlineUser', 'OnlineUser', '/monitor/online-user', 'monitor/online/index', 0, 'ri:user-voice-line', '', '', 0, 1, ''),
(9, 7, 1, 'menus.pureLoginLog', 'LoginLog', '/monitor/login-logs', 'monitor/logs/login/index', 0, 'ri:window-line', '', '', 0, 1, ''),
(10, 7, 1, 'menus.pureOperationLog', 'OperationLog', '/monitor/operation-logs', 'monitor/logs/operation/index', 0, 'ri:history-fill', '', '', 0, 1, ''),
(11, 7, 1, 'menus.pureSystemLog', 'SystemLog', '/monitor/system-logs', 'monitor/logs/system/index', 0, 'ri:file-search-line', '', '', 0, 1, ''),
(12, 0, 0, 'menus.purePermission', '', '/permission', NULL, 3, 'ep:lollipop', '', '', 0, 1, ''),
(13, 12, 1, 'menus.purePermissionPage', 'PermissionPage', '/permission/page/index', 'permission/page/index', 0, '', '', '', 0, 1, ''),
(14, 12, 0, 'menus.purePermissionButton', '', '/permission/button', NULL, 0, '', '', '', 0, 1, ''),
(15, 14, 1, 'menus.purePermissionButtonRouter', 'PermissionButtonRouter', '/permission/button/router', 'permission/button/index', 0, '', 'permission:btn:add,permission:btn:edit,permission:btn:delete', '', 0, 1, ''),
(16, 14, 1, 'menus.purePermissionButtonLogin', 'PermissionButtonLogin', '/permission/button/login', 'permission/button/perms', 0, '', '', '', 0, 1, ''),
(17, 0, 0, 'menus.pureExternalPage', '', '/iframe', NULL, 4, 'ri:links-fill', '', '', 0, 1, ''),
(18, 17, 0, 'menus.pureEmbeddedDoc', '', '/iframe/embedded', NULL, 0, '', '', '', 0, 1, ''),
(19, 18, 1, 'menus.pureColorHuntDoc', 'FrameColorHunt', '/iframe/colorhunt', NULL, 0, '', '', 'https://colorhunt.co/', 1, 1, ''),
(20, 18, 1, 'menus.pureUiGradients', 'FrameUiGradients', '/iframe/uigradients', NULL, 0, '', '', 'https://uigradients.com/', 1, 1, ''),
(21, 18, 1, 'menus.pureEpDoc', 'FrameEp', '/iframe/ep', NULL, 0, '', '', 'https://element-plus.org/zh-CN/', 1, 1, ''),
(22, 18, 1, 'menus.pureTailwindcssDoc', 'FrameTailwindcss', '/iframe/tailwindcss', NULL, 0, '', '', 'https://tailwindcss.com/docs/installation', 1, 1, ''),
(23, 18, 1, 'menus.pureVueDoc', 'FrameVue', '/iframe/vue3', NULL, 0, '', '', 'https://cn.vuejs.org/', 1, 1, ''),
(24, 18, 1, 'menus.pureViteDoc', 'FrameVite', '/iframe/vite', NULL, 0, '', '', 'https://cn.vitejs.dev/', 1, 1, ''),
(25, 18, 1, 'menus.purePiniaDoc', 'FramePinia', '/iframe/pinia', NULL, 0, '', '', 'https://pinia.vuejs.org/zh/index.html', 1, 1, ''),
(26, 18, 1, 'menus.pureRouterDoc', 'FrameRouter', '/iframe/vue-router', NULL, 0, '', '', 'https://router.vuejs.org/zh/', 1, 1, ''),
(27, 17, 0, 'menus.pureExternalDoc', '', '/iframe/external', NULL, 0, '', '', '', 0, 1, ''),
(28, 27, 1, 'menus.pureExternalLink', 'https://pure-admin.cn/', '/external', NULL, 0, '', '', '', 0, 1, ''),
(29, 27, 1, 'menus.pureUtilsLink', 'https://pure-admin-utils.netlify.app/', '/pureUtilsLink', NULL, 0, '', '', '', 0, 1, ''),
(30, 0, 0, 'menus.pureTabs', '', '/tabs', NULL, 5, 'ri:bookmark-2-line', '', '', 0, 1, ''),
(31, 30, 1, 'menus.pureTabs', 'Tabs', '/tabs/index', 'tabs/index', 0, '', '', '', 0, 1, ''),
(32, 30, 1, '', 'TabQueryDetail', '/tabs/query-detail', NULL, 0, '', '', '', 0, 0, '/tabs/index'),
(33, 30, 1, '', 'TabParamsDetail', '/tabs/params-detail/:id', 'params-detail', 0, '', '', '', 0, 0, '/tabs/index'),
(34, 0, 0, '官网管理', '', '/site', NULL, 6, 'ri:global-line', '', '', 0, 1, ''),
(35, 34, 1, '内容管理', 'SiteContent', '/site/content/index', 'site/content/index', 0, 'ri:article-line', '', '', 1, 1, ''),
(36, 34, 1, '类别管理', 'SiteCategory', '/site/category/index', 'site/category/index', 0, 'ri:price-tag-3-line', '', '', 0, 1, ''),
(37, 34, 1, '标签管理', 'SiteTag', '/site/tag/index', 'site/tag/index', 0, 'ri:price-tag-2-line', '', '', 0, 1, '');

-- Grant all imported menus to admin role
INSERT INTO `gin_role_menu` (`role_id`, `menu_id`)
SELECT 1, `id` FROM `gin_menu`;
