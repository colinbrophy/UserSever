package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
type User struct {
	Id int `db:"id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
}

var db* sqlx.DB

const databaseStr  = "dbname=users user=colin password=password port=8081 sslmode=disable"
func initDatabase() {
	db = sqlx.MustConnect("postgres", databaseStr)
}

func Users() (users []User) {
	db.Select(&users,"SELECT * FROM public.users")
	return
}

func UserFromDb(id int) (usr User) {
	db.Get(&usr,"SELECT * FROM public.users WHERE id=$1", id)
	return
}

func DeleteUsr(id int) {
	db.MustExec("DELETE FROM public.users WHERE id=$1", id)
}

func UpdateUsr(id int, FirstName string, LastName string) {
	db.MustExec(
		"UPDATE public.users SET first_name=$1, last_name=$2 WHERE id=$3",
		  FirstName, LastName, int(id))
}

func CreateUsr(FirstName string, LastName string) {
	db.MustExec(`INSERT INTO public.users (id, first_name, last_name)
	 VALUES (nextval('public.user_id_seq'), $1, $2)`,
		FirstName, LastName)
}