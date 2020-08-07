/*
 Navicat Premium Data Transfer

 Source Server         : dev_
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : goApi

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 07/08/2020 14:41:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info` (
  `rid` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(64) DEFAULT NULL,
  `avatar` varchar(128) DEFAULT NULL,
  `gender` int(2) NOT NULL,
  `birthday` varchar(64) DEFAULT NULL,
  `height` int(4) DEFAULT NULL,
  `weight` int(4) DEFAULT NULL,
  `install` varchar(64) DEFAULT NULL COMMENT '安装来源',
  `login_type` int(2) NOT NULL COMMENT '第三方登录方式',
  `sub` varchar(64) NOT NULL COMMENT '第三方登录 id',
  `userid` varchar(32) NOT NULL COMMENT 'app的userid',
  `email` varchar(255) DEFAULT '' COMMENT '用户的邮箱',
  `ccid` varchar(64) DEFAULT NULL COMMENT '专属客服id',
  `promote_user` varchar(32) DEFAULT NULL COMMENT '邀请人id',
  `is_customer_care` int(2) NOT NULL DEFAULT '0' COMMENT '是否是客服',
  `status` int(2) DEFAULT '1',
  `ctime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`rid`),
  UNIQUE KEY `idx_uid` (`userid`) USING BTREE,
  KEY `idx_sub` (`sub`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

