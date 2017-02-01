package auction

import (
	"context"
	"errors"
	"fmt"
)

type ChiChildCategory struct {
	ChiCategoryId       int    `xml:"urn:yahoo:jp:auc:categoryTree CategoryId,omitempty"`
	ChiCategoryIdPath   string `xml:"urn:yahoo:jp:auc:categoryTree CategoryIdPath,omitempty"`
	ChiCategoryName     string `xml:"urn:yahoo:jp:auc:categoryTree CategoryName,omitempty"`
	ChiCategoryPath     string `xml:"urn:yahoo:jp:auc:categoryTree CategoryPath,omitempty"`
	ChiDepth            int    `xml:"urn:yahoo:jp:auc:categoryTree Depth,omitempty"`
	ChiIsAdult          bool   `xml:"urn:yahoo:jp:auc:categoryTree IsAdult,omitempty"`
	ChiIsLeaf           bool   `xml:"urn:yahoo:jp:auc:categoryTree IsLeaf,omitempty"`
	ChiIsLeafToLink     bool   `xml:"urn:yahoo:jp:auc:categoryTree IsLeafToLink,omitempty"`
	ChiIsLink           bool   `xml:"urn:yahoo:jp:auc:categoryTree IsLink,omitempty"`
	ChiNumOfAuctions    int    `xml:"urn:yahoo:jp:auc:categoryTree NumOfAuctions,omitempty"`
	ChiOrder            int    `xml:"urn:yahoo:jp:auc:categoryTree Order,omitempty"`
	ChiParentCategoryId int    `xml:"urn:yahoo:jp:auc:categoryTree ParentCategoryId,omitempty"`
}

type ChiResult struct {
	ChiCategoryId       int                 `xml:"urn:yahoo:jp:auc:categoryTree CategoryId,omitempty"`
	ChiCategoryIdPath   string              `xml:"urn:yahoo:jp:auc:categoryTree CategoryIdPath,omitempty"`
	ChiCategoryName     string              `xml:"urn:yahoo:jp:auc:categoryTree CategoryName,omitempty"`
	ChiCategoryPath     string              `xml:"urn:yahoo:jp:auc:categoryTree CategoryPath,omitempty"`
	ChiChildCategory    []*ChiChildCategory `xml:"urn:yahoo:jp:auc:categoryTree ChildCategory,omitempty"`
	ChiChildCategoryNum int                 `xml:"urn:yahoo:jp:auc:categoryTree ChildCategoryNum,omitempty"`
	ChiDepth            int                 `xml:"urn:yahoo:jp:auc:categoryTree Depth,omitempty"`
	ChiIsAdult          bool                `xml:"urn:yahoo:jp:auc:categoryTree IsAdult,omitempty"`
	ChiIsLeaf           bool                `xml:"urn:yahoo:jp:auc:categoryTree IsLeaf,omitempty"`
	ChiIsLeafToLink     bool                `xml:"urn:yahoo:jp:auc:categoryTree IsLeafToLink,omitempty"`
	ChiIsLink           bool                `xml:"urn:yahoo:jp:auc:categoryTree IsLink,omitempty"`
	ChiOrder            int                 `xml:"urn:yahoo:jp:auc:categoryTree Order,omitempty"`
}

type ChiResultSet struct {
	Attr_firstResultPosition   string     `xml:"firstResultPosition,attr"`
	Attr_xsi_schemaLocation    string     `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	Attr_totalResultsAvailable string     `xml:"totalResultsAvailable,attr"`
	Attr_totalResultsReturned  string     `xml:"totalResultsReturned,attr"`
	Attr_xmlns                 string     `xml:"xmlns,attr"`
	Attr_xsi                   string     `xml:"xmlns xsi,attr"`
	ChiResult                  *ChiResult `xml:"urn:yahoo:jp:auc:categoryTree Result,omitempty"`
}

func (c *Client) GetCategoryList(ctx context.Context) ([]ChiChildCategory, error) {
	spath := fmt.Sprintf("/categoryTree" + "?appid=" + c.AuthToken)
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		c.Logger.Println("[ERROR] fail newRequest in GetCategoryIdList")
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Println("[ERROR] fail HTTPClient.Do in GetCategoryIdList")
		return nil, err
	}

	// status check
	if res.StatusCode >= 400 {
		c.Logger.Println("[ERROR] Response fail status code")
		return nil, err
	}

	var resultSet ChiResultSet
	if err := decodeBody(res, &resultSet); err != nil {
		c.Logger.Println("[ERROR] fail decodeBody in GetCategoryIdList")
		return nil, err
	}

	clist := []ChiChildCategory{}
	for _, category := range resultSet.ChiResult.ChiChildCategory {
		clist = append(clist, *category)
	}

	return clist, nil
}

func (c *Client) GetCategoryIdList(clist []ChiChildCategory) ([]int, error) {
	cids := []int{}

	for _, category := range clist {
		cids = append(cids, category.ChiCategoryId)
	}

	return cids, nil
}

func (c *Client) GetCategoryNameList(clist []ChiChildCategory) ([]string, error) {
	cNames := []string{}
	for _, category := range clist {
		cNames = append(cNames, category.ChiCategoryName)
	}

	return cNames, nil
}

func (c *Client) GetCategoryIdByName(clist []ChiChildCategory, categoryName string) (int, error) {
	for _, category := range clist {
		if category.ChiCategoryName == categoryName {
			return category.ChiCategoryId, nil
		}
	}

	return -1, errors.New("Not Found categoryName")
}
