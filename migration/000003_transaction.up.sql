CREATE TABLE IF NOT EXISTS `transactions` (
  `transaction_id` VARCHAR(255) NOT NULL,
  `customer_id` VARCHAR(45) NOT NULL,
  `nomor_kontrak` VARCHAR(255) DEFAULT NULL,
  `otr` DOUBLE DEFAULT NULL,
  `admin_fee` DOUBLE DEFAULT NULL,
  `jumlah_cicilan` BIGINT UNSIGNED DEFAULT NULL,
  `jumlah_bunga` DOUBLE DEFAULT NULL,
  `transaction_date` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`trx_id`),
  KEY `idx_transactions_customer` (`customer_id`),
  CONSTRAINT `fk_transactions_customer` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
