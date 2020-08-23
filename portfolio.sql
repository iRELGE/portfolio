-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1
-- Généré le : Dim 23 août 2020 à 20:49
-- Version du serveur :  10.4.13-MariaDB
-- Version de PHP : 7.2.32

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `portfolio`
--

-- --------------------------------------------------------

--
-- Structure de la table `comment`
--

CREATE TABLE `comment` (
  `ID` int(10) NOT NULL,
  `project_id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `Subject` varchar(255) DEFAULT NULL,
  `Note` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Structure de la table `project`
--

CREATE TABLE `project` (
  `ID` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `Description` varchar(255) DEFAULT NULL,
  `Photo` varchar(255) DEFAULT NULL,
  `Link` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `project`
--

INSERT INTO `project` (`ID`, `user_id`, `Title`, `Description`, `Photo`, `Link`) VALUES
(1, 2, 'title1', 'desc1', '2020-08-19-10-18-45Class Diagram1.jpg', 'link1'),
(2, 3, 'title1', 'desc1', '2020-08-19-10-35-06Class Diagram1.jpg', 'link1'),
(3, 3, 'title1', 'desc1', '2020-08-19-10-35-10Class Diagram1.jpg', 'link1'),
(4, 5, 'title1', 'desc1', '2020-08-19-11-09-35Class Diagram1.jpg', 'link1'),
(5, 6, 'title1', 'desc1', '2020-08-19-12-48-21Class Diagram1.jpg', 'link1'),
(6, 6, 'title1', 'desc1', '2020-08-19-15-12-49Class Diagram1.jpg', 'link1');

-- --------------------------------------------------------

--
-- Structure de la table `roles`
--

CREATE TABLE `roles` (
  `ID` int(10) NOT NULL,
  `Role` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `roles`
--

INSERT INTO `roles` (`ID`, `Role`) VALUES
(1, 'Admin'),
(2, 'User');

-- --------------------------------------------------------

--
-- Structure de la table `user`
--

CREATE TABLE `user` (
  `ID` int(10) NOT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Password` varchar(255) DEFAULT NULL,
  `Status` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `user`
--

INSERT INTO `user` (`ID`, `Email`, `Password`, `Status`) VALUES
(1, 'test@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1),
(2, 'test1@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1),
(3, 'test2@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1),
(4, 'test3@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1),
(5, 'test4@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1),
(6, 'test5@gmail.com', 'e755c10fb83db73fe371e1c3f976c17020f7a150859b20100cbfe253e21df431', 1);

-- --------------------------------------------------------

--
-- Structure de la table `userprofile`
--

CREATE TABLE `userprofile` (
  `ID` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Name` varchar(255) DEFAULT NULL,
  `LastName` varchar(255) DEFAULT NULL,
  `Photo` varchar(255) DEFAULT NULL,
  `Adress` varchar(255) DEFAULT NULL,
  `Phone` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `userprofile`
--

INSERT INTO `userprofile` (`ID`, `user_id`, `Email`, `Name`, `LastName`, `Photo`, `Adress`, `Phone`) VALUES
(1, 1, 'test@gmail.com', 'rabie', 'elgouaill', '2020-08-19-10-34-36WIN_20200811_11_37_29_Pro.jpg', 'marrakech azli', '06666666678'),
(2, 2, 'test1@gmail.com', '', '', '', '', ''),
(3, 3, 'test2@gmail.com', '', '', '', '', ''),
(4, 4, 'test3@gmail.com', '', '', '', '', ''),
(5, 5, 'test4@gmail.com', 'rabie', 'elgouaill', '2020-08-19-11-09-08WIN_20200811_11_37_29_Pro.jpg', 'marrakech azli', '06666666678'),
(6, 6, 'test5@gmail.com', 'rabie', 'elgouaill', '2020-08-19-12-48-02WIN_20200811_11_37_29_Pro.jpg', 'marrakech azli', '06666666678');

-- --------------------------------------------------------

--
-- Structure de la table `user_roles`
--

CREATE TABLE `user_roles` (
  `ID` int(10) NOT NULL,
  `roles_id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Déchargement des données de la table `user_roles`
--

INSERT INTO `user_roles` (`ID`, `roles_id`, `user_id`) VALUES
(6, 1, 6),
(1, 2, 1),
(2, 2, 2),
(3, 2, 3),
(4, 2, 4),
(5, 2, 5);

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `FKcomment868090` (`user_id`),
  ADD KEY `FKcomment881308` (`project_id`);

--
-- Index pour la table `project`
--
ALTER TABLE `project`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `FKproject842422` (`user_id`);

--
-- Index pour la table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`ID`);

--
-- Index pour la table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`);

--
-- Index pour la table `userprofile`
--
ALTER TABLE `userprofile`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `FKuserprofil492294` (`user_id`);

--
-- Index pour la table `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`ID`,`user_id`) USING BTREE,
  ADD KEY `FKuser_roles214972` (`user_id`),
  ADD KEY `FKuser_roles764398` (`roles_id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `comment`
--
ALTER TABLE `comment`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `project`
--
ALTER TABLE `project`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `roles`
--
ALTER TABLE `roles`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `user`
--
ALTER TABLE `user`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `userprofile`
--
ALTER TABLE `userprofile`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `ID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `comment`
--
ALTER TABLE `comment`
  ADD CONSTRAINT `FKcomment868090` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`),
  ADD CONSTRAINT `FKcomment881308` FOREIGN KEY (`project_id`) REFERENCES `project` (`ID`);

--
-- Contraintes pour la table `project`
--
ALTER TABLE `project`
  ADD CONSTRAINT `FKproject842422` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`);

--
-- Contraintes pour la table `userprofile`
--
ALTER TABLE `userprofile`
  ADD CONSTRAINT `FKuserprofil492294` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`);

--
-- Contraintes pour la table `user_roles`
--
ALTER TABLE `user_roles`
  ADD CONSTRAINT `FKuser_roles214972` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`),
  ADD CONSTRAINT `FKuser_roles764398` FOREIGN KEY (`roles_id`) REFERENCES `roles` (`ID`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
