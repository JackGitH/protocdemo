DROP TABLE IF EXISTS `events_client_address`;
CREATE TABLE `events_client_address` (
  `ID` char(64) NOT NULL,
  `ECLIENTIP` char(64) NOT NULL,
  `ECLIENTPORT` char(64) NOT NULL,
  `ETIME` bigint(20) NOT NULL,
  `REMARK` varchar(512) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;