# MyGameList site scraper

Site scraper written in Go

## MyGameList family

- [Backend](https://github.com/br3w0r/gamelist-backend)
- [Frontend](https://github.com/br3w0r/gamelist-frontend)
- [Site Scraper](https://github.com/br3w0r/gamelist-scraper)
- [Proto file](https://github.com/br3w0r/gamelist-proto)

## Running in development mode

Run `go get` and `go run main.go` and the scraper will be listening to localhost:8888

You can set amount of pages to scrape with `PAGES_TO_SCRAPE` environment variable.

## Docker building an running

### Build

```bash
# docker build -t scraper .
```

### Common run command

```bash
# docker run \
    --network gamelist --network-alias scraper \
    -e ADDRESS=0.0.0.0 \
    gamelist-scraper
```

This command will set the scraper's address to `scraper:8888` inside `gamelist` docker network and start to listen to it.

Then all that's left to do is to build and run the backend's docker image as follows in its readme.
