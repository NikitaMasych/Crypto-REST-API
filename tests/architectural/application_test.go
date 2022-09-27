package architectural

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestApplicationLayer(t *testing.T) {
	layers := append(deliveryLayer, infrastructureLayer...)
	archtest.Package(t, applicationLayer...).ShouldNotDependOn(layers...)
}
