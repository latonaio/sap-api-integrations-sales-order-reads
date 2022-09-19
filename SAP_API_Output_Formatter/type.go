package sap_api_output_formatter

type SalesOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	SalesOrder    string `json:"sales_order"`
	Deleted       string `json:"deleted"`
}

type Header struct {
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
	ToHeaderPartner                string `json:"to_Partner"`
	ToItem                         string `json:"to_Item"`
}

type Item struct {
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
	ToItemPartner               string `json:"to_Partner"`
	ToItemPricingElement        string `json:"to_PricingElement"`
	ToItemScheduleLine          string `json:"to_ScheduleLine"`
}

type ToHeaderPartner struct {
	SalesOrder      string `json:"SalesOrder"`
	PartnerFunction string `json:"PartnerFunction"`
	Customer        string `json:"Customer"`
	Supplier        string `json:"Supplier"`
}

type ToItem struct {
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
	ToItemPartner               string `json:"to_Partner"`
	ToItemPricingElement        string `json:"to_PricingElement"`
	ToItemScheduleLine          string `json:"to_ScheduleLine"`
}

type ToItemPartner struct {
	SalesOrder                  string `json:"SalesOrder"` 
	SalesOrderItem              string `json:"SalesOrderItem"`
	PartnerFunction             string `json:"PartnerFunction"` 
	Customer                    string `json:"Customer"`
	Supplier                    string `json:"Supplier"`
	Personnel                   string `json:"Personnel"` 
	ContactPerson               string `json:"ContactPerson"`
}

type ToItemPricingElement struct {
	SalesOrder                     string `json:"SalesOrder"`
	SalesOrderItem                 string `json:"SalesOrderItem"`
	PricingProcedureStep           string `json:"PricingProcedureStep"`
	PricingProcedureCounter        string `json:"PricingProcedureCounter"`
	ConditionType                  string `json:"ConditionType"`
	PriceConditionDeterminationDte string `json:"PriceConditionDeterminationDte"`
	ConditionCalculationType       string `json:"ConditionCalculationType"`
	ConditionBaseValue             string `json:"ConditionBaseValue"`
	ConditionRateValue             string `json:"ConditionRateValue"`
	ConditionCurrency              string `json:"ConditionCurrency"`
	ConditionQuantity              string `json:"ConditionQuantity"`
	ConditionQuantityUnit          string `json:"ConditionQuantityUnit"`
	ConditionCategory              string `json:"ConditionCategory"`
	PricingScaleType               string `json:"PricingScaleType"`
	ConditionRecord                string `json:"ConditionRecord"`
	ConditionSequentialNumber      string `json:"ConditionSequentialNumber"`
	TaxCode                        string `json:"TaxCode"`
	ConditionAmount                string `json:"ConditionAmount"`
	TransactionCurrency            string `json:"TransactionCurrency"`
	PricingScaleBasis              string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue       string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit        string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency    string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged     bool   `json:"ConditionIsManuallyChanged"`
}

type ToItemScheduleLine struct {
	SalesOrder                    string `json:"SalesOrder"`
	SalesOrderItem                string `json:"SalesOrderItem"`
	ScheduleLine                  string `json:"ScheduleLine"`
	RequestedDeliveryDate         string `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate         string `json:"ConfirmedDeliveryDate"`
	OrderQuantityUnit             string `json:"OrderQuantityUnit"`
	ScheduleLineOrderQuantity     string `json:"ScheduleLineOrderQuantity"`
	ConfdOrderQtyByMatlAvailCheck string `json:"ConfdOrderQtyByMatlAvailCheck"`
	DeliveredQtyInOrderQtyUnit    string `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit string `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	CorrectedQtyInOrderQtyUnit    string `json:"CorrectedQtyInOrderQtyUnit"`
	DelivBlockReasonForSchedLine  string `json:"DelivBlockReasonForSchedLine"`
}
