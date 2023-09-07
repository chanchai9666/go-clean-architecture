FROM golang:1.19.2-alpine3.16 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates git openssh-client

RUN mkdir -p /api
WORKDIR /api
ADD . /api
# RUN go install github.com/swaggo/swag/cmd/swag@latest
# RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
# RUN swag init
# RUN swag init -g cmd/api/main.go

# ENV GOPRIVATE=github.com/ace-energy/go-libmaster

# ARG ACCESS_TOKEN
# RUN git config --global url."https://$ACCESS_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# RUN go get -u github.com/ace-energy/go-libmaster

RUN go mod download
RUN go build -a -installsuffix cgo -o api


FROM alpine:3.12.0
RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates tzdata && \
  rm -rf /var/cache/*

COPY --from=builder /api/api .
COPY --from=builder /api/docs /docs

ADD configs /configs
# ADD /json /json

# EXPOSE 8800
EXPOSE 8318

# CMD ["./api"]
CMD ["sh", "-c", "/api"]
