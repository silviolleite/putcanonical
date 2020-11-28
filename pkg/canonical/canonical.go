package canonical

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type canonical struct {
	sku string
	id  string
}

type Canonicals struct {
	items []*canonical
	meli  *MeliService
	token string
}

func New(hc *http.Client, ls []string, token string) *Canonicals {
	m := NewClient(hc)
	canonicals := Canonicals{
		meli:  m.Meli,
		token: token,
	}
	for _, v := range ls {
		if v != "" {
			l := strings.Split(v, ",")
			canonicals.items = append(canonicals.items, &canonical{l[0], l[1]})
		}
	}

	return &canonicals
}

func (cn *Canonicals) Run() {
	var wg sync.WaitGroup
	wg.Add(len(cn.items))

	cn.printHeader()

	for _, c := range cn.items {
		go cn.process(c, &wg)
	}

	wg.Wait()
	fmt.Printf("\n\n****** Finished *******\n\n")
}

func (cn *Canonicals) process(c *canonical, wg *sync.WaitGroup) error {
	defer wg.Done()
	item, err := cn.meli.ItemsVariations(c.id)
	Check(err)
	item = setSKU(item, c.sku)

	payload, err := json.Marshal(item)
	Check(err)

	fmt.Printf("\n--------------------------------------------------------\nProcessing: sku: %s \nID: %s Payload: %s \n", c.sku, c.id, payload)

	err = cn.meli.PutSKU(c.id, cn.token, payload)
	Check(err)

	fmt.Printf("\n--------------------------------------------------------\n[SUCCESS]\nsku: %s \nID: %s \n", c.sku, c.id)
	return nil
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func setSKU(i *Item, sku string) (item *Item) {
	i.Sku = sku
	for k, _ := range i.Variations {
		i.Variations[k].Sku = sku
	}
	return i
}

func (cn *Canonicals) printHeader() {
	fmt.Printf("\n****************************************** *******\n")
	fmt.Printf("*** Start update Meli Items with Canonical SKU ***\n")
	fmt.Printf("****************************************** *******\n\n")

	fmt.Printf("Items to process: %d \n\n", len(cn.items))
	fmt.Printf("****************************************** *******\n")
}
