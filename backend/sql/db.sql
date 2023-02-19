
CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    `cust_id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `dob` date NOT NULL,
    `city` varchar(100) NOT NULL,
    `zipcode` varchar(10) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`cust_id`) 
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
INSERT INTO `customers` VALUES
    (2000, 'Sundar', '1983-11-08', 'Chennai', '602025', 1),
    (2001, 'Anushya', '1984-09-02', 'Chennai', '600001', 0),
    (2002, 'Shankya', '2010-02-18', 'Chennai', '600020', 1),
    (2003, 'Dhyan', '2013-09-30', 'Bangalore', '560048', 1);


DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `cust_id` int(11) NOT NULL,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `amount` float NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`account_id`),
    KEY `accounts_FK` (`cust_id`),
    CONSTRAINT `accounts_FK` FOREIGN KEY (`cust_id`) REFERENCES `customers` (`cust_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95476 DEFAULT CHARSET=latin1;

INSERT INTO `accounts` VALUES
    (95470, 2000, '2020-08-22 10:20:06', 'Saving', 6075, 1),
    (95471, 2001, '2021-08-22 10:20:06', 'Saving', 10076, 1),
    (95472, 2002, '2022-08-22 10:20:06', 'Checking', 111077, 1);

CREATE TABLE `transactions` (
    `tx_id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `amount` int(11) NOT NULL,
    `tx_type` varchar(10) NOT NULL,
    `tx_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`tx_id`),
    KEY `transactions_FK` (`account_id`),
    CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)    
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
