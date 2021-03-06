# create database `open_period_cards` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE `tb_admin` (
                         `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                         `phone` varchar(11) NOT NULL DEFAULT '',
                         `name` varchar(30) NOT NULL DEFAULT '',
                         `pwd` varchar(32) NOT NULL DEFAULT '',
                         `salt` varchar(32) NOT NULL DEFAULT '',
                         `mtime` int(11) NOT NULL DEFAULT '0',
                         `ctime` int(11) NOT NULL DEFAULT '0',
                         `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment 'cacheKey|phone,Name=name';
##########
CREATE TABLE `tb_merchant` (
                            `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                            `phone` varchar(11) NOT NULL DEFAULT '',
                            `shop_name` varchar(60) NOT NULL DEFAULT '',
                            `industry_name` varchar(30) NOT NULL DEFAULT '',
                            `pwd` varchar(32) NOT NULL DEFAULT '',
                            `salt` varchar(32) NOT NULL DEFAULT '',
                            `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '0-applied 1-passed 2-denied',
                            `mtime` int(11) NOT NULL DEFAULT '0',
                            `ctime` int(11) NOT NULL DEFAULT '0',
                            `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment 'cacheKey=phone';
##########

CREATE TABLE `tb_user` (
                               `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                               `phone` varchar(11) NOT NULL DEFAULT '',
                               `name` varchar(60) NOT NULL DEFAULT '',
                               `pwd` varchar(32) NOT NULL DEFAULT '',
                               `salt` varchar(32) NOT NULL DEFAULT '',
                               `mtime` int(11) NOT NULL DEFAULT '0',
                               `ctime` int(11) NOT NULL DEFAULT '0',
                               `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment 'cacheKey=phone';
##########

CREATE TABLE `tb_card_type` (
                           `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                           `merchant_id` int(11) NOT NULL DEFAULT '0',
                           `type` tinyint(3) NOT NULL DEFAULT '0' comment '0-day 1-week 2-month',
                           `period_times` tinyint(3) NOT NULL DEFAULT '0' comment '????????????????????????',
                           `total_times` tinyint(3) NOT NULL DEFAULT '0' comment '??????????????????',
                           `describe` varchar(60) NOT NULL DEFAULT '',
                           `mtime` int(11) NOT NULL DEFAULT '0',
                           `ctime` int(11) NOT NULL DEFAULT '0',
                           `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
##########
CREATE TABLE `tb_card_type_info_tpl` (
                                         `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                                         `name` varchar(60) NOT NULL DEFAULT '',
                                         `card_type_id` int(11) NOT NULL DEFAULT '0',
                                         `tpl` text NOT NULL,
                                         `describe` varchar(60) NOT NULL DEFAULT '',
                                         `mtime` int(11) NOT NULL DEFAULT '0',
                                         `ctime` int(11) NOT NULL DEFAULT '0',
                                         `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '???????????????????????????';
##########

CREATE TABLE `tb_card` (
                                `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                                `merchant_id` int(11) NOT NULL DEFAULT '0',
                                `user_id` int(11) NOT NULL DEFAULT '0',
                                `name` varchar(60) NOT NULL DEFAULT '',
                                `card_type_id` int(11) NOT NULL DEFAULT '0',
                                `serial_code` varchar(30) NOT NULL DEFAULT '',
                                `used` tinyint(3) NOT NULL DEFAULT '0' comment '0-not 1-yes',
                                `describe` varchar(60) NOT NULL DEFAULT '',
                                `mtime` int(11) NOT NULL DEFAULT '0',
                                `ctime` int(11) NOT NULL DEFAULT '0',
                                `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
##########



CREATE TABLE `tb_card_order` (
                           `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                           `merchant_id` int(11) NOT NULL DEFAULT '0',
                           `user_id` int(11) NOT NULL DEFAULT '0',
                           `card_id` int(11) NOT NULL DEFAULT '0' comment 'card_id ?????????0???????????????????????????????????????????????????',
                           `card_type_id` int(11) NOT NULL DEFAULT '0' comment '????????????',
                           `send_type` tinyint(3) NOT NULL DEFAULT '0' comment '??????',
                           `send_day_list` varchar(60) NOT NULL DEFAULT '0' comment '????????????????????????',
                           `period_send_times` int(11) NOT NULL DEFAULT '0' comment '???????????????????????????',
                           `total_send_times` int(11) NOT NULL DEFAULT '0' comment '?????????????????????',
                           `is_total_finished` tinyint(3) NOT NULL DEFAULT '0' comment '??????????????????????????????0-??? 1-???',
                           `is_period_finished` tinyint(3) NOT NULL DEFAULT '0' comment '?????????????????????????????????0-??? 1-???',
                           `last_send_time` int(11) NOT NULL DEFAULT '0' comment '??????????????????',
                           `describe` varchar(60) NOT NULL DEFAULT '',
                           `mtime` int(11) NOT NULL DEFAULT '0',
                           `ctime` int(11) NOT NULL DEFAULT '0',
                           `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
##########
CREATE TABLE `tb_card_order_info` (
                                 `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                                 `card_order_id` int(11) NOT NULL DEFAULT '0',
                                 `content` text NOT NULL,
                                 `describe` varchar(60) NOT NULL DEFAULT '',
                                 `mtime` int(11) NOT NULL DEFAULT '0',
                                 `ctime` int(11) NOT NULL DEFAULT '0',
                                 `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '???????????????????????????';
##########
CREATE TABLE `tb_card_order_delivery_log` (
                                      `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                                      `card_order_id` int(11) NOT NULL DEFAULT '0',
                                      `content` text NOT NULL,
                                      `describe` varchar(60) NOT NULL DEFAULT '',
                                      `mtime` int(11) NOT NULL DEFAULT '0',
                                      `ctime` int(11) NOT NULL DEFAULT '0',
                                      `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '???????????????????????????';
##########
CREATE TABLE `tb_card_order_change_log` (
                                              `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
                                              `card_order_id` int(11) NOT NULL DEFAULT '0',
                                              `change_log` text NOT NULL,
                                              `describe` varchar(60) NOT NULL DEFAULT '',
                                              `mtime` int(11) NOT NULL DEFAULT '0',
                                              `ctime` int(11) NOT NULL DEFAULT '0',
                                              `deleted` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '??????????????????????????????';
