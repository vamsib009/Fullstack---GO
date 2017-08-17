### Steps to run the project

#### Requires:

* ![golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

* ![github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

### How To Run

Create a new EC2 instance  in AWS
Create a new database in AWS RDS with a users table


```sql
CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120)
);
```

Go get both required packages listed below

```bash
go get golang.org/x/crypto/bcrypt

go get github.com/go-sql-driver/mysql
```

Inside of **signup.go** line **169** replace <example> with your own credentials

```go
db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
// Replace with
db, err = sql.Open("mysql", "myUsername:myPassword@/myDatabase")
```


### go run signup.go
