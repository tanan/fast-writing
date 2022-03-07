CREATE USER 'system' IDENTIFIED BY 'system';
GRANT ALL PRIVILEGES ON *.* TO 'system'@'%';
CREATE DATABASE fast_writing;

USE fast_writing;

create table user
(
  id binary(16) primary key default (UUID_TO_BIN(UUID())),
  name varchar(256) not null,
  email varchar(256) not null,
  last_updated datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
charset=utf8mb4;

create table writing_contents
(
  id bigint primary key AUTO_INCREMENT,
  user_id binary(16) not null,
  title varchar(256) not null,
  last_updated datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
charset=utf8mb4;

create table quiz
(
  id bigint primary key AUTO_INCREMENT,
  contents_id bigint not null,
  question varchar(256) not null,
  answer varchar(256) not null,
  last_updated datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
charset=utf8mb4;
