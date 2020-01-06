package main

import (
	"fmt"
	"model"
)


func main()  {

	var user model.User
	ret := user.GetById(2113894)
	if ret == false {
		fmt.Println("GetById fail.")
		return
	}

	fmt.Println(user.Account)
}
