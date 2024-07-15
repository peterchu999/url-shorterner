# URL Shorterner

Simple API for generating shortening API, inspired by [tinly.url](https://tinyurl.com/) and [bit.ly](https://bitly.com/). the main goals for this project were learning system design and golang implementation.

## Run The App

```sh
go run main.go
```

### Run with Hot reload

I encourage you to install [gin](https://github.com/codegangsta/gin?tab=readme-ov-file) it help with hot / live realoading.

```sh
gin --appPort 8080 --port 3000
```

- change the `appPor`t with the .env `PORT`, default gin server would serve at `8080` port
- `port` flag was used to setup the proxy server, default `3000`

## API

- ### Create short url

  - path: `/short`
  - method: `POST`
  - content-type: `application/json`
  - request body:
    - attribute:
      - _long-url_ : `url`
    - example:
      ```
      {
          "long-url": "https://google.com"
      }
      ```
  - response:
    - attribute:
      - _short-url_: `url`
    - example:
      ```
      {
          "short_url": "localhost:3000/t/h"
      }
      ```

- ### Redirect short url
  - path: `/:shortUrl`
  - method: `GET`
  - path parameter:
    - shortUrl : the short url generated from the `create short url` api

## Tech Stack

- Gin -> http library
- Monggo DB -> database

## Reference

- hot / live reloading : https://reyhansofian.github.io/en/post/live-reload-golang-development-with-gin/
- short url generation : https://www.geeksforgeeks.org/how-to-design-a-tiny-url-or-url-shortener/
- monggo db setup :
  - https://dasarpemrogramangolang.novalagung.com/A-mongodb.html
  - https://www.mongodb.com/developer/products/mongodb/build-go-web-application-gin-mongodb-help-ai/
