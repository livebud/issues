package controller

func Load() (*Controller, error) {
	return &Controller{}, nil
}

type Controller struct {
}

func (c *Controller) Index() string {
	return ""
}
