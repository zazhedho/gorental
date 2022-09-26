# Go-rent Golang

<h1 align="center">
  Vehicle-Rent RESTfull API With Gorilla/mux, Gorm
</h1>
<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png" width="400px" alt="Golang.jpg" /></p>
<p align="center">
    <a href="https://golang.org/" target="blank">More about Golang</a>
</p>

## 🔗 Description

This Backend Application is used for vehicle rental systems such as car rental, motorbikes, and bicycles. In the application, users can add, change, delete, and read the data of the vehicle they want to rent. In addition, users can also see the rental history. This application was built using the Golang programming language with the Gorilla / Mux Framework and uses GORM, a Database that is used using PostgreSQL.

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/zazhedho/gorental.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Run the app

```bash
go run main.go server
```

4. Add Env

```sh
  APP_PORT= Your Port
  JWT_KEYS= Your Secret Keys

  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
```

5. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration
# or
go run main.go migrate --down //for rollback
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [gorilla/mux](https://github.com/gorilla/mux): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS

## 🚀 About Me

- Linkedin : [Zaidus Zhuhur](https://www.linkedin.com/in/zaidus-zhuhur/)
