/*
 Navicat Premium Data Transfer

 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Schema         : lottery

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 02/04/2019 09:16:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lt_blackip
-- ----------------------------
DROP TABLE IF EXISTS `lt_blackip`;
CREATE TABLE `lt_blackip` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `blacktime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_blackip
-- ----------------------------
BEGIN;
INSERT INTO `lt_blackip` VALUES (1, '127.0.0.1', 1554194209, 0, 1553589409);
COMMIT;

-- ----------------------------
-- Table structure for lt_code
-- ----------------------------
DROP TABLE IF EXISTS `lt_code`;
CREATE TABLE `lt_code` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品ID，关联lt_gift表',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '虚拟券编码',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sys_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '状态，0正常，1作废，2已发放',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `gift_id` (`gift_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_code
-- ----------------------------
BEGIN;
INSERT INTO `lt_code` VALUES (1, 3, 'abc\r', 1532602694, 1553674427, 2);
INSERT INTO `lt_code` VALUES (2, 3, 'aa\r', 1532602694, 1553674435, 2);
INSERT INTO `lt_code` VALUES (3, 3, 'cs', 1532602694, 1553674437, 2);
INSERT INTO `lt_code` VALUES (4, 3, '332', 1532602970, 1553674438, 2);
INSERT INTO `lt_code` VALUES (6, 3, '123456', 1553588143, 1553674438, 2);
INSERT INTO `lt_code` VALUES (7, 3, '234567', 1553588143, 1553674778, 2);
INSERT INTO `lt_code` VALUES (9, 3, '123', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (10, 3, '1234', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (11, 3, '1235', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (12, 3, '12356', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (13, 3, '12311', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (14, 3, '123123123', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (15, 3, '123123', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (16, 3, '12332313', 1553842311, 0, 0);
INSERT INTO `lt_code` VALUES (18, 3, '111', 1553843964, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for lt_gift
-- ----------------------------
DROP TABLE IF EXISTS `lt_gift`;
CREATE TABLE `lt_gift` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `prize_num` int(11) NOT NULL DEFAULT '-1' COMMENT '奖品数量，0 无限量，>0限量，<0无奖品',
  `left_num` int(11) NOT NULL DEFAULT '0' COMMENT '剩余数量',
  `prize_code` varchar(50) NOT NULL DEFAULT '' COMMENT '0-9999表示100%，0-0表示万分之一的中奖概率',
  `prize_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发奖周期，D天',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图片',
  `displayorder` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '位置序号，小的排在前面',
  `gtype` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，0 虚拟币，1 虚拟券，2 实物-小奖，3 实物-大奖',
  `gdata` varchar(255) NOT NULL DEFAULT '' COMMENT '扩展数据，如：虚拟币数量',
  `time_begin` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `time_end` int(11) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `prize_data` mediumtext COMMENT '发奖计划，[[时间1,数量1],[时间2,数量2]]',
  `prize_begin` int(11) NOT NULL DEFAULT '0' COMMENT '发奖计划周期的开始',
  `prize_end` int(11) NOT NULL DEFAULT '0' COMMENT '发奖计划周期的结束',
  `sys_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '状态，0 正常，1 删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_gift
-- ----------------------------
BEGIN;
INSERT INTO `lt_gift` VALUES (1, 'T恤', 10, 98, '1-100', 30, 'https://p0.ssl.qhmsg.com/t016c44d161c478cfe0.png', 1, 2, '', 1532592420, 1564128420, '[[1554610050,1],[1554734790,1],[1555032210,1],[1555074270,1],[1555777950,1],[1556024670,1],[1556044050,1],[1556076210,1],[1556262030,1]]', 1553842230, 1556434230, 0, 1532592429, 1553842230, '::1');
INSERT INTO `lt_gift` VALUES (2, '手机N7', 2, 100, '0-0', 30, 'https://p0.ssl.qhmsg.com/t016ff98b934914aca6.png', 0, 3, '', 1532592420, 1564128420, '[[1556074110,1]]', 1553842230, 1556434230, 0, 1532592474, 1553842230, '');
INSERT INTO `lt_gift` VALUES (3, '手机充电器', 2000, 0, '200-1000', 1, 'https://p0.ssl.qhmsg.com/t01ec4648d396ad46bf.png', 3, 2, '', 1532592420, 1564128420, '', 1553846517, 1553932917, 0, 1532592558, 1553850422, '::1');
INSERT INTO `lt_gift` VALUES (4, '优惠券', 100, 99, '2000-5000', 1, 'https://p0.ssl.qhmsg.com/t01f84f00d294279957.png', 4, 1, '', 1532592420, 1564128420, '[[1554100660,1],[1554102280,1],[1554102820,1],[1554103240,1],[1554104680,1],[1554104740,1],[1554104800,1],[1554105400,1],[1554105700,1],[1554106240,1],[1554106360,1],[1554106660,1],[1554106840,1],[1554107020,1],[1554107740,1],[1554108220,1],[1554108520,1],[1554108760,2],[1554108880,1],[1554109420,1],[1554109540,1],[1554109660,1],[1554109960,1]]', 1554023620, 1554110020, 0, 1532599140, 1554023620, '::1');
COMMIT;

-- ----------------------------
-- Table structure for lt_result
-- ----------------------------
DROP TABLE IF EXISTS `lt_result`;
CREATE TABLE `lt_result` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品ID，关联lt_gift表',
  `gift_name` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `gift_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，同lt_gift. gtype',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `prize_code` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '抽奖编号（4位的随机数）',
  `gift_data` varchar(255) NOT NULL DEFAULT '' COMMENT '获奖信息',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '用户抽奖的IP',
  `sys_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '状态，0 正常，1删除，2作弊',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `gift_id` (`gift_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_result
-- ----------------------------
BEGIN;
INSERT INTO `lt_result` VALUES (1, 1, 'T恤', 2, 1, 'yifan', 1, '', 0, '', 0);
INSERT INTO `lt_result` VALUES (2, 3, '手机充电器', 2, 31280, 'admin-31280', 915, 'aa\r', 1553673796, '::1', 0);
INSERT INTO `lt_result` VALUES (3, 3, '手机充电器', 2, 31280, 'admin-31280', 963, '234567', 1553673855, '::1', 0);
INSERT INTO `lt_result` VALUES (4, 3, '手机充电器', 2, 31280, 'admin-31280', 904, 'abc\r', 1553674014, '::1', 0);
INSERT INTO `lt_result` VALUES (5, 3, '手机充电器', 2, 31280, 'admin-31280', 327, '123456', 1553674195, '::1', 0);
INSERT INTO `lt_result` VALUES (6, 3, '手机充电器', 2, 31280, 'admin-31280', 506, '234567', 1553674310, '::1', 0);
INSERT INTO `lt_result` VALUES (7, 3, '手机充电器', 2, 31280, 'admin-31280', 286, 'abc\r', 1553674356, '::1', 0);
INSERT INTO `lt_result` VALUES (8, 3, '手机充电器', 2, 31280, 'admin-31280', 984, 'aa\r', 1553674362, '::1', 0);
INSERT INTO `lt_result` VALUES (9, 3, '手机充电器', 2, 31280, 'admin-31280', 836, 'cs', 1553674366, '::1', 0);
INSERT INTO `lt_result` VALUES (10, 3, '手机充电器', 2, 31280, 'admin-31280', 927, '332', 1553674367, '::1', 0);
INSERT INTO `lt_result` VALUES (11, 3, '手机充电器', 2, 31280, 'admin-31280', 732, '123456', 1553674368, '::1', 0);
INSERT INTO `lt_result` VALUES (12, 3, '手机充电器', 2, 31280, 'admin-31280', 542, '234567', 1553674369, '::1', 0);
INSERT INTO `lt_result` VALUES (13, 3, '手机充电器', 2, 31280, 'admin-31280', 904, 'abc\r', 1553674427, '::1', 0);
INSERT INTO `lt_result` VALUES (14, 3, '手机充电器', 2, 31280, 'admin-31280', 739, 'aa\r', 1553674435, '::1', 0);
INSERT INTO `lt_result` VALUES (15, 3, '手机充电器', 2, 31280, 'admin-31280', 265, 'cs', 1553674437, '::1', 0);
INSERT INTO `lt_result` VALUES (16, 3, '手机充电器', 2, 31280, 'admin-31280', 915, '332', 1553674438, '::1', 0);
INSERT INTO `lt_result` VALUES (17, 3, '手机充电器', 2, 31280, 'admin-31280', 329, '123456', 1553674438, '::1', 0);
INSERT INTO `lt_result` VALUES (18, 3, '手机充电器', 2, 31280, 'admin-31280', 318, '234567', 1553674781, '::1', 0);
INSERT INTO `lt_result` VALUES (19, 4, '优惠券', 1, 31280, 'admin-31280', 2525, '', 1553678344, '::1', 0);
COMMIT;

-- ----------------------------
-- Table structure for lt_user
-- ----------------------------
DROP TABLE IF EXISTS `lt_user`;
CREATE TABLE `lt_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `blacktime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_user
-- ----------------------------
BEGIN;
INSERT INTO `lt_user` VALUES (1, 'wangyi', 0, '一凡Sir', '11111111111', 'abcdefg', 0, 1532595094, '');
COMMIT;

-- ----------------------------
-- Table structure for lt_userday
-- ----------------------------
DROP TABLE IF EXISTS `lt_userday`;
CREATE TABLE `lt_userday` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `day` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '日期，如：20180725',
  `num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '次数',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid_day` (`uid`,`day`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lt_userday
-- ----------------------------
BEGIN;
INSERT INTO `lt_userday` VALUES (1, 85084, 20190327, 3, 1553671872, 0);
INSERT INTO `lt_userday` VALUES (2, 50506, 20190327, 3, 1553671925, 0);
INSERT INTO `lt_userday` VALUES (3, 95957, 20190327, 10, 1553672173, 0);
INSERT INTO `lt_userday` VALUES (4, 30961, 20190327, 10, 1553672754, 0);
INSERT INTO `lt_userday` VALUES (5, 94636, 20190327, 26, 1553673060, 0);
INSERT INTO `lt_userday` VALUES (6, 31280, 20190327, 49, 1553673601, 0);
INSERT INTO `lt_userday` VALUES (7, 31280, 20190401, 3, 1554086450, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
