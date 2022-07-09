package user

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Jasoner struct {
	Id     string `bson:"_id"`
	Name   string
	Family string
	Age    int
	Pone string
}