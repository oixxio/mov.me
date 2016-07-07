-- phpMyAdmin SQL Dump
-- version 4.0.10deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: May 10, 2016 at 12:01 PM
-- Server version: 5.5.49-0ubuntu0.14.04.1
-- PHP Version: 5.5.9-1ubuntu4.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `movilizame`
--

-- --------------------------------------------------------

--
-- Table structure for table `apps`
--

CREATE TABLE IF NOT EXISTS `apps` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(200) COLLATE utf8_bin NOT NULL,
  `name` varchar(200) COLLATE utf8_bin NOT NULL,
  `title` varchar(200) COLLATE utf8_bin NOT NULL,
  `description` varchar(6000) COLLATE utf8_bin NOT NULL,
  `url_demo` varchar(200) COLLATE utf8_bin NOT NULL,
  `icon_path` varchar(200) COLLATE utf8_bin NOT NULL,
  `header_path` varchar(200) COLLATE utf8_bin NOT NULL,
  `screenshot_path_1` varchar(200) COLLATE utf8_bin NOT NULL,
  `screenshot_path_2` varchar(200) COLLATE utf8_bin NOT NULL,
  `screenshot_path_3` varchar(200) COLLATE utf8_bin NOT NULL,
  `url_video` varchar(200) COLLATE utf8_bin NOT NULL,
  `version` varchar(20) COLLATE utf8_bin NOT NULL,
  `version_notes` varchar(2000) COLLATE utf8_bin NOT NULL,
  `state` varchar(2000) COLLATE utf8_bin NOT NULL,
  `keywords` varchar(2000) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=6 ;

--
-- Dumping data for table `apps`
--

INSERT INTO `apps` (`ID`, `type`, `name`, `title`, `description`, `url_demo`, `icon_path`, `header_path`, `screenshot_path_1`, `screenshot_path_2`, `screenshot_path_3`, `url_video`, `version`, `version_notes`, `state`, `keywords`) VALUES
(1, 'app', 'facundomero', 'la mejor app del mundo', 'soluciona todos los problemas del mundo', '-', '-', '-', '-', '-', '-', '', '0.0', 'ni esta arrancada', 'development', 'awsome'),
(3, 'facundo', 'facundo', '', '', '', '', '', '', '', '', '', '', '', '', ''),
(4, 'guille', 'guille', '', '', '', '', '', '', '', '', '', '', '', '', ''),
(5, 'guille', 'guille', '', '', '', '', '', '', '', '', '', '', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `app_requirements`
--

CREATE TABLE IF NOT EXISTS `app_requirements` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `id_app` int(11) NOT NULL,
  `type` varchar(200) COLLATE utf8_bin NOT NULL,
  `description` varchar(200) COLLATE utf8_bin NOT NULL,
  `rules` varchar(500) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=7 ;

--
-- Dumping data for table `app_requirements`
--

INSERT INTO `app_requirements` (`ID`, `id_app`, `type`, `description`, `rules`) VALUES
(2, 2, 'game', 'dasd', 's}dñlmsldñgn'),
(3, 2, 'a', 'a', 'a'),
(4, 1, '', 's', ''),
(5, 2, '', 'facu', ''),
(6, 2, '', 'a', '');

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE IF NOT EXISTS `clients` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(200) COLLATE utf8_bin NOT NULL,
  `password` varchar(200) COLLATE utf8_bin NOT NULL,
  `name` varchar(200) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=3 ;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`ID`, `email`, `password`, `name`) VALUES
(1, 'facuompre@gmail.com', 'admin', 'facundo Ompre'),
(2, 'f', 'f', 'f');

-- --------------------------------------------------------

--
-- Table structure for table `client_features`
--

CREATE TABLE IF NOT EXISTS `client_features` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `id_app` int(11) NOT NULL,
  `screenshot_path` varchar(200) COLLATE utf8_bin NOT NULL,
  `title` varchar(200) COLLATE utf8_bin NOT NULL,
  `description` varchar(200) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=6 ;

--
-- Dumping data for table `client_features`
--

INSERT INTO `client_features` (`ID`, `id_app`, `screenshot_path`, `title`, `description`) VALUES
(2, 2, 'asgasdge', 'SAGSDG', 'GFSGSDsfgsg'),
(3, 2, 'a', 'a', 'a'),
(4, 1, 'f', 'f', 'f'),
(5, 1, '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `developers`
--

CREATE TABLE IF NOT EXISTS `developers` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(200) COLLATE utf8_bin NOT NULL,
  `password` varchar(200) COLLATE utf8_bin NOT NULL,
  `name` varchar(200) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=3 ;

--
-- Dumping data for table `developers`
--

INSERT INTO `developers` (`ID`, `email`, `password`, `name`) VALUES
(1, 'facu', 'facu', 'facu'),
(2, 'g', 'g', 'g');

-- --------------------------------------------------------

--
-- Table structure for table `plans`
--

CREATE TABLE IF NOT EXISTS `plans` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `id_developer` int(11) NOT NULL,
  `id_client` int(11) NOT NULL,
  `id_app` int(11) NOT NULL,
  `state` varchar(200) COLLATE utf8_bin NOT NULL,
  `type` int(11) NOT NULL,
  `requirements` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=6 ;

--
-- Dumping data for table `plans`
--

INSERT INTO `plans` (`ID`, `id_developer`, `id_client`, `id_app`, `state`, `type`, `requirements`) VALUES
(1, 1, 1, 1, 'w', 0, 0),
(2, 0, 0, 0, 'asd', 0, 0),
(3, 1, 1, 1, 'activo', 1, 12),
(4, 0, 0, 0, '', 0, 0),
(5, 0, 0, 0, '', 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_features`
--

CREATE TABLE IF NOT EXISTS `user_features` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `id_app` int(11) NOT NULL,
  `screenshot_path` varchar(200) COLLATE utf8_bin NOT NULL,
  `title` varchar(200) COLLATE utf8_bin NOT NULL,
  `description` varchar(200) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=13 ;

--
-- Dumping data for table `user_features`
--

INSERT INTO `user_features` (`ID`, `id_app`, `screenshot_path`, `title`, `description`) VALUES
(2, 2, 'asdñlvjnpjnpu', 'sdkf vkjdsfa vkj', '|dflkjsv .as vñkj'),
(3, 1, 'aaaaaaaaaa', 'sssssssssssss', 'gggggggggggggg'),
(4, 2, 'aaaaaaaaaaa', 'ssssssssssssss', 'dddddddddddddddddd'),
(5, 1, 'a', 'a', 'a'),
(6, 2, 's', 's', 's'),
(7, 2, 'e', 'e', 'e'),
(8, 2, 'e', 'e', 'e'),
(9, 2, 'e', 'e', 'e'),
(10, 2, 'e', 'e', 'e'),
(11, 1, 's', 's', 's'),
(12, 1, 't', 't', 't');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
