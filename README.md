## About the project and objectives
lorem ipsum from hell

## Objectives
ipsum lorem

## Running the system
In order to run this system, just run `docker-compose up` and make sure you have Docker running in your system. To shut down the system, run `docker-compose down`, you can even add the `-v` argument to this command to remove the volumes.

## Issues
For some reason, the web container would start up before the database was fully started up, even after adding `depends_on` in the Compose file. If you see "Connection refused" in the web container's logs, restart the compose file or the web container, it should connect to the database normally. I think it's an issue in the MySQL image.
