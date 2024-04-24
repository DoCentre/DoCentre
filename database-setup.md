

### 使用 docker 建立 MySQL
```bash
$ docker run --name docentre-sql \
-p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=root \
-e MYSQL_DATABASE=docentre \
-e MYSQL_USER=my_user \
-e MYSQL_PASSWORD=my_password \
-d mysql
```

此時已經可以跑後端 service

### 若需進入 MySQL 容器確認

```
$ docker exec -it docentre-sql bash
```

```
$ mysql -u root -p
Password: root
// or
$ mysql -u my_user -p
Password: my_password
```

### 可視化工具連接 MySQL connection string

- username/password: 
  - root/root or my_user/my_password
- database: docentre
- host: localhost
- port: 3306
