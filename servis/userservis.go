package servis

import(
	"myapp/model/user"


)

type UserServis interface{
	GetServis() ([]user.Jasoner,error )

}

type userServis struct{}


func NewUserServis()UserServis{
	return userServis{}
}



func(userServis) GetServis() ([]user.Jasoner,error ){
	listjson:=[]user.Jasoner{
		{
			Name :"hossein",
			Family : "soltan",
			Age : 18,
			Fone :"0203494039",

		},
		{
			Name :"syrus",
			Family : "soltan",
			Age : 56,
			Fone :"0203494039",

		},
		{
			Name :"nikoo",
			Family : "soltan",
			Age : 23,
			Fone :"0203494039",

		},
		{
			Name :"victor",
			Family : "soltan",
			Age : 54,
			Fone :"0203494039",

		},
		{
			Name :"kamran",
			Family : "soltan",
			Age : 56,
			Fone :"0203494039",

		},
		{
			Name :"sajad",
			Family : "soltan",
			Age : 45,
			Fone :"0203494039",

		},
		{
			Name :"sohil",
			Family : "soltan",
			Age : 34,
			Fone :"0203494039",

		},
	}

	return listjson,nil
}