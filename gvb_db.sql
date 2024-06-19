/*
 Navicat Premium Data Transfer

 Source Server         : localhost_mysql
 Source Server Type    : MySQL
 Source Server Version : 80020 (8.0.20)
 Source Host           : localhost:3306
 Source Schema         : gvb_db

 Target Server Type    : MySQL
 Target Server Version : 80020 (8.0.20)
 File Encoding         : 65001

 Date: 08/05/2024 11:23:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for advert_models
-- ----------------------------
DROP TABLE IF EXISTS `advert_models`;
CREATE TABLE `advert_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `href` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `image` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_show` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of advert_models
-- ----------------------------
INSERT INTO `advert_models` VALUES (1, '2024-02-28 21:58:29.780', '2024-02-29 10:58:14.015', '小米su7', 'https://www.xiaomiev.com/', 'http://qiniu.xiaoxinqj.top/gvb/20240228215501_小米su7.jpg', 1);
INSERT INTO `advert_models` VALUES (3, '2024-02-29 15:38:45.115', '2024-02-29 15:38:45.115', '比亚迪秦plus', 'https://bydauto.com.cn/pc/configCar?id=109&networkType=dynasty', 'http://qiniu.xiaoxinqj.top/gvb/20240228220025_比亚迪秦plus.jpg', 0);

-- ----------------------------
-- Table structure for banner_models
-- ----------------------------
DROP TABLE IF EXISTS `banner_models`;
CREATE TABLE `banner_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '\'图片路径\'',
  `hash` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '\'图片哈希\'',
  `name` varchar(38) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'图片名称\'',
  `image_type` bigint NULL DEFAULT 1 COMMENT '\'图片类型\'',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of banner_models
-- ----------------------------
INSERT INTO `banner_models` VALUES (4, '2024-02-27 16:52:44.239', '2024-02-27 16:52:44.239', 'http://qiniu.xiaoxinqj.top/gvb/20240227165243_a4.jpg', '7822a6008d890a2357aff37955478d3b', 'a4.jpg', 2);
INSERT INTO `banner_models` VALUES (5, '2024-02-27 16:52:44.323', '2024-02-27 16:52:44.323', 'http://qiniu.xiaoxinqj.top/gvb/20240227165244_a5.jpg', '15365249953e302b36f539e29f37f9aa', 'a5.jpg', 2);
INSERT INTO `banner_models` VALUES (6, '2024-02-28 16:36:48.344', '2024-02-28 16:36:48.344', 'http://qiniu.xiaoxinqj.top/gvb/20240228163647_a7.jpg', '1a3181ca1442d8954b9b7bd20a5e26f6', 'a7.jpg', 2);
INSERT INTO `banner_models` VALUES (7, '2024-02-28 16:39:38.622', '2024-02-28 16:39:38.622', 'https://qiniu.xiaoxinqj.top/gvb/20240228163938_a1.jpg', '4c078e4b389cd9ec26c358904cdf2010', 'a1.jpg', 2);
INSERT INTO `banner_models` VALUES (8, '2024-02-28 16:39:38.631', '2024-02-28 16:39:38.631', 'https://qiniu.xiaoxinqj.top/gvb/20240228163938_a3.jpg', '9a028a90f521d3903e8791a58b0e1da3', 'a3.jpg', 2);
INSERT INTO `banner_models` VALUES (9, '2024-02-28 17:30:59.861', '2024-02-28 17:30:59.861', 'https://qiniu.xiaoxinqj.top/gvb/20240228173059_google.png', '14b07f21274b69891b1ad6c8a6a1fb05', 'google.png', 2);
INSERT INTO `banner_models` VALUES (11, '2024-02-28 21:55:02.155', '2024-02-28 21:55:02.155', 'http://qiniu.xiaoxinqj.top/gvb/20240228215501_小米su7.jpg', 'f24ba05632d7d3244b79491ad6827fe1', '小米su7.jpg', 2);
INSERT INTO `banner_models` VALUES (12, '2024-02-28 22:00:26.290', '2024-02-28 22:00:26.290', 'http://qiniu.xiaoxinqj.top/gvb/20240228220025_比亚迪秦plus.jpg', '1e13dfe9c16ff3a1b406309fd220d50e', '比亚迪秦plus.jpg', 2);

-- ----------------------------
-- Table structure for chat_models
-- ----------------------------
DROP TABLE IF EXISTS `chat_models`;
CREATE TABLE `chat_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `nick_name` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `msg_type` tinyint NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_group` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat_models
-- ----------------------------
INSERT INTO `chat_models` VALUES (1, '2024-03-11 19:13:34.876', '2024-03-11 19:13:34.876', '普拉蒂尼完成了五子登科', 'static/chat_avatar/普.png', '普拉蒂尼完成了五子登科 加入聊天室', 1, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (2, '2024-03-11 19:14:40.718', '2024-03-11 19:14:40.718', '普拉蒂尼完成了五子登科', 'static/chat_avatar/普.png', '大家好', 2, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (3, '2024-03-11 19:15:00.219', '2024-03-11 19:15:00.219', '普拉蒂尼完成了五子登科', 'static/chat_avatar/普.png', '消息类型错误', 6, '127.0.0.1', '内网', 0);
INSERT INTO `chat_models` VALUES (4, '2024-03-11 19:16:10.339', '2024-03-11 19:16:10.339', '普斯卡什在武汉看电影', 'static/chat_avatar/普.png', '普斯卡什在武汉看电影 加入聊天室', 1, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (5, '2024-03-11 19:16:28.630', '2024-03-11 19:16:28.630', '普斯卡什在武汉看电影', 'static/chat_avatar/普.png', '你好', 2, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (6, '2024-03-11 19:16:46.138', '2024-03-11 19:16:46.138', '普斯卡什在武汉看电影', 'static/chat_avatar/普.png', '普斯卡什在武汉看电影 离开聊天室', 7, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (7, '2024-03-11 19:17:12.838', '2024-03-11 19:17:12.838', '普拉蒂尼完成了五子登科', 'static/chat_avatar/普.png', '消息不能为空', 6, '127.0.0.1', '内网', 0);
INSERT INTO `chat_models` VALUES (8, '2024-03-11 23:15:24.418', '2024-03-11 23:15:24.418', '罗纳尔多望穿秋水', 'static/chat_avatar/罗.png', '罗纳尔多望穿秋水 加入聊天室', 1, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (9, '2024-03-11 23:15:47.956', '2024-03-11 23:15:47.956', '巴乔一眼定情', 'static/chat_avatar/巴.png', '巴乔一眼定情 离开聊天室', 7, '127.0.0.1', '内网', 1);
INSERT INTO `chat_models` VALUES (10, '2024-03-11 23:16:03.161', '2024-03-11 23:16:03.161', '罗纳尔多望穿秋水', 'static/chat_avatar/罗.png', '罗纳尔多望穿秋水 离开聊天室', 7, '127.0.0.1', '内网', 1);

-- ----------------------------
-- Table structure for comment_models
-- ----------------------------
DROP TABLE IF EXISTS `comment_models`;
CREATE TABLE `comment_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `parent_comment_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '\'父评论ID\'',
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'评论内容\'',
  `digg_count` tinyint NULL DEFAULT 0 COMMENT '\'评论点赞量\'',
  `comment_count` tinyint NULL DEFAULT 0 COMMENT '\'评论量\'',
  `article_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'评论文章ID\'',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '\'评论用户ID\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_comment_models_sub_comments`(`parent_comment_id` ASC) USING BTREE,
  INDEX `fk_comment_models_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_comment_models_sub_comments` FOREIGN KEY (`parent_comment_id`) REFERENCES `comment_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_comment_models_user` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment_models
-- ----------------------------
INSERT INTO `comment_models` VALUES (1, '2024-03-08 15:08:31.681', '2024-03-10 17:18:45.240', NULL, '你说的对', 0, -2, 'ZhKLDo4Beq8OFDNuzYQB', 1);
INSERT INTO `comment_models` VALUES (3, '2024-03-08 15:19:24.011', '2024-03-08 15:20:49.084', 1, '我很棒棒哦', 0, 1, 'ZhKLDo4Beq8OFDNuzYQB', 1);
INSERT INTO `comment_models` VALUES (5, '2024-03-08 15:20:49.091', '2024-03-08 15:20:49.091', 3, '哈哈哈', 0, 0, 'ZhKLDo4Beq8OFDNuzYQB', 1);

-- ----------------------------
-- Table structure for feedback_models
-- ----------------------------
DROP TABLE IF EXISTS `feedback_models`;
CREATE TABLE `feedback_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `apply_content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_apply` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of feedback_models
-- ----------------------------

-- ----------------------------
-- Table structure for log_stash_models
-- ----------------------------
DROP TABLE IF EXISTS `log_stash_models`;
CREATE TABLE `log_stash_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `level` tinyint NULL DEFAULT NULL,
  `content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_stash_models
-- ----------------------------
INSERT INTO `log_stash_models` VALUES (1, '2024-03-12 11:24:21.604', '192.168.100.10', '内网', 1, '哈哈哈', 1);
INSERT INTO `log_stash_models` VALUES (2, '2024-03-12 11:34:44.373', '127.0.0.1', '内网', 3, '密码错误 admin jkl3', 0);
INSERT INTO `log_stash_models` VALUES (3, '2024-03-12 11:35:11.576', '127.0.0.1', '内网', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (4, '2024-03-12 16:30:27.126', '127.0.0.1', '内网', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (5, '2024-05-05 16:10:06.640', '127.0.0.1', '内网IP', 3, '1 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (6, '2024-05-05 16:15:27.503', '127.0.0.1', '内网IP', 3, '12 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (7, '2024-05-05 16:15:57.170', '127.0.0.1', '内网IP', 3, '1 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (8, '2024-05-05 16:19:15.511', '127.0.0.1', '内网IP', 3, '1 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (9, '2024-05-05 17:02:51.643', '127.0.0.1', '内网IP', 3, '1 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (10, '2024-05-05 17:03:03.578', '127.0.0.1', '内网IP', 3, '密码错误 admin 1234', 0);
INSERT INTO `log_stash_models` VALUES (11, '2024-05-05 17:03:07.479', '127.0.0.1', '内网IP', 3, '密码错误 admin 123456', 0);
INSERT INTO `log_stash_models` VALUES (12, '2024-05-05 17:03:12.742', '127.0.0.1', '内网IP', 3, '密码错误 admin 123', 0);
INSERT INTO `log_stash_models` VALUES (13, '2024-05-05 17:03:15.789', '127.0.0.1', '内网IP', 3, '密码错误 admin 111', 0);
INSERT INTO `log_stash_models` VALUES (14, '2024-05-05 17:03:58.278', '127.0.0.1', '内网IP', 3, '密码错误 admin 123456', 0);
INSERT INTO `log_stash_models` VALUES (15, '2024-05-05 17:04:07.484', '127.0.0.1', '内网IP', 3, '密码错误 admin 111111', 0);
INSERT INTO `log_stash_models` VALUES (16, '2024-05-05 17:05:29.345', '127.0.0.1', '内网IP', 3, '密码错误 admin 111111', 0);
INSERT INTO `log_stash_models` VALUES (17, '2024-05-05 17:05:41.783', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (18, '2024-05-05 17:06:29.711', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (19, '2024-05-05 17:06:44.901', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (20, '2024-05-05 17:09:04.610', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (21, '2024-05-05 17:10:11.646', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (22, '2024-05-05 17:12:14.444', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (23, '2024-05-05 17:13:43.655', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (24, '2024-05-05 17:29:06.525', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (25, '2024-05-05 17:29:53.594', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (26, '2024-05-05 17:31:53.882', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (27, '2024-05-05 17:46:07.705', '127.0.0.1', '内网IP', 3, 'adf 用户不存在', 0);
INSERT INTO `log_stash_models` VALUES (28, '2024-05-05 17:54:47.133', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);
INSERT INTO `log_stash_models` VALUES (29, '2024-05-07 15:14:02.579', '127.0.0.1', '内网IP', 2, 'admin 登录成功', 1);

-- ----------------------------
-- Table structure for login_data_models
-- ----------------------------
DROP TABLE IF EXISTS `login_data_models`;
CREATE TABLE `login_data_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `nick_name` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `token` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `device` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `login_type` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_login_data_models_user_model`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_login_data_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login_data_models
-- ----------------------------
INSERT INTO `login_data_models` VALUES (1, '2024-03-10 16:30:27.134', '2024-03-10 16:30:27.134', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTA0MDUwMjcuMTI1NjA1OCwiaXNzIjoicWlhbmppbiJ9.ab1opaqPDqy7pWRU9VGWvUL8Bqiw-DS3N_d33PxPPVI', 'web', '内网', 3);
INSERT INTO `login_data_models` VALUES (2, '2024-03-12 17:04:23.000', '2024-03-12 17:04:30.000', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTA0MDUwMjcuMTI1NjA1OCwiaXNzIjoicWlhbmppbiJ9.ab1opaqPDqy7pWRU9VGWvUL8Bqiw-DS3N_d33PxPPVI', 'web', '内网', 3);
INSERT INTO `login_data_models` VALUES (3, '2024-05-05 17:05:41.794', '2024-05-05 17:05:41.794', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzI3NDEuNzgyNTA5LCJpc3MiOiJxaWFuamluIn0.3NOQQolCw2lMCCTPmlNUn6Cztec772e2x9HqWbfcA3c', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (4, '2024-05-05 17:06:29.729', '2024-05-05 17:06:29.729', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzI3ODkuNzEwNjUyOCwiaXNzIjoicWlhbmppbiJ9.BEwFcohAn_zeNZNnqRIAmWLu5ATEHlpt-JSLOfFjxU0', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (5, '2024-05-05 17:06:44.919', '2024-05-05 17:06:44.919', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzI4MDQuOTAwNjgzLCJpc3MiOiJxaWFuamluIn0.EZ8hXhf__OdXNwLFRktXHqtouRSMijn6YDQ0mf0Rik0', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (6, '2024-05-05 17:09:04.619', '2024-05-05 17:09:04.619', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzI5NDQuNjA5MzkxMiwiaXNzIjoicWlhbmppbiJ9.M08Km38bITBwmr5z7m1rTwfGRv5lDIweCTQbTLIeVAY', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (7, '2024-05-05 17:10:11.654', '2024-05-05 17:10:11.654', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzMwMTEuNjQ1ODQ1LCJpc3MiOiJxaWFuamluIn0.PalEUYHeMlRGqfaZqALefKML9lBS4NiS2axlpBcCBwA', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (8, '2024-05-05 17:12:14.454', '2024-05-05 17:12:14.454', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzMxMzQuNDQ0Mjg4LCJpc3MiOiJxaWFuamluIn0.C2sSOqAg1Gaj6U6q28UsHjerwItF-JrWoxC2qi-ICMA', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (9, '2024-05-05 17:13:43.665', '2024-05-05 17:13:43.665', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzMyMjMuNjU0MDIsImlzcyI6InFpYW5qaW4ifQ.qO5aQFLqxPNiKHGM2qJy2HqNmp4aoegQkkMrnu2nYVs', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (10, '2024-05-05 17:29:06.533', '2024-05-05 17:29:06.533', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzQxNDYuNTI0MzYsImlzcyI6InFpYW5qaW4ifQ.evejydOfkbRi0Dd_Zip4dbRkkSIZQnbfwZ4W_KkDs2c', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (11, '2024-05-05 17:29:53.603', '2024-05-05 17:29:53.603', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzQxOTMuNTk0NDMyLCJpc3MiOiJxaWFuamluIn0.IQWq3AA0Ia_6oN8GZft3MSWhO8SZeVd5_Dj47WL3GhE', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (12, '2024-05-05 17:31:53.894', '2024-05-05 17:31:53.894', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzQzMTMuODgyNDY1MSwiaXNzIjoicWlhbmppbiJ9.NlWJ2G3Ufr_LgwO6M12-qfUnQpNdaNzT4CXGtFPiRQ0', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (13, '2024-05-05 17:54:47.143', '2024-05-05 17:54:47.143', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUwNzU2ODcuMTMyNzIwMiwiaXNzIjoicWlhbmppbiJ9.kDPpFwAmWSm0HXn44leL4V0VuND3dNJhZO3Fzl62rKo', 'web', '内网IP', 3);
INSERT INTO `login_data_models` VALUES (14, '2024-05-07 15:14:02.591', '2024-05-07 15:14:02.591', 1, '127.0.0.1', '管理员', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrX25hbWUiOiLnrqHnkIblkZgiLCJyb2xlIjoxLCJleHAiOjE3MTUyMzg4NDIuNTc4OTIwOCwiaXNzIjoicWlhbmppbiJ9.4J2yAb7bYhBBi94vyJS8kCOavnDKs2aqKYU__onq4Hw', 'web', '内网IP', 3);

-- ----------------------------
-- Table structure for menu_banner_models
-- ----------------------------
DROP TABLE IF EXISTS `menu_banner_models`;
CREATE TABLE `menu_banner_models`  (
  `menu_id` bigint UNSIGNED NOT NULL,
  `banner_id` bigint UNSIGNED NOT NULL,
  `sort` smallint NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`, `banner_id`) USING BTREE,
  INDEX `fk_menu_banner_models_banner_model`(`banner_id` ASC) USING BTREE,
  CONSTRAINT `fk_menu_banner_models_banner_model` FOREIGN KEY (`banner_id`) REFERENCES `banner_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_menu_banner_models_menu_model` FOREIGN KEY (`menu_id`) REFERENCES `menu_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu_banner_models
-- ----------------------------
INSERT INTO `menu_banner_models` VALUES (1, 4, 1);
INSERT INTO `menu_banner_models` VALUES (1, 5, 0);

-- ----------------------------
-- Table structure for menu_models
-- ----------------------------
DROP TABLE IF EXISTS `menu_models`;
CREATE TABLE `menu_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'菜单标题\'',
  `path` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'菜单路径\'',
  `slogan` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'标语\'',
  `abstract` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '\'简介\'',
  `abstract_time` bigint NULL DEFAULT NULL COMMENT '\'简介的切换时间\'',
  `banner_time` bigint NULL DEFAULT NULL COMMENT '\'菜单的切换时间\'',
  `sort` smallint NULL DEFAULT NULL COMMENT '\'菜单的顺序\'',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu_models
-- ----------------------------
INSERT INTO `menu_models` VALUES (1, '2024-02-29 17:38:55.163', '2024-02-29 20:17:53.832', '首页', 'index', '小新个人博客', '小新的个人博客\ngin-vue-blog', 5, 5, 1);

-- ----------------------------
-- Table structure for message_models
-- ----------------------------
DROP TABLE IF EXISTS `message_models`;
CREATE TABLE `message_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `send_user_id` bigint UNSIGNED NOT NULL,
  `send_user_nick_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `send_user_avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `receive_user_id` bigint UNSIGNED NOT NULL,
  `receive_user_nick_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `receive_user_avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_read` tinyint(1) NULL DEFAULT 0,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`, `send_user_id`, `receive_user_id`) USING BTREE,
  INDEX `fk_message_models_send_user_model`(`send_user_id` ASC) USING BTREE,
  INDEX `fk_message_models_receive_user_model`(`receive_user_id` ASC) USING BTREE,
  CONSTRAINT `fk_message_models_receive_user_model` FOREIGN KEY (`receive_user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_message_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message_models
-- ----------------------------
INSERT INTO `message_models` VALUES (1, '2024-03-04 09:30:31.312', '2024-03-04 09:30:31.312', 1, '管理员', '/static/avatar/default.jpg', 2, 'xiaoxin', '/static/avatar/default.jpg', 0, '你好');
INSERT INTO `message_models` VALUES (2, '2024-03-04 09:31:44.241', '2024-03-04 09:31:44.241', 2, 'xiaoxin', '/static/avatar/default.jpg', 1, '管理员', '/static/avatar/default.jpg', 0, '你好，我也好');
INSERT INTO `message_models` VALUES (3, '2024-03-04 10:10:57.293', '2024-03-04 10:10:57.293', 1, '管理员', '/static/avatar/default.jpg', 3, 'arong', '/static/avatar/default.jpg', 1, '原神害人');
INSERT INTO `message_models` VALUES (4, '2024-03-04 10:11:04.467', '2024-03-04 11:00:39.088', 3, 'arong', '/static/avatar/default.jpg', 1, '管理员', '/static/avatar/default.jpg', 1, '还真是');
INSERT INTO `message_models` VALUES (5, '2024-03-04 10:12:58.067', '2024-03-04 10:12:58.067', 2, 'xiaoxin', '/static/avatar/default.jpg', 3, 'arong', '/static/avatar/default.jpg', 0, '起了没');
INSERT INTO `message_models` VALUES (6, '2024-03-04 10:13:03.816', '2024-03-04 10:13:03.816', 2, 'xiaoxin', '/static/avatar/default.jpg', 3, 'arong', '/static/avatar/default.jpg', 0, '扫码');
INSERT INTO `message_models` VALUES (7, '2024-03-04 10:13:15.057', '2024-03-04 10:13:15.057', 3, 'arong', '/static/avatar/default.jpg', 2, 'xiaoxin', '/static/avatar/default.jpg', 0, '好扫');
INSERT INTO `message_models` VALUES (8, '2024-03-04 10:34:58.688', '2024-03-04 11:00:39.095', 3, 'arong', '/static/avatar/default.jpg', 1, '管理员', '/static/avatar/default.jpg', 1, '害nm');

-- ----------------------------
-- Table structure for tag_models
-- ----------------------------
DROP TABLE IF EXISTS `tag_models`;
CREATE TABLE `tag_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'标签名称\'',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag_models
-- ----------------------------
INSERT INTO `tag_models` VALUES (1, '2024-03-03 23:25:15.485', '2024-03-13 18:23:46.934', 'java1');
INSERT INTO `tag_models` VALUES (3, '2024-03-05 22:39:48.965', '2024-03-05 22:39:48.965', 'go');
INSERT INTO `tag_models` VALUES (4, '2024-03-05 22:39:55.493', '2024-03-05 22:39:55.493', '后端');

-- ----------------------------
-- Table structure for user_collect_models
-- ----------------------------
DROP TABLE IF EXISTS `user_collect_models`;
CREATE TABLE `user_collect_models`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `article_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`user_id`, `article_id`) USING BTREE,
  CONSTRAINT `fk_user_collect_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_collect_models
-- ----------------------------
INSERT INTO `user_collect_models` VALUES ('2024-03-11 18:40:10.000', 1, 'ZhKLDo4Beq8OFDNuzYQB');

-- ----------------------------
-- Table structure for user_models
-- ----------------------------
DROP TABLE IF EXISTS `user_models`;
CREATE TABLE `user_models`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `nick_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'昵称\'',
  `user_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'用户名\'',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'密码\'',
  `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'头像\'',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'邮箱\'',
  `tel` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'电话\'',
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'地址\'',
  `token` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'令牌\'',
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '\'IP\'',
  `role` tinyint NULL DEFAULT 1 COMMENT '\'角色\'',
  `sign_status` bigint NULL DEFAULT NULL COMMENT '\'登录方式\'',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_models
-- ----------------------------
INSERT INTO `user_models` VALUES (1, '2024-02-20 16:04:14.106', '2024-03-03 11:44:06.936', '管理员', 'admin', '$2a$04$9M2Iks.xT9v1rPypFwbNS.K9S.DJYXOHnAfOX0oaUO3zAU.gNZ.eK', '/static/avatar/default.jpg', 'admin@qq.com', '18372754601', '内网地址', '', '127.0.0.1', 1, 3);
INSERT INTO `user_models` VALUES (2, '2024-03-01 17:57:29.645', '2024-03-03 18:52:02.211', 'xiaoxin', 'xiaoxin', '$2a$04$V7X4Mppiom4mYlpZ57ObTueqfVerxvZHAAM91Li.7RhqArim3XyQ6', '/static/avatar/default.jpg', '1429030919@qq.com', '18372754601', '内网地址', '', '127.0.0.1', 2, 3);
INSERT INTO `user_models` VALUES (3, '2024-03-03 23:17:06.076', '2024-03-03 23:17:06.076', 'arong', '阿荣', '$2a$04$GljAO4j8bQrFoX3Yh6tF0.yTrQiE5dgaWi3txP.TUxvSsIETyHRVO', '/static/avatar/default.jpg', '', '', '内网地址', '', '127.0.0.1', 2, 3);

SET FOREIGN_KEY_CHECKS = 1;
