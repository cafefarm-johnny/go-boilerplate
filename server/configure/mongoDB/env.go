package mongoDB

import "os"

type mongoEnv struct {
	uri      string
	db       string
	username string
	password string
}

func newMongoEnv() *mongoEnv {
	env := new(mongoEnv)

	profile := os.Getenv("profile")

	switch profile {
	case "dev":
		break
	case "stage":
		break
	case "prod":
		break
	default:
		env.uri = "mongodb://localhost:27017"
		env.db = "testdb"
		env.username = "npcdja"
		env.password = "1234"
		break
	}

	return env
}
