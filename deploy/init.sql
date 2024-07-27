-- 创建数据库 db
CREATE DATABASE IF NOT EXISTS db;

-- 使用数据库 db
USE db;

-- 创建 users 表
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) ;

-- 插入一行初始数据
INSERT INTO users (username, password) VALUES ('admin', '$2a$10$ze.d0x2I8BuqrxMqoTPy8enhspVsHaMwivu/.kyoCZPSYp9/BKEvK');
