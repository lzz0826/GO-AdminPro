CREATE DATABASE admin;

USE admin;

/**
  1. Administrator Account
 */
CREATE TABLE IF NOT EXISTS `admin_admin` (
  `id` varchar(40) NOT NULL COMMENT 'Admin user ID',
  `channel_id` varchar(40) DEFAULT NULL,
  `username` varchar(40) NOT NULL COMMENT 'Username',
  `admin_name` varchar(40) DEFAULT NULL COMMENT 'Administrator name',
  `nickname` varchar(40) DEFAULT NULL,
  `account_status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status',
  `login_ip` varchar(40) DEFAULT NULL COMMENT 'Login IP',
  `login_time` timestamp NULL DEFAULT NULL COMMENT 'Latest login date',
  `memo` text COMMENT 'Memo',
  `creator_id` varchar(40) DEFAULT NULL COMMENT 'Creator ID (admin ID)',
  `updater_id` varchar(40) DEFAULT NULL COMMENT 'Updater ID',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  `two_factor_code` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_username_chId` (`username`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert data for administrator accounts
INSERT INTO admin.admin_admin
(id, channel_id, username, admin_name, nickname, account_status, login_ip, login_time, memo, creator_id, updater_id, update_time, create_time, two_factor_code)
VALUES('1', 'channel001', 'admin', 'Highest', 'Big Boss', 0, '192.168.1.1', '2024-01-12 09:58:00', 'Some memo', '999', '999', '2024-01-12 09:58:00', '2024-01-12 09:58:00', '123456');
INSERT INTO admin.admin_admin
(id, channel_id, username, admin_name, nickname, account_status, login_ip, login_time, memo, creator_id, updater_id, update_time, create_time, two_factor_code)
VALUES('2', 'channel002', 'manager', 'Manager', 'Second in Command', 0, '192.168.1.1', '2024-01-12 09:58:00', 'Some memo', '999', '999', '2024-01-12 09:58:00', '2024-01-14 09:58:00', '123456');


/**
  2. Permissions for Administrators
 */
CREATE TABLE IF NOT EXISTS `admin_admin_permit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(40) NOT NULL COMMENT 'Admin ID',
  `permit_id` varchar(40) NOT NULL COMMENT 'Permission ID',
  `creator_id` varchar(40) NOT NULL COMMENT 'Creator ID',
  `updater_id` varchar(40) NOT NULL COMMENT 'Updater ID',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_adminid_permitid` (`admin_id`,`permit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



/**
  3. Roles for Administrators
 */
CREATE TABLE IF NOT EXISTS `admin_admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` varchar(40) NOT NULL COMMENT 'Associated admin ID',
  `role_id` varchar(40) NOT NULL COMMENT 'Associated role ID',
  `creator_id` varchar(40) DEFAULT NULL COMMENT 'Creator ID',
  `updater_id` varchar(40) DEFAULT NULL COMMENT 'Updater ID',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_adminid_roleid` (`admin_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert data for roles associated with administrators
INSERT INTO admin.admin_admin_role
(id, admin_id, role_id, creator_id, updater_id, create_time, update_time)
VALUES(1, '1', '1', '999', '999', '2024-01-14 14:39:16', '2024-01-14 14:39:16');
INSERT INTO admin.admin_admin_role
(id, admin_id, role_id, creator_id, updater_id, create_time, update_time)
VALUES(2, '2', '2', '999', '999', '2024-01-12 10:14:18', '2024-01-12 10:14:18');




/**
  4. Administrator Account Passwords
 */
CREATE TABLE IF NOT EXISTS `admin_admintoken` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(40) NOT NULL COMMENT 'Admin ID',
  `token_type` int(11) DEFAULT NULL COMMENT 'Token type',
  `token` text NOT NULL COMMENT 'Token',
  `expire_time` timestamp NULL DEFAULT NULL COMMENT 'Expiration time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'Update time',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  `creator_id` varchar(40) DEFAULT NULL,
  `updater_id` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_admin_id` (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert data for administrator account passwords
INSERT INTO admin.admin_admintoken
(id, admin_id, token_type, token, expire_time, update_time, create_time, creator_id, updater_id)
VALUES(1, '1', 1, '$2a$10$dU6nEGsh7QA4/caU3WZC9ODBThGZ8f/p9a3Q66LHH0UQzqKnANpva
', '2024-01-12 11:22:52', '2024-01-12 11:22:52', '2024-01-12 11:22:52', '999', '999');
INSERT INTO admin.admin_admintoken
(id, admin_id, token_type, token, expire_time, update_time, create_time, creator_id, updater_id)
VALUES(2, '2', 1, '$2a$10$dU6nEGsh7QA4/caU3WZC9ODBThGZ8f/p9a3Q66LHH0UQzqKnANpva
', '2024-01-12 11:22:52', '2024-01-12 10:22:52', '2024-01-12 10:22:52', '999', '999');




/**
  5. Permissions
 */
CREATE TABLE IF NOT EXISTS `admin_permit` (
  `id` varchar(40) NOT NULL COMMENT 'ID',
  `permit_key` varchar(40) NOT NULL COMMENT 'Used by Spring for identification',
  `permit_name` varchar(40) NOT NULL COMMENT 'Permission name',
  `memo` text COMMENT 'Memo',
  `permit_desc` text COMMENT 'Description',
  `sort` int(11) DEFAULT NULL COMMENT 'Sorting',
  `creator_id` varchar(40) NOT NULL COMMENT 'Creator ID (admin ID)',
  `updater_id` varchar(40) NOT NULL COMMENT 'Updater ID (admin ID)',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permit_permit_key_uindex` (`permit_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert permission data
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('1', 'AddAdmin', 'AddAdmin', 'Add administrator', 'Add administrator', 1, '999', '999', '2024-01-12 10:31:53', '2024-01-12 10:31:53');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('10', 'AddRolePermits', 'AddRolePermits', 'Add permissions to a role', 'Add permissions to a role', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('11', 'AddAdminRoles', 'AddAdminRoles', 'Add roles to an administrator', 'Add roles to an administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('12', 'AddAdminPermits', 'AddAdminPermits', 'Add permissions to an administrator', 'Add permissions to an administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('13', 'GetAllAdminList', 'GetAllAdminList', 'Query all administrators', 'Query all administrators', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('14', 'GetAdminRole', 'GetAdminRole', 'Query roles for a specified administrator', 'Query roles for a specified administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('15', 'GetAdminExtraPermits', 'GetAdminExtraPermits', 'Query additional permissions for a specified administrator', 'Query additional permissions for a specified administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('2', 'AddRole', 'AddRole', 'Add role', 'Add role', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('3', 'GetAdminAllPermits', 'GetAdminAllPermits', 'Query all permissions for a specified administrator (including roles)', 'Query all permissions for a specified administrator (including roles)', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('4', 'RemoveRolePermits', 'RemoveRolePermits', 'Remove permissions associated with a role', 'Remove permissions associated with a role', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('5', 'RemoveAdminPermits', 'RemoveAdminPermits', 'Remove additional permissions for an administrator', 'Remove additional permissions for an administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('6', 'RemoveAdminRoles', 'RemoveAdminRoles', 'Remove roles assigned to an administrator', 'Remove roles assigned to an administrator', 1, '999', '999', '2024-01-14 14:51:43', '2024-01-14 14:51:43');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('7', 'GetRolePermits', 'GetRolePermits', 'Get role permissions', 'Get role permissions', 1, '999', '999', '2024-01-12 10:31:53', '2024-01-12 10:31:53');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('8', 'GetAllRoleList', 'GetAllRoleList', 'Query all roles', 'Query all roles', 1, '999', '999', '2024-01-12 10:32:21', '2024-01-12 10:32:21');
INSERT INTO admin.admin_permit
(id, permit_key, permit_name, memo, permit_desc, sort, creator_id, updater_id, update_time, create_time)
VALUES('9', 'GetAllPermitList', 'GetAllPermitList', 'Query all permissions', 'Query all permissions', 1, '999', '999', '2024-01-12 10:32:21', '2024-01-12 10:32:21');




/**
  6. Roles
 */
CREATE TABLE IF NOT EXISTS `admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_key` varchar(40) NOT NULL COMMENT 'Used by Spring for identification',
  `role_name` varchar(40) NOT NULL COMMENT 'Role name',
  `sort` int(11) DEFAULT NULL COMMENT 'Sorting',
  `role_status` int(11) NOT NULL DEFAULT '0' COMMENT 'Role status',
  `memo` text COMMENT 'Memo',
  `creator_id` varchar(40) NOT NULL COMMENT 'Creator ID (admin ID)',
  `updater_id` varchar(40) NOT NULL COMMENT 'Updater ID (admin ID)',
  `role_desc` text COMMENT 'Description',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_role_role_key_uindex` (`role_key`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert role data
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(1, 'ADMIN', 'Administrator', 0, 0, NULL, '1', '1', 'Administrator', '2023-03-17 13:13:44', '2023-03-17 13:13:43');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(2, 'SUPER_MANAGER', 'Super Manager', 0, 0, NULL, '1', '1', 'Super Manager', '2023-04-25 09:40:15', '2023-04-25 09:40:15');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(3, 'NORMAL_MANAGER', 'Normal Manager', 0, 0, NULL, '1', '1', 'Normal Manager', '2023-04-25 09:43:19', '2023-04-25 09:43:19');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(4, 'USER', 'Regular User', 0, 0, NULL, '1', '1', 'Regular User', '2023-04-25 09:44:10', '2023-04-25 09:44:10');
INSERT INTO admin.admin_role
(id, role_key, role_name, sort, role_status, memo, creator_id, updater_id, role_desc, update_time, create_time)
VALUES(5, 'TESTKEY', 'TESTNAME', 0, 1, '', '0', '0', '', '2024-01-16 09:18:50', '2024-01-16 09:18:50');



/**
  7. Mapping between Roles and Permissions
 */
CREATE TABLE IF NOT EXISTS `admin_role_permit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(40) NOT NULL COMMENT 'Role ID',
  `permit_id` varchar(40) NOT NULL COMMENT 'Permission ID',
  `creator_id` varchar(40) NOT NULL COMMENT 'Creator ID',
  `updater_id` varchar(40) NOT NULL COMMENT 'Updater ID',
  `create_time` timestamp NOT NULL COMMENT 'Creation time',
  `update_time` timestamp NOT NULL COMMENT 'Update time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_roleid_permitid` (`role_id`,`permit_id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert data for the mapping between Roles and Permissions
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(1, '1', '1', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(2, '1', '2', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(3, '1', '3', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(4, '1', '4', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(5, '1', '5', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(6, '1', '6', '999', '999', '2024-01-12 10:42:59', '2024-01-12 10:42:59');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(7, '1', '7', '999', '999', '2024-01-22 07:30:42', '2024-01-22 07:30:42');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(8, '1', '8', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(9, '1', '9', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(10, '1', '10', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(11, '1', '11', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(12, '1', '12', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(13, '1', '13', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(14, '1', '14', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(15, '1', '15', '999', '999', '2024-01-12 10:42:33', '2024-01-12 10:42:33');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(16, '2', '10', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(17, '2', '11', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(18, '2', '12', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(19, '2', '13', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(20, '2', '14', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(21, '2', '15', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(22, '2', '2', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(23, '2', '3', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(24, '2', '4', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(25, '2', '5', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(26, '2', '6', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(27, '2', '7', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(28, '2', '8', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(29, '2', '9', '1', '1', '2024-01-24 13:35:47', '2024-01-24 13:35:47');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(30, '3', '13', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(31, '3', '14', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(32, '3', '15', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(33, '3', '3', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(34, '3', '7', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(35, '3', '8', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');
INSERT INTO admin.admin_role_permit
(id, role_id, permit_id, creator_id, updater_id, create_time, update_time)
VALUES(36, '3', '9', '1', '1', '2024-01-24 13:39:35', '2024-01-24 13:39:35');



CREATE TABLE IF NOT EXISTS `account_payee_check` (
                                       `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
                                       `uid` INT COMMENT '用户id',
                                       `type` INT COMMENT '审核类型1-银行卡,2-支付保',
                                       `description` VARCHAR(255) COMMENT '审核内容',
                                       `status` INT COMMENT '状态(0-审核中,1-审核成功,2-审核失败)',
                                       `check_id` INT COMMENT '审核人id',
                                       `check_time` DATETIME COMMENT '审核时间',
                                       `update_time` DATETIME COMMENT '更新时间',
                                       `created_time` DATETIME COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account Payee Check Table';