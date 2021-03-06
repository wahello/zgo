-- -------------------------------------------------------
-- build by cmd/db/mysql/mysql.go
-- time: 2020-09-16 14:56:51 CST
-- -------------------------------------------------------
-- 表结构
-- -------------------------------------------------------
-- 账户实体
CREATE TABLE `zgo_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `pid` int(11) DEFAULT NULL COMMENT '上级账户',
  `account` varchar(255) NOT NULL COMMENT '账户',
  `account_typ` tinyint(4) DEFAULT '1' COMMENT '账户类型',
  `account_kid` varchar(64) DEFAULT NULL COMMENT '账户归属平台',
  `password` varchar(255) DEFAULT NULL COMMENT '登录密码',
  `password_salt` varchar(255) DEFAULT NULL COMMENT '密码盐值',
  `password_type` varchar(16) DEFAULT NULL COMMENT '密码方式',
  `verify_secret` varchar(255) DEFAULT NULL COMMENT '校验密钥',
  `user_id` int(11) NOT NULL COMMENT '用户标识',
  `role_id` int(11) DEFAULT NULL COMMENT '角色标识',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `description` varchar(255) DEFAULT NULL COMMENT '账户描述',
  `oa2_token` varchar(1024) DEFAULT NULL COMMENT 'oauth2令牌',
  `oa2_expired` int(11) DEFAULT NULL COMMENT 'oauth2过期时间',
  `oa2_refresh` varchar(2048) DEFAULT NULL COMMENT '刷新令牌',
  `oa2_scope` varchar(255) DEFAULT NULL COMMENT '授权作用域',
  `oa2_fake` varchar(1024) DEFAULT NULL COMMENT '伪造令牌',
  `oa2_client` int(11) DEFAULT NULL COMMENT '客户端上次',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `string_2` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `string_3` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  `number_2` int(11) DEFAULT NULL COMMENT '备用字段',
  `number_3` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_account(`account`,`account_typ`,`account_kid`),
  INDEX idx_account_client(`oa2_client`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方登陆实体
CREATE TABLE `zgo_oauth2_platform` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `kid` varchar(64)  NOT NULL COMMENT '三方标识',
  `platform` varchar(32) NOT NULL COMMENT '平台标识',
  `app_id` varchar(255) DEFAULT NULL COMMENT '应用标识',
  `app_secret` varchar(1024) DEFAULT NULL COMMENT '应用密钥',
  `avatar` varchar(255) DEFAULT NULL COMMENT '平台头像',
  `description` varchar(255) DEFAULT NULL COMMENT '平台描述',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `signin` tinyint(4) DEFAULT 0 COMMENT '可登陆',
  `agent_id` varchar(255) DEFAULT NULL COMMENT '代理商标识',
  `agent_secret` varchar(1024) DEFAULT NULL COMMENT '代理商密钥',
  `suite_id` varchar(255) DEFAULT NULL COMMENT '套件标识',
  `suite_secret` varchar(1024) DEFAULT NULL COMMENT '套件密钥',
  `authorize_url` varchar(1024) DEFAULT NULL COMMENT '认证地址',
  `token_url` varchar(1024) DEFAULT NULL COMMENT '令牌地址',
  `profile_url` varchar(1024) DEFAULT NULL COMMENT '个人资料地址',
  `signin_url` varchar(128) DEFAULT NULL COMMENT '重定向回调地址',
  `js_secret` varchar(255) DEFAULT NULL COMMENT 'javascript密钥',
  `state_secret` varchar(255) DEFAULT NULL COMMENT '回调state密钥',
  `callback` tinyint(4) DEFAULT 0 COMMENT '是否支持回调',
  `cb_domain` varchar(128) DEFAULT NULL COMMENT '默认域名',
  `cb_scheme` varchar(16) DEFAULT 'https' COMMENT '默认协议',
  `cb_encrypt` tinyint(4) DEFAULT 0 COMMENT '回调是否加密',
  `cb_token` varchar(255) DEFAULT NULL COMMENT '加密令牌',
  `cb_encoding` varchar(255) DEFAULT NULL COMMENT '加密编码',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_oauth2_platform_kid(`kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方通信实体
CREATE TABLE `zgo_oauth2_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `oauth2_id` int(11) DEFAULT NULL COMMENT '平台',
  `token_kid` varchar(64) DEFAULT NULL COMMENT '角色标识',
  `access_token` varchar(1024) DEFAULT NULL COMMENT '访问令牌',
  `expires_in` int(11) DEFAULT 7200 COMMENT '有限期间隔',
  `expires_time` timestamp DEFAULT NULL COMMENT '凭据过期时间',
  `refresh_token` varchar(1024) DEFAULT NULL COMMENT '刷新令牌',
  `refresh_count` int(11) DEFAULT 0 COMMENT '刷新次数',
  `sync_lock` timestamp DEFAULT NULL COMMENT '同步锁',
  `call_count` int(11) DEFAULT 0 COMMENT '调用次数',
  `token_type` varchar(32) DEFAULT NULL COMMENT '令牌类型',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  INDEX idx_oauth2_token_tkid(`token_kid`),
  INDEX idx_oauth2_token_exp_time(`expires_time`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方授权实体
CREATE TABLE `zgo_oauth2_client` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `kid` varchar(64)  NOT NULL COMMENT '客户端标识',
  `audience` varchar(255) DEFAULT NULL COMMENT '令牌接受平台',
  `issuer` varchar(255) DEFAULT NULL COMMENT '令牌签发平台',
  `expired` int(11) DEFAULT 7200 COMMENT '令牌有效期',
  `token_type` varchar(32) DEFAULT 'JWT' COMMENT '令牌类型',
  `token_method` varchar(32) DEFAULT 'HS512' COMMENT '令牌方法',
  `token_secret` varchar(255) NOT NULL COMMENT '令牌密钥',
  `token_getter` varchar(32) DEFAULT NULL COMMENT '令牌获取方法',
  `signin_url` varchar(2048) DEFAULT NULL COMMENT '登陆地址',
  `signin_force` tinyint(4) DEFAULT 0 COMMENT '强制跳转登陆',
  `signin_check` tinyint(4) DEFAULT 0 COMMENT '登陆确认',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_oauth2_client_kid(`kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方授权实体
CREATE TABLE `zgo_account_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `account_id` int(11) NOT NULL COMMENT '账户标识',
  `token_kid` varchar(64) DEFAULT NULL COMMENT '角色标识',
  `client_id` int(11) DEFAULT NULL COMMENT '客户端标识',
  `client_kid` varchar(64) DEFAULT NULL COMMENT '客户端标识',
  `user_kid` varchar(64) DEFAULT NULL COMMENT '用户标识',
  `role_kid` varchar(64) DEFAULT NULL COMMENT '角色标识',
  `last_ip` varchar(64) DEFAULT NULL COMMENT '上次登陆IP',
  `last_at` timestamp DEFAULT NULL COMMENT '上次登陆时间',
  `limit_exp` timestamp DEFAULT NULL COMMENT '限制登陆',
  `limit_key` varchar(255) DEFAULT NULL COMMENT '限制登陆',
  `mode` varchar(16) DEFAULT 'signin' COMMENT '方式',
  `expires_at` int(11) DEFAULT NULL COMMENT '授权有效期',
  `access_token` varchar(2048) DEFAULT NULL COMMENT '访问令牌',
  `refresh_token` varchar(128) DEFAULT NULL COMMENT '刷新令牌',
  `refresh_count` int(11) DEFAULT 0 COMMENT '刷新次数',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  INDEX idx_account_token_ukid(`user_kid`),
  INDEX idx_account_token_rkid(`role_kid`),
  INDEX idx_oauth2_account_expires(`expires_at`),
  INDEX idx_oauth2_account_rtid(`refresh_token`),
  INDEX idx_account_token_aid(`account_id`),
  INDEX idx_account_token_tkid(`token_kid`),
  INDEX idx_account_token_cid(`client_id`),
  INDEX idx_account_token_ckid(`client_kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 用户实体
CREATE TABLE `zgo_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `kid` varchar(64) DEFAULT NULL COMMENT '唯一标识',
  `name` varchar(64) DEFAULT NULL COMMENT '用户名',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `delete` tinyint(4) DEFAULT 0 COMMENT '删除标识',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_user_kid(`kid`),
  INDEX idx_user_name(`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 用户安全
CREATE TABLE `zgo_user_security` (
  `id` int(11) NOT NULL COMMENT '唯一标识',
  `pwd_account` int(11) DEFAULT NULL COMMENT '密码账户',
  `3rd_account` int(11) DEFAULT NULL COMMENT '第三方账户',
  `mfa_secret` varchar(1024) DEFAULT NULL COMMENT 'mfa密钥',
  `tel_secret` varchar(16) DEFAULT NULL COMMENT '密保电话',
  `bak_email` varchar(128) DEFAULT NULL COMMENT '备用邮箱',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  INDEX idx_user_pwd_account_id(`pwd_account`),
  INDEX idx_user_3rd_account_id(`3rd_account`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 角色实体
CREATE TABLE `zgo_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `kid` varchar(64) DEFAULT NULL COMMENT '唯一标识',
  `name` varchar(64) DEFAULT NULL COMMENT '角色名',
  `description` varchar(128) DEFAULT NULL COMMENT '角色描述',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `domain` varchar(255) DEFAULT NULL COMMENT '域',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_role_kid(`kid`),
  UNIQUE udx_role_name(`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 角色角色实体
CREATE TABLE `zgo_role_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `owner_id` int(11) DEFAULT NULL COMMENT '父节点标识',
  `child_id` int(11) DEFAULT NULL COMMENT '子节点标识',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 用户角色实体
CREATE TABLE `zgo_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `user_id` int(11) DEFAULT NULL COMMENT '账户标识',
  `role_id` int(11) DEFAULT NULL COMMENT '客户端标识',
  `expired` int(11) DEFAULT NULL COMMENT '授权有效期',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 资源实体
CREATE TABLE `zgo_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `resource` varchar(64) DEFAULT NULL COMMENT '资源名',
  `domain` varchar(255) DEFAULT NULL COMMENT '域名',
  `methods` varchar(64) DEFAULT NULL COMMENT '方法',
  `path` varchar(255) DEFAULT NULL COMMENT '路径',
  `netmask` varchar(64) DEFAULT NULL COMMENT '网络掩码',
  `allow` tinyint(4) DEFAULT NULL COMMENT '允许vs拒绝',
  `description` varchar(128) DEFAULT NULL COMMENT '描述',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  INDEX idx_resource_name(`resource`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 资源角色实体
CREATE TABLE `zgo_resource_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `role_id` int(11) DEFAULT NULL COMMENT '角色',
  `resource` varchar(64) DEFAULT NULL COMMENT '资源名',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 资源角色实体
CREATE TABLE `zgo_resource_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `user_id` int(11) DEFAULT NULL COMMENT '用户',
  `resource` varchar(64) DEFAULT NULL COMMENT '资源名',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 菜单实体
CREATE TABLE `zgo_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `kid` varchar(32) DEFAULT NULL COMMENT '唯一标识',
  `name` varchar(64) DEFAULT NULL COMMENT '菜单名称',
  `local` varchar(128) DEFAULT NULL COMMENT '菜单名称',
  `sequence` tinyint(4) DEFAULT 64 COMMENT '排序值',
  `group` varchar(64) DEFAULT NULL COMMENT '菜单分组',
  `group_local` varchar(64) DEFAULT NULL COMMENT '菜单分组',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `router` varchar(255) DEFAULT NULL COMMENT '访问路由',
  `memo` varchar(255) DEFAULT NULL COMMENT '备注',
  `domain` varchar(255) DEFAULT NULL COMMENT '域名',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_menu_kid(`kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 角色自定义菜单实体
CREATE TABLE `zgo_menu_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `pid` int(11) DEFAULT NULL COMMENT '父节点',
  `kid` varchar(32) DEFAULT NULL COMMENT '唯一标识',
  `name` varchar(64) DEFAULT NULL COMMENT '菜单名称',
  `local` varchar(128) DEFAULT NULL COMMENT '菜单名称',
  `sequence` tinyint(4) DEFAULT 64 COMMENT '排序值',
  `role_id` int(11) DEFAULT NULL COMMENT '角色 ID',
  `role_kid` varchar(64) DEFAULT NULL COMMENT '角色 UID',
  `menu_id` int(11) DEFAULT NULL COMMENT '菜单 ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  UNIQUE udx_menu_role_kid(`kid`),
  INDEX idx_parent_id(`pid`),
  INDEX idx_role_kid(`role_kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 角色自定义菜单实体
CREATE TABLE `zgo_menu_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `role_id` int(11) DEFAULT NULL COMMENT '角色 ID',
  `role_kid` varchar(64) DEFAULT NULL COMMENT '角色 UID',
  `user_id` int(11) DEFAULT NULL COMMENT '用户 ID',
  `user_kid` varchar(64) DEFAULT NULL COMMENT '用户 UID',
  `menu_ids` varchar(4096) DEFAULT NULL COMMENT '菜单 ID',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  INDEX idx_user_kid(`user_kid`),
  INDEX idx_role_kid(`role_kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 菜单事件实体
CREATE TABLE `zgo_menu_action` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `menu_id` int(11) DEFAULT NULL COMMENT '菜单 ID',
  `role_id` int(11) DEFAULT NULL COMMENT '角色 ID',
  `role_kid` varchar(64) DEFAULT NULL COMMENT '角色 UID',
  `code` varchar(64) DEFAULT NULL COMMENT '动作编号',
  `name` varchar(64) DEFAULT NULL COMMENT '动作名称',
  `disable` tinyint(4) DEFAULT 0 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  INDEX idx_role_kid(`role_kid`),
  INDEX idx_menu_action_code(`code`),
  INDEX idx_menu_action_name(`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 系统标签实体
CREATE TABLE `zgo_tag_system` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 系统标签实体
CREATE TABLE `zgo_tag_system_owner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `type` tinyint(4) DEFAULT NULL COMMENT '标签类型',
  `owner_id` int(11) DEFAULT NULL COMMENT '归属id',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  INDEX idx_tag_system_owner_id(`owner_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 通用标签实体
CREATE TABLE `zgo_tag_custom` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `type` tinyint(4) DEFAULT 0 COMMENT '标签类型',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 通用标签实体
CREATE TABLE `zgo_tag_custom_owner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `type` tinyint(4) DEFAULT NULL COMMENT '标签类型',
  `owner_id` int(11) DEFAULT NULL COMMENT '归属id',
  `deleted` tinyint(4) DEFAULT 0 COMMENT '删除标识',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  INDEX idx_tag_custom_owner_id(`owner_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 国际化实体
CREATE TABLE `zgo_i18n_language` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `mid` varchar(255) DEFAULT NULL COMMENT 'message id',
  `lang` varchar(16) DEFAULT NULL COMMENT '语言',
  `description` varchar(64) DEFAULT NULL COMMENT '描述',
  `left_delim` varchar(16) DEFAULT NULL COMMENT '定界符',
  `right_delim` varchar(16) DEFAULT NULL COMMENT '定界符',
  `zero` varchar(255) DEFAULT NULL COMMENT 'zero',
  `one` varchar(255) DEFAULT NULL COMMENT 'one',
  `two` varchar(255) DEFAULT NULL COMMENT 'two',
  `few` varchar(255) DEFAULT NULL COMMENT 'few',
  `many` varchar(255) DEFAULT NULL COMMENT 'many',
  `other` varchar(255) DEFAULT NULL COMMENT 'other',
  `status` tinyint(4) DEFAULT 1 COMMENT '状态',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  UNIQUE udx_i18n_message_id(`mid`,`lang`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 用户详情实体
CREATE TABLE `zgo_user_detail` (
  `id` int(11) NOT NULL COMMENT '唯一标识',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `mfa_secret` varchar(1024) DEFAULT NULL COMMENT 'MFA密钥',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 员工详情实体
CREATE TABLE `zgo_user_employee` (
  `id` int(11) NOT NULL COMMENT '唯一标识',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `mfa_secret` varchar(1024) DEFAULT NULL COMMENT 'MFA密钥',
  `deleted` tinyint(4) DEFAULT 0 COMMENT '逻辑删除',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 顾客详情实体
CREATE TABLE `zgo_user_customer` (
  `id` int(11) NOT NULL COMMENT '唯一标识',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `mfa_secret` varchar(1024) DEFAULT NULL COMMENT 'MFA密钥',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 用户消息实体
CREATE TABLE `zgo_user_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `from_id` int(64) DEFAULT NULL COMMENT '发送消息',
  `to_id` int(64) DEFAULT NULL COMMENT '发送消息',
  `kid` varchar(64) DEFAULT NULL COMMENT '索引',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `datetime` timestamp DEFAULT NULL COMMENT '日期',
  `type` varchar(16) DEFAULT NULL COMMENT '类型',
  `read` tinyint(4) DEFAULT NULL COMMENT '已读',
  `description` TEXT DEFAULT NULL COMMENT '描述',
  `clickClose` tinyint(4) DEFAULT NULL COMMENT '关闭',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  UNIQUE udx_user_message_kid(`kid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方消息实体
CREATE TABLE `zgo_oauth2_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------
-- 第三方消息转发(forward)实体
CREATE TABLE `zgo_oauth2_message_next` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `created_at` timestamp DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp DEFAULT NULL COMMENT '更新时间',
  `version` int(11) DEFAULT 0 COMMENT '数据版本',
  `string_1` varchar(255) DEFAULT NULL COMMENT '备用字段',
  `number_1` int(11) DEFAULT NULL COMMENT '备用字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
-- -------------------------------------------------------

-- -------------------------------------------------------
-- 表外键
-- -------------------------------------------------------
ALTER TABLE `zgo_account`
ADD CONSTRAINT `fk_account_kid` FOREIGN KEY (`account_kid`)  REFERENCES `zgo_oauth2_platform` (`kid`),
ADD CONSTRAINT `fk_account_user` FOREIGN KEY (`user_id`)  REFERENCES `zgo_user` (`id`),
ADD CONSTRAINT `fk_account_role` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`);

ALTER TABLE `zgo_oauth2_token`
ADD CONSTRAINT `fk_oauth2_token_id` FOREIGN KEY (`oauth2_id`)  REFERENCES `zgo_oauth2_platform` (`id`);

ALTER TABLE `zgo_user_security`
ADD CONSTRAINT `fk_user_security` FOREIGN KEY (`id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_role_role`
ADD CONSTRAINT `fk_role_owner_id` FOREIGN KEY (`owner_id`)  REFERENCES `zgo_role` (`id`),
ADD CONSTRAINT `fk_role_child_id` FOREIGN KEY (`child_id`)  REFERENCES `zgo_role` (`id`);

ALTER TABLE `zgo_user_role`
ADD CONSTRAINT `fk_role_role_id` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`),
ADD CONSTRAINT `fk_role_user_id` FOREIGN KEY (`user_id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_resource_role`
ADD CONSTRAINT `fk_resource_role_id` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`),
ADD CONSTRAINT `fk_resource_role_res` FOREIGN KEY (`resource`)  REFERENCES `zgo_resource` (`resource`);

ALTER TABLE `zgo_resource_user`
ADD CONSTRAINT `fk_resource_user_id` FOREIGN KEY (`user_id`)  REFERENCES `zgo_user` (`id`),
ADD CONSTRAINT `fk_resource_user_res` FOREIGN KEY (`resource`)  REFERENCES `zgo_resource` (`resource`);

ALTER TABLE `zgo_menu_role`
ADD CONSTRAINT `fk_menu_role_role_id` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`),
ADD CONSTRAINT `fk_menu_role_menu_id` FOREIGN KEY (`menu_id`)  REFERENCES `zgo_menu` (`id`);

ALTER TABLE `zgo_menu_user`
ADD CONSTRAINT `fk_menu_user_role_id` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`),
ADD CONSTRAINT `fk_menu_user_user_id` FOREIGN KEY (`user_id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_menu_action`
ADD CONSTRAINT `fk_menu_action_menu_id` FOREIGN KEY (`menu_id`)  REFERENCES `zgo_menu` (`id`),
ADD CONSTRAINT `fk_menu_action_role_id` FOREIGN KEY (`role_id`)  REFERENCES `zgo_role` (`id`);

ALTER TABLE `zgo_user_detail`
ADD CONSTRAINT `fk_user_detail` FOREIGN KEY (`id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_user_employee`
ADD CONSTRAINT `fk_user_employee` FOREIGN KEY (`id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_user_customer`
ADD CONSTRAINT `fk_user_customer` FOREIGN KEY (`id`)  REFERENCES `zgo_user` (`id`);

ALTER TABLE `zgo_user_message`
ADD CONSTRAINT `fk_u_msg_from_id` FOREIGN KEY (`from_id`)  REFERENCES `zgo_user` (`id`),
ADD CONSTRAINT `fk_u_msg_to_id` FOREIGN KEY (`to_id`)  REFERENCES `zgo_user` (`id`);

-- -------------------------------------------------------
-- -------------------------------------------------------
-- insert into 
-- -------------------------------------------------------
-- INSERT INTO `zgo_account`(`id`, `pid`, `account`, `account_typ`, `account_kid`, `password`, `password_salt`, `password_type`, `verify_secret`, `user_id`, `role_id`, `status`, `description`, `oa2_token`, `oa2_expired`, `oa2_refresh`, `oa2_scope`, `oa2_fake`, `oa2_client`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `string_2`, `string_3`, `number_1`, `number_2`, `number_3`) VALUES ();
-- INSERT INTO `zgo_oauth2_platform`(`id`, `kid`, `platform`, `app_id`, `app_secret`, `avatar`, `description`, `status`, `signin`, `agent_id`, `agent_secret`, `suite_id`, `suite_secret`, `authorize_url`, `token_url`, `profile_url`, `signin_url`, `js_secret`, `state_secret`, `callback`, `cb_domain`, `cb_scheme`, `cb_encrypt`, `cb_token`, `cb_encoding`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_oauth2_token`(`id`, `oauth2_id`, `token_kid`, `access_token`, `expires_in`, `expires_time`, `refresh_token`, `refresh_count`, `sync_lock`, `call_count`, `token_type`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_oauth2_client`(`id`, `kid`, `audience`, `issuer`, `expired`, `token_type`, `token_method`, `token_secret`, `token_getter`, `signin_url`, `signin_force`, `signin_check`, `status`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_account_token`(`id`, `account_id`, `token_kid`, `client_id`, `client_kid`, `user_kid`, `role_kid`, `last_ip`, `last_at`, `limit_exp`, `limit_key`, `mode`, `expires_at`, `access_token`, `refresh_token`, `refresh_count`, `status`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_user`(`id`, `kid`, `name`, `status`, `delete`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_user_security`(`id`, `pwd_account`, `3rd_account`, `mfa_secret`, `tel_secret`, `bak_email`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_role`(`id`, `kid`, `name`, `description`, `status`, `domain`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_role_role`(`id`, `owner_id`, `child_id`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_user_role`(`id`, `user_id`, `role_id`, `expired`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_resource`(`id`, `resource`, `domain`, `methods`, `path`, `netmask`, `allow`, `description`, `status`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_resource_role`(`id`, `role_id`, `resource`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_resource_user`(`id`, `user_id`, `resource`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_menu`(`id`, `kid`, `name`, `local`, `sequence`, `group`, `group_local`, `icon`, `router`, `memo`, `domain`, `status`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_menu_role`(`id`, `pid`, `kid`, `name`, `local`, `sequence`, `role_id`, `role_kid`, `menu_id`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_menu_user`(`id`, `role_id`, `role_kid`, `user_id`, `user_kid`, `menu_ids`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_menu_action`(`id`, `menu_id`, `role_id`, `role_kid`, `code`, `name`, `disable`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_tag_system`(`id`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_tag_system_owner`(`id`, `type`, `owner_id`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_tag_custom`(`id`, `type`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_tag_custom_owner`(`id`, `type`, `owner_id`, `deleted`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_i18n_language`(`id`, `mid`, `lang`, `description`, `left_delim`, `right_delim`, `zero`, `one`, `two`, `few`, `many`, `other`, `status`, `creator`, `created_at`, `updated_at`, `version`) VALUES ();
-- INSERT INTO `zgo_user_detail`(`id`, `nickname`, `avatar`, `mfa_secret`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_user_employee`(`id`, `nickname`, `avatar`, `mfa_secret`, `deleted`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_user_customer`(`id`, `nickname`, `avatar`, `mfa_secret`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_user_message`(`id`, `from_id`, `to_id`, `kid`, `avatar`, `title`, `datetime`, `type`, `read`, `description`, `clickClose`, `status`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_oauth2_message`(`id`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();
-- INSERT INTO `zgo_oauth2_message_next`(`id`, `creator`, `created_at`, `updated_at`, `version`, `string_1`, `number_1`) VALUES ();

-- -------------------------------------------------------
-- -------------------------------------------------------
-- drop table 
-- -------------------------------------------------------
-- ALTER TABLE `zgo_account`
-- DROP FOREIGN KEY `fk_account_role`,
-- DROP FOREIGN KEY `fk_account_kid`,
-- DROP FOREIGN KEY `fk_account_user`;
-- ALTER TABLE `zgo_oauth2_token`
-- DROP FOREIGN KEY `fk_oauth2_token_id`;
-- ALTER TABLE `zgo_user_security`
-- DROP FOREIGN KEY `fk_user_security`;
-- ALTER TABLE `zgo_role_role`
-- DROP FOREIGN KEY `fk_role_owner_id`,
-- DROP FOREIGN KEY `fk_role_child_id`;
-- ALTER TABLE `zgo_user_role`
-- DROP FOREIGN KEY `fk_role_user_id`,
-- DROP FOREIGN KEY `fk_role_role_id`;
-- ALTER TABLE `zgo_resource_role`
-- DROP FOREIGN KEY `fk_resource_role_id`,
-- DROP FOREIGN KEY `fk_resource_role_res`;
-- ALTER TABLE `zgo_resource_user`
-- DROP FOREIGN KEY `fk_resource_user_id`,
-- DROP FOREIGN KEY `fk_resource_user_res`;
-- ALTER TABLE `zgo_menu_role`
-- DROP FOREIGN KEY `fk_menu_role_role_id`,
-- DROP FOREIGN KEY `fk_menu_role_menu_id`;
-- ALTER TABLE `zgo_menu_user`
-- DROP FOREIGN KEY `fk_menu_user_role_id`,
-- DROP FOREIGN KEY `fk_menu_user_user_id`;
-- ALTER TABLE `zgo_menu_action`
-- DROP FOREIGN KEY `fk_menu_action_menu_id`,
-- DROP FOREIGN KEY `fk_menu_action_role_id`;
-- ALTER TABLE `zgo_user_detail`
-- DROP FOREIGN KEY `fk_user_detail`;
-- ALTER TABLE `zgo_user_employee`
-- DROP FOREIGN KEY `fk_user_employee`;
-- ALTER TABLE `zgo_user_customer`
-- DROP FOREIGN KEY `fk_user_customer`;
-- ALTER TABLE `zgo_user_message`
-- DROP FOREIGN KEY `fk_u_msg_from_id`,
-- DROP FOREIGN KEY `fk_u_msg_to_id`;
-- 
-- DROP TABLE IF EXISTS `zgo_account`;
-- DROP TABLE IF EXISTS `zgo_oauth2_platform`;
-- DROP TABLE IF EXISTS `zgo_oauth2_token`;
-- DROP TABLE IF EXISTS `zgo_oauth2_client`;
-- DROP TABLE IF EXISTS `zgo_account_token`;
-- DROP TABLE IF EXISTS `zgo_user`;
-- DROP TABLE IF EXISTS `zgo_user_security`;
-- DROP TABLE IF EXISTS `zgo_role`;
-- DROP TABLE IF EXISTS `zgo_role_role`;
-- DROP TABLE IF EXISTS `zgo_user_role`;
-- DROP TABLE IF EXISTS `zgo_resource`;
-- DROP TABLE IF EXISTS `zgo_resource_role`;
-- DROP TABLE IF EXISTS `zgo_resource_user`;
-- DROP TABLE IF EXISTS `zgo_menu`;
-- DROP TABLE IF EXISTS `zgo_menu_role`;
-- DROP TABLE IF EXISTS `zgo_menu_user`;
-- DROP TABLE IF EXISTS `zgo_menu_action`;
-- DROP TABLE IF EXISTS `zgo_tag_system`;
-- DROP TABLE IF EXISTS `zgo_tag_system_owner`;
-- DROP TABLE IF EXISTS `zgo_tag_custom`;
-- DROP TABLE IF EXISTS `zgo_tag_custom_owner`;
-- DROP TABLE IF EXISTS `zgo_i18n_language`;
-- DROP TABLE IF EXISTS `zgo_user_detail`;
-- DROP TABLE IF EXISTS `zgo_user_employee`;
-- DROP TABLE IF EXISTS `zgo_user_customer`;
-- DROP TABLE IF EXISTS `zgo_user_message`;
-- DROP TABLE IF EXISTS `zgo_oauth2_message`;
-- DROP TABLE IF EXISTS `zgo_oauth2_message_next`;

-- -------------------------------------------------------