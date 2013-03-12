package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	DB_NAME = "webdemo"
	DB_USER = "root"
	DB_PASS = ""
)

type User struct {
	id      int    //`json:"uid"`
	name    string //`json:"userName"`
	age     string //`json:"departname"`
	addtime string //`json:"created"`
}

type Admin struct {
	admin_id       int
	admin_name     string
	admin_password string
}

func OpenDB() *sql.DB {
	//    db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/webdemo?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

//func UserById(id int) *Admin {
//	db := OpenDB()
//	defer db.Close()
//	fmt.Print(id)
//	row := db.QueryRow("SELECT * FROM `userinfo` WHERE uid=?", id)
//	admin := new(Admin)
//	row.Scan(&admin.Uid, &admin.UserName, &admin.Created)
//	return admin
//}

func DelectById() {
	db := OpenDB()
	defer db.Close()
	db.Query("DELETE FROM userinfo WHERE uid =2")
}

func FindAdmin() {
	db := OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `admin`")
	if err != nil {
	}

	admin := new(Admin)
	for rows.Next() {
		err = rows.Scan(&admin.admin_id, &admin.admin_name, &admin.admin_password)
		fmt.Print(admin.admin_password + "\t\n")

	}
}

func FindUser(){
	db := OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `user2`")
	if err != nil {
	}

	user := new(User)
	for rows.Next() {
		err = rows.Scan(&user.id, &user.name, &user.age)
		fmt.Print(user.name + "\n")
	}
}

func main() {
	FindUser()
}
