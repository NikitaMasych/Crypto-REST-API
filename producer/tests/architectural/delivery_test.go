package architectural

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestDeliveryLayer(t *testing.T) {
	archtest.Package(t, deliveryLayer...).ShouldNotDependOn(infrastructureLayer...)
}
