package responses

type ToItemPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesOrder                  string `json:"SalesOrder"` 
			SalesOrderItem              string `json:"SalesOrderItem"`
			PartnerFunction             string `json:"PartnerFunction"` 
			Customer                    string `json:"Customer"`
			Supplier                    string `json:"Supplier"`
			Personnel                   string `json:"Personnel"` 
			ContactPerson               string `json:"ContactPerson"`
		} `json:"results"`
	} `json:"d"`
}
