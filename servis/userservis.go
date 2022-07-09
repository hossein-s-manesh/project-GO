package servis

import (
	"fmt"
	"log"
	"myapp/model/user"
	"myapp/repository"
)

type UserServis interface{
	GetServis() ([]user.Jasoner,error )

}

type userServis struct{}


func NewUserServis()UserServis{
	return userServis{}
}



func(userServis) GetServis() ([]user.Jasoner,error ){
	userRepository :=repository.NewUserRipository()
	userList,err:=userRepository.GetUserList()
	user,err:=userRepository.GetById("62b3662376d085ebba7ce737")
	if err !=nil{
		log.Fatalln(err)
	}
	fmt.Println(user)
	return userList,err
}