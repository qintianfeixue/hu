/*
 Navicat Premium Data Transfer

 Source Server         : 我的电脑
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : www_yy_com

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 31/03/2021 18:04:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 204 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dept
-- ----------------------------
INSERT INTO `dept` VALUES (100, 0, '0', '若依科技', 0, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:54:39');
INSERT INTO `dept` VALUES (101, 100, '0,100', '深圳总公司', 1, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:54:45');
INSERT INTO `dept` VALUES (102, 100, '0,100', '长沙分公司', 2, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:55:08');
INSERT INTO `dept` VALUES (103, 101, '0,100,101', '研发部门', 1, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:54:49');
INSERT INTO `dept` VALUES (104, 101, '0,100,101', '市场部门', 2, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:54:53');
INSERT INTO `dept` VALUES (105, 101, '0,100,101', '测试部门', 3, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:54:57');
INSERT INTO `dept` VALUES (106, 101, '0,100,101', '财务部门', 4, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:55:00');
INSERT INTO `dept` VALUES (107, 101, '0,100,101', '运维部门', 5, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:55:04');
INSERT INTO `dept` VALUES (108, 102, '0,100,102', '市场部门', 1, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:55:12');
INSERT INTO `dept` VALUES (109, 102, '0,100,102', '财务部门', 2, '若依', '15888888888', 'ry@qq.com', '1', '0', 'admin', '2018-03-16 11:33:00', '', '2020-06-30 14:55:18');
INSERT INTO `dept` VALUES (200, 100, '', '大数据', 1, '小刘', '18888888888', 'liou@qq.com', '1', '0', '', NULL, '', NULL);
INSERT INTO `dept` VALUES (201, 100, '', '开发', 1, '老李', '18888888888', 'li@qq.com', '1', '0', 'admin', '2020-06-19 10:14:06', '', '2020-06-30 14:55:22');
INSERT INTO `dept` VALUES (202, 108, '', '外勤', 1, '小a', '18888888888', 'aa@qq.com', '1', '0', '', '2020-06-19 14:59:44', '', '2020-06-30 14:55:15');
INSERT INTO `dept` VALUES (203, 108, '', '行政', 0, 'aa', '18888888888', 'aa@qq.com', '0', '0', '', '2020-07-03 11:44:57', '', NULL);

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件（没用）',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `createtime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updatetime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重 越低越靠前',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0停用 1正常',
  `always_show` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态（没用）0隐藏 1显示',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块（没用）',
  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID（没用）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 117 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, 0, 'system/config', '系统配置', 'system', '', 'Admin tips', 0, 1497429920, 1497430320, 0, 1, 1, 'config', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (2, 0, 'system/auth', '权限管理', 'peoples', '', '', 0, 1497429920, 1592380524, 0, 1, 1, 'system/auth', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (3, 0, 'system/monitor', '系统监控', 'monitor', '', '', 0, 1497429920, 1592380560, 0, 1, 1, 'monitor', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (4, 0, 'system/cms', 'cms管理', 'form', '', '', 0, 1497429920, 1592380573, 0, 1, 1, 'cms', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (5, 1, 'system/config/dict/list', '字典管理', 'dict', '', '', 1, 1497429920, 1592380469, 0, 1, 1, 'dict/list', 'system/config/dict/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (6, 1, 'system/config/params/list', '参数管理', 'date-range', '', '', 1, 1497429920, 1592380510, 0, 1, 1, 'params/list', 'system/config/params/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (8, 2, 'system/auth/menuList', '菜单管理', 'nested', '', '', 1, 1497429920, 1592380548, 0, 1, 1, 'menuList', 'system/auth/menuList', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (21, 5, 'system/config/dict/delete', '删除字典', '', '', '', 2, 1592363629, 1592363629, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (16, 8, 'system/auth/addMenu', '添加菜单', '', '', '', 2, 1592357208, 1600396699, 0, 1, 0, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (17, 8, 'system/auth/editMenu', '修改菜单', '', '', '', 2, 1592357231, 1592357274, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (18, 8, 'system/auth/deleteMenu', '删除菜单', '', '', '', 2, 1592357268, 1592357268, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (19, 5, 'system/config/dict/add', '添加字典', '', '', '', 2, 1592363585, 1592363585, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (20, 5, 'system/config/dict/edit', '修改字典', '', '', '', 2, 1592363604, 1592363604, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (22, 5, 'system/config/dict/dataList', '字典数据管理', '', '', '', 2, 1592363790, 1592365455, 0, 1, 1, 'dataList', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (23, 5, 'system/config/dict/dataAdd', '添加字典数据', '', '', '', 2, 1592380398, 1593411393, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (24, 5, 'system/config/dict/dataEdit', '修改字典数据', '', '', '', 2, 1592380417, 1593411398, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (25, 5, 'system/config/dict/dataDelete', '删除字典数据', '', '', '', 2, 1592380449, 1593411403, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (26, 2, 'system/auth/roleList', '角色管理', 'logininfor', '', '', 1, 1592385260, 1592385291, 0, 1, 1, 'role', 'system/auth/roleList', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (27, 26, 'system/auth/addRole', '添加角色', '', '', '', 2, 1592389821, 1592389821, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (28, 2, 'system/auth/dept', '部门管理', 'peoples', '', '', 1, 1592449845, 1592449845, 0, 1, 1, 'dept', 'system/auth/dept', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (29, 26, 'system/auth/editRole', '修改角色', '', '', '', 2, 1592469153, 1592469153, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (30, 26, 'system/auth/statusSetRole', '设置角色状态', '', '', '', 2, 1592469201, 1592469201, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (31, 26, 'system/auth/deleteRole', '删除角色', '', '', '', 2, 1592469234, 1592469234, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (32, 4, 'system/cms/menu/list', '栏目管理', 'tree-table', '', '', 1, 1592469318, 1592469318, 0, 1, 1, 'menu', 'system/cms/menu/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (33, 2, 'system/auth/post', '岗位管理', 'tab', '', '', 1, 1592558968, 1592558968, 0, 1, 1, 'post', 'system/auth/post', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (34, 3, 'system/monitor/online/list', '在线用户', 'cascader', '', '', 1, 1593328511, 1593328511, 0, 1, 1, 'online', 'system/monitor/online/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (35, 4, 'system/cms/news/list', '文章管理', 'log', '', '', 1, 1593331969, 1593331991, 0, 1, 1, 'news', 'system/cms/news/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (36, 3, 'system/monitor/job', '定时任务', 'clipboard', '', '', 1, 1593332305, 1593332434, 0, 1, 1, 'job', 'system/monitor/job', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (37, 3, 'system/monitor/server', '服务监控', 'dict', '', '', 1, 1593419394, 1593419394, 0, 1, 1, 'server', 'system/monitor/server', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (38, 3, 'system/monitor/logininfor', '登录日志', 'chart', '', '', 1, 1593423378, 1593423505, 0, 1, 1, 'logininfor', 'system/monitor/logininfor', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (39, 3, 'system/monitor/operlog', '操作日志', 'dashboard', '', '', 1, 1593485097, 1593485097, 0, 1, 1, 'operlog', 'system/monitor/operlog', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (40, 2, 'system/auth/userList', '用户管理', 'user', '', '', 1, 1593572523, 1593572523, 0, 1, 1, 'user', 'system/auth/userList', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (41, 6, 'system/config/params/add', '添加参数', '', '', '', 2, 1593684331, 1593684331, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (42, 6, 'system/config/params/edit', '修改参数', '', '', '', 2, 1593684351, 1593684351, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (43, 6, 'system/config/params/delete', '删除参数', '', '', '', 2, 1593684470, 1593684470, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (44, 28, 'system/dept/addDept', '添加部门', '', '', '', 2, 1593738070, 1593738070, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (45, 28, 'system/dept/editDept', '修改部门', '', '', '', 2, 1593738097, 1593738097, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (46, 28, 'system/dept/delDept', '删除部门', '', '', '', 2, 1593738125, 1593738125, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (47, 33, 'system/post/add', '添加岗位', '', '', '', 2, 1593738444, 1593738444, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (48, 33, 'system/post/edit', '修改岗位', '', '', '', 2, 1593738567, 1593738567, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (49, 33, 'system/post/delete', '删除岗位', '', '', '', 2, 1593738590, 1593738590, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (50, 40, 'system/auth/addUser', '添加用户', '', '', '', 2, 1593738798, 1593738798, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (51, 40, 'system/auth/editUser', '修改用户', '', '', '', 2, 1593738950, 1593738950, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (52, 40, 'system/auth/resetUserPwd', '密码重置', '', '', '', 2, 1593739001, 1593739001, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (53, 40, 'system/auth/changeUserStatus', '状态设置', '', '', '', 2, 1593739079, 1593739079, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (54, 40, 'system/auth/deleteAdmin', '删除用户', '', '', '', 2, 1593739113, 1593739113, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (55, 34, 'system/monitor/online/forceLogout', '强制退出', '', '', '', 2, 1593739201, 1593739201, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (56, 36, 'system/monitor/job/add', '添加任务', '', '', '', 2, 1593740041, 1593740041, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (57, 36, 'system/monitor/job/edit', '修改任务', '', '', '', 2, 1593740062, 1593740062, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (58, 36, 'system/monitor/job/start', '开启任务', '', '', '', 2, 1593740105, 1593740105, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (59, 36, 'system/monitor/job/stop', '停止任务', '', '', '', 2, 1593740139, 1593740139, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (60, 36, 'system/monitor/job/delete', '删除任务', '', '', '', 2, 1593740165, 1593740165, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (61, 38, 'system/monitor/loginlog/delete', '删除', '', '', '', 2, 1593740342, 1593740342, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (62, 38, 'system/monitor/loginlog/clear', '清空', '', '', '', 2, 1593740359, 1593740359, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (63, 39, 'system/monitor/operlog/delete', '删除', '', '', '', 2, 1593740422, 1593740422, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (64, 39, 'system/monitor/operlog/clear', '清空', '', '', '', 2, 1593740434, 1593740434, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (65, 32, 'system/cms/menu/add', '添加栏目', '', '', '', 2, 1593740504, 1593740504, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (66, 32, 'system/cms/menu/edit', '修改栏目', '', '', '', 2, 1593740521, 1593740521, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (67, 32, 'system/cms/menu/sort', '栏目排序', '', '', '', 2, 1593740549, 1593740549, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (68, 32, 'system/cms/menu/delete', '删除栏目', '', '', '', 2, 1593740568, 1593740568, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (69, 35, 'system/cms/news/add', '添加文章', '', '', '', 2, 1593740691, 1593740691, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (70, 35, 'system/cms/news/edit', '修改文章', '', '', '', 2, 1593740711, 1593740711, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (71, 35, 'system/cms/news/delete', '删除文章', '', '', '', 2, 1593740730, 1593740730, 0, 1, 1, '', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (72, 0, 'system/model', '模型管理', 'table', '', '', 0, 1593742999, 1593742999, 0, 1, 1, 'model', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (73, 72, 'system/model/category/list', '模型分类', 'tree-table', '', '', 1, 1593743065, 1593743065, 0, 1, 1, 'category', 'system/model/category/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (74, 72, 'system/model/info/list', '模型列表', 'list', '', '', 1, 1593743131, 1594781057, 0, 1, 1, 'list', 'system/model/info/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (75, 0, 'system/tools', '系统工具', 'server', '', '', 0, 1594016328, 1594016328, 0, 1, 1, 'system/tools', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (76, 75, 'system/tools/build', '表单构建', 'build', '', '', 1, 1594016392, 1594016808, 0, 1, 1, 'build', 'system/tools/build', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (77, 75, 'system/tools/gen', '代码生成', 'code', '', '', 1, 1594016637, 1594016637, 0, 1, 1, 'gen', 'system/tools/gen', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (78, 0, 'system/plug', '扩展管理', 'logininfor', '', '', 0, 1594169636, 1605862005, 0, 1, 1, 'system/plug', '', 1, 'sys_admin', 0);
INSERT INTO `permission` VALUES (79, 78, 'system/plug/ad', '广告管理', 'color', '', '', 0, 1594169691, 1595410089, 0, 1, 1, 'adManage', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (80, 79, 'system/plug/ad/type/list', '广告位管理', 'nested', '', '', 1, 1594169783, 1595405904, 0, 1, 1, 'adtype', 'system/plug/ad/type/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (81, 79, 'system/plug/ad/info/list', '广告列表', 'list', '', '', 1, 1594169949, 1596418803, 0, 1, 1, 'adlist', 'system/plug/ad/info/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (82, 75, 'system/tools/api', '系统接口', 'guide', '', '', 1, 1594951684, 1594951684, 0, 1, 1, 'api', 'system/tools/api', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (83, 78, 'system/plug/link', '友情链接', 'cascader', '', '', 0, 1595381634, 1595403122, 0, 1, 1, 'link', '', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (84, 83, 'system/plug/link/type/list', '分类管理', 'component', '', '', 1, 1595381717, 1595381717, 0, 1, 1, 'type', 'system/plug/link/type/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (85, 83, 'system/plug/link/info/list', '链接管理', 'list', '', '', 1, 1595381754, 1595381754, 0, 1, 1, 'info', 'system/plug/link/info/list', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (86, 1, 'system/config/webSet', '站点设置', 'system', '', '', 1, 1596420340, 1596420340, 0, 1, 1, '/webSet', 'system/config/webSet', 0, 'sys_admin', 0);
INSERT INTO `permission` VALUES (106, 4, 'cms', 'cms前端展示', 'education', '', '', 1, 1604040725, 1604040798, 0, 1, 1, 'http://localhost:8200/cms', 'system/cms/news/list', 1, '', 0);
INSERT INTO `permission` VALUES (107, 0, 'system/wf', '流程管理', 'cascader', '', '', 0, 1606967359, 1606967359, 0, 1, 1, 'system/wf', '', 0, '', 0);
INSERT INTO `permission` VALUES (108, 107, 'system/wf/flow/list', '工作流列表', 'component', '', '', 1, 1606967458, 1606967516, 0, 1, 1, 'flow/list', 'system/wf/flow/list', 0, '', 0);
INSERT INTO `permission` VALUES (109, 108, 'system/wf/flow/design', '设计流程', '', '', '', 2, 1606967544, 1607074912, 0, 1, 1, '', '', 0, '', 0);
INSERT INTO `permission` VALUES (110, 107, 'system/wf/flow/news', '测试业务', 'job', '', '', 1, 1606967599, 1606967599, 0, 1, 1, 'flow/news', 'system/wf/flow/news', 0, '', 0);
INSERT INTO `permission` VALUES (111, 107, 'system/wf/flow/monitoring', '流程监控', 'eye-open', '', '', 1, 1606967670, 1606967680, 0, 1, 1, 'monitoring', 'system/wf/flow/monitoring', 0, '', 0);
INSERT INTO `permission` VALUES (112, 0, 'system/plugin/blog', '简单博客管理', 'education', '', '', 0, 1607400193, 1607400193, 0, 1, 1, 'blog', '', 0, '', 0);
INSERT INTO `permission` VALUES (113, 112, 'system/plugin/blog/log/list', '日志管理', 'log', '', '', 1, 1607400265, 1607400265, 0, 1, 1, 'log', 'plugin/blog/log/list', 0, '', 0);
INSERT INTO `permission` VALUES (114, 112, 'system/plugin/blog/classification/list', '分类管理', 'tab', '', '', 1, 1607400305, 1607400305, 0, 1, 1, 'classification', 'plugin/blog/classification/list', 0, '', 0);
INSERT INTO `permission` VALUES (115, 112, 'system/plugin/blog/comment/list', '评论管理', 'date', '', '', 1, 1607400346, 1607400346, 0, 1, 1, 'comment', 'plugin/blog/comment/list', 0, '', 0);
INSERT INTO `permission` VALUES (116, 112, 'plugin/blog', '博客前端展示', 'eye-open', '', '', 1, 1607400394, 1607400394, 0, 1, 1, 'http://localhost:8200/plugin/blog', 'plugin/blog/log/list', 1, '', 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `list_order` float NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 1, 1582773740, 1593680183, 0, '超级管理员', '备注', 3);
INSERT INTO `role` VALUES (2, 1, 1582773740, 1606698503, 0, '普通管理员', '备注', 3);
INSERT INTO `role` VALUES (3, 1, 1582773740, 0, 0, '站点管理员', '站点管理人员', 3);
INSERT INTO `role` VALUES (4, 1, 1582773740, 0, 0, '初级管理员', '初级管理员', 3);
INSERT INTO `role` VALUES (5, 1, 1582773740, 0, 0, '高级管理员', '高级管理员', 3);
INSERT INTO `role` VALUES (8, 1, 1582773740, 1593506004, 0, '区级管理员', '', 2);

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) NULL DEFAULT NULL COMMENT 'role表，角色id',
  `permission_id` int(10) NULL DEFAULT NULL COMMENT 'permission表，权限id',
  `create_time` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和权限管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (1, 1, 1, NULL);
INSERT INTO `role_permission` VALUES (2, 1, 2, NULL);
INSERT INTO `role_permission` VALUES (3, 1, 3, NULL);
INSERT INTO `role_permission` VALUES (4, 2, 1, NULL);
INSERT INTO `role_permission` VALUES (5, 1, 5, NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `user_password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录密码',
  `create_time` int(30) NULL DEFAULT NULL COMMENT '注册时间',
  `sex` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像',
  `user_status` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户状态；0:禁用,1:正常',
  `last_login_time` int(13) NULL DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '最后登录ip',
  `dept_id` int(30) UNSIGNED NULL DEFAULT NULL COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `is_admin` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否后台管理员；1:是，0:否',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '18687455555', 'admin', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (2, 'admin1', '18687455555', 'admin1', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');
INSERT INTO `user` VALUES (3, 'admin', '18687455555', 'admin', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (4, 'admin1', '18687455555', 'admin1', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');
INSERT INTO `user` VALUES (5, 'shen', '18687455555', 'shen', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (6, 'shen1', '18687455555', 'shen1', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');
INSERT INTO `user` VALUES (7, 'kkkkkk', '18687455555', 'kkkkkk', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (8, 'kkkkkk1', '18687455555', 'kkkkkk1', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');
INSERT INTO `user` VALUES (9, 'kkkkkk2', '18687455555', 'kkkkkk2', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (10, 'kkkkkk3', '18687455555', 'kkkkkk3', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');
INSERT INTO `user` VALUES (11 'kkkkkk4', '18687455555', 'kkkkkk5', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '0');
INSERT INTO `user` VALUES (12, 'kkkkkk6', '18687455555', 'kkkkkk6', 0, '1', 'IJ1xz+Wve+ZONVMFfXJQMw==', '1', 1606698803, '1', 1, '1606698803', '1');


-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NULL DEFAULT NULL COMMENT 'user表，用户id',
  `role_id` int(10) NULL DEFAULT NULL COMMENT 'role表，角色id',
  `create_time` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` int(11) NULL DEFAULT NULL COMMENT '修改时间',
  `delete_time` int(11) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 1, 1616997466, NULL, NULL);
INSERT INTO `user_role` VALUES (2, 1, 2, 1617001601, NULL, NULL);
INSERT INTO `user_role` VALUES (3, 2, 4, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
