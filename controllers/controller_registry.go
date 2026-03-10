package controllers

type ControllerRegistry struct {
	Redirect RedirectController
}

func NewControllerRegistry(redirect RedirectController) *ControllerRegistry {
	return &ControllerRegistry{Redirect: redirect}
}
