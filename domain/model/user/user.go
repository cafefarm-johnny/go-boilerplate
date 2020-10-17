package user

type User struct {
	Id   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
	Team Team   `bson:"team" json:"team"`
}

type Team struct {
	Name string `bson:"name" json:"name"`
}
