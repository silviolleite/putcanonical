package canonical

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MeliService service

type Variation struct {
	ID  int    `json:"id"`
	Sku string `json:"seller_custom_field"`
}

type Item struct {
	Sku        string       `json:"seller_custom_field"`
	Variations []*Variation `json:"variations"`
}

func (m *MeliService) Items(ID string) (value string, err error) {
	path := items + ID

	req, err := m.client.NewRequest(http.MethodGet, path, nil, "")
	if err != nil {
		return value, err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return value, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (m *MeliService) ItemsVariations(ID string) (item *Item, err error) {
	path := items + ID

	req, err := m.client.NewRequest(http.MethodGet, path, nil, "")
	if err != nil {
		return item, err
	}

	q := req.URL.Query()
	q.Set("attributes", "variations")

	req.URL.RawQuery = q.Encode()

	resp, err := m.client.Do(req)
	if err != nil {
		return item, err
	}

	defer resp.Body.Close()

	item = &Item{}

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return item, err
	}

	return item, nil
}

func (m *MeliService) PutSKU(ID, token string, payload []byte) error {
	path := items + ID

	req, err := m.client.NewRequest(http.MethodPut, path, bytes.NewReader(payload), token)
	if err != nil {
		return err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status code: ", resp.StatusCode)

	return nil
}
