package models

type Tweet struct {
	Id      string `json:"id"`
	Handle  string `json:"handle"`
	Search   string `json:"search"`
	Body    string `json:"body"`
	Likes   int    `json:"likes"`
	Retweet int    `json:"retweet"`
}




