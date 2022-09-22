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


# Project structure
```
.
├── Assignment_specification.pdf
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   ├── config_test.go
│   └── env_names.go
├── Dockerfile
├── go.mod
├── go.sum
├── pkg
│   ├── cache
│   │   ├── cache.go
│   │   ├── cache_test.go
│   │   └── utils.go
│   ├── crypto
│   │   ├── binance_rate.go
│   │   ├── binance_service.go
│   │   ├── coinapi_rate.go
│   │   ├── coinapi_service.go
│   │   ├── coinbase_rate.go
│   │   ├── coinbase_service.go
│   │   ├── cryptochain.go
│   │   ├── crypto_provider.go
│   │   └── crypto_test.go
│   ├── emails
│   │   └── email_processing.go
│   ├── errors
│   │   └── errors.go
│   ├── handlers
│   │   └── handlers.go
│   ├── logger
│   │   └── logger.go
│   ├── model
│   │   └── user.go
│   ├── platform
│   │   ├── file_assurance.go
│   │   └── platform_test.go
│   ├── presentation
│   │   └── json_presenter.go
│   ├── repository
│   │   ├── file_repository.go
│   │   ├── interface.go
│   │   └── repository_test.go
│   └── routes
│       └── routes.go
└── README.md

```
Good reviewing!
