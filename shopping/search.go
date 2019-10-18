package shopping

import (
	"fmt"
	"net/url"
)

// type Chiroot struct {
// 	ChiResultSet *ChiResultSet `xml:"urn:yahoo:jp:itemSearch ResultSet,omitempty"`
// }

// type ChiSearchCharityOption struct {
// 	ChiProportion int `xml:"urn:yahoo:jp:itemSearch Proportion,omitempty"`
// }

// type ChiImage struct {
// 	Attr_height string `xml:" height,attr"`
// 	Attr_width  string `xml:" width,attr"`
// 	Text        string `xml:",chardata"`
// }
type SrchImage struct {
	Id     string `xml:"Id,omitempty"`
	Small  string `xml:"Small,omitempty"`
	Medium string `xml:"Medium,omitempty"`
}
type SrchExtImage struct {
	Url    string `xml:"Url,omitempty"`
	Width  string `xml:"Width,omitempty"`
	Height string `xml:"Height,omitempty"`
}
type SrchReview struct {
	Rate  string `xml:"Rate,omitempty"`
	Count string `xml:"Count,omitempty"`
	Url   string `xml:"Url,omitempty"`
}
type SrchAffiliate struct {
	Rate float64 `xml:"Rate,omitempty"`
}
type SrchPrice struct {
	Currency string `xml:"currency,attr"`
	Value    int64  `xml:",chardata"`
}
type SrchPriceLabel struct {
	Currency            bool   `xml:"currency,attr"`
	FixedPrice          int64  `xml:FixedPrice,omitempty`
	DefaultPrice        int64  `xml:DefaultPrice,omitempty`
	SalePrice           int64  `xml:SalePrice,omitempty`
	PremiumPriceStatus  int64  `xml:PremiumPriceStatus,omitempty`
	PremiumPrice        int64  `xml:PremiumPrice,omitempty`
	PremiumDiscountType string `xml:PremiumDiscountType,omitempty`
	PremiumDiscountRate int64  `xml:PremiumDiscountRate,omitempty`
	PeriodStart         string `xml:PeriodStart,omitempty`
	PeriodEnd           string `xml:PeriodEnd,omitempty`
}
type SrchPoint struct {
	Amount        int64 `xml:"Amount,omitempty"`
	Times         int64 `xml:"Times,omitempty"`
	PremiumAmount int64 `xml:"PremiumAmount,omitempty"`
	PremiumTimes  int64 `xml:"PremiumTimes,omitempty"`
}
type SrchShipping struct {
	Code int64  `xml:"Code,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type SrchNode struct {
	Id   string `xml:"Id,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type SrchCategory struct {
	Current *SrchNode `xml:"Current,omitempty"`
}
type SrchCategoryIdPath struct {
	Category []SrchNode `xml:"Category,omitempty"`
}
type SrchBrand struct {
	Id string `xml:"Id,omitempty"`
}
type SrchBrands struct {
	Name string     `xml:"Name,omitempty"`
	Path *SrchBrand `xml:"Path,omitempty"`
}
type SrchMethod struct {
	Code int64  `xml:"Code,omitempty"`
	Name string `xml:"Name,omitempty"`
}
type SrchPayment struct {
	Method []SrchMethod `xml:"Method,omitempty"`
}
type SrchRatings struct {
	Rate       float64 `xml:"Rate,omitempty"`
	Count      int64   `xml:"Count,omitempty"`
	Total      int64   `xml:"Total,omitempty"`
	DetailRate float64 `xml:"DetailRate,omitempty"`
}
type StoreImage struct {
	Id     string `xml:"Id,omitempty"`
	Medium string `xml:"Medium,omitempty"`
}
type SrchStore struct {
	Id           string       `xml:"Id,omitempty"`
	Name         string       `xml:"Name,omitempty"`
	Url          string       `xml:"Url,omitempty"`
	IsBestStore  string       `xml:"IsBestStore,omitempty"`
	Payment      *SrchPayment `xml:"Payment,omitempty"`
	Ratings      *SrchRatings `xml:"Ratings,omitempty"`
	Image        *StoreImage  `xml:"Image,omitempty"`
	IsPMallStore bool         `xml:"IsPMallStore,omitempty"`
}
type SrchDeliveryinfo struct {
	Area     string `xml:"Area,omitempty"`
	Deadline string `xml:"Deadline,omitempty"`
	Day      int64  `xml:"Day,omitempty"`
}
type Request struct {
	Query string `xml:"Query"`
}
type Hit struct {
	Index          int64               `xml:"index,attr"`
	Name           string              `xml:"Name,omitempty"`
	Description    string              `xml:"Description,omitempty"`
	Headline       string              `xml:"Headline,omitempty"`
	Url            string              `xml:"Url,omitempty"`
	Availability   string              `xml:"Availability,omitempty"`
	Code           string              `xml:"Code,omitempty"`
	Image          *SrchImage          `xml:"Image,omitempty"`
	ExImage        *SrchExtImage       `xml:"ExImage,omitempty"`
	Review         *SrchReview         `xml:"Review,omitempty"`
	Affiliate      *SrchAffiliate      `xml:"Affiliate,omitempty"`
	Price          *SrchPrice          `xml:"Price,omitempty"`
	PriceLabel     *SrchPriceLabel     `xml:"PriceLabel,omitempty"`
	Point          *SrchPoint          `xml:"Point,omitempty"`
	Shipping       *SrchShipping       `xml:"Shipping,omitempty"`
	Category       *SrchCategory       `xml:"Category,omitempty"`
	CategoryIdPath *SrchCategoryIdPath `xml:"CategoryIdPath,omitempty"`
	Brands         *SrchBrands         `xml:"Brands,omitempty"`
	JanCode        string              `xml:"JanCode,omitempty"`
	Model          string              `xml:"Model,omitempty"`
	IsbnCode       string              `xml:"IsbnCode,omitempty"`
	ReleaseDate    string              `xml:"ReleaseDate,omitempty"`
	Store          *SrchStore          `xml:"Store,omitempty"`
	IsAdult        int64               `xml:"IsAdult,omitempty"`
	Deliveryinfo   *SrchDeliveryinfo   `xml:"Deliveryinfo,omitempty"`
}
type SearchResult struct {
	Hits    []*Hit   `xml:"Hit,omitempty"`
	Request *Request `xml:"Request,omitempty"`
}
type SearchResultSet struct {
	Attr_firstResultPosition   string        `xml:"firstResultPosition,attr"`
	Attr_xsi_schemaLocation    string        `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string        `xml:"totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string        `xml:"totalResultsReturned,attr"`
	Attr_xmlns                 string        `xml:"xmlns,attr"`
	Attr_xsi                   string        `xml:"xmlns xsi,attr"`
	SearchResult               *SearchResult `xml:"Result,omitempty"`
}

func (c *Client) GetShoppingItemListBySearch(keyword string) ([]Hit, error) {
	queryWord := url.QueryEscape(keyword)
	if queryWord == "" {
		return nil, fmt.Errorf("empty keyword")
	}
	// queryWord, err := url.Parse(keyword)
	// if err != nil {
	// 	return nil, err
	// }

	spath := fmt.Sprintf("/itemSearch" + "?appid=" + c.AppID + "&query=" + queryWord)
	req, err := c.newRequest("GET", spath, nil)
	if err != nil {
		fmt.Println("[ERROR] fail newRequest in GetShoppingItemsList")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] fail HTTPClient.Do in GetShoppingItemsList")
		return nil, err
	}

	// status check
	if res.StatusCode >= 400 {
		fmt.Println("[ERROR] Response fail status code", res.StatusCode)
		fmt.Printf("%+v\n", res.Uncompressed)
		return nil, err
	}

	var resultSet SearchResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		fmt.Println("[ERROR] fail decodeBody in GetShoppingItemsList")
		return nil, err
	}

	ilist := []Hit{}
	for _, item := range resultSet.SearchResult.Hits {
		ilist = append(ilist, *item)
	}

	return ilist, nil
}
