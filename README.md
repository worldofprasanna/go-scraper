# go-scraper

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Maintainability](https://api.codeclimate.com/v1/badges/7a6c3b78f3c0130d1bd5/maintainability)](https://codeclimate.com/github/worldofprasanna/go-scraper/maintainability)

> Scrape NSE website to get the board meeting information

## Table of Contents

- [go-scraper](#go-scraper)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
  - [Install Using Docker](#install-using-docker)
  - [Maintainers](#maintainers)
  - [Contributing](#contributing)
  - [License](#license)

## Install

```
# To run the unit test
./bin/test

# This needs go version to be >= 1.11 because of go modules dependency
./bin/build

# To start the server
./app

```

## Install Using Docker

```
# Build the docker image
docker build -t nse_scrapper .

# Run the docker container
docker run -p 8080:8080 nse_scrapper
```

## Maintainers

[@worldofprasanna](https://github.com/worldofprasanna)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2019 Prasanna V
