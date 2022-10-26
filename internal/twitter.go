package internal

import "github.com/MikeMwita/tweeterapp/internal/common"

type Twitter struct {
	client common.SocialInterface
}

func NewClient() (common.SocialInterface, error) {

}