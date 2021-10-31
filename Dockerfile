FROM golang:1.17.1 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todoapp

FROM alpine:3.11.3

WORKDIR /

COPY --from=build /todoapp /todoapp
COPY --from=build /app/repository/migrations ./repository/migrations

ENTRYPOINT [ "/todoapp" ]