# # Specify the base image for the go app.
FROM golang:1.16-alpine
# # Specify that we now need to excute any commands in this directory
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
# COPY . .
RUN go mod download
COPY *.go ./
# RUN go build -o main main.go
RUN go build -o /docker-gs-ping
RUN echo "build pass"
EXPOSE 3020

CMD ["/docker-gs-ping"]

# docker build --tag docker-gs-ping:1.2.0 .
# docker run -d --name backend-project -p 3020:3020 docker-gs-ping:1.2.0