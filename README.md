# goRestApi
This is a very basic example of a restFul Api in Golang.

## Usage
This package depends on beego (https://beego.me/)
```
go get -u github.com/astaxie/beego
```
```
go get -u github.com/beego/bee
```
```
go get github.com/tkanos/gonfig
```
```
go get github.com/go-sql-driver/mysql
```
```
cd /path/to/go/src/streamApi
```
```
cp secrets.yml.example secrets.yml
```

Edit secrets.yml as needed (for database access).
...and finally:

```
bee run -gendoc=true -downdoc=true
```
