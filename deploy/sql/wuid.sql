CREATE SCHEMA IF NOT EXISTS `im-zero` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `im-zero`;
CREATE TABLE `wuid` (
                        `h` int(10) NOT NULL AUTO_INCREMENT,
                        `x` tinyint(4) NOT NULL DEFAULT '0',
                        PRIMARY KEY (`x`),
                        UNIQUE KEY `h` (`h`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;