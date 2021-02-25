CREATE DATABASE IF NOT EXISTS `database`;
USE `database`;

CREATE TABLE ADS(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    price DECIMAL(10, 2),
    data_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IMAGES(
    id INT PRIMARY KEY AUTO_INCREMENT,
    ref VARCHAR(512) NOT NULL,
    data_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ads_id INT NOT NULL,
    FOREIGN KEY (ads_id) REFERENCES ADS (id)
);

-- INSERT INTO ADS(title,description,price) VALUES("title0","disc0",0.0);
-- INSERT INTO ADS(title,description,price) VALUES("title1","disc1",1.0);
-- INSERT INTO ADS(title,description,price) VALUES("title2","disc2",2.0);

-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref0", 1);
-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref1", 1);
-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref2", 1);
-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref3", 2);
-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref4", 2);
-- INSERT INTO IMAGES(ref,ads_id) VALUES("ref5", 3);