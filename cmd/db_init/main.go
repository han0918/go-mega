package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sample/model"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("afei", "abc123", "111@qq.com")
	model.AddUser("jason", "abc123", "22@qq.com")

	u1, _ := model.GetUserByUsername("afei")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, "hello about me")

	u2, _ := model.GetUserByUsername("jason")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u1.Follow(u2.Username)
}
