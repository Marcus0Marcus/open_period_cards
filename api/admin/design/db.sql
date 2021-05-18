# create database `open_period_cards` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE `tb_admin` (
                         `id` int(11) NOT NULL,
                         `phone` varchar(11) NOT NULL DEFAULT '',
                         `name` varchar(30) NOT NULL DEFAULT '',
                         `pwd` varchar(32) NOT NULL DEFAULT '',
                         `mtime` int(11) NOT NULL DEFAULT '0',
                         `ctime` int(11) NOT NULL DEFAULT '0',
                         `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `tb_merchant` (
                            `id` int(11) NOT NULL,
                            `phone` varchar(11) NOT NULL DEFAULT '',
                            `shop_name` varchar(60) NOT NULL DEFAULT '',
                            `pwd` varchar(32) NOT NULL DEFAULT '',
                            `mtime` int(11) NOT NULL DEFAULT '0',
                            `ctime` int(11) NOT NULL DEFAULT '0',
                            `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `tb_user` (
                               `id` int(11) NOT NULL,
                               `phone` varchar(11) NOT NULL DEFAULT '',
                               `name` varchar(60) NOT NULL DEFAULT '',
                               `pwd` varchar(32) NOT NULL DEFAULT '',
                               `mtime` int(11) NOT NULL DEFAULT '0',
                               `ctime` int(11) NOT NULL DEFAULT '0',
                               `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `tb_card_type` (
                           `id` int(11) NOT NULL,
                           `type` tinyint(3) NOT NULL DEFAULT '0' comment '0-day 1-week 2-month 3-year',
                           `times` tinyint(3) NOT NULL DEFAULT '0',
                           `describe` varchar(60) NOT NULL DEFAULT '',
                           `mtime` int(11) NOT NULL DEFAULT '0',
                           `ctime` int(11) NOT NULL DEFAULT '0',
                           `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE `tb_card` (
                                `id` int(11) NOT NULL,
                                `name` varchar(60) NOT NULL DEFAULT '',
                                `card_type_id` int(11) NOT NULL DEFAULT '0',
                                `serial_code` varchar(30) NOT NULL DEFAULT '',
                                `used` tinyint(3) NOT NULL DEFAULT '0' comment '0-not 1-yes',
                                `describe` varchar(60) NOT NULL DEFAULT '',
                                `mtime` int(11) NOT NULL DEFAULT '0',
                                `ctime` int(11) NOT NULL DEFAULT '0',
                                `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
