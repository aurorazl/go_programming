create database qps

create table qps.user (
    `id` BIGINT AUTO_INCREMENT,
    `name` varchar(256),
    `age` TINYINT,
    `favorite` varchar(256),
    `salary` Int,
    PRIMARY KEY (id),
    index name
)

DELIMITER $$	--更改结束符
CREATE PROCEDURE autoinsert(IN num INT)
BEGIN
    DECLARE i INT DEFAULT 1;			--declare声明，set设置值
    WHILE(i<num) DO					--类似的函数有if，repeat until，loop_label:LOOP
    INSERT INTO qps.user (NAME,age,favorite,salary) VALUES(REPLACE(UUID(),'-',''),FLOOR(RAND()*100),REPLACE(UUID(),'-',''),FLOOR(RAND()*10000));
    SET i = i+1;
END WHILE;
END $$
DELIMITER ;	--改回来

call autoinsert(1000000);

alter table qps.user add index (name)

explain select count(1) from `qps`.`user` WHERE favorite like 'run%'

alter table qps.user add index(favorite)