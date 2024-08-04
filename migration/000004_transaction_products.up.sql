CREATE TABLE IF NOT EXISTS `transaction_products` (
  `transaction_id` VARCHAR(255) NOT NULL,
  `product_id` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`transaction_id`, `product_id`),
  KEY `idx_transaction_products_transaction` (`transaction_id`),
  KEY `idx_transaction_products_product` (`product_id`),
  CONSTRAINT `fk_transaction_products_transaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`transaction_id`),
  CONSTRAINT `fk_transaction_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
