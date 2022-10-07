package crypto

import (
	"GenesisTask/pkg/application"
	"fmt"
	"time"

	"gopkg.in/resty.v0"
)

var logger application.Logger

func ComposeProviderResponseLog(timestamp time.Time, provider string, resp *resty.Response) string {
	return fmt.Sprintf("%s response at %s: %s, %s", provider,
		timestamp.String(), resp.Status(), string(resp.Body))
}
