package auction

import (
	"context"
	"fmt"
	"time"
)

type ChiItemroot struct {
	ChiItemResultSet *ChiItemResultSet `xml:"urn:yahoo:jp:auc:categoryLeaf ResultSet,omitempty"`
}

type ChiItemImage struct {
	Attr_height string `xml:" height,attr"`
	Attr_width  string `xml:" width,attr"`
	Text        string `xml:",chardata"`
}

type ChiItem struct {
	ChiAuctionID      string                `xml:"urn:yahoo:jp:auc:categoryLeaf AuctionID,omitempty"`
	ChiAuctionItemUrl string                `xml:"urn:yahoo:jp:auc:categoryLeaf AuctionItemUrl,omitempty"`
	ChiBidOrBuy       float64               `xml:"urn:yahoo:jp:auc:categoryLeaf BidOrBuy,omitempty"`
	ChiBids           int                   `xml:"urn:yahoo:jp:auc:categoryLeaf Bids,omitempty"`
	ChiCharityOption  *ChiItemCharityOption `xml:"urn:yahoo:jp:auc:categoryLeaf CharityOption,omitempty"`
	ChiCurrentPrice   float64               `xml:"urn:yahoo:jp:auc:categoryLeaf CurrentPrice,omitempty"`
	ChiEndTime        time.Time             `xml:"urn:yahoo:jp:auc:categoryLeaf EndTime,omitempty"`
	ChiItemImage      *ChiItemImage         `xml:"urn:yahoo:jp:auc:categoryLeaf Image,omitempty"`
	ChiIsAdult        bool                  `xml:"urn:yahoo:jp:auc:categoryLeaf IsAdult,omitempty"`
	ChiIsReserved     bool                  `xml:"urn:yahoo:jp:auc:categoryLeaf IsReserved,omitempty"`
	ChiItemUrl        string                `xml:"urn:yahoo:jp:auc:categoryLeaf ItemUrl,omitempty"`
	ChiItemOption     *ChiItemOption        `xml:"urn:yahoo:jp:auc:categoryLeaf Option,omitempty"`
	ChiItemSeller     *ChiItemSeller        `xml:"urn:yahoo:jp:auc:categoryLeaf Seller,omitempty"`
	ChiTitle          string                `xml:"urn:yahoo:jp:auc:categoryLeaf Title,omitempty"`
}

type ChiItemCharityOption struct {
	ChiProportion int `xml:"urn:yahoo:jp:auc:categoryLeaf Proportion,omitempty"`
}

type ChiItemOption struct {
	ChiBuynowIcon        string `xml:"urn:yahoo:jp:auc:categoryLeaf BuynowIcon,omitempty"`
	ChiEasyPaymentIcon   string `xml:"urn:yahoo:jp:auc:categoryLeaf EasyPaymentIcon,omitempty"`
	ChiIsBackGroundColor bool   `xml:"urn:yahoo:jp:auc:categoryLeaf IsBackGroundColor,omitempty"`
	ChiIsBold            bool   `xml:"urn:yahoo:jp:auc:categoryLeaf IsBold,omitempty"`
	ChiIsCharity         bool   `xml:"urn:yahoo:jp:auc:categoryLeaf IsCharity,omitempty"`
	ChiIsOffer           bool   `xml:"urn:yahoo:jp:auc:categoryLeaf IsOffer,omitempty"`
	ChiNewIcon           string `xml:"urn:yahoo:jp:auc:categoryLeaf NewIcon,omitempty"`
	ChiNewItemIcon       string `xml:"urn:yahoo:jp:auc:categoryLeaf NewItemIcon,omitempty"`
	ChiStarClubIcon      string `xml:"urn:yahoo:jp:auc:categoryLeaf StarClubIcon,omitempty"`
}

type ChiItemResult struct {
	ChiCategoryPath string     `xml:"urn:yahoo:jp:auc:categoryLeaf CategoryPath,omitempty"`
	ChiItem         []*ChiItem `xml:"urn:yahoo:jp:auc:categoryLeaf Item,omitempty"`
}

type ChiItemResultSet struct {
	Attr_firstResultPosition   string         `xml:"firstResultPosition,attr" `
	Attr_xsi_schemaLocation    string         `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string         `xml:"totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string         `xml:"totalResultsReturned,attr"`
	Attr_xmlns                 string         `xml:"xmlns,attr"`
	Attr_xsi                   string         `xml:"xmlns xsi,attr"`
	ChiItemResult              *ChiItemResult `xml:"urn:yahoo:jp:auc:categoryLeaf Result,omitempty"`
}

type ChiItemSeller struct {
	ChiItemSellerId  string `xml:"urn:yahoo:jp:auc:categoryLeaf Id,omitempty"`
	ChiItemListUrl   string `xml:"urn:yahoo:jp:auc:categoryLeaf ItemListUrl,omitempty"`
	ChiItemRatingUrl string `xml:"urn:yahoo:jp:auc:categoryLeaf RatingUrl,omitempty"`
}

func (c *Client) GetItemsList(ctx context.Context, categoryId int) ([]ChiItem, error) {
	spath := fmt.Sprintf("/categoryLeaf" + "?appid=" + c.AuthToken)
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

	if res.StatusCode >= 400 {
		c.Logger.Println("[ERROR] Response fail status code")
		return nil, err
	}

	var resultSet ChiItemResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		c.Logger.Println("[ERROR] fail decodeBody in GetItemsList")
		return nil, err
	}

	ilist := []ChiItem{}
	for _, item := range resultSet.ChiItemResult.ChiItem {
		ilist = append(ilist, *item)
	}

	return ilist, nil
}
