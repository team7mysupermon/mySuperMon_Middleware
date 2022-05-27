[![Go Report Card](https://goreportcard.com/badge/github.com/team7mysupermon/mySuperMon_Middleware)](https://goreportcard.com/report/github.com/team7mysupermon/mySuperMon_Middleware)

# mySuperMon_Middleware

This middleware was created to have an easy to set up link between MySuperMon and Prometheus.

This middleware helps you, the user, moniter your database. Through MySuperMon and Prometheus all the relevant information about different tasks performed on your database will be monitored and saved, and this information is easily accessable through the Prometheus and MySuperMon dashboard.

Further down this document, you can find a guide on how to install, run and use this middleware.

You must have a MySuperMon account to use this middleware. You can create an account on their [website](https://mysupermon.com/).

## How to Install

To install the middleware locally, you must have docker and docker compose installed and do the following:

Download the docker compose file from the release.

Open the directory in a terminal where the docker compose file is.

Write the following command:
```docker-compose pull```

This will download the docker images locally.

## How to Run

To start program open a terminal and navigate to the folder containing the docker compose file.
Write following command:
```docker-compose up```

Before proceding login. To login, see : [Login](#login)

## How to Use

When the docker image is running, it is running on the local port **8999**, which is the port you can use to start and stop a MySuperMon recording.

Also, the image will open the port **9090** that can be used to access information about the recording through **Prometheus.**

Once the middleware is up and running, you can do the following API calls, API calls can be made through the address-bar in the browser:

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

## **Prometheus**

### **Accessing metrics**

*Please remember to login before hand. See subsection [Login](#login)*
Access prometheus dashboard (in browser) on path: http://mymiddelware.localhost:9090/
Access mySuperMon custom metrics in txt format on path: http://localhost:9091/metrics

## **Grafana**

Access Grafana on path: http://localhost:3000/

*OBS! Beware that first time users of Grafana needs to login with credentials: {uname}: admin, {password}: admin*

### **Steps to connect Prometheus to Grafana**

- Press *Add datasource*
- Select *Prometheus* as the type
- Fill out the form, with the following info:

    **HTTP**
    - Name: whatever you wanna call it
    - URL:
    - Access:
    some fields that doesn't matter...
    
    **Auth**
    - Basic auth: on
    the rest should be left off...
    
    **Basic Auth Details**
    - User: *Username for mySuperMon*
    - Password: *Password for mySuperMon*
    
    **Alerting**
    - Scrape interval: 5s
    
    All the remaining fields should be left untouched

- Press: *Save & test*, and pray to god it works. If green yes, if red no.
- Access metrics in explore
- See Grafana tutorials for more

## Swagger

Once the middleware is up and running, swagger documentation will be up on the following page: [http://localhost:8999/swagger/index.html#/](http://localhost:8999/swagger/index.html#/)

### How to use

When the swagger page is opened the API endpoints can be tested by opening a tab and pressing the “try it out” button. Fill out the required information and press execute.
