# Morning!

This is technical assignment for the Software Engineering school by Genesis&&KMA

Made with golang (gin, in particular) and love 🤗

For the task I registered new google account genesisbriefingnm@gmail.com, enabled 2-factor authentication and thereby got the opportunity to generate app password (needed due to the google policy regarding to the less secure apps).
Solution uses google SMTP server for sending emails and because this is a free version, it allows to send no more than 100 letters a day.

Running and listening on the localhost:8080 port 

# Docker: 
```bash
$ docker build -t genapp .

$ docker run -d -p 8080:8080 genapp
```

# Project architecture:
![http://url/to/img.png](https://github.com/GenesisEducationKyiv/hw1-se-school_2022-code-review-NikitaMasych/blob/hw6/docs/Architecture.png)

# Project structure
```
.
├── cmd
│   ├── main.go
│   └── setup.go
├── config
│   ├── config.go
│   └── config_test.go
├── Dockerfile
├── docs
│   └── Architecture.png
├── go.mod
├── go.sum
├── pkg
│   ├── application
│   │   ├── contracts.go
│   │   ├── email_ucase.go
│   │   ├── rate_ucase.go
│   │   └── subscription_ucase.go
│   ├── delivery
│   │   ├── handlers
│   │   │   ├── handlers.go
│   │   │   ├── rate.go
│   │   │   ├── sendEmails.go
│   │   │   └── subscribe.go
│   │   └── presentors
│   │       └── json_presenter.go
│   ├── domain
│   │   ├── models
│   │   │   ├── currency_pair.go
│   │   │   ├── currency_rate.go
│   │   │   ├── email_address.go
│   │   │   ├── subscription.go
│   │   │   └── user.go
│   │   └── services
│   │       ├── email_service.go
│   │       ├── rate_service.go
│   │       └── subscription_service.go
│   ├── errors
│   │   └── errors.go
│   ├── infrastructure
│   │   ├── crypto
│   │   │   ├── binance_provider.go
│   │   │   ├── coinapi_provider.go
│   │   │   ├── coinbase_provider.go
│   │   │   ├── crypto_test.go
│   │   │   ├── providers_chain.go
│   │   │   └── providers_urls.go
│   │   ├── email
│   │   │   └── email_sender.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   └── storage
│   │       ├── cache
│   │       │   ├── go-cache
│   │       │   │   ├── cache.go
│   │       │   │   └── cache_test.go
│   │       │   └── redis
│   │       │       ├── cache.go
│   │       │       └── cache_test.go
│   │       └── subscription_repository
│   │           ├── file_subscription_repository.go
│   │           └── file_subscription_repository_test.go
│   └── utils
│       ├── file_assurance.go
│       └── utils_test.go
├── README.md
└── tests
    └── architectural
        ├── application_test.go
        ├── delivery_test.go
        ├── domain_test.go
        ├── infrastructure_test.go
        ├── layer_names.go
        └── package_names.go

```
