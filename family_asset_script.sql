CREATE SCHEMA `family_asset` ;

use family_asset;
-- Query without GORM
-- SET FOREIGN_KEY_CHECKS = 0;
-- DROP TABLE if exists families;
-- DROP TABLE if exists assets;
-- SET FOREIGN_KEY_CHECKS = 1;
-- CREATE TABLE IF NOT EXISTS families (
--   ID INT AUTO_INCREMENT PRIMARY KEY,
--   NAME VARCHAR(255) DEFAULT 'paijo',
--   SEX ENUM('male','female') NOT NULL DEFAULT 'female',
--   DESCENDANT_ID INT,
--   FOREIGN KEY (DESCENDANT_ID) REFERENCES families(ID)
-- );

-- CREATE TABLE IF NOT EXISTS assets (
--   ID INT AUTO_INCREMENT PRIMARY KEY,
--   NAME VARCHAR(255) DEFAULT 'barang X',
--   FAMILY_ID INT,
--   FOREIGN KEY (FAMILY_ID) REFERENCES families(ID)
-- );

INSERT INTO families (name, sex, DESCENDANT_ID) VALUES
('Bani', 'male', NULL),
('Budi', 'male', 1),
('Nida', 'female', 1),
('Andi', 'male', 1),
('Sigit', 'male', 1),
('Hari', 'male', 2),
('Siti', 'female', 2),
('Bila', 'female', 3),
('Lesti', 'female', 3),
('Diki', 'male', 4),
('Doni', 'male', 5),
('Toni', 'male', 5);

INSERT INTO assets (NAME, FAMILY_ID) VALUES
('Samsung Universe 9',2),
('Samsung Galaxy Book',2),
('iPhone 9',6),
('iPhone X',7),
('Huawei P30',3),
('Samsung Universe 9',8),
('Huawei P30',9),
('iPhone X',9),
('Samsung Universe 9 ',4),
('Samsung Galaxy Book',10),
('Huawei P30',5),
('iPhone X',11);