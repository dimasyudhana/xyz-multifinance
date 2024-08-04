CREATE FUNCTION IF NOT EXISTS id(length INT) 
RETURNS CHAR(15) 
DETERMINISTIC
BEGIN
    DECLARE chars VARCHAR(26) DEFAULT 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    DECLARE result VARCHAR(15) DEFAULT '';
    DECLARE i INT DEFAULT 0;
    WHILE i < length DO
        SET result = CONCAT(result, SUBSTRING(chars, FLOOR(1 + (RAND() * LENGTH(chars))), 1));
        SET i = i + 1;
    END WHILE;
    RETURN result;
END;

INSERT INTO xyz_multifinance.products
(product_id, category, sku, name, amount, stock, tenor, bunga, fee, created_at, updated_at, deleted_at)
VALUES

-- White Goods - Refrigerator LG
(id(15), 'white goods', 'WG201', 'Refrigerator LG', 5000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG201', 'Refrigerator LG', 5000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG201', 'Refrigerator LG', 5000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG201', 'Refrigerator LG', 5000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- White Goods - Washing Machine Samsung
(id(15), 'white goods', 'WG202', 'Washing Machine Samsung', 3000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG202', 'Washing Machine Samsung', 3000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG202', 'Washing Machine Samsung', 3000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG202', 'Washing Machine Samsung', 3000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- White Goods - Microwave Panasonic
(id(15), 'white goods', 'WG203', 'Microwave Panasonic', 2000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG203', 'Microwave Panasonic', 2000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG203', 'Microwave Panasonic', 2000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'white goods', 'WG203', 'Microwave Panasonic', 2000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Motor - Honda Vario
(id(15), 'motor', 'MOT301', 'Honda Vario', 20000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT301', 'Honda Vario', 20000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT301', 'Honda Vario', 20000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT301', 'Honda Vario', 20000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Motor - Yamaha NMAX
(id(15), 'motor', 'MOT302', 'Yamaha NMAX', 25000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT302', 'Yamaha NMAX', 25000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT302', 'Yamaha NMAX', 25000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT302', 'Yamaha NMAX', 25000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Motor - Suzuki GSX-R150
(id(15), 'motor', 'MOT303', 'Suzuki GSX-R150', 28000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT303', 'Suzuki GSX-R150', 28000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT303', 'Suzuki GSX-R150', 28000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'motor', 'MOT303', 'Suzuki GSX-R150', 28000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Mobil - Mitsubishi XPorce
(id(15), 'mobil', 'CAR701', 'Mitsubishi XPorce', 200000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR701', 'Mitsubishi XPorce', 200000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR701', 'Mitsubishi XPorce', 200000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR701', 'Mitsubishi XPorce', 200000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Mobil - Honda Mobilio
(id(15), 'mobil', 'CAR702', 'Honda Mobilio', 350000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR702', 'Honda Mobilio', 350000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR702', 'Honda Mobilio', 350000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR702', 'Honda Mobilio', 350000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL),

-- Mobil - Suzuki Jimny
(id(15), 'mobil', 'CAR703', 'Suzuki Jimny', 220000000, 3, 1, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR703', 'Suzuki Jimny', 220000000, 3, 2, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR703', 'Suzuki Jimny', 220000000, 3, 3, 2, 2500, '2020-01-01', '2020-01-01', NULL),
(id(15), 'mobil', 'CAR703', 'Suzuki Jimny', 220000000, 3, 6, 2, 2500, '2020-01-01', '2020-01-01', NULL);
