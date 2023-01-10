-- 创建数据库 
CREATE DATABASE `gin_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

/*
 用户相关的表
 */
-- 用户表
CREATE TABLE user (
    `id` INT AUTO_INCREMENT COMMENT '用户ID',
    `realname` VARCHAR(32) NULL COMMENT '真实姓名',
    `nickname` VARCHAR(32) NOT NULL COMMENT '昵称',
    `gender` TINYINT(1) NULL COMMENT '性别: 0:女,1:男',
    `age` INT(11) NULL COMMENT '年龄',
    `birthday` VARCHAR(20) NULL COMMENT '出生日期',
    `avatar` VARCHAR(50) NULL COMMENT '用户头像URL',
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

-- 用户位置 - 待定
CREATE TABLE user_location (
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
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户位置';

-- 用户头像表 - 待定
CREATE TABLE user_avatar (
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
-- 用户登录Token表
CREATE TABLE user_login_token (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` VARCHAR(10) NOT NULL COMMENT '用户ID',
    `token` VARCHAR(50) NOT NULL COMMENT 'Token 信息',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户登录Token表-用于登录';

/*
 网络相关的表
 */
-- 网络请求日志表
CREATE TABLE http_log (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT NULL COMMENT '请求用户ID',
    `trace_id` VARCHAR(32) NULL COMMENT '请求traceId',
    `status_code` INT(10) NOT NULL COMMENT '请求状态码',
    `method` VARCHAR(10) NOT NULL COMMENT '请求方法',
    `path` VARCHAR(500) NOT NULL COMMENT '请求地址路径',
    `query` VARCHAR(500) NULL COMMENT '请求参数',
    `body` VARCHAR(500) NULL COMMENT '请求体/响应体',
    `remote_addr` VARCHAR(64) NOT NULL COMMENT '请求IP',
    `user_agent` VARCHAR(100) NOT NULL COMMENT '用户代理',
    `cost` INT(20) NOT NULL COMMENT '耗时,纳秒',
    `htpp_type` VARCHAR(64) NOT NULL COMMENT '日志类型:REQ/RSP',
    `note` VARCHAR(255) NULL COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '网络请求日志';

/*
 授权相关的表
 */
-- 用户Token信息表
CREATE TABLE user_token (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT(20) NOT NULL COMMENT '用户ID',
    `token` VARCHAR(50) NOT NULL COMMENT 'Token 信息',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_token_uni_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户Token信息表-用于授权';

-- 用户Token的API授权
CREATE TABLE user_token_api_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_token_id` INT(11) NOT NULL COMMENT '用户Token信息表ID',
    `uri` VARCHAR(200) NOT NULL COMMENT '请求地址路径',
    `expire` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权到期时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_token_api_auth_uni_user_token_id` (`user_token_id`, `uri`),
    CONSTRAINT `user_token_api_auth_user_token_id` FOREIGN KEY (`user_token_id`) REFERENCES `user_token` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户Token的API授权';