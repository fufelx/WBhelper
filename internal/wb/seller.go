package wb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetSellerInfo() (*SellerInfo, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/api/v1/seller-info", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info SellerInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
