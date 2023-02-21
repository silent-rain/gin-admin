/*API授权相关的表 */
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

-- API_Token_URI授权表,基于Token授权
CREATE TABLE api_token_uri_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `api_token_id` INT(11) NOT NULL COMMENT 'API-Token令牌ID',
    `uri` VARCHAR(200) NOT NULL COMMENT '请求地址路径',
    `expire` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权到期时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_api_token_uri_auth_api_token_id_uri` (`api_token_id`, `uri`),
    CONSTRAINT `api_token_uri_auth_api_token_id` FOREIGN KEY (`api_token_id`) REFERENCES `api_token` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'API_Token_URI授权表';

-- api_token_role API_Token角色关联表,基于角色授权
-- api_uri_role API_URI角色关联表