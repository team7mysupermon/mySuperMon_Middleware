[![Go Report Card](https://goreportcard.com/badge/github.com/team7mysupermon/mySuperMon_Middleware)](https://goreportcard.com/report/github.com/team7mysupermon/mySuperMon_Middleware)

# **mySuperMon_Middleware**

This middleware was created for complementing mySuperMon’s monitoring functionalities by integrating Prometheus.

All the relevant information from monitoring database tests will be saved and be easily accessible via Prometheus’ and mySuperMon’s dashboards.

Information regarding installing and operating the middleware is available later in this document.

Utilizing this software requires a mySuperMon account, which can be created at their [website](https://mysupermon.com/).

## **How to Install**

## **How to Run**

## **How to Use**

Once the docker image is running, local port **8999** is used for operating mySuperMon recordings.

Additionally, the image will open port **9090** which is used for accessing information about the recording via **Prometheus.**

Once the middleware is up and running, the following API calls are available:

### **Login**

```yaml
localhost:8999/Login/{MySuperMon Username}/{MySuperMon Password}
```

This **must** **be the first** API call, or it will not be possible to start or stop a recording.

Once you have logged in, you can start and stop recordings without having to log in again.

If the image is ever shut down, you must log in again when you restart the program.

### **Start Recording**

```yaml
localhost:8999/Start/{Usecase name}/{Application Identifier}
```
`Usecase name` can be anything that you choose.

`Application Identifier` can be found in mySuperMon, under *Applications* → *Application*

### **Stop Recording**

```yaml
localhost:8999/Start/{Usecase name}/{Application Identifier}
```

`Usecase name` **must** be the same as the name used to start the recording.

The `application Identifier` **must** be the same as the application identifier used to start the recording.

## **Using Prometheus Dashboard**
### **Starting program**
To start program open directory in terminal to `mySuperMon_Middelware` folder.
Write following command:
[first time]:
```docker-compose up --build```
[after first time]:
```docker-compose up```

Before proceding login. To login, see : [Login]

### **Accessing metrics**
Access prometheus dashboard (in browser) on path: http://mymiddelware.localhost:9090/
Access mySuperMon custom metrics in txt format on path: http://localhost:9091/metrics
