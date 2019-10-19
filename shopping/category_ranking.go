package shopping

import (
	"fmt"
)

type RankImage struct {
	Id     string `xml:"Id,omitempty"`
	Small  string `xml:"Small,omitempty"`
	Medium string `xml:"Medium,omitempty"`
}
type RankReview struct {
	Rate  string `xml:"Rate,omitempty"`
	Count string `xml:"Count,omitempty"`
	Url   string `xml:"Url,omitempty"`
}
type RankStore struct {
	Id   string `xml:"Id,omitempty"`
	Name string `xml:"Name,omitempty"`
	Url  string `xml:"Url,omitempty"`
}
type RankingInfo struct {
	LastModified string `xml:"LastModified,omitempty"`
	StartDate    string `xml:"StartDate,omitempty"`
	EndDate      string `xml:"EndDate,omitempty"`
	CategoryId   int64  `xml:"CategoryId,omitempty"`
	Gender       string `xml:"Gender,omitempty"`
	Generation   string `xml:"Generation,omitempty"`
	Period       string `xml:"Period,omitempty"`
}
type Ranking struct {
	Rank   int64       `xml:"rank,attr"`
	Vector string      `xml:"vector,attr"`
	Type   string      `xml:"type,attr"`
	Name   string      `xml:"Name,omitempty"`
	Code   string      `xml:"Code,omitempty"`
	Url    string      `xml:"Url,omitempty"`
	Image  *RankImage  `xml:"Image,omitempty"`
	Review *RankReview `xml:"Review,omitempty"`
	Store  *RankStore  `xml:"Store,omitempty"`
}
type RankingResult struct {
	RankingData []*Ranking   `xml:"RankingData,omitempty"`
	RankingInfo *RankingInfo `xml:"RankingInfo,omitempty"`
}
type RankingResultSet struct {
	Attr_firstResultPosition   string         `xml:"firstResultPosition,attr"`
	Attr_xsi_schemaLocation    string         `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string         `xml:"totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string         `xml:"totalResultsReturned,attr"`
	Attr_xmlns                 string         `xml:"xmlns,attr"`
	Attr_xsi                   string         `xml:"xmlns xsi,attr"`
	RankingResult              *RankingResult `xml:"Result,omitempty"`
}

func (c *Client) GetShoppingCategoryRanking(categoryId string, offset int64) ([]Ranking, error) {
	// queryWord := url.QueryEscape(keyword)
	if categoryId == "" {
		return nil, fmt.Errorf("empty category id")
	}

	spath := fmt.Sprintf("/categoryRanking?appid=%s&category_id=%s&offset=%d", c.AppID, categoryId, offset)
	req, err := c.newRequest("GET", spath, nil)
	if err != nil {
		fmt.Println("[ERROR] fail newRequest in GetShoppingCategoryRanking")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println("[ERROR] fail HTTPClient.Do in GetShoppingCategoryRanking")
		return nil, err
	}

	// status check
	if res.StatusCode >= 400 {
		fmt.Println("[ERROR] Response fail status code", res.StatusCode)
		fmt.Printf("%+v\n", res.Uncompressed)
		return nil, err
	}

	var resultSet RankingResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		fmt.Println("[ERROR] fail decodeBody in GetShoppingCategoryRanking")
		return nil, err
	}

	ilist := []Ranking{}
	for _, item := range resultSet.RankingResult.RankingData {
		ilist = append(ilist, *item)
	}

	return ilist, nil
}
