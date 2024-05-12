DROP DATABASE IF EXISTS sample;
CREATE DATABASE sample;

DROP USER IF EXISTS 'upd_user';
CREATE USER 'upd_user'@'%' IDENTIFIED BY 'upd_user';
GRANT ALL PRIVILEGES ON sample.* TO 'upd_user'@'%';

DROP USER IF EXISTS 'ref_user';
CREATE USER 'ref_user'@'%' IDENTIFIED BY 'ref_user';
GRANT SELECT ON sample.* TO 'ref_user'@'%';

USE sample;

