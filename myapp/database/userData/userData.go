package userData

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID         int    `db:"id"`
	Name       string `db:"user_name"`
	Password   string `db:"password"`
	Authority  string `db:"authority"`
	InvalidFlg bool   `db:"invalid_flg"`
}

type Authority struct {
	Authority string `db:"authority"`
	Name      string `db:"name"`
}

func AddUser(name string, password string, authority string) {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO gogo.users (user_name, password, authority, invalid_flg) VALUES ($1, $2, $3, $4)", name, password, authority, false)
	tx.Commit()
}

func GetAllUsers() []User {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	users := []User{}
	db.Select(&users, "SELECT * FROM gogo.users")

	return users
}

func GetAllAuthorities() []Authority {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	authorities := []Authority{}
	db.Select(&authorities, "SELECT * FROM gogo.authorities")

	return authorities
}

func GetUserByName(name string) User {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	user := User{}
	err = db.Get(&user, "SELECT * FROM gogo.users WHERE user_name=$1", name)
	return user
}

func GetUserById(id int) User {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	user := User{}
	err = db.Get(&user, "SELECT * FROM gogo.users WHERE id=$1", id)
	return user
}

func Truncate() {
	db, err := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres dbname=godb sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("TRUNCATE TABLE gogo.users")
	tx.Commit()
}
