CREATE TABLE IF NOT EXISTS `products` (
  `product_id` VARCHAR(45) NOT NULL,
  `category` VARCHAR(45) NOT NULL,
  `sku` VARCHAR(45) NOT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  `amount` DOUBLE DEFAULT NULL,
  `stock` BIGINT UNSIGNED DEFAULT NULL,
  `tenor` FLOAT DEFAULT NULL,
  `bunga` BIGINT UNSIGNED DEFAULT NULL,
  `fee` DOUBLE DEFAULT NULL,
  `created_at` DATETIME(3) DEFAULT NULL,
  `updated_at` DATETIME(3) DEFAULT NULL,
  `deleted_at` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`product_id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;