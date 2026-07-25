package aerogram

import "log"

type Client struct {
	Token string
}

func New(token string) {
	log.Println(token)
}
