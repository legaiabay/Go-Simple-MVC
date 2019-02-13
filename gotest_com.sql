-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2.1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Feb 13, 2019 at 08:05 PM
-- Server version: 5.7.25-0ubuntu0.16.04.2
-- PHP Version: 7.0.32-0ubuntu0.16.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gotest_com`
--

-- --------------------------------------------------------

--
-- Table structure for table `4su_todos`
--

CREATE TABLE `4su_todos` (
  `id` int(11) NOT NULL,
  `author` varchar(255) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `id_group` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `4su_todos`
--

INSERT INTO `4su_todos` (`id`, `author`, `note`, `id_group`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Abay', 'I don\'t know what to do', 1, '2019-01-25 00:00:19', '2019-01-25 00:00:19', '2019-02-13 00:29:06'),
(2, 'Abay', 'I have to do something!', 1, '2019-01-25 00:18:40', '2019-01-25 00:18:40', '2019-02-13 00:29:37'),
(3, 'Abay', 'I have to do something!', 2, '2019-01-25 05:47:07', '2019-01-25 05:47:07', '2019-02-13 00:29:42'),
(4, 'Abay', 'I have to do something!', 3, '2019-01-25 05:49:00', '2019-01-25 05:49:00', '2019-01-25 05:49:12'),
(5, 'Abay', 'I have to do something!', 4, '2019-01-25 05:49:41', '2019-01-25 05:49:41', '2019-01-25 05:49:55'),
(6, 'Abay', 'Buat testing ahahahahahahahahah', 4, '2019-02-12 19:53:06', '2019-02-12 19:53:06', '2019-02-13 01:49:17'),
(7, 'Abay', 'aku kesal testing ahahahahahahahahah', 4, '2019-02-12 21:18:11', '2019-02-12 21:18:11', '2019-02-12 23:28:07'),
(15, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-12 23:25:20', '2019-02-12 23:25:20', '2019-02-13 00:30:08'),
(16, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 00:01:34', '2019-02-13 00:01:34', '2019-02-13 00:30:11'),
(17, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 00:01:35', '2019-02-13 00:01:35', '2019-02-13 02:00:01'),
(18, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 00:01:37', '2019-02-13 00:01:37', '2019-02-13 00:29:48'),
(19, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 00:01:38', '2019-02-13 00:01:38', '2019-02-13 01:24:28'),
(20, 'Abay', 'baru inapp', 4, '2019-02-13 01:29:09', '2019-02-13 01:29:09', '2019-02-13 03:16:05'),
(21, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:21', '2019-02-13 01:38:21', '2019-02-13 02:00:04'),
(22, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:22', '2019-02-13 01:38:22', '2019-02-13 03:16:03'),
(23, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:23', '2019-02-13 01:38:23', '2019-02-13 03:14:40'),
(24, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:24', '2019-02-13 01:38:24', '2019-02-13 02:24:46'),
(25, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:25', '2019-02-13 01:38:25', '2019-02-13 01:49:18'),
(26, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:26', '2019-02-13 01:38:26', '2019-02-13 02:03:41'),
(27, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 01:38:26', '2019-02-13 01:38:26', '2019-02-13 02:23:46'),
(28, 'Abay', 'baru inapp', 4, '2019-02-13 01:39:01', '2019-02-13 01:39:01', '2019-02-13 01:53:08'),
(29, 'Abay', 'baru inapp', 4, '2019-02-13 01:41:12', '2019-02-13 01:41:12', '2019-02-13 01:53:08'),
(30, 'Abay', 'baru inapp', 4, '2019-02-13 01:43:27', '2019-02-13 01:43:27', '2019-02-13 01:53:07'),
(31, 'Abay', 'baru inapp', 4, '2019-02-13 01:48:19', '2019-02-13 01:48:19', '2019-02-13 01:53:06'),
(32, 'Abay', 'baru inapp', 4, '2019-02-13 01:48:21', '2019-02-13 01:48:21', '2019-02-13 01:51:16'),
(33, 'Abay', 'baru inapp', 4, '2019-02-13 01:50:58', '2019-02-13 01:50:58', '2019-02-13 01:51:01'),
(34, 'Abay', 'baru inapp', 4, '2019-02-13 01:53:03', '2019-02-13 01:53:03', '2019-02-13 01:53:06'),
(35, 'Abay', 'baru inapp', 4, '2019-02-13 01:53:12', '2019-02-13 01:53:12', '2019-02-13 01:53:26'),
(36, 'Abay', 'baru inapp', 4, '2019-02-13 01:54:45', '2019-02-13 01:54:45', '2019-02-13 02:23:33'),
(37, 'Abay', 'baru inapp', 4, '2019-02-13 01:56:32', '2019-02-13 01:56:32', '2019-02-13 02:23:34'),
(38, 'Abay', 'baru inapp', 4, '2019-02-13 01:59:19', '2019-02-13 01:59:19', '2019-02-13 02:23:36'),
(39, 'Abay', 'baru inapp', 4, '2019-02-13 01:59:22', '2019-02-13 01:59:22', '2019-02-13 02:23:38'),
(40, 'Abay', 'baru inapp', 4, '2019-02-13 03:16:06', '2019-02-13 03:16:06', '2019-02-13 03:16:12'),
(41, 'Abay', 'baru inapp', 4, '2019-02-13 03:16:08', '2019-02-13 03:16:08', '2019-02-13 04:44:54'),
(42, 'Abay', 'baru inapp', 4, '2019-02-13 03:16:09', '2019-02-13 03:16:09', '2019-02-13 03:26:42'),
(43, 'Abay', 'baru inapp', 4, '2019-02-13 03:16:09', '2019-02-13 03:16:09', '2019-02-13 03:25:46'),
(44, 'Abay', 'baru inapp', 4, '2019-02-13 03:26:46', '2019-02-13 03:26:46', '2019-02-13 04:47:00'),
(45, 'Abay', 'baru inapp', 4, '2019-02-13 03:26:47', '2019-02-13 03:26:47', '2019-02-13 03:27:59'),
(46, 'Abay', 'baru inapp', 4, '2019-02-13 03:28:00', '2019-02-13 03:28:00', '2019-02-13 04:57:34'),
(47, 'Abay', 'sssssllspepepkkxnndn\nDjdjdndmuusksksl', 4, '2019-02-13 04:58:52', '2019-02-13 05:31:48', NULL),
(48, 'Abay', 'Buat baru tessssss', 4, '2019-02-13 05:06:38', '2019-02-13 05:06:46', '2019-02-13 05:07:21'),
(49, 'Abay', 'Buddddsfggggjsjjssjs', 4, '2019-02-13 05:07:30', '2019-02-13 05:31:53', NULL),
(50, 'Asrul', 'pepepepepepe kesal testing ahahahahahahahahah', 4, '2019-02-13 05:26:08', '2019-02-13 05:32:00', '2019-02-13 05:32:02');

-- --------------------------------------------------------

--
-- Table structure for table `4su_todo_groups`
--

CREATE TABLE `4su_todo_groups` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `4su_todo_groups`
--

INSERT INTO `4su_todo_groups` (`id`, `title`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'todo list', '2019-01-24 23:54:04', '2019-01-24 23:54:04', '2019-01-25 05:43:59'),
(2, 'todo list', '2019-01-25 05:46:49', '2019-01-25 05:46:49', '2019-01-25 05:47:30'),
(3, 'todo list', '2019-01-25 05:48:43', '2019-01-25 05:48:43', '2019-01-25 05:49:12'),
(5, 'todo list', '2019-01-25 05:50:36', '2019-01-25 05:50:36', '2019-01-25 05:50:45');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `4su_todos`
--
ALTER TABLE `4su_todos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `4su_todo_groups`
--
ALTER TABLE `4su_todo_groups`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `4su_todos`
--
ALTER TABLE `4su_todos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=51;
--
-- AUTO_INCREMENT for table `4su_todo_groups`
--
ALTER TABLE `4su_todo_groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
