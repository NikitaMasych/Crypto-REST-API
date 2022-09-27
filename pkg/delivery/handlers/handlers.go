package handlers

type Handlers struct {
	RateHandler
	SubscribeHandler
	SendRateEmailsHandler
}

func NewHandlers(h1 *RateHandler,
	h2 *SubscribeHandler,
	h3 *SendRateEmailsHandler) *Handlers {
	return &Handlers{*h1, *h2, *h3}
}
