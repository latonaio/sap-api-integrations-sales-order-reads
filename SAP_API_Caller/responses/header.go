package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			SalesOrder                     string `json:"SalesOrder"`
			SalesOrderType                 string `json:"SalesOrderType"`
			SalesOrganization              string `json:"SalesOrganization"`
			DistributionChannel            string `json:"DistributionChannel"`
			OrganizationDivision           string `json:"OrganizationDivision"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOffice                    string `json:"SalesOffice"`
			SalesDistrict                  string `json:"SalesDistrict"`
			SoldToParty                    string `json:"SoldToParty"`
			CreationDate                   string `json:"CreationDate"`
			LastChangeDate                 string `json:"LastChangeDate"`
			ExternalDocumentID             string `json:"ExternalDocumentID"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			PurchaseOrderByCustomer        string `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderDate      string `json:"CustomerPurchaseOrderDate"`
			SalesOrderDate                 string `json:"SalesOrderDate"`
			TotalNetAmount                 string `json:"TotalNetAmount"`
			OverallDeliveryStatus          string `json:"OverallDeliveryStatus"`
			TotalBlockStatus               string `json:"TotalBlockStatus"`
			OverallOrdReltdBillgStatus     string `json:"OverallOrdReltdBillgStatus"`
			OverallSDDocReferenceStatus    string `json:"OverallSDDocReferenceStatus"`
			TransactionCurrency            string `json:"TransactionCurrency"`
			SDDocumentReason               string `json:"SDDocumentReason"`
			PricingDate                    string `json:"PricingDate"`
			PriceDetnExchangeRate          string `json:"PriceDetnExchangeRate"`
			RequestedDeliveryDate          string `json:"RequestedDeliveryDate"`
			ShippingCondition              string `json:"ShippingCondition"`
			CompleteDeliveryIsDefined      bool   `json:"CompleteDeliveryIsDefined"`
			ShippingType                   string `json:"ShippingType"`
			HeaderBillingBlockReason       string `json:"HeaderBillingBlockReason"`
			DeliveryBlockReason            string `json:"DeliveryBlockReason"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			CustomerPriceGroup             string `json:"CustomerPriceGroup"`
			PriceListType                  string `json:"PriceListType"`
			CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
			PaymentMethod                  string `json:"PaymentMethod"`
			ReferenceSDDocument            string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentCategory    string `json:"ReferenceSDDocumentCategory"`
			CustomerAccountAssignmentGroup string `json:"CustomerAccountAssignmentGroup"`
			AccountingExchangeRate         string `json:"AccountingExchangeRate"`
			CustomerGroup                  string `json:"CustomerGroup"`
			AdditionalCustomerGroup1       string `json:"AdditionalCustomerGroup1"`
			AdditionalCustomerGroup2       string `json:"AdditionalCustomerGroup2"`
			AdditionalCustomerGroup3       string `json:"AdditionalCustomerGroup3"`
			AdditionalCustomerGroup4       string `json:"AdditionalCustomerGroup4"`
			AdditionalCustomerGroup5       string `json:"AdditionalCustomerGroup5"`
			CustomerTaxClassification1     string `json:"CustomerTaxClassification1"`
			TotalCreditCheckStatus         string `json:"TotalCreditCheckStatus"`
			BillingDocumentDate            string `json:"BillingDocumentDate"`
			ToHeaderPartner                struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
