package internal

import (
	"github.com/MikeMwita/tweeterapp.git/internal/common"
	"log"
)

type Twitter struct {
	client common.SocialInterface
}

func NewClient() (common.SocialInterface, error) {
	if err != nil {
		log.Printf("unable to find the common controller")
	}
}
