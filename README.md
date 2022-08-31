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
├── config
│   ├── config.go
│   └── env_names.go
├── crypto
│   └── crypto.go
├── Dockerfile
├── emails
│   ├── email_data.go
│   └── email_processing.go
├── errors
│   └── errors.go
├── go.mod
├── go.sum
├── main.go
├── platform
│   └── subscribed_emails.txt
├── README.md
└── routes
    └── routes.go
```
Good reviewing!
