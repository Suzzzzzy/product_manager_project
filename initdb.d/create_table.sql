DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `id` int(11) NOT NULL,
                         `password` varchar(45) NOT NULL,
                         `phone_number` int(11) NOT NULL,
                         `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                         `deleted_at` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `category` varchar(50) NOT NULL,
                           `price` float NOT NULL,
                           `cost` float NOT NULL,
                           `name` varchar(50) NOT NULL,
                           `description` varchar(255) NOT NULL,
                           `barcode` varchar(255) NOT NULL,
                           `expiration_date` datetime NOT NULL,
                           `size` varchar(50) NOT NULL,
                           `user_id` int(11) NOT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

