/* API授权相关的表 */
-- Http协议接口管理表
CREATE TABLE sys_http_api (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `name` VARCHAR(50) NOT NULL COMMENT '接口名称',
    `method` VARCHAR(50) NOT NULL COMMENT '请求类型',
    `uri` VARCHAR(50) NOT NULL COMMENT 'URI',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'Http协议接口管理表';

-- 用户Token令牌表
CREATE TABLE sys_user_token (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `user_id` INT(20) NOT NULL COMMENT '用户ID',
    `token` VARCHAR(50) NOT NULL COMMENT 'Token信息',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_sys_user_token_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '用户Token令牌表';

-- Token令牌Http协议API接口鉴权表
CREATE TABLE sys_token_http_api_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `token_id` INT(11) NOT NULL COMMENT 'Token令牌ID',
    `http_api_id` INT(11) NOT NULL COMMENT 'Http协议接口ID',
    `expire` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权到期时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`token_id`, `http_api_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'Token令牌Http协议API接口鉴权表';

-- 角色与Http协议接口关联表
CREATE TABLE sys_role_http_api_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `role_id` INT(11) NOT NULL COMMENT '角色ID',
    `http_api_id` INT(11) NOT NULL COMMENT 'Http协议接口ID',
    `expire` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权到期时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`token_id`, `http_api_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '角色与Http协议接口关联表';