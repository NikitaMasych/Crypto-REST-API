# Description:
This is REST-API based project written using golang and gin framework.
Four services, MySQL, Redis cache and RabbitMQ as a message broker communication and deployment are set up using docker compose.
In producer, when subscribing specific user, is created request to the service orders to register new customer via saga-based approach, redirecting requests to the customers service. For this purpose, project uses transaction manager dtm. Implemented using DDD and onion-based architecture.

# Use cases:
* Subscribe particular email for rate digest on a specified currency pair.
* Get rate for a specified pair.
* Send emails with corresponding rates to subscribed emails.

# URL Paths:
```
    http://localhost:8080/api/rate       -> POST
    http://localhost:8080/api/subscribe  -> POST 
    http://localhost:8080/api/sendEmails -> POST
```

# Cache:
For caching rates purpose project uses Redis database, which is connected via docker-compose. Default expiration time is set for 5 minutes, however could be modified in .env file, configuring appropriate value.

# Message broker:
This project uses RabbitMQ as a message broker.
There are two independent services: the one for producing and another for consuming log messages. 
Forth one publishes logs in three queues, accordingly to log types: "debug", "error" and "info" and the second one constantly reads "error" queue and outputs using CLI.

# Deployment using docker compose: 
```
$ docker compose up --build
```

# Producer architecture:
![http://url/to/img.png](https://github.com/GenesisEducationKyiv/hw1-se-school_2022-code-review-NikitaMasych/blob/hw6/docs/Architecture.png)

# Project structure
```
.
├── consumer
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── customers
│   ├── cmd
│   │   ├── main.go
│   │   └── setup.go
│   ├── config
│   │   └── config.go
│   ├── delivery
│   │   ├── handlers
│   │   │   ├── handlers.go
│   │   │   ├── register_customer_compensate.go
│   │   │   └── register_customer.go
│   │   └── routes
│   │       └── init_routes.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── storage
│   │   └── database_utils.go
│   └── types
│       └── types.go
├── docker-compose.yml
├── orders
│   ├── cmd
│   │   └── main.go
│   ├── config
│   │   └── config.go
│   ├── delivery
│   │   ├── handlers
│   │   │   └── create_customer_handler.go
│   │   └── routes
│   │       └── init_routes.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── types
│       └── types.go
├── producer
│   ├── cmd
│   │   ├── main.go
│   │   └── setup.go
│   ├── config
│   │   ├── config.go
│   │   └── config_test.go
│   ├── Dockerfile
│   ├── docs
│   │   └── Architecture.png
│   ├── go.mod
│   ├── go.sum
│   ├── pkg
│   │   ├── application
│   │   │   ├── contracts.go
│   │   │   ├── email_ucase.go
│   │   │   ├── rate_ucase.go
│   │   │   └── subscription_ucase.go
│   │   ├── delivery
│   │   │   ├── handlers
│   │   │   │   ├── handlers.go
│   │   │   │   ├── rate.go
│   │   │   │   ├── sendEmails.go
│   │   │   │   └── subscribe.go
│   │   │   └── presentors
│   │   │       └── json_presenter.go
│   │   ├── domain
│   │   │   ├── models
│   │   │   │   ├── currency_pair.go
│   │   │   │   ├── currency_rate.go
│   │   │   │   ├── email_address.go
│   │   │   │   ├── subscription.go
│   │   │   │   └── user.go
│   │   │   └── services
│   │   │       ├── email_service.go
│   │   │       ├── rate_service.go
│   │   │       └── subscription_service.go
│   │   ├── errors
│   │   │   └── errors.go
│   │   ├── infrastructure
│   │   │   ├── crypto
│   │   │   │   ├── binance_provider.go
│   │   │   │   ├── coinapi_provider.go
│   │   │   │   ├── coinbase_provider.go
│   │   │   │   ├── crypto_test.go
│   │   │   │   ├── logger_aux.go
│   │   │   │   ├── providers_chain.go
│   │   │   │   └── providers_urls.go
│   │   │   ├── customers
│   │   │   │   └── customer_creation.go
│   │   │   ├── email
│   │   │   │   └── email_sender.go
│   │   │   ├── logger
│   │   │   │   ├── constructor.go
│   │   │   │   ├── logtypes
│   │   │   │   │   └── logtypes.go
│   │   │   │   ├── rabbitmq
│   │   │   │   │   └── rabbitmq_logger.go
│   │   │   │   └── txt
│   │   │   │       └── txt_logger.go
│   │   │   └── storage
│   │   │       ├── cache
│   │   │       │   ├── go-cache
│   │   │       │   │   ├── cache.go
│   │   │       │   │   └── cache_test.go
│   │   │       │   └── redis
│   │   │       │       ├── cache.go
│   │   │       │       └── cache_test.go
│   │   │       └── subscription_repository
│   │   │           ├── file_subscription_repository.go
│   │   │           └── file_subscription_repository_test.go
│   │   └── utils
│   │       ├── file_assurance.go
│   │       └── utils_test.go
│   └── tests
│       └── architectural
│           ├── application_test.go
│           ├── delivery_test.go
│           ├── domain_test.go
│           ├── infrastructure_test.go
│           ├── layer_names.go
│           └── package_names.go
└── README.md

```
