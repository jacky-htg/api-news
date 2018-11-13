-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2.1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Nov 13, 2018 at 06:21 PM
-- Server version: 5.7.24-0ubuntu0.16.04.1
-- PHP Version: 7.0.32-0ubuntu0.16.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `api_news`
--

-- --------------------------------------------------------

--
-- Table structure for table `access`
--

CREATE TABLE `access` (
  `id` int(10) UNSIGNED NOT NULL,
  `parent_id` int(10) UNSIGNED DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `access`
--

INSERT INTO `access` (`id`, `parent_id`, `name`, `created_at`, `updated_at`) VALUES
(1, NULL, 'root', '2018-11-09 08:18:05', '2018-11-09 08:18:05'),
(2, 1, 'topics', '2018-11-11 16:30:34', '2018-11-11 16:30:34'),
(3, 2, 'topics.store', '2018-11-11 16:31:44', '2018-11-11 16:31:44'),
(4, 2, 'topics.update', '2018-11-11 16:31:44', '2018-11-11 16:31:44'),
(5, 2, 'topics.destroy', '2018-11-11 16:32:05', '2018-11-11 16:32:05'),
(6, 1, 'news', '2018-11-11 16:32:57', '2018-11-11 16:32:57'),
(7, 6, 'news.store', '2018-11-11 16:33:38', '2018-11-11 16:33:38'),
(8, 6, 'news.update', '2018-11-11 16:33:38', '2018-11-11 16:33:38'),
(9, 6, 'news.publish', '2018-11-11 16:34:00', '2018-11-11 16:34:00'),
(10, 6, 'news.destroy', '2018-11-11 16:34:00', '2018-11-11 16:34:00');

-- --------------------------------------------------------

--
-- Table structure for table `access_groups`
--

CREATE TABLE `access_groups` (
  `id` int(10) UNSIGNED NOT NULL,
  `group_id` int(10) UNSIGNED NOT NULL,
  `access_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `access_groups`
--

INSERT INTO `access_groups` (`id`, `group_id`, `access_id`) VALUES
(1, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `groups`
--

CREATE TABLE `groups` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `groups`
--

INSERT INTO `groups` (`id`, `title`, `created_at`, `updated_at`) VALUES
(1, 'Editor', '2018-11-11 04:07:47', '2018-11-11 04:07:47'),
(2, 'Writer', '2018-11-11 16:35:03', '2018-11-11 16:35:03');

-- --------------------------------------------------------

--
-- Table structure for table `news`
--

CREATE TABLE `news` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `image` varchar(512) DEFAULT NULL,
  `image_caption` varchar(512) DEFAULT NULL,
  `status` varchar(1) NOT NULL DEFAULT 'D',
  `publish_date` datetime DEFAULT NULL,
  `writer` int(10) UNSIGNED NOT NULL,
  `editor` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `news`
--

INSERT INTO `news` (`id`, `title`, `slug`, `content`, `image`, `image_caption`, `status`, `publish_date`, `writer`, `editor`, `created_at`, `updated_at`) VALUES
(10, 'News Kedua', 'news-kedua', '', 'https://theshonet-assets.s3.ap-southeast-1.amazonaws.com/images/filemanager/shared/b6589fc6ab0dc82cf12099d1c2d40ab994e8410c2djhMu1mJJ4EMLsENQzr.png', 'BERKACALAH JAKARTA', 'X', '2018-11-11 17:48:52', 1, 1, '2018-11-11 15:06:59', '2018-11-11 18:01:41');

-- --------------------------------------------------------

--
-- Table structure for table `news_topics`
--

CREATE TABLE `news_topics` (
  `id` int(10) UNSIGNED NOT NULL,
  `news_id` int(10) UNSIGNED NOT NULL,
  `topic_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `news_topics`
--

INSERT INTO `news_topics` (`id`, `news_id`, `topic_id`) VALUES
(1, 10, 1),
(2, 10, 2);

-- --------------------------------------------------------

--
-- Table structure for table `topics`
--

CREATE TABLE `topics` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `topics`
--

INSERT INTO `topics` (`id`, `title`, `created_at`, `updated_at`) VALUES
(1, 'Pilpres', '2018-11-11 08:37:41', '2018-11-11 08:37:41'),
(2, 'Nasional', '2018-11-11 08:38:34', '2018-11-11 08:38:34'),
(3, 'Jakarta', '2018-11-11 10:39:04', '2018-11-11 10:39:04');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `group_id` int(10) UNSIGNED NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '0',
  `phone_number` varchar(255) DEFAULT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `biography` text,
  `birthdate` datetime DEFAULT NULL,
  `gender` varchar(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `group_id`, `is_active`, `phone_number`, `photo`, `biography`, `birthdate`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'I\'m Editor', 'editor@gmail.com', '$2y$10$LfV5oMDKIwhM6biUCH6U/e2Sc5BbwSaVs2xHbcgjH7S7KhbmwX.li', 1, 1, NULL, NULL, NULL, NULL, 'M', '2018-11-11 04:06:31', '2018-11-11 04:06:31'),
(2, 'I\'m Writer', 'writer@gmail.com', '$2y$10$LfV5oMDKIwhM6biUCH6U/e2Sc5BbwSaVs2xHbcgjH7S7KhbmwX.li', 2, 1, NULL, NULL, NULL, NULL, 'F', '2018-11-11 16:36:33', '2018-11-11 16:36:33');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `access`
--
ALTER TABLE `access`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `access_name_unique` (`name`(250));

--
-- Indexes for table `access_groups`
--
ALTER TABLE `access_groups`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UNIQUE` (`group_id`,`access_id`);

--
-- Indexes for table `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`id`),
  ADD KEY `groups_title_index` (`title`(250));

--
-- Indexes for table `news`
--
ALTER TABLE `news`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `slug_i` (`slug`(250)),
  ADD KEY `title_i` (`title`(250));

--
-- Indexes for table `news_topics`
--
ALTER TABLE `news_topics`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UNIQUE` (`news_id`,`topic_id`);

--
-- Indexes for table `topics`
--
ALTER TABLE `topics`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `categories_title_index` (`title`(250)) USING BTREE;

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`(250)),
  ADD KEY `relasi_i` (`group_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `access`
--
ALTER TABLE `access`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT for table `access_groups`
--
ALTER TABLE `access_groups`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `groups`
--
ALTER TABLE `groups`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `news`
--
ALTER TABLE `news`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT for table `news_topics`
--
ALTER TABLE `news_topics`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `topics`
--
ALTER TABLE `topics`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
