package common

type SocialInterface interface {
	Auth() error
}
type defaultSocialInterface struct {

	//what do we need for this --->this is where we inject the adapters

}
