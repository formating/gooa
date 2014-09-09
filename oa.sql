-- phpMyAdmin SQL Dump
-- version 4.2.2
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: 2014-09-04 09:45:15
-- 服务器版本： 5.6.17
-- PHP Version: 5.5.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `oa`
--

-- --------------------------------------------------------

--
-- 表的结构 `b_company`
--

CREATE TABLE IF NOT EXISTS `b_company` (
`cid` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `phone` varchar(12) DEFAULT NULL,
  `phone1` varchar(12) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `createdate` int(10) NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `b_company`
--

INSERT INTO `b_company` (`cid`, `name`, `phone`, `phone1`, `address`, `createdate`) VALUES
(1, '测试', '010-11111111', NULL, '北京市海淀区', 1401231231);

-- --------------------------------------------------------

--
-- 表的结构 `b_user`
--

CREATE TABLE IF NOT EXISTS `b_user` (
`uid` int(11) unsigned NOT NULL,
  `cid` int(11) NOT NULL,
  `ip_address` varbinary(16) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(80) NOT NULL,
  `salt` varchar(40) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `forgotten_password_code` varchar(40) DEFAULT NULL,
  `forgotten_password_time` int(11) unsigned DEFAULT NULL,
  `remember_code` varchar(40) DEFAULT NULL,
  `created_on` int(11) unsigned NOT NULL,
  `last_login` int(11) unsigned DEFAULT NULL,
  `active` tinyint(1) unsigned DEFAULT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `b_user`
--

INSERT INTO `b_user` (`uid`, `cid`, `ip_address`, `username`, `password`, `salt`, `email`, `forgotten_password_code`, `forgotten_password_time`, `remember_code`, `created_on`, `last_login`, `active`) VALUES
(1, 1, '12121212', 'username', 'pass', 'salt', 'baiyuxiong@126.com', '123', 1231231231, '123123', 1231231231, 1231231231, 1);

-- --------------------------------------------------------

--
-- 表的结构 `b_user_groups`
--

CREATE TABLE IF NOT EXISTS `b_user_groups` (
`gid` mediumint(8) unsigned NOT NULL,
  `cid` int(11) NOT NULL COMMENT '单位ID',
  `name` varchar(20) NOT NULL,
  `description` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `b_user_profile`
--

CREATE TABLE IF NOT EXISTS `b_user_profile` (
  `uid` int(11) NOT NULL,
  `gender` tinyint(4) NOT NULL COMMENT '1 男 2 女',
  `name` varchar(32) NOT NULL,
  `avatar` varchar(256) NOT NULL,
  `avatar_thumb1` varchar(256) NOT NULL,
  `avatar_thumb2` varchar(256) NOT NULL,
  `avatar_thumb3` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `b_user_users_groups_rel`
--

CREATE TABLE IF NOT EXISTS `b_user_users_groups_rel` (
`id` int(11) unsigned NOT NULL,
  `cid` int(11) NOT NULL,
  `uid` int(11) unsigned NOT NULL,
  `gid` mediumint(8) unsigned NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `b_company`
--
ALTER TABLE `b_company`
 ADD PRIMARY KEY (`cid`);

--
-- Indexes for table `b_user`
--
ALTER TABLE `b_user`
 ADD PRIMARY KEY (`uid`), ADD KEY `fk_user_cid_idx` (`cid`);

--
-- Indexes for table `b_user_groups`
--
ALTER TABLE `b_user_groups`
 ADD PRIMARY KEY (`gid`), ADD KEY `cid` (`cid`);

--
-- Indexes for table `b_user_profile`
--
ALTER TABLE `b_user_profile`
 ADD KEY `uid` (`uid`);

--
-- Indexes for table `b_user_users_groups_rel`
--
ALTER TABLE `b_user_users_groups_rel`
 ADD PRIMARY KEY (`id`), ADD KEY `cid_uid` (`cid`,`uid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `b_company`
--
ALTER TABLE `b_company`
MODIFY `cid` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `b_user`
--
ALTER TABLE `b_user`
MODIFY `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `b_user_groups`
--
ALTER TABLE `b_user_groups`
MODIFY `gid` mediumint(8) unsigned NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `b_user_users_groups_rel`
--
ALTER TABLE `b_user_users_groups_rel`
MODIFY `id` int(11) unsigned NOT NULL AUTO_INCREMENT;
--
-- 限制导出的表
--

--
-- 限制表 `b_user`
--
ALTER TABLE `b_user`
ADD CONSTRAINT `fk_user_cid` FOREIGN KEY (`cid`) REFERENCES `b_company` (`cid`) ON DELETE NO ACTION ON UPDATE NO ACTION;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
