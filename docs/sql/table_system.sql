-- 创建数据库 
CREATE DATABASE `gin_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

/*
 用户相关的表
 */
-- 用户表
CREATE TABLE sys_user (
    `id` INT AUTO_INCREMENT COMMENT '用户ID',
    `realname` VARCHAR(32) NULL COMMENT '真实姓名',
    `nickname` VARCHAR(32) NOT NULL COMMENT '昵称',
    `gender` TINYINT(1) NULL COMMENT '0: 保密,1: 女,2: 男',
    `age` INT(11) NULL COMMENT '年龄',
    `birthday` VARCHAR(20) NULL COMMENT '出生日期',
    `avatar` VARCHAR(60) NULL COMMENT '用户头像URL',
    `phone` VARCHAR(20) NOT NULL COMMENT '手机号码',
    `email` VARCHAR(50) NULL COMMENT '邮件',
    `intro` VARCHAR(200) NULL COMMENT '介绍',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `password` VARCHAR(50) NOT NULL COMMENT '密码',
    `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户';

-- 角色表
CREATE TABLE sys_role (
    `id` INT AUTO_INCREMENT COMMENT '角色ID',
    `name` VARCHAR(20) UNIQUE NOT NULL COMMENT '角色名称',
    `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '角色状态,0:停用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '角色表';

-- 用户角色关联表
CREATE TABLE sys_user_role_rel (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT(10) NOT NULL COMMENT '用户ID',
    `role_id` INT(10) NOT NULL COMMENT '角色ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    CONSTRAINT `user_role_rel_user_id` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `user_role_rel_role_id` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户角色关联表';

-- 菜单表
CREATE TABLE sys_menu (
    `id` INT AUTO_INCREMENT COMMENT '菜单ID',
    `parent_id` INT(20) NULL COMMENT '父菜单ID',
    `title` VARCHAR(20) NOT NULL COMMENT '菜单名称',
    `icon` VARCHAR(20) NULL COMMENT '菜单图标',
    `el_svg_icon` VARCHAR(20) NULL COMMENT 'Element菜单图标',
    `menu_type` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '菜单类型,0:菜单,1:按钮',
    `open_type` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '打开方式,0:组件,1:内链,2:外链',
    `path` VARCHAR(500) NULL COMMENT '路由地址',
    `name` VARCHAR(50) NULL COMMENT '路由别名',
    `component` VARCHAR(500) NULL COMMENT '组件路径',
    `redirect` VARCHAR(500) NULL COMMENT '路由重定向',
    `link` VARCHAR(500) NULL COMMENT '链接地址:内链地址/外链地址',
    `target` VARCHAR(500) NULL COMMENT '链接地址跳转方式, component/_blank/_self',
    `permission` VARCHAR(200) NULL COMMENT '权限标识',
    `hidden` TINYINT(1) NULL DEFAULT 1 COMMENT '是否隐藏,0:显示,1:隐藏',
    `always_show` TINYINT(1) NULL DEFAULT 1 COMMENT '始终显示根菜单,0:显示,1:隐藏',
    `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态,0:停用,1:启用',
    `create_user_id` INT NULL COMMENT '创建菜单用户ID',
    `update_user_id` INT NULL COMMENT '更新菜单用户ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '菜单表';

-- 角色菜单关联表
CREATE TABLE sys_role_menu_rel (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `role_id` INT(10) NOT NULL COMMENT '角色ID',
    `menu_id` INT(10) NOT NULL COMMENT '菜单ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    CONSTRAINT `sys_role_menu_rel_role_id` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `sys_role_menu_rel_menu_id` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '角色菜单关联表';

-- user表触发器，更新其他表冗余字段
/*
 CREATE TRIGGER trigger_update_user
 AFTER
 UPDATE
 ON `user` FOR EACH ROW BEGIN
 
 IF NEW.nickname != OLD.nickname THEN 
 -- 更新 api_token.nickname 字段
 UPDATE
 api_token
 SET
 nickname = NEW.nickname
 WHERE
 user_id = NEW.id;
 END IF;
 END;
 */
-- 用户地理位置 - 待定
CREATE TABLE sys_user_location (
    `id` INT AUTO_INCREMENT COMMENT '位置ID',
    `user_id` VARCHAR(10) NOT NULL COMMENT '用户ID',
    `province_code` VARCHAR(10) NULL COMMENT '省',
    `city_code` VARCHAR(10) NULL COMMENT '市',
    `district_code` VARCHAR(10) NULL COMMENT '区',
    `address` VARCHAR(200) NULL COMMENT '居住地址',
    `ad_code` VARCHAR(10) NULL COMMENT '地理编号',
    `lng` VARCHAR(20) NULL COMMENT '城市坐标中心点经度 （ * 1e6 ） ： 如果是中国 ， 此值是 1e7',
    `lat` VARCHAR(20) NULL COMMENT '城市坐标中心点纬度 （ * 1e6 ）',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    CONSTRAINT `user_location_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户地理位置';

-- 用户头像表 - 待定, 可上传至服务器中
CREATE TABLE sys_user_avatar (
    `id` INT AUTO_INCREMENT COMMENT '头像ID',
    `user_id` VARCHAR(10) NOT NULL COMMENT '用户ID',
    `data` LONGBLOB NULL COMMENT '头像数据',
    `hash` VARCHAR(50) NULL COMMENT '头像hash值',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户头像';

/* 待定
 - 部门 岗位 职级
 */
-- 用户登录表 - 待定，可在请求日志表中解析出来
CREATE TABLE sys_user_login (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` VARCHAR(10) NOT NULL COMMENT '用户ID',
    `token` VARCHAR(50) NOT NULL COMMENT 'Token 信息',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户登录表-用于登录';

/*
 日志相关表
 */
-- 网络请求日志表
CREATE TABLE sys_http_log (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT NULL COMMENT '请求用户ID',
    `parent_trace_id` VARCHAR(32) NULL COMMENT '上游请求traceId',
    `trace_id` VARCHAR(32) NULL COMMENT '请求traceId',
    `status_code` INT(10) NOT NULL COMMENT '请求状态码',
    `method` VARCHAR(10) NOT NULL COMMENT '请求方法',
    `path` VARCHAR(500) NOT NULL COMMENT '请求地址路径',
    `query` VARCHAR(500) NULL COMMENT '请求参数',
    `body` TEXT NULL COMMENT '请求体/响应体',
    `remote_addr` VARCHAR(64) NOT NULL COMMENT '请求IP',
    `user_agent` VARCHAR(256) NOT NULL COMMENT '用户代理',
    `cost` INT(20) NOT NULL COMMENT '耗时,纳秒',
    `htpp_type` VARCHAR(64) NOT NULL COMMENT '日志类型:REQ/RSP',
    `note` VARCHAR(255) NULL COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '网络请求日志';

-- 系统日志表
CREATE TABLE sys_system_log (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT NULL COMMENT '请求用户ID',
    `trace_id` VARCHAR(32) NULL COMMENT '请求traceId',
    `level` VARCHAR(10) NOT NULL COMMENT '日志级别',
    `caller_line` VARCHAR(100) NOT NULL COMMENT '日志发生位置',
    `error_code` INT(10) NULL COMMENT '业务错误码',
    `error_msg` VARCHAR(500) NOT NULL COMMENT '业务错误信息',
    `TEXT` VARCHAR(100) NULL COMMENT '日志消息',
    `extend` VARCHAR(500) NOT NULL COMMENT '日志扩展信息/json',
    `note` VARCHAR(255) NULL COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '系统日志';

-- 应用配置表
CREATE TABLE sys_config (
    `id` INT AUTO_INCREMENT COMMENT '配置ID',
    `parent_id` INT(11) DEFAULT NULL COMMENT '父节点ID',
    `name` VARCHAR(32) NOT NULL COMMENT '配置名称',
    `key` VARCHAR(32) NOT NULL UNIQUE COMMENT '配置参数(英文)',
    `value` TEXT COMMENT '配置参数值',
    `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
    `note` VARCHAR(32) DEFAULT NULL COMMENT '配置描述',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '应用配置表';