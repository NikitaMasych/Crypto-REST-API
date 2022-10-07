package architectural

var (
	domainLayer         = []string{domainModels, domainServices}
	applicationLayer    = []string{application}
	deliveryLayer       = []string{handlers, presentors}
	infrastructureLayer = []string{crypto, email, logger, cache, storage}
)
