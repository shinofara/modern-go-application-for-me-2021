create database if not exists odai character set utf8mb4 collate utf8mb4_bin;
use odai;

drop table if exists wallets;
create table if not exists wallets
(
  id      int unsigned not null primary key auto_increment,
  balance int not null
) character set utf8mb4 collate utf8mb4_bin;


drop table if exists stocks;
create table if not exists stocks
(
  id          int unsigned not null primary key auto_increment,
  item_id     int not null,
  amount      int          not null
) character set utf8mb4 collate utf8mb4_bin;

insert into wallets (balance)
values (20000);

insert into stocks (item_id, amount)
values (1, 100),
(1, 100),
(1, 100),
(2, 200),
(3, 1000),
(3, 1000),
(3, 1000),
(3, 1000),
(3, 1000),
(3, 1000),
(4, 10),
(4, 10);
