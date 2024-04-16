
mysql in docker 
```bash
$ docker run --name docentre-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
$ docker exec -it docentre-mysql bash
```

使用 MySQL 的 root 管理者帳號登入：

```
$ mysql -u root //無密碼
$ mysql -u root -p //需要密碼
```

```
-- 新增資料庫
CREATE DATABASE IF NOT EXISTS `docentre`;
-- 新增 my_user/my_password，並授予 docentre 資料庫所有權限
CREATE USER 'my_user'@'localhost' IDENTIFIED BY 'my_password';
GRANT ALL PRIVILEGES ON docentre.* TO 'my_user'@'localhost';
```

退出 root 改用 my_user 登入
```
$ mysql -u my_user -p
```

connection string
```
jdbc:mysql://localhost:3306/docentre
```
