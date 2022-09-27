package architectural

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestInfrastructureLayer(t *testing.T) {
	archtest.Package(t, infrastructureLayer...).ShouldNotDependOn(deliveryLayer...)
}
