package shopping

import (
	"fmt"
)

// type ChiItemroot struct {
// 	ChiItemResultSet *ChiItemResultSet `xml:"urn:yahoo:jp:auc:categoryLeaf ResultSet,omitempty"`
// }

type InfoImage struct {
	Id     string `xml:"Id,omitempty"`
	Small  string `xml:"Small,omitempty"`
	Medium string `xml:"Medium,omitempty"`
}
type InfoReview struct {
	Rate  string `xml:"Rate,omitempty"`
	Count string `xml:"Count,omitempty"`
	Url   string `xml:"Url,omitempty"`
}
type InfoPrice struct {
	Currency string `xml:"currency,attr"`
	Value    int64  `xml:",chardata"`
}
type InfoPriceLabel struct {
	TaxIncluded  bool   `xml:"taxIncluded,attr"`
	FixedPrice   int64  `xml:FixedPrice,omitempty`
	DefaultPrice int64  `xml:DefaultPrice,omitempty`
	SalePrice    int64  `xml:SalePrice,omitempty`
	MemberPrice  int64  `xml:MemberPrice,omitempty`
	PeriodStart  string `xml:PeriodStart,omitempty`
	PeriodEnd    string `xml:PeriodEnd,omitempty`
}
type InfoPoint struct {
	Amount          int64 `xml:"Amount,omitempty"`
	Times           int64 `xml:"Times,omitempty"`
	PremiumAmount   int64 `xml:"PremiumAmount,omitempty"`
	PremiumTimes    int64 `xml:"PremiumTimes,omitempty"`
	PremiumCpAmount int64 `xml:"PremiumCpAmount,omitempty"`
	PremiumCpTimes  int64 `xml:"PremiumCpTimes,omitempty"`
	AppCpAmount     int64 `xml:"AppCpAmount,omitempty"`
	AppCpTimes      int64 `xml:"AppCpTimes,omitempty"`
	PreAppCpAmount  int64 `xml:"PreAppCpAmount,omitempty"`
	PreAppCpTimes   int64 `xml:"PreAppCpTimes,omitempty"`
}
type InfoShipping struct {
	Code int64  `xml:"Code,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type InfoNode struct {
	Id   string `xml:"Id,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type InfoProductCategory struct {
	ID int64 `xml:"ID,omitempty"`
}
type InfoMethod struct {
	Code int64  `xml:"Code,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type InfoPayment struct {
	Method []InfoMethod `xml:"Method,omitempty"`
}
type InfoRatings struct {
	Rate       float64 `xml:"Rate,omitempty"`
	Count      int64   `xml:"Count,omitempty"`
	Total      int64   `xml:"Total,omitempty"`
	DetailRate float64 `xml:"DetailRate,omitempty"`
}
type InfoStoreImage struct {
	Id     string `xml:"Id,omitempty"`
	Medium string `xml:"Medium,omitempty"`
}
type InfoStorePoint struct {
	Grant  bool `xml:"Grant,omitempty"`
	Accept bool `xml:"Accept,omitempty"`
}
type InfoPrefectures struct {
	Prefecture []InfoNode `xml:"Prefecture,omitempty"`
}
type InfoArea struct {
	Code        string            `xml:"Code,omitempty"`
	Name        string            `xml:"Name,omitempty"`
	Prefectures []InfoPrefectures `xml:"Prefectures,omitempty"`
}
type InfoAreas struct {
	Area *InfoArea `xml:"Area,omitempty"`
}
type InfoDelivery struct {
	Areas      []InfoAreas `xml:"Areas,omitempty"`
	Deadline   string      `xml:"Deadline,omitempty"`
	Conditions string      `xml:"Conditions,omitempty"`
}
type InfoStore struct {
	Id               string          `xml:"Id,omitempty"`
	Name             string          `xml:"Name,omitempty"`
	Url              string          `xml:"Url,omitempty"`
	SellerType       string          `xml:"SellerType,omitempty"`
	InventoryMessage string          `xml:"InventoryMessage,omitempty"`
	ToolType         string          `xml:"ToolType,omitempty"`
	Ratings          *InfoRatings    `xml:"Ratings,omitempty"`
	Image            *InfoStoreImage `xml:"Image,omitempty"`
	Point            *InfoStorePoint `xml:"Point,omitempty"`
	SameDayDelivery  *InfoDelivery   `xml:"SameDayDelivery,omitempty"`
	ExpressDelivery  *InfoDelivery   `xml:"ExpressDelivery,omitempty"`
}
type InfoCode struct {
	Code string `xml:"Code"`
}
type InfoItemCode struct {
	Codes []InfoCode `xml:"Codes"`
}
type InfoInventory struct {
	SubCode        string `xml:"SubCode"`
	Order          string `xml:"Order"`
	Availability   string `xml:"Availability"`
	Quantity       int64  `xml:"Quantity"`
	AllowOverdraft int64  `xml:"AllowOverdraft"`
}
type InfoInventories struct {
	Inventory []InfoInventory `xml:"Inventory"`
}

type InfoHit struct {
	Code                  string               `xml:"Code,omitempty"`
	Name                  string               `xml:"Name,omitempty"`
	Url                   string               `xml:"Url,omitempty"`
	Condition             string               `xml:"Condition,omitempty"`
	Headline              string               `xml:"Headline,omitempty"`
	Caption               string               `xml:"Caption,omitempty"`
	Abstract              string               `xml:"Abstract,omitempty"`
	Additional1           string               `xml:"Additional1,omitempty"`
	Additional2           string               `xml:"Additional2,omitempty"`
	Additional3           string               `xml:"Additional3,omitempty"`
	SpAdditional          string               `xml:"SpAdditional,omitempty"`
	ProductCategory       *InfoProductCategory `xml:"ProductCategory,omitempty"`
	IsBargain             bool                 `xml:"IsBargain,omitempty"`
	OriginalPriceEvidence string               `xml:"OriginalPriceEvidence,omitempty"`
	Description           string               `xml:"Description,omitempty"`
	ReleaseDate           string               `xml:"ReleaseDate,omitempty"`
	JanCode               string               `xml:"JanCode,omitempty"`
	Image                 *InfoImage           `xml:"Image,omitempty"`
	RelatedImages         []InfoImage          `xml:"RelatedImages,omitempty"`
	Review                *InfoReview          `xml:"Review,omitempty"`
	Price                 *InfoPrice           `xml:"Price,omitempty"`
	PriceLabel            *InfoPriceLabel      `xml:"PriceLabel,omitempty"`
	ShipWeight            float64              `xml:"ShipWeight,omitempty"`
	SaleLimit             int64                `xml:"SaleLimit,omitempty"`
	Inventories           *InfoInventories     `xml:"Inventories,omitempty"`
	Point                 *InfoPoint           `xml:"Point,omitempty"`
	Payment               *InfoPayment         `xml:"Payment,omitempty"`
	Shipping              *InfoShipping        `xml:"Shipping,omitempty"`
	Store                 *InfoStore           `xml:"Store,omitempty"`
	Order                 string               `xml:"Order,omitempty"`
	IsAdult               int64                `xml:"IsAdult,omitempty"`
	IsCarBodySeller       int64                `xml:"IsCarBodySeller,omitempty"`
	Availability          string               `xml:"Availability,omitempty"`
}
type InfoResult struct {
	ItemCode *InfoItemCode `xml:"ItemCode,omitempty"`
	Hits     []*InfoHit    `xml:"Hit,omitempty"`
}
type InfoResultSet struct {
	Attr_xmlns                 string      `xml:"xmlns,attr"`
	Attr_xsi                   string      `xml:"xmlns xsi,attr"`
	Attr_xsi_schemaLocation    string      `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_firstResultPosition   string      `xml:"firstResultPosition,attr"`
	Attr_totalResultsAvailable string      `xml:"totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string      `xml:"totalResultsReturned,attr"`
	Result                     *InfoResult `xml:"Result,omitempty"`
}

func (c *Client) GetShoppingItemInfo(itemCode string) ([]InfoHit, error) {
	spath := fmt.Sprintf("/itemLookup?appid=%s&itemcode=%s&responsegroup=large", c.AppID, itemCode)
	req, err := c.newRequest("GET", spath, nil)
	if err != nil {
		fmt.Println("[ERROR] fail newRequest in GetShoppingItemInfo")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] fail HTTPClient.Do in GetShoppingItemInfo")
		return nil, err
	}

	if res.StatusCode >= 400 {
		fmt.Println("[ERROR] Response fail status code: ", res.StatusCode)
		return nil, err
	}

	var resultSet InfoResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		fmt.Println("[ERROR] fail decodeBody in GetShoppingItemInfo")
		return nil, err
	}

	ilist := []InfoHit{}
	for _, item := range resultSet.Result.Hits {
		ilist = append(ilist, *item)
	}

	return ilist, nil
}
