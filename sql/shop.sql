DROP TABLE IF EXISTS `shop_admin`;
CREATE TABLE IF NOT EXISTS `shop_admin`(
  `uid` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '管理员账号',
  `password` CHAR(32) NOT NULL DEFAULT '' COMMENT '管理员密码',
  `email` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '管理员电子邮箱',
  `login_time` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '登录时间',
  `login_ip` BIGINT NOT NULL DEFAULT '0' COMMENT '登录IP',
  `create_time` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modify_time` INT UNSIGNED NOT NULL  DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY(`uid`),
  unique (`email`),
  unique (`username`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `shop_admin`(username,password,email,create_time) VALUES('mok', md5('123'), '1005914310@qq.com', UNIX_TIMESTAMP());

