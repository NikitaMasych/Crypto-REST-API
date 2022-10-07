package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsConfigPresent(t *testing.T) {
	require.NotEmpty(t, ServerUrl)
	require.NotEmpty(t, CoinApiKey)
	require.NotEmpty(t, GenesisProvider)
	require.NotEmpty(t, EmailAddress)
	require.NotEmpty(t, EmailPassword)
	require.NotEmpty(t, StorageFile)
	require.NotEmpty(t, DebugLogFile)
	require.NotEmpty(t, ErrorsLogFile)
	require.NotEmpty(t, InfoLogFile)
	require.NotEmpty(t, SMTPHost)
	assert.NotEqual(t, SMTPPort, 0)
	require.NotEmpty(t, CacheHost)
	assert.NotEqual(t, CacheDb, 0)
	assert.NotEqual(t, CacheDurationMins, 0)
	require.NotEmpty(t, CurrencyPairSeparator)
	require.NotEmpty(t, EmailAddressSeparator)
}
