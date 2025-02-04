# URL Shorterner

Simple API for generating shortening API, inspired by [tinly.url](https://tinyurl.com/) and [bit.ly](https://bitly.com/). the main goals for this project were learning system design and golang implementation.

visit [project wiki](https://github.com/peterchu999/url-shorterner/wiki) to learn more

## Run The App

```sh
go run main.go
```

### Run with Hot reload

I encourage you to install [gin](https://github.com/codegangsta/gin?tab=readme-ov-file) it help with hot / live realoading.
**NOTE** remember to change the mongodb uri in `.env`

```sh
EXPORT gin --appPort 8080 --port 3000
```

### Run with Docker Compose

The docker and docker compose file are already setup so just run the script below to make sure the apps is running

**NOTE** remember to change the mongodb uri in `.env`

```sh
docker compose up
```

## API

<details>
<summary> Create short url</summary>

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
- response: - attribute: - _short-url_: `url` - example:
`{
          "short_url": "localhost:3000/t/h"
      }`
</details>

<details>
<summary> Redirect short url</summary>
  - path: `/:shortUrl`
  - method: `GET`
  - path parameter:
    - shortUrl : the short url generated from the `create short url` api
</details>

## Tech Stack

- Gin -> http library
- Monggo DB -> database

## Reference

- hot / live reloading : https://reyhansofian.github.io/en/post/live-reload-golang-development-with-gin/
- short url generation : https://www.geeksforgeeks.org/how-to-design-a-tiny-url-or-url-shortener/
- monggo db setup :
  - https://dasarpemrogramangolang.novalagung.com/A-mongodb.html
  - https://www.mongodb.com/developer/products/mongodb/build-go-web-application-gin-mongodb-help-ai/
