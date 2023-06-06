DROP TABLE IF EXISTS `web_news`;
CREATE TABLE `web_news` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '新闻标题',
  `created_on` varchar(100) DEFAULT '' COMMENT '时间',
  `content` text COMMENT '内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='新闻管理';
select * from web_news;

DROP TABLE IF EXISTS `web_member`;
CREATE TABLE `web_member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(100) DEFAULT '' COMMENT '身份',
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `name` varchar(100) DEFAULT '' COMMENT '姓名',
  `phone` varchar(100) DEFAULT '' COMMENT '电话',
  `mail` varchar(100) DEFAULT '' COMMENT '邮箱',
  `research` varchar(100) DEFAULT '' COMMENT '研究方向',
  `achievement` text COMMENT '成果',
  `introduction` text COMMENT '内容',
  `state` varchar(100) DEFAULT '可用' COMMENT '状态',
 
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='成员管理';
select * from web_member;

DROP TABLE IF EXISTS `web_image`;
CREATE TABLE `web_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `content` text COMMENT '内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='相册管理';
select * from web_image;

DROP TABLE IF EXISTS `web_achievement`;
CREATE TABLE `web_achievement` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `category` varchar(100) DEFAULT '' COMMENT '类别',
  `address` varchar(100) DEFAULT '' COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='成果管理';
select * from web_achievement;

DROP TABLE IF EXISTS `web_manager`;
CREATE TABLE `web_manager` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `identity` varchar(50) DEFAULT '' COMMENT '身份',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `web_manager` (`id`, `username`, `password`,`identity`) VALUES ('1001', '1', '1','系统管理员');

DROP TABLE IF EXISTS `web_article`;
CREATE TABLE `web_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '论文标题',
  `journal` varchar(100) DEFAULT '' COMMENT '期刊',
  `author` varchar(100) DEFAULT '' COMMENT '第一作者',
  `authors` varchar(100) DEFAULT '' COMMENT '其他作者',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `link` varchar(100) DEFAULT '' '详情页链接',
  `papercode` varchar(100) DEFAULT '' '代码',
  `theyear` varchar(10) DEFAULT '' COMMENT '论文年份',
  `abstract` text COMMENT '摘要',
  `state` varchar(10) DEFAULT '草稿' COMMENT '论文状态',
  `hide` varchar(10) DEFAULT '0' COMMENT '是否隐藏',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='论文管理';
select * from web_article;

DROP TABLE IF EXISTS `web_demo`;
CREATE TABLE web_demo (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
   `title` varchar(100) DEFAULT '' COMMENT 'demo标题',
   `date` varchar(100) DEFAULT '' COMMENT '时间',
   `content` text COMMENT '内容',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='demo管理';
SELECT * FROM web_demo

DROP TABLE IF EXISTS `web_project`;
CREATE TABLE `web_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '项目名称',
  `link` varchar(100) DEFAULT '' COMMENT '项目链接',
  `state` varchar(50) DEFAULT '在研' COMMENT '项目状态',
  `theyear` varchar(10) DEFAULT '' COMMENT '项目年份',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
select * from web_project;

DROP TABLE IF EXISTS `web_activity`;
CREATE TABLE `web_activity` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '活动标题',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `content` text COMMENT '内容',
  `kind` varchar(50) COMMENT '活动类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='活动管理';
select * from web_activity;

DROP TABLE IF EXISTS `web_standards`;
CREATE TABLE `web_standards` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '技术标准标题',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `content` text COMMENT '内容',
  `kind` varchar(50) COMMENT '类型',
  `state` varchar(50) COMMENT '申请状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='技术标准管理';
select * from web_standards;

DROP TABLE IF EXIST `web_casbin_rule`;
CREATE TABLE `web_casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL COMMENT '规则类型',
  `v0` varchar(100) DEFAULT NULL COMMENT '角色ID',
  `v1` varchar(100) DEFAULT NULL COMMENT 'api路径',
  `v2` varchar(100) DEFAULT NULL COMMENT 'api访问方法',
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限规则表';
/*插入操作casbin api的权限规则*/
INSERT INTO `web_casbin_rule`(`p_type`, `v0`, `v1`, `v2`) VALUES ('p', 'admin', '/api/v1/casbin', 'POST');
INSERT INTO `web_casbin_rule`(`p_type`, `v0`, `v1`, `v2`) VALUES ('p', 'admin', '/api/v1/casbin/list', 'GET');