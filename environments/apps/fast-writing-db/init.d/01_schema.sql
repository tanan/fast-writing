CREATE USER 'system' IDENTIFIED BY 'system';
GRANT ALL PRIVILEGES ON *.* TO 'system'@'%';
CREATE DATABASE fast_writing;

USE fast_writing;

create table user
(
  id varchar(32) primary key,
  name varchar(256) not null,
  email varchar(256) not null,
  last_updated datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
)
charset=utf8mb4;

create table writing_contents
(
  id bigint primary key AUTO_INCREMENT,
  user_id varchar(32) not null,
  title varchar(256) not null,
  description varchar(256) not null,
  scope varchar(16) not null,
  last_updated datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  quiz json
)
charset=utf8mb4;
