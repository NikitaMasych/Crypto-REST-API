package architectural

const (
	domainModels   = "producer/pkg/domain/models"
	domainServices = "producer/pkg/domain/services"
	application    = "producer/pkg/application"
	handlers       = "producer/pkg/delivery/handlers"
	presentors     = "producer/pkg/delivery/presentors"
	crypto         = "producer/pkg/infrastructure/crypto"
	email          = "producer/pkg/infrastructure/email"
	logger         = "producer/pkg/infrastructure/logger"
	cache          = "producer/pkg/infrastructure/storage/cache"
	storage        = "producer/pkg/infrastructure/storage/emails_repository"
)
