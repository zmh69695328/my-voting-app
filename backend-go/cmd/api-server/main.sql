/*
 Navicat Premium Data Transfer

 Source Server         : backend-go-voting
 Source Server Type    : SQLite
 Source Server Version : 3035005 (3.35.5)
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3035005 (3.35.5)
 File Encoding         : 65001

 Date: 18/10/2023 16:46:18
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS "team";
CREATE TABLE "team" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "teamname" text,
  "workname" TEXT,
  "order" TEXT,
  "group" TEXT
);

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF EXISTS "vote";
CREATE TABLE "vote" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "teamid" INTEGER NOT NULL,
  "score" INTEGER NOT NULL,
  "date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Auto increment value for vote
-- ----------------------------

PRAGMA foreign_keys = true;
