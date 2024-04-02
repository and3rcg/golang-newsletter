## About the project and objectives
lorem ipsum from hell

## Objectives
ipsum lorem

## Running the system
In order to run this system, just run `docker-compose up` and make sure you have Docker running in your system. To shut down the system, run `docker-compose down`, you can even add the `-v` argument to this command to remove the volumes.

## Known issues
1. For some reason, the web container would start up before the database one was fully started up, even after adding `depends_on` in the Compose file. If you see `Connection refused` in the web container's logs, restart the compose file or the web container, it should connect to the database normally. I think it's an issue in the MySQL image.
2. According to the MailerSend API, trial accounts can send e-mails to up to 10 recipients at once. However, whenever I tried to send an e-mail to more than one recipient at a time, I would get an error 422, indicating that my plan didn't support that. So, I had to literally send one e-mail at a time.
