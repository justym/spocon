package pkg

import (
	"os"
)

var (
	clientID     string
	clientSecret string
)

// これでもいいんだけど、init関数で読んであげれば、最初に呼ばれるからそっちのが良さそうかな
//　二つならまだいいけど、もしこれ以上増えるならgithub.com/kelseyhightower/envconfigを使うことを考えるといいかも！
//LoadEnv loads value from .env file
func LoadEnv() {
	clientID = os.Getenv("SPOTIFY_ID")
	clientSecret = os.Getenv("SPOTIFY_KEY")
}

//GetValues return clientID clientSecret
func GetValues() (string, string) {
	return clientID, clientSecret
}
