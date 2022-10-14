package handlers

type Handlers struct {
	*RegisterCustomerHandler
	*RegisterCustomerCompensateHandler
}

func NewHandlers(h1 *RegisterCustomerHandler, h2 *RegisterCustomerCompensateHandler) *Handlers {
	return &Handlers{h1, h2}
}
