/*
 Navicat MySQL Data Transfer

 Source Server         : 00
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : go

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 12/09/2021 20:43:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for person
-- ----------------------------
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person`  (
  `uid` int(10) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NULL DEFAULT NULL,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `department` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of person
-- ----------------------------
INSERT INTO `person` VALUES (1, '2019-07-06 11:45:20', 'lktbz', '123456', '技术部', '123456@163.com');
INSERT INTO `person` VALUES (2, '2019-07-06 11:45:20', 'johny', '123456', '技术部', '12@163.com');
INSERT INTO `person` VALUES (3, '2019-07-06 11:45:20', 'johnys', '123456', '技术部', '56@163.com');
INSERT INTO `person` VALUES (4, '2019-07-06 11:45:20', 'lk', '123456', '技术部', '6@163.com');
INSERT INTO `person` VALUES (5, '2019-07-06 11:45:20', 'tom', '123456', '技术部', '1@163.com');
INSERT INTO `person` VALUES (6, '2019-07-06 11:45:20', 'jack', '123456', '技术部', '2@163.com');
INSERT INTO `person` VALUES (7, '2019-07-06 11:45:20', 'mic', '123456', '技术部', '456@163.com');
INSERT INTO `person` VALUES (8, '2019-07-06 11:45:20', 'james', '123456', '技术部', '34@163.com');
INSERT INTO `person` VALUES (9, '2019-07-06 11:45:20', 'gyus', '123456', '技术部', '3@163.com');
INSERT INTO `person` VALUES (10, '2019-07-06 11:45:20', 'fres', '123456', '技术部', '12@163.com');
INSERT INTO `person` VALUES (11, '2019-07-06 11:45:20', 'shd', '123456', '技术部', '1451456@163.com');
INSERT INTO `person` VALUES (12, '2019-07-06 11:45:20', 'keys', '123456', '技术部', '1qqq@163.com');

SET FOREIGN_KEY_CHECKS = 1;
