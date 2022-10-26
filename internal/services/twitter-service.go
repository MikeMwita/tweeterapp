package services

type TwitterService interface {
	Auth() error
}

type DefaultTwService struct {
}
