/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : couples_subtotal

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 29/09/2022 22:14:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dynamic_records
-- ----------------------------
DROP TABLE IF EXISTS `dynamic_records`;
CREATE TABLE `dynamic_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_visible` tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '对方是否可见（可见-1，不可见-2）',
  `content` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '动态内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_dynamic_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dynamic_records
-- ----------------------------

-- ----------------------------
-- Table structure for photo_albums
-- ----------------------------
DROP TABLE IF EXISTS `photo_albums`;
CREATE TABLE `photo_albums`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_photo_albums_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of photo_albums
-- ----------------------------

-- ----------------------------
-- Table structure for photo_records
-- ----------------------------
DROP TABLE IF EXISTS `photo_records`;
CREATE TABLE `photo_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `img_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '照片路径',
  `album_id` bigint(20) NOT NULL COMMENT '相册id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_photo_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of photo_records
-- ----------------------------

-- ----------------------------
-- Table structure for sys_usrs
-- ----------------------------
DROP TABLE IF EXISTS `sys_usrs`;
CREATE TABLE `sys_usrs`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `account_number` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户账号',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `birthday_data` bigint(20) NULL DEFAULT NULL COMMENT '出生日期',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户邮箱',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_usrs_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_usrs
-- ----------------------------

-- ----------------------------
-- Table structure for user_relationships
-- ----------------------------
DROP TABLE IF EXISTS `user_relationships`;
CREATE TABLE `user_relationships`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `update_by` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `person_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '个人id',
  `couple_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '情侣id',
  `memorial_date` bigint(20) NULL DEFAULT NULL COMMENT '纪念日',
  `rela_status` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '关系状态（恋爱中-1，已解绑-2）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_relationships_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_relationships
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
