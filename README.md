# DogWatch

[![Build Status](https://travis-ci.org/gugahoi/dogwatch.svg?branch=master)](https://travis-ci.org/gugahoi/dogwatch)
[![Coverage Status](https://coveralls.io/repos/github/gugahoi/dogwatch/badge.svg?branch=master)](https://coveralls.io/github/gugahoi/dogwatch?branch=master)

DogWatch is a DogNZB cli tool and library to interact with DogNZB's Watchlists. It supports adding, removing and listing items in both TV and Movies watchlists.

## Usage

### CLI

```bash
# List all movies in watchlist
dogwatch list movies --apikey SOME-API-KEY

# List all series in watchlist
dogwatch list tv --apikey SOME-API-KEY

# Alternatively set `DOGNZB_API` env variable instead if using `--apikey` flag
export DOGNZB_API="SOME-API-KEY"

# Add a movie to the watchlist
dogwatch add movie tt123456

# Add a seties to the watchlist
dogwatch add tv 123456

# remove a movie from the watchlist
dogwatch remove movie tt123456

# remove a series from the watchlist
dogwatch remove tv 123456
```

### Library

```go
package main

import github.com/gugahoi/dogwatch/pkg/dognzb

func main() {
    d := dognzb.New("some-api-key")

    // list movies
    d.List(dognzb.Movies)

    // add movie
    d.Add(dognzb.Movies, "tt123455")
}
```

## Development

This project uses [`dep`](https://github.com/golang/dep) as the dependency manager.

```bash
# download deps
dep ensure

# build
go build *.go -o dogwatch

# profit!
./dogwatch --help
```
