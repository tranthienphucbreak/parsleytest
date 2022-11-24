# ------------------------------------------------------------------------
#  The first stage container, for setting up and downloading dependencies
# ------------------------------------------------------------------------
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY database.db ./
COPY settings.yaml ./

RUN go mod download

RUN go install github.com/mattn/go-sqlite3

COPY . .

WORKDIR /app/cmd/webapp

RUN go build -o /app/patient -tags purego

EXPOSE 18443

WORKDIR /app

RUN chmod 0440 /app/settings.yaml
RUN chmod 0440 /app/database.db

CMD ["/app/patient"]
