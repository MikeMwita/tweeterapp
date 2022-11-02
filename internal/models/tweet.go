package models

type Tweet struct {
	Id      string `json:"id"`
	Handle  string `json:"handle"`
	Body    string `json:"body"`
	Likes   int    `json:"likes"`
	Retweet int    `json:"retweet"`
}
