package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesOrder                  string `json:"SalesOrder"`
			SalesOrderItem              string `json:"SalesOrderItem"`
			SalesOrderItemCategory      string `json:"SalesOrderItemCategory"`
			SalesOrderItemText          string `json:"SalesOrderItemText"`
			PurchaseOrderByCustomer     string `json:"PurchaseOrderByCustomer"`
			Material                    string `json:"Material"`
			MaterialByCustomer          string `json:"MaterialByCustomer"`
			PricingDate                 string `json:"PricingDate"`
			BillingPlan                 string `json:"BillingPlan"`
			RequestedQuantity           string `json:"RequestedQuantity"`
			RequestedQuantityUnit       string `json:"RequestedQuantityUnit"`
			OrderQuantityUnit           string `json:"OrderQuantityUnit"`
			ConfdDelivQtyInOrderQtyUnit string `json:"ConfdDelivQtyInOrderQtyUnit"`
			ItemGrossWeight             string `json:"ItemGrossWeight"`
			ItemNetWeight               string `json:"ItemNetWeight"`
			ItemWeightUnit              string `json:"ItemWeightUnit"`
			ItemVolume                  string `json:"ItemVolume"`
			ItemVolumeUnit              string `json:"ItemVolumeUnit"`
			TransactionCurrency         string `json:"TransactionCurrency"`
			NetAmount                   string `json:"NetAmount"`
			MaterialGroup               string `json:"MaterialGroup"`
			MaterialPricingGroup        string `json:"MaterialPricingGroup"`
			BillingDocumentDate         string `json:"BillingDocumentDate"`
			Batch                       string `json:"Batch"`
			ProductionPlant             string `json:"ProductionPlant"`
			StorageLocation             string `json:"StorageLocation"`
			DeliveryGroup               string `json:"DeliveryGroup"`
			ShippingPoint               string `json:"ShippingPoint"`
			ShippingType                string `json:"ShippingType"`
			DeliveryPriority            string `json:"DeliveryPriority"`
			IncotermsClassification     string `json:"IncotermsClassification"`
			TaxAmount                   string `json:"TaxAmount"`
			ProductTaxClassification1   string `json:"ProductTaxClassification1"`
			MatlAccountAssignmentGroup  string `json:"MatlAccountAssignmentGroup"`
			CostAmount                  string `json:"CostAmount"`
			CustomerPaymentTerms        string `json:"CustomerPaymentTerms"`
			CustomerGroup               string `json:"CustomerGroup"`
			SalesDocumentRjcnReason     string `json:"SalesDocumentRjcnReason"`
			ItemBillingBlockReason      string `json:"ItemBillingBlockReason"`
			WBSElement                  string `json:"WBSElement"`
			ProfitCenter                string `json:"ProfitCenter"`
			AccountingExchangeRate      string `json:"AccountingExchangeRate"`
			ReferenceSDDocument         string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem     string `json:"ReferenceSDDocumentItem"`
			SDProcessStatus             string `json:"SDProcessStatus"`
			DeliveryStatus              string `json:"DeliveryStatus"`
			OrderRelatedBillingStatus   string `json:"OrderRelatedBillingStatus"`
			ToItemPartner               struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
			ToItemScheduleLine struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ScheduleLine"`
		} `json:"results"`
	} `json:"d"`
}
