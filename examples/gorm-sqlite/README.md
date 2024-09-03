# GORM examples

Install sqlite3 on Ubuntu:

```bash
sudo apt update
sudo apt install sqlite3
sqlite3 --version
```

```bash
go get .
```

```bash
rm -f test.db
go main.go
```

```bash
$ sqlite3
SQLite version 3.31.1 2020-01-27 19:55:54
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .open test.db
sqlite> SELECT * from users;
1|2023-09-09 15:41:39.4630017+08:00|2023-09-09 15:41:39.4630017+08:00||1|alex
sqlite> .quit
$
```
