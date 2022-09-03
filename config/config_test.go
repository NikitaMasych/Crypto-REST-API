package config

import "testing"

func TestThatConfigIsPresent(t *testing.T) {
	cfg := Get()

	if cfg.ServerURL == "" {
		t.Error("no server url present")
	}
	if cfg.CryptoApiFormatUrl == "" {
		t.Error("no crypto api format url present")
	}
	if cfg.BaseCurrency == "" {
		t.Error("no base currency present")
	}
	if cfg.QuotedCurrency == "" {
		t.Error("no quoted currency present")
	}
	if cfg.EmailAddress == "" {
		t.Error("no email address present")
	}
	if cfg.EmailPassword == "" {
		t.Error("no email password present")
	}
	if cfg.StorageFile == "" {
		t.Error("no storage file path present")
	}
	if cfg.SMTPHost == "" {
		t.Error("no SMTP host present")
	}
	if cfg.SMTPPort == 0 {
		t.Error("no SMTP port present")
	}
}
