USE MYSQL_DATABASE;
DROP TABLE IF EXISTS t1;
CREATE TABLE t1(x INT PRIMARY KEY auto_increment);
INSERT INTO t1 () VALUES (),(),(),(),(),(),(),(),(),(),(),(),();
DELIMITER $$
DROP PROCEDURE IF EXISTS data$$
CREATE PROCEDURE data()
BEGIN
DECLARE COUNTER_X INT;
SET COUNTER_X = 1;
SELECT 'start';
WHILE COUNTER_X <= 24 DO
INSERT INTO t1 (x) SELECT x + (SELECT count(*) FROM t1) FROM t1;
SELECT COUNTER_X;
SET COUNTER_X = COUNTER_X + 1;
END WHILE;
SELECT 'finish';
END$$
DELIMITER ;
CALL data();