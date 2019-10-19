# go-yahoo

go-yahoo is go client for [Yahoo Web API](http://developer.yahoo.co.jp/)

## Support Web API

- [WIP] yahoo auction(ヤフオク!)
- [WIP] yahoo shopping(ヤフショッピング!)

## installation

```
go get github.com/suisun2015/go-yahoo
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

  "github.com/suisun2015/go-yahoo/auction"
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

Search "jewelry" in [ヤフショッピング!](http://shopping.yahoo.co.jp/)

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-yahoo/shopping"
)

func main() {
	// logger := log.Logger{}
  c, _ := shopping.NewClient(nil)
  /* keyword search
	keyword := "vaio"  
  items, _ := c.GetShoppingItemListBySearch(keyword)
  */
  /* item information
	code := "creeam_a01556"
  items, err := c.GetShoppingItemInfo(code)
  */
	category := "1635"
	items, err := c.GetShoppingCategoryRanking(category, 1)
	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		result, _ := json.Marshal(items)
		// fmt.Printf("%+v\n", items)
		fmt.Println(string(result))
	}
}
```
