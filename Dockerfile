FROM golang:1.17.6-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init modname
RUN go mod tidy
RUN go build -o server .
CMD ["/app/server"]


# FROM golang:alpine AS build
# RUN apk --no-cache add gcc g++ make git
# WORKDIR /go/src/app
# COPY . .
# RUN go mod init webserver
# RUN go mod tidy
# RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

# FROM alpine:3.13
# RUN apk --no-cache add ca-certificates
# WORKDIR /usr/bin
# COPY --from=build /go/src/app/bin /go/bin
# EXPOSE 80
# ENTRYPOINT /go/bin/web-app --port 80