CREATE USER 'contoro' IDENTIFIED BY 'contoro';
GRANT ALL PRIVILEGES ON *.* TO 'contoro'@'%';
CREATE DATABASE contract;

USE contract;

create table contract
(
  id binary(16) primary key default (UUID_TO_BIN(UUID(), TRUE)),
  name varchar(256) not null,
  description varchar(256) null,
  note varchar(256) null,
  registration_time datetime default CURRENT_TIMESTAMP,
  update_time datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  start_date datetime not null,
  end_date datetime not null
)
charset=utf8mb4;

create table category
(
  id bigint not null AUTO_INCREMENT primary key,
  name varchar(256) unique not null
)
charset=utf8mb4;

create table contract_category
(
  id binary(16) not null,
  category_id bigint not null,
  primary key (id, category_id)
)
charset=utf8mb4;

create table editor
(
  id binary(16) primary key default (UUID_TO_BIN(UUID(), TRUE)),
  name varchar(256) unique not null
)
charset=utf8mb4;

create table contract_editor
(
  id binary(16) not null,
  editor_id binary(16) not null,
  primary key (id, editor_id)
)
charset=utf8mb4;


create table contractor
(
  id binary(16) primary key default (UUID_TO_BIN(UUID(), TRUE)),
  name varchar(256) unique not null
)
charset=utf8mb4;

create table contract_contractor
(
  id binary(16) not null,
  contractor_id binary(16) not null,
  primary key (id, contractor_id)
)
charset=utf8mb4;
