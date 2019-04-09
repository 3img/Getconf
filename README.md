# Getconf
get conf from local file

> If you has file like this format:
```
;Annotation here, start from ;
[MYSQL]
DBNAME = comments
DBHOST = 10.10.0.211
DBPORT = 30521
DBUSER = root
DBPASS = 123456

[LISTEN]
HOST = 127.0.0.1
PORT = 8080
```

> And the file name is 'config.conf', you can use Getconf to get the value from this file. code like this:
```
c, err := Getconf.NewFileConf("./config.ini")
if err != nil {
  fmt.Println(err)
  return
}

fmt.Println("LISTEN.PORT")
```

