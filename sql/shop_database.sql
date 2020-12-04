drop database if exists shop;
drop user if exists 'shop'@'%';
-- 支持emoji：需要mysql数据库参数： character_set_server=utf8mb4
create database shop default character set utf8mb4 collate utf8mb4_unicode_ci;
use shop;
create user 'shop'@'%' identified by 'Shop12#$';
grant all privileges on shop.* to 'shop'@'%';
flush privileges;