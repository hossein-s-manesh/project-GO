package servis

import(
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
	return userList,err
}