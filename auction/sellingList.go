package auction

import (
	"context"
	"fmt"
	"time"
)

type ChiSellingListroot struct {
	ChiResultSet *ChiResultSet `xml:"urn:yahoo:jp:auc:sellingList ResultSet,omitempty"`
}

type ChiSellingListImage struct {
	Attr_height string `xml:" height,attr"`
	Attr_width  string `xml:" width,attr"`
	Text        string `xml:",chardata"`
}

type ChiSellingListItem struct {
	ChiAuctionID      string                       `xml:"urn:yahoo:jp:auc:sellingList AuctionID,omitempty"`
	ChiAuctionItemUrl string                       `xml:"urn:yahoo:jp:auc:sellingList AuctionItemUrl,omitempty"`
	ChiBidOrBuy       float64                      `xml:"urn:yahoo:jp:auc:sellingList BidOrBuy,omitempty"`
	ChiBids           int                          `xml:"urn:yahoo:jp:auc:sellingList Bids,omitempty"`
	ChiCharityOption  *ChiSellingListCharityOption `xml:"urn:yahoo:jp:auc:sellingList CharityOption,omitempty"`
	ChiCurrentPrice   float64                      `xml:"urn:yahoo:jp:auc:sellingList CurrentPrice,omitempty"`
	ChiEndTime        time.Time                    `xml:"urn:yahoo:jp:auc:sellingList EndTime,omitempty"`
	ChiImage          *ChiSellingListImage         `xml:"urn:yahoo:jp:auc:sellingList Image,omitempty"`
	ChiIsReserved     bool                         `xml:"urn:yahoo:jp:auc:sellingList IsReserved,omitempty"`
	ChiItemUrl        string                       `xml:"urn:yahoo:jp:auc:sellingList ItemUrl,omitempty"`
	ChiOption         *ChiSellingListOption        `xml:"urn:yahoo:jp:auc:sellingList Option,omitempty"`
	ChiTitle          string                       `xml:"urn:yahoo:jp:auc:sellingList Title,omitempty"`
}

type ChiSellingListOption struct {
	ChiBuynowIcon        string `xml:"urn:yahoo:jp:auc:sellingList BuynowIcon,omitempty"`
	ChiEasyPaymentIcon   string `xml:"urn:yahoo:jp:auc:sellingList EasyPaymentIcon,omitempty"`
	ChiIsBackGroundColor bool   `xml:"urn:yahoo:jp:auc:sellingList IsBackGroundColor,omitempty"`
	ChiIsBold            bool   `xml:"urn:yahoo:jp:auc:sellingList IsBold,omitempty"`
	ChiIsCharity         bool   `xml:"urn:yahoo:jp:auc:sellingList IsCharity,omitempty"`
	ChiIsOffer           bool   `xml:"urn:yahoo:jp:auc:sellingList IsOffer,omitempty"`
	ChiNewItemIcon       string `xml:"urn:yahoo:jp:auc:sellingList NewItemIcon,omitempty"`
}

type ChiSellingListCharityOption struct {
	ChiProportion int `xml:"urn:yahoo:jp:auc:categoryLeaf Proportion,omitempty"`
}

type ChiSellingListRating struct {
	ChiIsDeleted   bool `xml:"urn:yahoo:jp:auc:sellingList IsDeleted,omitempty"`
	ChiIsSuspended bool `xml:"urn:yahoo:jp:auc:sellingList IsSuspended,omitempty"`
	ChiPoint       int  `xml:"urn:yahoo:jp:auc:sellingList Point,omitempty"`
}

type ChiSellingListResult struct {
	ChiSellingListItem   []*ChiSellingListItem `xml:"urn:yahoo:jp:auc:sellingList Item,omitempty"`
	ChiSellingListSeller *ChiSellingListSeller `xml:"urn:yahoo:jp:auc:sellingList Seller,omitempty"`
}

type ChiSellingListResultSet struct {
	Attr_firstResultPosition   string                `xml:" firstResultPosition,attr"`
	Attr_xsi_schemaLocation    string                `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string                `xml:" totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string                `xml:" totalResultsReturned,attr"`
	Attr_xmlns                 string                `xml:" xmlns,attr"`
	Attr_xsi                   string                `xml:"xmlns xsi,attr"`
	ChiSellingListResult       *ChiSellingListResult `xml:"urn:yahoo:jp:auc:sellingList Result,omitempty"`
}

type ChiSellingListSeller struct {
	ChiAboutUrl          string                `xml:"urn:yahoo:jp:auc:sellingList AboutUrl,omitempty"`
	ChiId                string                `xml:"urn:yahoo:jp:auc:sellingList Id,omitempty"`
	ChiItemListUrl       string                `xml:"urn:yahoo:jp:auc:sellingList ItemListUrl,omitempty"`
	ChiSellingListRating *ChiSellingListRating `xml:"urn:yahoo:jp:auc:sellingList Rating,omitempty"`
	ChiRatingUrl         string                `xml:"urn:yahoo:jp:auc:sellingList RatingUrl,omitempty"`
}

func (c *Client) getSellingList(ctx context.Context, sellerID string) (*ChiSellingListResultSet, error) {
	spath := fmt.Sprintf("/sellingList" + "?appid=" + c.AuthToken + "&sellerID=" + sellerID)
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		c.Logger.Println("[ERROR] fail newRequest in GetItemListbySellerID")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Println("[ERROR] fail HTTPClient.Do in GetItemListbySellerID")
		return nil, err
	}

	// status check
	if res.StatusCode >= 400 {
		c.Logger.Println("[ERROR] Response fail status code")
		return nil, err
	}

	var resultSet *ChiSellingListResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		c.Logger.Println("[ERROR] fail decodeBody in GetItemListbySellerID")
		return nil, err
	}

	return resultSet, nil
}

func (c *Client) GetItemListBySellerID(ctx context.Context, sellerID string) ([]ChiSellingListItem, error) {
	resultSet, err := c.getSellingList(ctx, sellerID)
	if err != nil {
		c.Logger.Println("[ERROR] getSellingList is fail")
		return nil, err
	}

	list := []ChiSellingListItem{}
	for _, item := range resultSet.ChiSellingListResult.ChiSellingListItem {
		list = append(list, *item)
	}

	return list, nil
}

func (c *Client) GetSellerInfomation(ctx context.Context, sellerID string) (*ChiSellingListSeller, error) {
	resultSet, err := c.getSellingList(ctx, sellerID)
	if err != nil {
		c.Logger.Println("[ERROR] getSellingList is fail")
		return nil, err
	}

	return resultSet.ChiSellingListResult.ChiSellingListSeller, nil
}
