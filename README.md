## About the project and objectives
Since I got a task at my job to implement the system's background worker and a service for sending e-mails, I've decided to try this process in a personal project, and this is what I came up with: a very basic system where users can create newsletters and send messages to the users registered to them. 

The objective of the project is to be a simple proof of concept on how to implement a basic e-mail service in a background worker, and for that, I decided to go with [MailerSend for Go](https://github.com/mailersend/mailersend-go) and [Asynq](https://github.com/hibiken/asynq). Of course, since this is a simple project, it doesn't have many concerns for security, it is merely a showcase of what can be done with Go.

A very interesting detail that I noticed was that sending e-mails via MailerSend directly in an endpoint would usually take around 600 ms for the API to respond. After putting that in the background worker, the API would return the response in about 10 ms. This emphasizes the importance of having a good background worker system running in your app, so that longer tasks, especially ones that depend on external responses, don't get the main thread stuck. Imagine if a newsletter had a thousand registered users. It would take ten minutes (600,000 ms) for the API's main thread to get unlocked and process the next request(s). Yeah, that gets pretty crazy pretty fast.

## Running the system
Make sure you have Docker installed in your machine. In order to run this system, you'll need to create the `.env` file. Create a copy of `.env example` and change all fields. The only fields you won't have to create some random string will be the MailerSend domain and MailerSend API key. 

After that, just run `docker-compose up` and make sure you have Docker running in your system. To shut down the system, run `docker-compose down`, you can even add the `-v` argument to this command to remove the volumes.

As for MailerSend, you'll need to create a Trial account and generate an API key for the trial domain they provide. After that, it is strongly recommended to go to the Domains page and uncheck the "Track Clicks" option, since it breaks the links in the HTML for the e-mails.

## Known issues
1. For some reason, the web container would start up before the database one was fully started up, even after adding `depends_on` in the Compose file. If you see `Connection refused` in the web container's logs, restart the compose file or the web container, it should connect to the database normally. I think it's an issue in the MySQL image.
2. According to the MailerSend API, trial accounts can send e-mails to up to 10 recipients at once. However, whenever I tried to send an e-mail to more than one recipient at a time, I would get an error 422, indicating that my plan didn't support that. So, I had to literally send one e-mail at a time.
