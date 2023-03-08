/* API授权相关的表 */
-- Http协议接口管理表
CREATE TABLE api_http (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `name` VARCHAR(50) NOT NULL COMMENT '接口名称',
    `method` VARCHAR(50) NOT NULL COMMENT '请求类型',
    `uri` VARCHAR(50) NOT NULL COMMENT 'URI资源',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'Http协议接口管理表';

-- 角色与Http协议接口关联表
CREATE TABLE api_role_http_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `role_id` INT(11) NOT NULL COMMENT '角色ID',
    `http_api_id` INT(11) NOT NULL COMMENT 'Http协议接口ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`token_id`, `http_api_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '角色与Http协议接口关联表';

-- Token令牌Http协议API接口鉴权表, 用户主动申请接口，待定
CREATE TABLE api_token_http_auth (
    `id` INT AUTO_INCREMENT COMMENT '自增ID',
    `token_id` INT(11) NOT NULL COMMENT 'Token令牌ID',
    `http_api_id` INT(11) NOT NULL COMMENT 'Http协议接口ID',
    `expire` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权到期时间',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`token_id`, `http_api_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT 'Token令牌Http协议API接口鉴权表';