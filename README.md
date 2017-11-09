# DogWatch

[![Build Status](https://travis-ci.org/gugahoi/dogwatch.svg?branch=master)](https://travis-ci.org/gugahoi/dogwatch)
[![Coverage Status](https://coveralls.io/repos/github/gugahoi/dogwatch/badge.svg?branch=master)](https://coveralls.io/github/gugahoi/dogwatch?branch=master)

DogWatch is a DogNZB cli tool and library to interact with DogNZB's Watchlists. It supports adding, removing and listing items in both TV and Movies watchlists.

## Installation

Simply download the latest release from the [releases page](https://github.com/gugahoi/dogwatch/releases)

```bash
# download the latest tar
curl -LOs https://github.com/gugahoi/dogwatch/releases/download/0.1.1/dogwatch_0.1.1_darwin_amd64.tar.gz

# uncompress and put the binary into /usr/local/bin
tar -C /usr/local/bin -xvf dogwatch*.tar.gz dogwatch

# profit!
dogwatch version
```

## Usage

### CLI

```bash
# List all movies in watchlist
dogwatch list movies --api SOME-API-KEY

# List all series in watchlist
dogwatch list tv --api SOME-API-KEY

# Alternatively set `DOGNZB_API` env variable instead if using `--api` flag
export DOGNZB_API="SOME-API-KEY"

# Add movies to the watchlist
dogwatch add movies tt123456 tt567890

# Add series to the watchlist
dogwatch add tv 123456 098767

# remove movies from the watchlist
dogwatch remove movies tt123456 ...

# remove series from the watchlist
dogwatch remove tv 123456 ...
```

### Library

```go
package main

import github.com/gugahoi/dogwatch/pkg/dognzb

func main() {
    d := dognzb.New("some-api-key")

    // list movies
    items, err := d.List(dognzb.Movies)
    if err != nil {...}
    for _, item := range items {
        fmt.Println(item.Title, item.GetID())
    }

    // add movie
    q, err := d.Add(dognzb.Movies, "tt123455")
    if err != nil {...}
}
```

## Development

This project uses [`dep`](https://github.com/golang/dep) as the dependency manager.

```bash
# download deps
dep ensure

# build
make build

# yeah!
./build/dogwatch version
```
