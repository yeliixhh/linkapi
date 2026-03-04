-- 用户表
drop table if exists sys_user;
create table sys_user (
                          id varchar(36) not null primary key default uuid_generate_v4(),
                          username varchar(32) not null,
                          password varchar(255) not null,
                          nickname varchar(255) not null,
                          is_admin int not null default 0,
                          created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
comment on table sys_user is '用户表';
comment on column sys_user.id is '主键id';
comment on column sys_user.username is '用户名';
comment on column sys_user.password is '密码';
comment on column sys_user.nickname is '用户名';
comment on column  sys_user.is_admin is '是否为管理员 0 为普通用户 1 为管理员';

create unique index idx_username on sys_user(username);


-- token 表
drop table if exists sys_token;
create table sys_token (
                id varchar(36) not null primary key default uuid_generate_v4(),
                user_id varchar(36) not null,
                token_type varchar(50),
                token text,
                is_revoked bool default false,
                expires_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
comment on table sys_token is '用户token表';
comment on column sys_token.token_type is 'token类型access_token与refresh_token';
create index idx_user_id on sys_token(user_id);
create index idx_token on sys_token(token);

