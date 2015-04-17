LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (username,email,password,name,birthday,gender,md5,activation,token,token_expiration) VALUES ('admin', 'admin@goyangi.github.io','$2a$10$voqxhv08H2eWHbLJo2rEeO1GwGlg8ZLW3Y8348aqe0XBqVgEZxGOu','Goyangi', '2014-12-01 04:05:20',2,'10d17498672e2dd040e8c0cf5a337a61',1,'168355cf5b6d31827c694260ab24e3bc3e990290ca94c7c30c6489ae1c1f212c','2999-12-31 00:00:00');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
SET @userId = LAST_INSERT_ID();

INSERT IGNORE INTO role (name,description) VALUES ('admin','administrator');
SET @roleId = LAST_INSERT_ID();
INSERT INTO users_roles (user_id,role_id) VALUES(@userId, @roleId);

INSERT IGNORE INTO role (name,description) VALUES ('user','general user');
SET @roleId = LAST_INSERT_ID();
INSERT INTO users_roles (user_id,role_id) VALUES(@userId, @roleId);
