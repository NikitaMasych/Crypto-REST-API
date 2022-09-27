package architectural

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestDomainLayer(t *testing.T) {
	layers := append(applicationLayer, infrastructureLayer...)
	layers = append(layers, deliveryLayer...)
	archtest.Package(t, domainLayer...).ShouldNotDependOn(layers...)
}
