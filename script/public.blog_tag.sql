CREATE TABLE blog_tag
(
    id          SERIAL PRIMARY KEY,
    name        varchar(100),
    created_on  timestamp,                --COMMENT '创建时间'
    created_by  varchar(100) DEFAULT '',  --COMMENT '创建人'
    modified_on timestamp,                --COMMENT '修改时间'
    modified_by varchar(100) DEFAULT '',  --COMMENT '修改人'
    deleted_on  timestamp,                --COMMENT '删除时间'
    is_del      smallint     DEFAULT '0', --COMMENT '是否删除 0 为未删除、1 为已删除'
    state       smallint     DEFAULT '1'  --COMMENT '状态 0 为禁用、1 为启用'
);

CREATE TABLE blog_article
(
    id              SERIAL PRIMARY KEY,
    title           varchar(100) DEFAULT '',  --COMMENT '文章标题'
    description     varchar(255) DEFAULT '',  -- COMMENT '文章简述'
    cover_image_url varchar(255) DEFAULT '',--COMMENT '封面图片地址'
    content         text,                     --COMMENT '文章内容'
    created_on      timestamp,                --COMMENT '创建时间'
    created_by      varchar(100) DEFAULT '',  --COMMENT '创建人'
    modified_on     timestamp,                --COMMENT '修改时间'
    modified_by     varchar(100) DEFAULT '',  --COMMENT '修改人'
    deleted_on      timestamp,                --COMMENT '删除时间'
    is_del          smallint     DEFAULT '0', --COMMENT '是否删除 0 为未删除、1 为已删除'
    state           smallint     DEFAULT '1'  --COMMENT '状态 0 为禁用、1 为启用'
);

CREATE TABLE blog_article_tag
(
    id          SERIAL PRIMARY KEY,
    article_id  integer NOT NULL,             --COMMENT '文章 ID'
    tag_id      integer NOT NULL DEFAULT '0', --COMMENT '标签 ID'
    created_on  timestamp,                    --COMMENT '创建时间'
    created_by  varchar(100)     DEFAULT '',  --COMMENT '创建人'
    modified_on timestamp,                    --COMMENT '修改时间'
    modified_by varchar(100)     DEFAULT '',  --COMMENT '修改人'
    deleted_on  timestamp,                    --COMMENT '删除时间'
    is_del      smallint         DEFAULT '0'  --COMMENT '是否删除 0 为未删除、1 为已删除'
);

CREATE TABLE blog_auth
(
    id          SERIAL PRIMARY KEY,
    app_key     varchar(20)  DEFAULT '',
    app_secret  varchar(50)  DEFAULT '',
    created_on  timestamp,               --COMMENT '创建时间'
    created_by  varchar(100) DEFAULT '', --COMMENT '创建人'
    modified_on timestamp,               --COMMENT '修改时间'
    modified_by varchar(100) DEFAULT '', --COMMENT '修改人'
    deleted_on  timestamp,               --COMMENT '删除时间'
    is_del      smallint     DEFAULT '0' --COMMENT '是否删除 0 为未删除、1 为已删除'
);
INSERT INTO blog_auth
(id, app_key, app_secret, created_on, created_by, modified_on, modified_by, deleted_on, is_del)
VALUES (1, 'xuhy', '123456', (select (CURRENT_TIMESTAMP)), 'xuhy', NULL, '', NULL, 0);