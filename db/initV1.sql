CREATE DATABASE admin;

USE admin;

/**
  1. 管理員帳號
 */
CREATE TABLE IF NOT EXISTS `admin_admin` (
  `id` varchar(40) NOT NULL COMMENT '管理用戶ID',
  `channel_id` varchar(40) DEFAULT NULL,
  `username` varchar(40) NOT NULL COMMENT '帳號',
  `admin_name` varchar(40) DEFAULT NULL COMMENT '管理員名稱',
  `nickname` varchar(40) DEFAULT NULL,
  `account_status` int(11) NOT NULL DEFAULT '0' COMMENT '狀態',
  `login_ip` varchar(40) DEFAULT NULL COMMENT '登入IP',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '最新一次的登入日期',
  `memo` text COMMENT '備註',
  `creator_id` varchar(40) DEFAULT NULL COMMENT '創建者id(admin id)',
  `updater_id` varchar(40) DEFAULT NULL COMMENT '更新者id',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  `two_factor_code` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_username_chId` (`username`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 插入管理員帳號数据
INSERT INTO admin.admin_admin
(id, channel_id, username, admin_name, nickname, account_status, login_ip, login_time, memo, creator_id, updater_id, update_time, create_time, two_factor_code)
VALUES('0', 'channel001', 'admin', '最高', '大老', 0, '192.168.1.1', '2024-01-12 09:58:00', 'Some memo', '999', '999', '2024-01-12 09:58:00', '2024-01-12 09:58:00', '123456');
INSERT INTO admin.admin_admin
(id, channel_id, username, admin_name, nickname, account_status, login_ip, login_time, memo, creator_id, updater_id, update_time, create_time, two_factor_code)
VALUES('1', 'channel002', 'manager', '主管', '二把手', 0, '192.168.1.1', '2024-01-12 09:58:00', 'Some memo', '999', '999', '2024-01-12 09:58:00', '2024-01-14 09:58:00', '123456');


/**
  2. 管理員對應的權限
 */
CREATE TABLE IF NOT EXISTS `admin_admin_permit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(40) NOT NULL COMMENT 'admin id',
  `permit_id` varchar(40) NOT NULL COMMENT 'permit id',
  `creator_id` varchar(40) NOT NULL COMMENT '創建者id',
  `updater_id` varchar(40) NOT NULL COMMENT '更新者id',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_adminid_permitid` (`admin_id`,`permit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



/**
  3. 管理員對應的腳色
 */
CREATE TABLE IF NOT EXISTS `admin_admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `admin_id` varchar(40) NOT NULL COMMENT '對應的admin id',
  `role_id` varchar(40) NOT NULL COMMENT '對應的腳色id',
  `creator_id` varchar(40) DEFAULT NULL COMMENT '創建者id',
  `updater_id` varchar(40) DEFAULT NULL COMMENT '更新者id',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_adminid_roleid` (`admin_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- 插入管理員對應的腳色数据
INSERT INTO admin.admin_admin_role
(id, admin_id, role_id, creator_id, updater_id, create_time, update_time)
VALUES(1, '0', '0', '999', '999', '2024-01-14 14:39:16', '2024-01-14 14:39:16');
INSERT INTO admin.admin_admin_role
(id, admin_id, role_id, creator_id, updater_id, create_time, update_time)
VALUES(2, '0', '1', '999', '999', '2024-01-12 10:14:18', '2024-01-12 10:14:18');



/**
  4. 管理員帳號密碼
 */
CREATE TABLE IF NOT EXISTS `admin_admintoken` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(40) NOT NULL COMMENT '代理id',
  `token_type` int(11) DEFAULT NULL COMMENT 'token類型',
  `token` text NOT NULL COMMENT 'token',
  `expire_time` timestamp NULL DEFAULT NULL COMMENT '過期時間',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新時間',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  `creator_id` varchar(40) DEFAULT NULL,
  `updater_id` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_admin_id` (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='提供於管理員帳號密碼';

-- 插入管理員帳號密碼数据

INSERT INTO admin.admin_admintoken
(id, admin_id, token_type, token, expire_time, update_time, create_time, creator_id, updater_id)
VALUES(0, '0', 1, '$2a$10$dU6nEGsh7QA4/caU3WZC9ODBThGZ8f/p9a3Q66LHH0UQzqKnANpva
', '2024-01-12 11:22:52', '2024-01-12 10:22:52', '2024-01-12 10:22:52', '999', '999');
INSERT INTO admin.admin_admintoken
(id, admin_id, token_type, token, expire_time, update_time, create_time, creator_id, updater_id)
VALUES(1, '1', 1, '$2a$10$dU6nEGsh7QA4/caU3WZC9ODBThGZ8f/p9a3Q66LHH0UQzqKnANpva
', '2024-01-12 11:22:52', '2024-01-12 11:22:52', '2024-01-12 11:22:52', '999', '999');



/**
  5. 權限
 */
CREATE TABLE IF NOT EXISTS `admin_permit` (
  `id` varchar(40) NOT NULL COMMENT 'id',
  `permit_key` varchar(40) NOT NULL COMMENT '給spring定位用',
  `permit_name` varchar(40) NOT NULL COMMENT '權限名稱',
  `memo` text COMMENT '備註',
  `permit_desc` text COMMENT '說明',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `creator_id` varchar(40) NOT NULL COMMENT '創建者id(admin id)',
  `updater_id` varchar(40) NOT NULL COMMENT '更新者id(admin id)',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permit_permit_key_uindex` (`permit_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 插入權限数据

INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('0', 'AddAdmin', 'AddAdmin', '添加管理員', '添加管理員', 1, '999', '999', '2024-01-12 10:31:53', '2024-01-12 10:31:53');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('1', 'AddRole', 'AddRole', '添加角色', '添加角色', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('10', 'GetAdminAllPermits', 'GetAdminAllPermits', '查詢指定管理員所有的權限(包含角色)', '查詢指定管理員所有的權限(包含角色)', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('11', 'RemoveRolePermits', 'RemoveRolePermits', '移除角色所屬的權限', '移除角色所屬的權限', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('12', 'RemoveAdminPermits', 'RemoveAdminPermits', '移除管理員額外的權限', '移除管理員額外的權限', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('13', 'RemoveAdminRoles', 'RemoveAdminRoles', '移除管理員的角色', '移除管理員的角色', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('14', 'GetRolePermits', 'GetRolePermits', '取得角色權限', '取得角色權限', 1, '999', '999', '2024-01-12 10:31:53', '2024-01-12 10:31:53');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('2', 'GetAllRoleList', 'GetAllRoleList', '查詢所有角色', '查詢所有角色', 1, '999', '999', '2024-01-12 10:32:21', '2024-01-12 10:32:21');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('3', 'GetAllPermitList', 'GetAllPermitList', '查詢所有權限', '查詢所有權限', 1, '999', '999', '2024-01-12 10:32:21', '2024-01-12 10:32:21');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('4', 'AddRolePermits', 'AddRolePermits', '角色添加權限', '角色添加權限', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('5', 'AddAdminRoles', 'AddAdminRoles', '為管理員添加腳色', '為管理員添加腳色', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('6', 'AddAdminPermits', 'AddAdminPermits', '為管理員添加權限', '為管理員添加權限', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('7', 'GetAllAdminList', 'GetAllAdminList', '查詢所有管理員', '查詢所有管理員', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('8', 'GetAdminRole', 'GetAdminRole', '查詢指定管理員的角色', '查詢指定管理員的角色', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('9', 'GetAdminExtraPermits', 'GetAdminExtraPermits', '查詢指定管理員額外的權限', '查詢指定管理員額外的權限', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');




/**
  6. 腳色
 */
CREATE TABLE IF NOT EXISTS `admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_key` varchar(40) NOT NULL COMMENT '給spring定位用',
  `role_name` varchar(40) NOT NULL COMMENT '腳色名稱',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `role_status` int(11) NOT NULL DEFAULT '0' COMMENT '腳色狀態',
  `memo` text COMMENT '備註',
  `creator_id` varchar(40) NOT NULL COMMENT '創建者id(admin id)',
  `updater_id` varchar(40) NOT NULL COMMENT '更新者id(admin id)',
  `role_desc` text COMMENT '說明',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_role_role_key_uindex` (`role_key`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- 插入腳色数据
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(0, 'ADMIN', '管理員', 0, 0, NULL, '1', '1', '管理員', '2023-03-17 13:13:44', '2023-03-17 13:13:43');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(1, 'SUPER_MANAGER', '超級管理員', 0, 0, NULL, '1', '1', '超級管理員', '2023-04-25 09:40:15', '2023-04-25 09:40:15');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(2, 'NORMAL_MANAGER', '普通管理員', 0, 0, NULL, '1', '1', '普通管理員', '2023-04-25 09:43:19', '2023-04-25 09:43:19');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(3, 'USER', '普通用戶', 0, 0, NULL, '1', '1', '普通用戶', '2023-04-25 09:44:10', '2023-04-25 09:44:10');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(4, 'TESTKEY', 'TESTNAME', 0, 1, '', '0', '0', '', '2024-01-16 09:18:50', '2024-01-16 09:18:50');




/**
  7. 腳色與權限的對應關係
 */
CREATE TABLE IF NOT EXISTS `admin_role_permit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(40) NOT NULL COMMENT '腳色id',
  `permit_id` varchar(40) NOT NULL COMMENT '權限id',
  `creator_id` varchar(40) NOT NULL COMMENT '創建者id',
  `updater_id` varchar(40) NOT NULL COMMENT '更新者id',
  `create_time` timestamp NOT NULL COMMENT '創建時間',
  `update_time` timestamp NOT NULL COMMENT '更新時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roleid_permitid` (`role_id`,`permit_id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4;

-- 插入腳色與權限的對應關係数据
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(0, '0', '0', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(1, '0', '1', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(2, '0', '2', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(3, '0', '3', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(4, '0', '4', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(5, '0', '5', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(6, '0', '6', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(7, '0', '7', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(8, '0', '8', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(9, '0', '9', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(10, '0', '10', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(11, '0', '11', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(12, '0', '12', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(13, '0', '13', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(14, '0', '14', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(15, '1', '1', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(16, '1', '10', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(17, '1', '11', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(18, '1', '12', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(19, '1', '13', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(20, '1', '14', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(21, '1', '2', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(22, '1', '3', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(23, '1', '4', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(24, '1', '5', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(25, '1', '6', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(26, '1', '7', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(27, '1', '8', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(28, '1', '9', '0', '0', '2024-01-24 09:43:45', '2024-01-24 09:43:45');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(29, '2', '10', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(30, '2', '2', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(31, '2', '3', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(32, '2', '7', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(33, '2', '8', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(34, '2', '9', '0', '0', '2024-01-24 09:49:09', '2024-01-24 09:49:09');
