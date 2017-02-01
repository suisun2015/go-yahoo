package auction

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"
)

type Chiroot struct {
	ChiResultSet *ChiResultSet `xml:"urn:yahoo:jp:auc:search ResultSet,omitempty"`
}

type ChiSearchCharityOption struct {
	ChiProportion int `xml:"urn:yahoo:jp:auc:search Proportion,omitempty"`
}

type ChiImage struct {
	Attr_height string `xml:" height,attr"`
	Attr_width  string `xml:" width,attr"`
	Text        string `xml:",chardata"`
}

type ChiSearchResult struct {
	ChiItem      []*ChiSearchItem `xml:"urn:yahoo:jp:auc:search Item,omitempty"`
	ChiUnitsWord string           `xml:"urn:yahoo:jp:auc:search UnitsWord,omitempty"`
}

type ChiSearchResultSet struct {
	Attr_firstResultPosition   string           `xml:" firstResultPosition,attr"`
	Attr_xsi_schemaLocation    string           `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string           `xml:" totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string           `xml:" totalResultsReturned,attr"`
	Attr_xmlns                 string           `xml:" xmlns,attr"`
	Attr_xsi                   string           `xml:"xmlns xsi,attr"`
	ChiSearchResult            *ChiSearchResult `xml:"urn:yahoo:jp:auc:search Result,omitempty"`
}

type ChiSearchItem struct {
	ChiAuctionID        string                  `xml:"urn:yahoo:jp:auc:search AuctionID,omitempty"`
	ChiAuctionItemUrl   string                  `xml:"urn:yahoo:jp:auc:search AuctionItemUrl,omitempty"`
	ChiBidOrBuy         float64                 `xml:"urn:yahoo:jp:auc:search BidOrBuy,omitempty"`
	ChiBids             int                     `xml:"urn:yahoo:jp:auc:search Bids,omitempty"`
	ChiCategoryId       int                     `xml:"urn:yahoo:jp:auc:search CategoryId,omitempty"`
	ChiCharityOption    *ChiSearchCharityOption `xml:"urn:yahoo:jp:auc:search CharityOption,omitempty"`
	ChiCurrentPrice     float64                 `xml:"urn:yahoo:jp:auc:search CurrentPrice,omitempty"`
	ChiEndTime          time.Time               `xml:"urn:yahoo:jp:auc:search EndTime,omitempty"`
	ChiImage            *ChiSearchImage         `xml:"urn:yahoo:jp:auc:search Image,omitempty"`
	ChiIsAdult          bool                    `xml:"urn:yahoo:jp:auc:search IsAdult,omitempty"`
	ChiIsReserved       bool                    `xml:"urn:yahoo:jp:auc:search IsReserved,omitempty"`
	ChiItemUrl          string                  `xml:"urn:yahoo:jp:auc:search ItemUrl,omitempty"`
	ChiOption           *ChiSearchOption        `xml:"urn:yahoo:jp:auc:search Option,omitempty"`
	ChiOriginalImageNum int                     `xml:"urn:yahoo:jp:auc:search OriginalImageNum,omitempty"`
	ChiSeller           *ChiSearchSeller        `xml:"urn:yahoo:jp:auc:search Seller,omitempty"`
	ChiTitle            string                  `xml:"urn:yahoo:jp:auc:search Title,omitempty"`
}

type ChiSearchImage struct {
	Attr_height string `xml:" height,attr"`
	Attr_width  string `xml:" width,attr"`
	Text        string `xml:",chardata"`
}

type ChiSearchOption struct {
	ChiBuynowIcon        string `xml:"urn:yahoo:jp:auc:search BuynowIcon,omitempty"`
	ChiEasyPaymentIcon   string `xml:"urn:yahoo:jp:auc:search EasyPaymentIcon,omitempty"`
	ChiIsBackGroundColor bool   `xml:"urn:yahoo:jp:auc:search IsBackGroundColor,omitempty"`
	ChiIsBold            bool   `xml:"urn:yahoo:jp:auc:search IsBold,omitempty"`
	ChiIsCharity         bool   `xml:"urn:yahoo:jp:auc:search IsCharity,omitempty"`
	ChiIsOffer           bool   `xml:"urn:yahoo:jp:auc:search IsOffer,omitempty"`
	ChiNewIcon           string `xml:"urn:yahoo:jp:auc:search NewIcon,omitempty"`
	ChiNewItemIcon       string `xml:"urn:yahoo:jp:auc:search NewItemIcon,omitempty"`
	ChiStoreIcon         string `xml:"urn:yahoo:jp:auc:search StoreIcon,omitempty"`
}

type ChiSearchSeller struct {
	ChiItemSellerId  string `xml:"urn:yahoo:jp:auc:search Id,omitempty"`
	ChiItemListUrl   string `xml:"urn:yahoo:jp:auc:search ItemListUrl,omitempty"`
	ChiItemRatingUrl string `xml:"urn:yahoo:jp:auc:search RatingUrl,omitempty"`
}

func (c *Client) GetItemsListBySearch(ctx context.Context, keyword string) ([]ChiSearchItem, error) {
	queryWord, err := url.Parse(keyword)
	if err != nil {
		return nil, err
	}

	spath := fmt.Sprintf("/search" + "?appid=" + c.AuthToken + "&query=" + queryWord.String())
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		c.Logger.Println("[ERROR] fail newRequest in GetItemsList")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Println("[ERROR] fail HTTPClient.Do in GetItemsList")
		return nil, err
	}

	// status check
	if res.StatusCode >= 400 {
		c.Logger.Println("[ERROR] Response fail status code")
		return nil, err
	}

	var resultSet ChiSearchResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		c.Logger.Println("[ERROR] fail decodeBody in GetItemsList")
		log.Println("[ERROR] fail decodeBody in GetItemsList")
		return nil, err
	}

	ilist := []ChiSearchItem{}
	for _, item := range resultSet.ChiSearchResult.ChiItem {
		ilist = append(ilist, *item)
	}

	return ilist, nil
}
