
# Introduction
The Vermouth Backend is a crucial component designed to serve the API for the Vermouth Frontend. Its primary function is to fetch and provide data derived from the anomaly detection results of various elements of the THEOS satellite. This backend system ensures seamless data retrieval and communication between the anomaly detection system and the frontend interface.


# Initialization and Setup

the initialization and setup of a Go project that uses the Gin web framework and the go-pg library for PostgreSQL database interactions. Let's break down each step in detail.

### 1. Initialize the Go Module
`$ go mod init vermouth-backend`

### 2. Install Libraries
- `$ github.com/gin-gonic/gin` : Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance.
- `$ github.com/go-pg/pg/v10'` : go-pg is an ORM for PostgreSQL in Go.


#### Additional Libraries
- `$ github.com/gin-contrib/cors` : This library provides middleware for handling Cross-Origin Resource Sharing (CORS), which is crucial for web applications that interact with resources from different origins.
- `github.com/gin-gonic/gin@v1.8.1` : This command installs a specific version (v1.8.1) of the Gin framework, ensuring that your project uses a stable and expected version of Gin.

### 3. Run the Go Project
`$ go run main.go`

# API services
### 1. GET : anomalyweekly
![Reference Image recommendation user](/img/anomalyweekly.png)

### 2. GET : reportdailyfilter
![Reference Image recommendation user](/img/postreportdaily.png)

### 3. GET : reportalldaily_proviousmonth
![Reference Image recommendation user](/img/reportalldaily_proviousmonth.png)

### 4. POST : postreportdaily
![Reference Image recommendation user](/img/postreportdaily.png)

### 5. GET : THEOSchartfilter
![Reference Image recommendation user](/img/THEOSchartfilter.png)

### 6. GET: THEOS_chartanomaly/:satname/:tmname/:freq"
![Reference Image recommendation user](/img/THEOS_chartanomaly.png)

### 7. GET : GET_THEOSDownload_staticscsv/:satname/:tmname/:freq/:analysis_table/:anomaly_table/:start_utc/:end_utc
![Reference Image recommendation user](/img/GET_THEOSDownload_staticscsv.png)