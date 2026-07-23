package main

import "fmt"

//Interface
type IDbConnect interface {
	connect()
}

type dBConnect struct {
	coonect IDbConnect
}

func (db dBConnect) DBConnect() {
	db.coonect.connect()
}

type mSQLConnect struct {
	connectionstring string
}

func (msql mSQLConnect) connect() {
	fmt.Println("MYSQL Connection")
}

type postConnect struct {
	connectionstring string
}

func (post postConnect) connect() {
	fmt.Println("Postgres Connection")
}

type mongoConnect struct {
	connectionstring string
}

func (mongo mongoConnect) connect() {
	fmt.Println("Mongo Connection")
}

func main() {

	mysql := mSQLConnect{"MYSQl connection"}
	con := dBConnect{mysql}
	con.DBConnect()

	postsql := postConnect{"Postgres connection"}
	con1 := dBConnect{postsql}
	con1.DBConnect()

	mongo := mongoConnect{"MongoDB connection"}
	con = dBConnect{mongo}
	con.DBConnect()
}
