-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: gvb_db
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `advert_models`
--

DROP TABLE IF EXISTS `advert_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `advert_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '广告标题',
  `href` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '跳转链接',
  `images` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '图片',
  `is_show` tinyint(1) DEFAULT NULL COMMENT '是否展示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `advert_models`
--

LOCK TABLES `advert_models` WRITE;
/*!40000 ALTER TABLE `advert_models` DISABLE KEYS */;
INSERT INTO `advert_models` VALUES (2,'2023-03-05 14:55:11.274','2024-05-08 21:14:14.335','gvb后端教程','https://www.bilibili.com/video/BV1f24y1G72C','https://qiniu.xiaoxinqj.fun/static/my/img/20230928233304.png',0),(3,'2024-05-08 16:25:22.000','2024-05-08 16:25:25.000','小米su7','https://www.xiaomiev.com/','https://qiniu.xiaoxinqj.fun/gvb/20240228215501_xiaomisu7.jpg',1),(4,'2024-05-08 20:41:30.080','2024-05-08 20:41:30.080','比亚迪秦plus','https://bydauto.com.cn/pc/configCar?id=109&networkType=dynasty','https://qiniu.xiaoxinqj.fun/gvb/20240228220025_byd_QinPlus.jpg',1),(5,'2024-05-11 00:27:16.000','2024-05-11 00:27:18.000','原神','https://ys.mihoyo.com/','/uploads/file/管理员/ys.png',1);
/*!40000 ALTER TABLE `advert_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auto_reply_models`
--

DROP TABLE IF EXISTS `auto_reply_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auto_reply_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(32) DEFAULT NULL,
  `mode` bigint DEFAULT NULL,
  `rule` varchar(64) DEFAULT NULL,
  `reply_content` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auto_reply_models`
--

LOCK TABLES `auto_reply_models` WRITE;
/*!40000 ALTER TABLE `auto_reply_models` DISABLE KEYS */;
INSERT INTO `auto_reply_models` VALUES (1,'2024-05-10 15:06:28.244','2024-05-10 15:08:13.862','规则1',2,'你是谁？','我是小新千问');
/*!40000 ALTER TABLE `auto_reply_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `banner_models`
--

DROP TABLE IF EXISTS `banner_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `banner_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '图片路径',
  `hash` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '图片的hash值',
  `name` varchar(38) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图片名称',
  `image_type` bigint DEFAULT '1' COMMENT '图片的类型，本地还是七牛,1本地，2七牛，默认是1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `banner_models`
--

LOCK TABLES `banner_models` WRITE;
/*!40000 ALTER TABLE `banner_models` DISABLE KEYS */;
INSERT INTO `banner_models` VALUES (7,'2023-02-20 20:16:57.782','2023-02-20 20:16:57.782','/uploads/file/38.jpg','fe2d77826fa9327431000afa23fdfb1d','38.jpg',1),(10,'2023-03-04 20:59:46.682','2023-03-04 20:59:46.682','/uploads/file/215953KgCPr.jpg','548bb4a8dc5448b83ad91861fb7751a2','215953KgCPr.jpg',1),(12,'2023-03-04 21:02:24.590','2023-03-04 21:02:24.590','/uploads/file/235628ZxKUe.jpg','1faa43ac634d5775ecb854494ae865ef','235628ZxKUe.jpg',1),(36,'2024-05-09 21:52:19.697','2024-05-09 21:52:19.697','http://qiniu.xiaoxinqj.fun/gvb/20240509215219__a4.jpg','7822a6008d890a2357aff37955478d3b','a4.jpg',2),(37,'2024-05-09 21:52:19.846','2024-05-09 21:52:19.846','http://qiniu.xiaoxinqj.fun/gvb/20240509215219__xw.jpg','8513fc467217aad1665b965371d590b6','xw.jpg',2);
/*!40000 ALTER TABLE `banner_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `big_model_chat_models`
--

DROP TABLE IF EXISTS `big_model_chat_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `big_model_chat_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `session_id` bigint unsigned DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `content` longtext,
  `bot_content` longtext,
  `role_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_big_model_session_models_chat_list` (`session_id`),
  KEY `fk_big_model_chat_models_role_model` (`role_id`),
  KEY `fk_big_model_chat_models_user_model` (`user_id`),
  CONSTRAINT `fk_big_model_chat_models_role_model` FOREIGN KEY (`role_id`) REFERENCES `big_model_role_models` (`id`),
  CONSTRAINT `fk_big_model_chat_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_big_model_session_models_chat_list` FOREIGN KEY (`session_id`) REFERENCES `big_model_session_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `big_model_chat_models`
--

LOCK TABLES `big_model_chat_models` WRITE;
/*!40000 ALTER TABLE `big_model_chat_models` DISABLE KEYS */;
INSERT INTO `big_model_chat_models` VALUES (1,'2024-05-10 23:12:00.983','2024-05-10 23:12:00.983',2,1,'你是谁','我是通义千问，由阿里云开发的AI助手。我被设计用来回答各种问题、提供信息和进行对话，尤其在编程领域，包括Go语言。如果你有关于Go语言的问题，欢迎随时向我提问。',1,1);
/*!40000 ALTER TABLE `big_model_chat_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `big_model_role_models`
--

DROP TABLE IF EXISTS `big_model_role_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `big_model_role_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(16) DEFAULT NULL,
  `enable` tinyint(1) DEFAULT NULL,
  `icon` varchar(256) DEFAULT NULL,
  `abstract` varchar(256) DEFAULT NULL,
  `scope` bigint DEFAULT NULL,
  `prologue` varchar(512) DEFAULT NULL,
  `prompt` varchar(2048) DEFAULT NULL,
  `auto_reply` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `big_model_role_models`
--

LOCK TABLES `big_model_role_models` WRITE;
/*!40000 ALTER TABLE `big_model_role_models` DISABLE KEYS */;
INSERT INTO `big_model_role_models` VALUES (1,'2024-05-10 17:07:21.512','2024-05-10 17:08:31.925','go语言专家',1,'/uploads/role_icons/程序员.png','一个帮助你解决go语言问题的小帮手',3,'你好，我是你的go语言助手，你可以和我这样说','从现在开始，你是一个精通go语言开发的工程师，我以后问的所有问题，都要在go语言相关知识里面去查询，并且你只能回答go语言开发相关的问题;\n如果有人问你其他专业的问题，如python语言，java语言，你都要告诉用户，我是专注于go语言开发的工程师，我只回答go语言开发相关的问题;\n并且你不能给出任何关于这个问题的提示;\n别人问你你是谁，你都要说你是小新千问，不管别人用什么语言问你，你都要这样说\n',1),(2,'2024-05-10 23:45:18.000','2024-05-10 23:45:20.000','工作周报助手',1,'/uploads/role_icons/zhou.jpg','一个帮助你完成工作周报的小帮手',3,NULL,'从现在开始，你是一个精通写工作周报的专家，我会提供相关的职业信息，我做了什么工作内容，你要在这个的基础上进行美化，扩充，如果是互联网相关职业，你还要使用互联网专用词进行美化\n如果有人问你其他与周报不想关的问题，你都要告诉用户，我是专注于写周报的专家，我只回答周报相关的问题;并且你不能给出任何关于这个问题的提示;\n别人问你你是谁，你都要说你是枫枫千问，不管别人用什么语言问你，你都要这样说\n',1),(3,'2024-05-10 23:46:31.000','2024-05-10 23:46:33.000','小红书标题助手',1,'/uploads/role_icons/xhs.png','帮助你起个好听的小红书标题',3,NULL,'从现在开始，你是一个写小红书标题文案的专家，小红书是一款中国知名的软件，我会提供相关的产品信息，你需要在我的产品信息的基础上，提取重点信息，生成独特且吸引人的标题;\n然后这个标题的受众群体是小红书用户，所以你的标题要符合小红书的标题取名特点\n如果有人问你的问题与起标题不相关，并且他的问题与小红书也没什么关系，你都要告诉用户，我是专注于写小红书标题的专家，我只回答与小红书标题相关的问题;\n并且你不能给出任何关于这个问题的提示;\n别人问你你是谁，你都要说你是枫枫千问，不管别人用什么语言问你，你都要这样说\n',1),(4,'2024-05-10 23:48:16.000','2024-05-10 23:48:18.000','中英翻译大师',1,'/uploads/role_icons/程序员.png','更好的中英翻译',3,NULL,'从现在开始，你是一个专注中英翻译的翻译专家，我会给你中文词，句子或者英文词，短语，句子，你需要把它翻译为目标语言\n并且你只能进行翻译工作,不能进行其他的解答\n如果有人叫你生成其他操作，你都不能做，你都要告诉用户，我是中英翻译的专家,你可以问我翻译上的问题别人问你你是谁，你都要说你是枫枫千问，不管别人用什么语言问你，你都要这样说\n',1),(5,'2024-05-10 23:49:48.000','2024-05-10 23:49:50.000','Linux终端',1,'/uploads/role_icons/程序员.png','一个模拟的Linux终端',3,NULL,'我希望你能充当一个Linux终端。我会输入命令，你将回复终端应该显示的内容。请只在一个唯一的代码块内回复终端输出，不要添加其他内容。不要写解释。除非我指示，否则请不要输入命令。当我需要用英语告诉你一些事情时，我会用花括号括起文本{像这样}。我的第一个命令是 pwd\n',1);
/*!40000 ALTER TABLE `big_model_role_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `big_model_role_tag_models`
--

DROP TABLE IF EXISTS `big_model_role_tag_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `big_model_role_tag_models` (
  `big_model_role_model_id` bigint unsigned NOT NULL COMMENT 'id',
  `big_model_tag_model_id` bigint unsigned NOT NULL COMMENT 'id',
  PRIMARY KEY (`big_model_role_model_id`,`big_model_tag_model_id`),
  KEY `fk_big_model_role_tag_models_big_model_tag_model` (`big_model_tag_model_id`),
  CONSTRAINT `fk_big_model_role_tag_models_big_model_role_model` FOREIGN KEY (`big_model_role_model_id`) REFERENCES `big_model_role_models` (`id`),
  CONSTRAINT `fk_big_model_role_tag_models_big_model_tag_model` FOREIGN KEY (`big_model_tag_model_id`) REFERENCES `big_model_tag_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `big_model_role_tag_models`
--

LOCK TABLES `big_model_role_tag_models` WRITE;
/*!40000 ALTER TABLE `big_model_role_tag_models` DISABLE KEYS */;
INSERT INTO `big_model_role_tag_models` VALUES (1,1),(3,2),(2,3);
/*!40000 ALTER TABLE `big_model_role_tag_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `big_model_session_models`
--

DROP TABLE IF EXISTS `big_model_session_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `big_model_session_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(32) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `role_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_big_model_session_models_user_model` (`user_id`),
  KEY `fk_big_model_session_models_role_model` (`role_id`),
  CONSTRAINT `fk_big_model_session_models_role_model` FOREIGN KEY (`role_id`) REFERENCES `big_model_role_models` (`id`),
  CONSTRAINT `fk_big_model_session_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `big_model_session_models`
--

LOCK TABLES `big_model_session_models` WRITE;
/*!40000 ALTER TABLE `big_model_session_models` DISABLE KEYS */;
INSERT INTO `big_model_session_models` VALUES (2,'2024-05-10 22:20:06.133','2024-05-10 22:20:06.133','新的会话',1,1),(3,'2024-05-10 23:58:59.732','2024-05-10 23:58:59.732','新的会话',1,2);
/*!40000 ALTER TABLE `big_model_session_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `big_model_tag_models`
--

DROP TABLE IF EXISTS `big_model_tag_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `big_model_tag_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `title` varchar(16) DEFAULT NULL,
  `color` varchar(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `big_model_tag_models`
--

LOCK TABLES `big_model_tag_models` WRITE;
/*!40000 ALTER TABLE `big_model_tag_models` DISABLE KEYS */;
INSERT INTO `big_model_tag_models` VALUES (1,'2024-05-10 15:49:40.506','2024-05-10 15:51:37.688','网络编程','blue'),(2,'2024-05-10 23:52:18.000','2024-05-10 23:52:21.000','营销文案','red'),(3,'2024-05-10 16:03:27.401','2024-05-10 16:03:27.401','效率工具','green');
/*!40000 ALTER TABLE `big_model_tag_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chat_models`
--

DROP TABLE IF EXISTS `chat_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chat_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `nick_name` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内容',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址',
  `msg_type` tinyint DEFAULT NULL COMMENT '消息类型',
  `is_group` tinyint(1) DEFAULT NULL COMMENT '是否是群组消息',
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=724 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chat_models`
--

LOCK TABLES `chat_models` WRITE;
/*!40000 ALTER TABLE `chat_models` DISABLE KEYS */;
INSERT INTO `chat_models` VALUES (709,'2023-11-12 18:42:05.334','2023-11-12 18:42:05.334','个性的卡卡','uploads/chat_avatar/个.png','个性的卡卡 进入聊天室','127.0.0.1','内网地址',1,1,NULL),(710,'2023-11-12 18:42:09.914','2023-11-12 18:42:09.914','个性的卡卡','uploads/chat_avatar/个.png','你好','127.0.0.1','内网地址',2,1,NULL),(711,'2023-11-12 18:42:10.866','2023-11-12 18:42:10.866','个性的卡卡','uploads/chat_avatar/个.png','个性的卡卡 离开聊天室','127.0.0.1','内网地址',7,1,NULL),(712,'2024-05-09 20:28:50.154','2024-05-09 20:28:50.154','复杂的柯南','uploads/chat_avatar/复.png','复杂的柯南 进入聊天室','127.0.0.1','内网地址',1,1,NULL),(713,'2024-05-09 20:28:55.090','2024-05-09 20:28:55.090','复杂的柯南','uploads/chat_avatar/复.png','dd','127.0.0.1','内网地址',2,1,NULL),(714,'2024-05-09 20:30:26.844','2024-05-09 20:30:26.844','苏牙完成了五子登科','uploads/chat_avatar/苏.png','苏牙完成了五子登科 进入聊天室','127.0.0.1','内网地址',1,1,NULL),(715,'2024-05-09 20:30:29.469','2024-05-09 20:30:29.469','苏牙完成了五子登科','uploads/chat_avatar/苏.png','fa','127.0.0.1','内网地址',2,1,NULL),(716,'2024-05-09 20:30:38.890','2024-05-09 20:30:38.890','苏牙完成了五子登科','uploads/chat_avatar/苏.png','苏牙完成了五子登科 离开聊天室','127.0.0.1','内网地址',7,1,NULL),(717,'2024-05-09 20:30:44.639','2024-05-09 20:30:44.639','复杂的柯南','uploads/chat_avatar/复.png','复杂的柯南 离开聊天室','127.0.0.1','内网地址',7,1,NULL),(718,'2024-05-11 00:20:01.894','2024-05-11 00:20:01.894','雅典娜掐指一算','uploads/chat_avatar/雅.png','雅典娜掐指一算 进入聊天室','127.0.0.1','内网地址',1,1,NULL),(719,'2024-05-11 00:20:09.893','2024-05-11 00:20:09.893','雅典娜掐指一算','uploads/chat_avatar/雅.png','arong来了','127.0.0.1','内网地址',2,1,NULL),(720,'2024-05-11 00:21:06.168','2024-05-11 00:21:06.168','雅典娜掐指一算','uploads/chat_avatar/雅.png','雅典娜掐指一算 离开聊天室','127.0.0.1','内网地址',7,1,NULL),(721,'2024-05-11 00:21:28.292','2024-05-11 00:21:28.292','苗条的范巴斯滕','uploads/chat_avatar/苗.png','苗条的范巴斯滕 进入聊天室','127.0.0.1','内网地址',1,1,NULL),(722,'2024-05-11 00:21:34.790','2024-05-11 00:21:34.790','苗条的范巴斯滕','uploads/chat_avatar/苗.png','qj来了','127.0.0.1','内网地址',2,1,NULL),(723,'2024-05-11 00:21:40.636','2024-05-11 00:21:40.636','苗条的范巴斯滕','uploads/chat_avatar/苗.png','苗条的范巴斯滕 离开聊天室','127.0.0.1','内网地址',7,1,NULL);
/*!40000 ALTER TABLE `chat_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_models`
--

DROP TABLE IF EXISTS `comment_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `parent_comment_id` bigint unsigned DEFAULT NULL COMMENT '父评论id',
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论内容',
  `digg_count` tinyint DEFAULT '0' COMMENT '点赞数',
  `comment_count` tinyint DEFAULT '0' COMMENT '子评论数',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '关联的用户id',
  `article_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文章id',
  PRIMARY KEY (`id`),
  KEY `fk_comment_models_user` (`user_id`),
  KEY `fk_comment_models_sub_comments` (`parent_comment_id`),
  CONSTRAINT `fk_comment_models_sub_comments` FOREIGN KEY (`parent_comment_id`) REFERENCES `comment_models` (`id`),
  CONSTRAINT `fk_comment_models_user` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_models`
--

LOCK TABLES `comment_models` WRITE;
/*!40000 ALTER TABLE `comment_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `feedback_models`
--

DROP TABLE IF EXISTS `feedback_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `feedback_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feedback_models`
--

LOCK TABLES `feedback_models` WRITE;
/*!40000 ALTER TABLE `feedback_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `feedback_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_models`
--

DROP TABLE IF EXISTS `log_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `ip` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'ip',
  `addr` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '地址',
  `level` bigint DEFAULT NULL COMMENT '等级',
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '详情',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `user_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '用户名',
  `service_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '服务名',
  `status` tinyint(1) DEFAULT NULL COMMENT '登录状态',
  `type` bigint DEFAULT NULL COMMENT '日志类型，1登录，2操作，3运行',
  `readStatus` tinyint(1) DEFAULT '0' COMMENT '阅读状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_models`
--

LOCK TABLES `log_models` WRITE;
/*!40000 ALTER TABLE `log_models` DISABLE KEYS */;
INSERT INTO `log_models` VALUES (365,'2024-05-08 20:30:02.974','2024-05-08 20:30:02.974','127.0.0.1','内网地址',0,'用户名密码错误','123456',0,'admin','',0,1,0),(366,'2024-05-08 20:30:06.972','2024-05-08 20:30:06.972','127.0.0.1','内网地址',0,'用户名密码错误','123',0,'admin','',0,1,0),(367,'2024-05-08 20:30:11.526','2024-05-08 20:30:11.526','127.0.0.1','内网地址',0,'用户名密码错误','111',0,'admin','',0,1,0),(368,'2024-05-08 20:32:05.815','2024-05-08 20:32:05.815','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(369,'2024-05-09 10:41:55.792','2024-05-09 10:41:55.792','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(370,'2024-05-09 20:47:27.445','2024-05-09 20:47:27.445','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(371,'2024-05-09 23:55:12.637','2024-05-09 23:55:12.637','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(372,'2024-05-10 00:18:39.325','2024-05-10 00:18:39.335','127.0.0.1','内网地址',1,'用户 admin 注销登录','<div class=\"log_request_header\">\n	<div class=\"log_request_body\">\n		<pre class=\"log_json_body\">{\"Accept\":[\"*/*\"],\"Accept-Encoding\":[\"gzip, deflate, br\"],\"Cache-Control\":[\"no-cache\"],\"Connection\":[\"keep-alive\"],\"Content-Length\":[\"0\"],\"Postman-Token\":[\"ff0bf52e-b5c5-44a6-ae19-db5aee50f703\"],\"Token\":[\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQyOTEyLjYzNjUyOCwiaXNzIjoicWlhbmppbiJ9.D7tRphdOVOD4pmVNdZHrO5COuojo1l3lmo7yKNE5-n0\"],\"User-Agent\":[\"PostmanRuntime/7.38.0\"]}</pre>\n	</div>\n</div>\n\n<div class=\"log_response\">\n	<pre class=\"log_json_body\">{\"code\":0,\"data\":{},\"msg\":\"注销成功\"}</pre>\n</div>\n',1,'admin','',0,2,0),(373,'2024-05-10 00:19:01.216','2024-05-10 00:19:01.216','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(374,'2024-05-10 00:20:35.638','2024-05-10 00:20:35.638','127.0.0.1','内网地址',0,'用户名密码错误','jkl',0,'xiaoxin','',0,1,0),(375,'2024-05-10 00:21:05.160','2024-05-10 00:21:05.160','127.0.0.1','内网地址',0,'用户名密码错误','jkl',0,'xiaoxin','',0,1,0),(376,'2024-05-10 00:21:08.649','2024-05-10 00:21:08.649','127.0.0.1','内网地址',0,'用户名密码错误','123',0,'xiaoxin','',0,1,0),(377,'2024-05-10 00:21:18.760','2024-05-10 00:21:18.760','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(378,'2024-05-10 00:21:40.334','2024-05-10 00:21:40.334','127.0.0.1','内网地址',0,'用户名密码错误','jkl',0,'xiaoxin','',0,1,0),(379,'2024-05-10 00:21:57.457','2024-05-10 00:21:57.457','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(380,'2024-05-10 00:23:28.939','2024-05-10 00:23:28.939','127.0.0.1','内网地址',0,'登录成功','--',24,'xiaoxin','',1,1,0),(381,'2024-05-10 00:38:55.762','2024-05-10 00:38:55.762','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(382,'2024-05-10 00:51:30.059','2024-05-10 00:51:30.059','127.0.0.1','内网地址',0,'登录成功','--',24,'xiaoxin','',1,1,0),(383,'2024-05-10 10:30:09.815','2024-05-10 10:30:09.815','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0),(384,'2024-05-10 10:50:02.961','2024-05-10 10:50:02.961','127.0.0.1','内网地址',0,'登录成功','--',1,'admin','',1,1,0);
/*!40000 ALTER TABLE `log_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_stash_models`
--

DROP TABLE IF EXISTS `log_stash_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_stash_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `level` tinyint DEFAULT NULL,
  `content` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_stash_models`
--

LOCK TABLES `log_stash_models` WRITE;
/*!40000 ALTER TABLE `log_stash_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_stash_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_data_models`
--

DROP TABLE IF EXISTS `login_data_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_data_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `nick_name` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `token` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'token',
  `device` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录失败',
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址',
  `login_type` tinyint DEFAULT NULL COMMENT '登录方式，1QQ，3邮箱',
  PRIMARY KEY (`id`),
  KEY `fk_login_data_models_user_model` (`user_id`),
  CONSTRAINT `fk_login_data_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=226 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `login_data_models`
--

LOCK TABLES `login_data_models` WRITE;
/*!40000 ALTER TABLE `login_data_models` DISABLE KEYS */;
INSERT INTO `login_data_models` VALUES (214,'2024-05-08 20:32:05.827','2024-05-08 20:32:05.827',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1MzQ0MzI1LjgxNDY2LCJpc3MiOiJxaWFuamluIn0.E8KWsKIr5F0qmfYru2SyREfgAn5nWPhlTFmXjbJMlt4','','内网地址',3),(215,'2024-05-09 10:41:55.801','2024-05-09 10:41:55.801',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1Mzk1MzE1Ljc5MTM1NywiaXNzIjoicWlhbmppbiJ9.a5FNtKyFXzMyTHrqyGTzoPNrYR756FBAteQHcbca6Bs','','内网地址',3),(216,'2024-05-09 20:47:27.454','2024-05-09 20:47:27.454',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDMxNjQ3LjQ0MzkzNiwiaXNzIjoicWlhbmppbiJ9.651oCh8GBi9PPNFF7DZoJphlVGUNB8_wFOOzcOceO38','','内网地址',3),(217,'2024-05-09 23:55:12.647','2024-05-09 23:55:12.647',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQyOTEyLjYzNjUyOCwiaXNzIjoicWlhbmppbiJ9.D7tRphdOVOD4pmVNdZHrO5COuojo1l3lmo7yKNE5-n0','','内网地址',3),(218,'2024-05-10 00:19:01.223','2024-05-10 00:19:01.223',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQ0MzQxLjIxNTYsImlzcyI6InFpYW5qaW4ifQ.Cc6p5E0LwlEj_-HfgUqPyQyZruR47hSkWFyRDJhI1hY','','内网地址',3),(219,'2024-05-10 00:21:18.778','2024-05-10 00:21:18.778',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQ0NDc4Ljc1OTc3NSwiaXNzIjoicWlhbmppbiJ9.OreTIWt5d8NJxi1RO7tL4B3z1JWDPLfVpEQfWArKUf0','','内网地址',3),(220,'2024-05-10 00:21:57.476','2024-05-10 00:21:57.476',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQ0NTE3LjQ1NzQwMywiaXNzIjoicWlhbmppbiJ9.sG7m9crR-Wd0NIf4G6RzOQI_QkFXNP21PyXi7OK-Wao','','内网地址',3),(221,'2024-05-10 00:23:28.957','2024-05-10 00:23:28.957',24,'127.0.0.1','小新','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InhpYW94aW4iLCJuaWNrX25hbWUiOiLlsI_mlrAiLCJyb2xlIjoyLCJ1c2VyX2lkIjoyNCwiZXhwIjoxNzE1NDQ0NjA4LjkzODA3MzIsImlzcyI6InFpYW5qaW4ifQ.-nUATaTkddncEMCRODlHPjRKEStynbLWMvDg1M_ixL4','','内网地址',3),(222,'2024-05-10 00:38:55.779','2024-05-10 00:38:55.779',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDQ1NTM1Ljc2MTg3MTgsImlzcyI6InFpYW5qaW4ifQ.1PxxTB2ySWSBAEYimIxpkwYjOVrKA_CQm8CN-ynGtts','','内网地址',3),(223,'2024-05-10 00:51:30.067','2024-05-10 00:51:30.067',24,'127.0.0.1','小新','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InhpYW94aW4iLCJuaWNrX25hbWUiOiLlsI_mlrAiLCJyb2xlIjoyLCJ1c2VyX2lkIjoyNCwiZXhwIjoxNzE1NDQ2MjkwLjA1ODIzOSwiaXNzIjoicWlhbmppbiJ9.9TXfY5YzHuhtQTnFL_qMHpqoZqbztrPUBBjJ7tuSEhE','','内网地址',3),(224,'2024-05-10 10:30:09.825','2024-05-10 10:30:09.825',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDgxMDA5LjgxNDU1OSwiaXNzIjoicWlhbmppbiJ9.jM-wIeMJOnAuXonrYQo2LQ-zHxZjKksmFqgRKUrIbYU','','内网地址',3),(225,'2024-05-10 10:50:02.978','2024-05-10 10:50:02.978',1,'127.0.0.1','管理员','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwibmlja19uYW1lIjoi566h55CG5ZGYIiwicm9sZSI6MSwidXNlcl9pZCI6MSwiZXhwIjoxNzE1NDgyMjAyLjk2MDA0NywiaXNzIjoicWlhbmppbiJ9.zkYB-MdaWxXtLEw0tbgA8GutjMAUudB4wh6vkgN-CF8','','内网地址',3);
/*!40000 ALTER TABLE `login_data_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_banner_models`
--

DROP TABLE IF EXISTS `menu_banner_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_banner_models` (
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单的id',
  `banner_id` bigint unsigned DEFAULT NULL COMMENT 'banner图的id',
  `sort` smallint DEFAULT NULL COMMENT '序号',
  KEY `fk_menu_banner_models_menu_model` (`menu_id`),
  KEY `fk_banner_models_menus_banner` (`banner_id`),
  CONSTRAINT `fk_banner_models_menus_banner` FOREIGN KEY (`banner_id`) REFERENCES `banner_models` (`id`),
  CONSTRAINT `fk_menu_banner_models_banner_model` FOREIGN KEY (`banner_id`) REFERENCES `banner_models` (`id`),
  CONSTRAINT `fk_menu_banner_models_menu_model` FOREIGN KEY (`menu_id`) REFERENCES `menu_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_banner_models`
--

LOCK TABLES `menu_banner_models` WRITE;
/*!40000 ALTER TABLE `menu_banner_models` DISABLE KEYS */;
INSERT INTO `menu_banner_models` VALUES (2,12,1),(1,7,1),(1,10,0);
/*!40000 ALTER TABLE `menu_banner_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu_models`
--

DROP TABLE IF EXISTS `menu_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单标题',
  `path` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单路径',
  `slogan` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'slogan',
  `abstract` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '简介，按照换行去切割为数组',
  `abstract_time` bigint DEFAULT NULL COMMENT '简介的切换时间',
  `banner_time` bigint DEFAULT NULL COMMENT 'banner图的切换时间',
  `sort` smallint DEFAULT NULL COMMENT '顺序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_models`
--

LOCK TABLES `menu_models` WRITE;
/*!40000 ALTER TABLE `menu_models` DISABLE KEYS */;
INSERT INTO `menu_models` VALUES (1,'2023-03-05 16:33:13.472','2024-05-09 22:34:27.620','首页','/','首页','小新的个人博客\n主打轻量化',7,7,7),(2,'2023-03-05 17:44:54.251','2023-10-31 23:03:29.033','新闻','/news','新闻','关注国家大事\n新闻',7,7,6),(3,'2023-03-23 22:58:51.425','2023-10-31 23:03:29.034','搜索','/search','文章搜索','文章搜索',7,7,5),(4,'2023-03-23 22:59:43.538','2023-10-31 23:03:29.036','聊天室','/chat','聊天室','聊天室\n大家一起嗨',7,7,4),(5,'2023-03-23 23:03:39.602','2023-10-31 23:03:29.037','官方文档','http://docs.fengfengzhidao.com/','官方文档','官方文档',7,7,1),(7,'2023-10-31 23:03:21.195','2023-10-31 23:04:06.609','网站关于','/article/3BpBhosBT6PNFiwnnFW8','','',7,7,2);
/*!40000 ALTER TABLE `menu_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_models`
--

DROP TABLE IF EXISTS `message_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `send_user_id` bigint unsigned NOT NULL,
  `send_user_nick_name` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '发送人昵称',
  `send_user_avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '发送人头像',
  `rev_user_id` bigint unsigned NOT NULL,
  `rev_user_nick_name` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '接收人昵称',
  `rev_user_avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '接收人头像',
  `is_read` tinyint(1) DEFAULT '0' COMMENT '接收人是否查看',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '消息内容',
  PRIMARY KEY (`id`,`send_user_id`,`rev_user_id`),
  KEY `fk_message_models_send_user_model` (`send_user_id`),
  KEY `fk_message_models_rev_user_model` (`rev_user_id`),
  CONSTRAINT `fk_message_models_rev_user_model` FOREIGN KEY (`rev_user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_message_models_send_user_model` FOREIGN KEY (`send_user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_models`
--

LOCK TABLES `message_models` WRITE;
/*!40000 ALTER TABLE `message_models` DISABLE KEYS */;
INSERT INTO `message_models` VALUES (3,'2024-05-10 00:51:43.408','2024-05-10 00:51:43.408',24,'小新','/uploads/avatar/default.png',1,'管理员','/uploads/file/管理员/images_20231005232530.png',0,'你的初始密码为：xxxxx');
/*!40000 ALTER TABLE `message_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_models`
--

DROP TABLE IF EXISTS `tag_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `title` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签的名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_models`
--

LOCK TABLES `tag_models` WRITE;
/*!40000 ALTER TABLE `tag_models` DISABLE KEYS */;
INSERT INTO `tag_models` VALUES (1,'2023-02-22 23:17:53.000','2022-02-22 23:18:03.000','python'),(2,'2023-02-22 23:17:54.000','2023-02-22 23:18:02.000','后端'),(3,'2023-02-22 21:17:55.000','2023-02-22 23:18:07.000','部署'),(4,'2023-02-22 23:50:06.000','2023-02-22 23:50:05.000','gin'),(5,'2023-03-06 00:18:03.847','2023-03-06 00:18:03.847','go');
/*!40000 ALTER TABLE `tag_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_collect_models`
--

DROP TABLE IF EXISTS `user_collect_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_collect_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  `article_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文章的es id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '收藏的时间',
  PRIMARY KEY (`id`),
  KEY `fk_user_collect_models_user_model` (`user_id`),
  CONSTRAINT `fk_user_collect_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_collect_models`
--

LOCK TABLES `user_collect_models` WRITE;
/*!40000 ALTER TABLE `user_collect_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_collect_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `nick_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `user_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `tel` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '地址',
  `token` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '其他平台的唯一id',
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `role` tinyint DEFAULT '1' COMMENT '权限，1管理员，2普通用户，3游客',
  `sign_status` bigint DEFAULT NULL COMMENT '注册来源，1qq，3邮箱',
  `integral` bigint DEFAULT '0' COMMENT '我的积分',
  `sign` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '我的签名',
  `link` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '我的链接地址',
  `scope` bigint DEFAULT '0' COMMENT '我的积分',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_models`
--

LOCK TABLES `user_models` WRITE;
/*!40000 ALTER TABLE `user_models` DISABLE KEYS */;
INSERT INTO `user_models` VALUES (1,'2023-02-19 17:23:04.543','2024-05-10 23:58:59.743','管理员','admin','$2a$04$t7xvcoBnSogrbfbB48LLSe1yzRnbypRKjmF.Jpyjzq5.KylmIw7j.','/uploads/file/管理员/images_20231005232530.png','1429030919@qq.com','','内网地址','','127.0.0.1',1,3,0,'这是我的花火','http://www.xiaoxinqj.fun',86),(24,'2023-10-09 21:07:41.207','2024-05-10 00:27:07.209','小新','xiaoxin','$2a$04$BUD3kA.36Bc3mnuKD4T7W.W.XF4OPvOJxBCd6scBwsWglJ502T1MG','/uploads/avatar/default.png','1429030919@qq.com','','内网地址','','127.0.0.1',2,3,0,'123','',0);
/*!40000 ALTER TABLE `user_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_scope_models`
--

DROP TABLE IF EXISTS `user_scope_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_scope_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `user_id` bigint unsigned DEFAULT NULL,
  `scope` bigint DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_scope_models`
--

LOCK TABLES `user_scope_models` WRITE;
/*!40000 ALTER TABLE `user_scope_models` DISABLE KEYS */;
INSERT INTO `user_scope_models` VALUES (1,'2024-05-10 11:33:14.288','2024-05-10 11:33:14.288',1,100,1);
/*!40000 ALTER TABLE `user_scope_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-11 10:17:20
