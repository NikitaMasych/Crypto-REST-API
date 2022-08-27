Morning!

This is technical assignment for the Software Engineering school by Genesis&&KMA

Made with golang (gin, in particular) and love 🤗

I tried to comment code, in order to make it more readable, however for any questions you can leave me a message:
  * @Just_law_abiding_citizen - telegram 
  * nikitamasich152@gmail.com - email

For the task I registered new google account genesisbriefingnm@gmail.com, enabled 2-factor authentication and thereby got the opportunity to generate app password (needed due to the google policy regarding to the less secure apps).
Solution uses google SMTP server for sending emails and because this is a free version, it allows to send no more than 100 letters a day.

I added code 400 response in POST request /subscribe - if provided email is not valid (though no verification whether it actually exists)

Running and listening on the localhost:8080 port 

Via docker: 

$ docker build -t genapp .

$ docker run -d -p 8080:8080 genapp

Good reviewing!
