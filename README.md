[![Go Report Card](https://goreportcard.com/badge/github.com/team7mysupermon/mySuperMon_Middleware)](https://goreportcard.com/report/github.com/team7mysupermon/mySuperMon_Middleware)

# mySuperMon_Middleware
This middleware was created to have an easy to set up link between MySuperMon and Prometheus.

This middleware helps you, the user, moniter your database. Through MySuperMon and Prometheus all the relevant information about different tasks performed on your database will be monitored and saved, and this information is easily accessable through the Prometheus and MySuperMon dashboard.

Further down this document, you can find a guide on how to install, run and use this middleware.

You must have a MySuperMon account to use this middleware. You can create an account on their [website](https://mysupermon.com/).

## How to Install

## How to Run

## How to Use

When the docker image is running, it is running on the local port **8999**, which is the port you can use to start and stop a MySuperMon recording.

Also, the image will open the port **9090** that can be used to access information about the recording through **Prometheus.**

Once the middleware is up and running, you can do the following API calls:

### Login

```
localhost:8999/Login/{MySuperMon Username}/{MySuperMon Password}
```

This call must be made as the first API call, or it will not be possible to start or stop a recording.

Once you have logged in, you can start and stop recordings without having to log in again.

If the image is ever shut down, you must log in again when you restart the program.

### Start Recording

```
localhost:8999/Start/{Usecase name}/{Application Identifier}
```

**Usecase name** can be anything that you choose.

**Application Identifier** can be found in MySuperMon, under *Applications* and *Application Management.*

### Stop Recording

```
localhost:8999/Start/{Usecase name}/{Application Identifier}
```

**Usecase name** has to be the same as the name used to start the recording.

**Application Identifier** has to be the same as the application identifier used to start the recording.
