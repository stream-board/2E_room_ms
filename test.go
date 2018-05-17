package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
  db, err := gorm.Open("mysql", "roomsUser:123@tcp(192.168.99.103:3306)/rooms?charset=utf8&parseTime=True&loc=Local")
  if err == nil{
  	fmt.Println("NO HAY ERROR")
  	fmt.Println(err)
  }else{
  	fmt.Println("HAY ERROR")
  	fmt.Println(err)
  }
  defer db.Close()
}