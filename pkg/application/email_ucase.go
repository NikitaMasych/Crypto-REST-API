package application

type EmailSenderRepository struct {
	storage    EmailAddressesStorage
	sender     EmailSender
	exchanger  RateRepository
	pairSource PairSource
}

func NewEmailSenderRepository(storage EmailAddressesStorage,
	sender EmailSender, exchanger RateRepository, pairSource PairSource) *EmailSenderRepository {
	return &EmailSenderRepository{storage, sender, exchanger, pairSource}
}

func (r *EmailSenderRepository) SendRateEmails() error {
	addresses := r.storage.GetAll()
	pair := r.pairSource.GetPair()
	rate, err := r.exchanger.GetRate(pair)
	if err != nil {
		return err
	}
	r.sender.SendRateEmails(rate, addresses)
	return nil
}
