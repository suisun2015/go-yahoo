# go-yahoo

go-yahoo is go client for [Yahoo Web API](http://developer.yahoo.co.jp/)

## Support Web API

- [WIP] yahoo auction(ヤフオク!)

## installation

```
go get github.com/whywaita/go-yahoo
```

go-yahoo need Application ID(auth token).
Please read [This page](http://developer.yahoo.co.jp/start/).

## Usage

Search "jewelry" in [ヤフオク!](http://auctions.yahoo.co.jp/)

```go
package main

import (
  "context"
  "fmt"
  "log"

  "github.com/whywaita/go-yahoo/auction"
)

const (
  token = "YOUR-TOKEN"
)

func main() {
  ctx := context.Background()

  client, err := auction.NewClient(token, nil)
  if err != nil {
    log.Println(err)
    log.Fatal("fail new client")
  }

  _, err = client.GetCategoryList(ctx)
  if err != nil {
    log.Println(err)
    log.Fatal("fail GetCategoryIdList")
  }

  list, err := client.GetItemsListBySearch(ctx, "jewelry")
  if err != nil {
    log.Println(err)
    log.Fatal("fail GetItemsListBySearch")
  }
  fmt.Println(list)
}
```
