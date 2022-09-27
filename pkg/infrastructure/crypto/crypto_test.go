package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatBinanceProviderReturnsRate(t *testing.T) {
	p := NewBinanceProvider()
	s := config.NewConfigPairSource()
	pair := s.GetPair()
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinbaseProviderReturnsRate(t *testing.T) {
	p := NewCoinbaseProvider()
	s := config.NewConfigPairSource()
	pair := s.GetPair()
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinApiProviderReturnsRate(t *testing.T) {
	p := NewCoinApiProvider()
	s := config.NewConfigPairSource()
	pair := s.GetPair()
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func provideArrange() {
	platform.EnsureFileExists(config.Get().LoggerFile)
}
