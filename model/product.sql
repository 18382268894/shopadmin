create table if not exists `shop_product`(
  `Pid` int(11) unsigned auto_increment not null comment '商品id',
  `Cid` mediumint(11) unsigned  not null comment '商品分类id',
  `Name` varchar(30)  not null comment '商品名称',
  `Summary` varchar(255)  not null comment '商品摘要',
  `Num`  int unsigned  not null default 0 comment '商品库存',
  `Price` decimal(10,2) unsigned  not null comment '商品价格',
  `Pics` varchar(500)   default '' comment '商品图片存放位置',
  CreateTime timestamp not null default current_timestamp,
  ModifyTime timestamp on update current_timestamp,
  primary key (`Pid`),
  foreign key (`Cid`) references `shop_category`(`Cid`) on UPDATE cascade on delete restrict
)engine = InnoDB default charset=utf8mb4 row_format dynamic;