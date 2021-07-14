FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY entity/ entity/
COPY format/ format/
COPY helpers/ helpers/
COPY proto/ proto/
COPY scraper/ scraper/

RUN go build -o /scraper

EXPOSE 8080

CMD ["/scraper"]
