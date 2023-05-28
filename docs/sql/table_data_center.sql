/*数据中心*/
-- 字典维度表
CREATE TABLE dc_dict (
    `id` INT AUTO_INCREMENT COMMENT '字典ID',
    `name` VARCHAR(50) NOT NULL COMMENT '字典名称',
    `code` VARCHAR(50) NOT NULL COMMENT '字典编码',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '字典维度表';

-- 字典数据表
CREATE TABLE dc_dict_data (
    `id` INT AUTO_INCREMENT COMMENT '字典项ID',
    `dict_id` INT(11) NOT NULL COMMENT '字典维度ID',
    `name` VARCHAR(50) NOT NULL COMMENT '字典项名称',
    `value` VARCHAR(50) NOT NULL COMMENT '字典项值',
    `note` VARCHAR(200) NULL COMMENT '备注',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用,0:禁用,1:启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    CONSTRAINT `dc_dict_data_dict_id` FOREIGN KEY (`dict_id`) REFERENCES `dc_dict` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '字典数据表';