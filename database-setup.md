使用 MySQL 的 root 管理者帳號登入：

```
$ mysql -u root -p
```

```
-- 新增資料庫
CREATE DATABASE IF NOT EXISTS `docentre`;

-- 新增 my_user/my_password，並授予 docentre 資料庫所有權限
CREATE USER 'my_user'@'localhost' IDENTIFIED BY 'my_password';
GRANT ALL PRIVILEGES ON docentre.* TO 'my_user'@'localhost';

```

```
$ mysql -u my_user -p
```

-- connection string
```
jdbc:mysql://localhost:3306/docentre
```