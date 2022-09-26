# sap-api-integrations-sales-order-reads  
sap-api-integrations-sales-order-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 受注データを取得するマイクロサービスです。    
sap-api-integrations-sales-order-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-sales-order-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview

## 動作環境  
sap-api-integrations-sales-order-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-sales-order-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-order-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview  
* APIサービス名(=baseURL): API_SALES_ORDER_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-order-reads には、次の API をコールするためのリソースが含まれています。  

* A_SalesOrder（受注 - ヘッダ）※受注の詳細データを取得するために、ToHeaderPartner、ToItem、と合わせて利用されます。
* A_SalesOrderItem（受注 - 明細）※受注明細の詳細データを取得するために、ToItemPricingElement、ToItemScheduleLine、と合わせて利用されます。
* ToHeaderPartner（受注 - ヘッダ取引先）
* ToItem（受注 - 明細）
* ToItemPricingElement（受注 - 明細価格条件）
* ToItemScheduleLine（受注 - 明細納入日程行）

## API への 値入力条件 の 初期値
sap-api-integrations-sales-order-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.SalesOrder.SalesOrder（受注番号）
* inoutSDC.SalesOrder.SalesOrderItem.SalesOrderItem（受注明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "SAPSalesOrderReads",
	"accepter": ["Header"],
	"sales_order": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPSalesOrderReads",
	"accepter": ["All"],
	"sales_order": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetSalesOrder(salesOrder, salesOrderItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(salesOrder)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(salesOrder, salesOrderItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 受注 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"SalesOrder" ～ "ToItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-order-reads/SAP_API_Caller/caller.go#L60",
	"function": "sap-api-integrations-sales-order-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"SalesOrder": "1",
			"SalesOrderType": "OR1",
			"SalesOrganization": "0001",
			"DistributionChannel": "01",
			"OrganizationDivision": "01",
			"SalesGroup": "",
			"SalesOffice": "",
			"SalesDistrict": "000001",
			"SoldToParty": "1",
			"CreationDate": "2022-09-10",
			"LastChangeDate": "",
			"ExternalDocumentID": "",
			"LastChangeDateTime": "2022-09-10T18:02:13+09:00",
			"PurchaseOrderByCustomer": "Test",
			"CustomerPurchaseOrderDate": "",
			"SalesOrderDate": "2022-09-10",
			"TotalNetAmount": "1000.00",
			"OverallDeliveryStatus": "",
			"TotalBlockStatus": "",
			"OverallOrdReltdBillgStatus": "",
			"OverallSDDocReferenceStatus": "",
			"TransactionCurrency": "EUR",
			"SDDocumentReason": "",
			"PricingDate": "2022-09-10",
			"PriceDetnExchangeRate": "",
			"RequestedDeliveryDate": "2022-09-12",
			"ShippingCondition": "01",
			"CompleteDeliveryIsDefined": false,
			"ShippingType": "",
			"HeaderBillingBlockReason": "",
			"DeliveryBlockReason": "",
			"IncotermsClassification": "FH",
			"CustomerPriceGroup": "01",
			"PriceListType": "",
			"CustomerPaymentTerms": "0001",
			"PaymentMethod": "",
			"ReferenceSDDocument": "",
			"ReferenceSDDocumentCategory": "",
			"CustomerAccountAssignmentGroup": "01",
			"AccountingExchangeRate": "0.00000",
			"CustomerGroup": "",
			"AdditionalCustomerGroup1": "",
			"AdditionalCustomerGroup2": "",
			"AdditionalCustomerGroup3": "",
			"AdditionalCustomerGroup4": "",
			"AdditionalCustomerGroup5": "",
			"CustomerTaxClassification1": "",
			"TotalCreditCheckStatus": "",
			"BillingDocumentDate": "",
			"to_Partner": "http://100.21.57.120:8080/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('1')/to_Partner",
			"to_Item": "http://100.21.57.120:8080/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('1')/to_Item"
		}
	],
	"time": "2022-09-10T19:38:33+09:00"
}

```

## API サービス名
【Sales Order > DPFM_API_SALES_ORDER_SRV】

## 利用可能な API タイプ
| READS | CREATES | UPDATES | DELETES | CANCELS | GETS PDF | 
| ----- | ------- | ------- | ------- | ------- | -------- | 
| ●    | ●      | ●      | ●      | ●      | ●       | 
|       | 

## API 名
【Sales Order Header > A_SalesOrder】

## ＜項目＞
| Property                                                               | Description                                                                                                                                                                                                    | 
| ---------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | 
| BusinessPartner                                                        | 売り手企業のビジネスパートナーコード。ビジネスパートナーマスタから指定する。                                                                                                                                   | 
| SalesOrder                                                             | 売り手企業が受注した受注番号                                                                                                                                                                                   | 
| SalesOrderType                                                         | 受注タイプ。次から選択する。OR:標準受注、ST:仕入先直送                                             | 
| SalesOrganization                                                      | 販売組織。固定値”0001”を指定する。                                             | 
| DistributionChannel                                                    | 流通チャネル。次から選択する。DS:直接販売、EC:EC販売                                                 | 
| OrganizationDivision                                                   | 製品部門。次の事業領域から選択する。MF:製造、RT:小売、TR:卸売                                              | 
| SalesGroup                                                             | 営業グループ。固定値”001”を指定する。                                              | 
| SalesOffice                                                            | 営業所。固定値”0001”を指定する。                                             | 
| SalesDistrict                                                          | 販売地域。ビジネスパートナーマスタの販売エリアの販売地域データから提案される。変更の必要があれば、販売地域マスタから選択して指定する。                                                                         | 
| SoldToParty                                                            | 買い手企業のビジネスパートナーコード。ビジネスパートナーマスタから指定する。                                                                                                                                   | 
| CreationDate                                                           | 作成日。自動生成される。                                                                                                                                                                                       | 
| LastChangeDate                                                         | 最終変更日時。自動生成される。                                                                                                                                                                                 | 
| PurchaseExternalDocumentID                                             | 買い手企業側の周辺業務システム内における発注番号。なお、それ以外の文書I Dがある場合には、その番号。                                                                                                            | 
| SalesExternalDocumentID                                                | 買い手企業側の周辺業務システム内における発注番号。なお、それ以外の文書I Dがある場合には、その番号。                                                                                                            | 
| PurchaseOrderByCustomer                                                | 買い手企業側の周辺業務システム内における発注番号。                                                                                                                                                             | 
| SalesOrderBySupplier                                                   | 売り手企業側の周辺業務システム内における受注番号。                                                                                                                                                             | 
| CustomerPurchaseOrderDate                                              | 買い手企業側の周辺業務システム内における発注日付。                                                                                                                                                             | 
| SalesOrderDate                                                         | 売り手企業側の周辺業務システム内における受注日付。                                                                                                                                                             | 
| TotalNetAmount                                                         | 受注総額。買い手企業側の周辺業務システム内における発注総額を連携する。                                                                                                                                         | 
| OverallDeliveryStatus                                                  | 出荷状況。以下のステータスが連携される。未出荷：NP、部分出荷完了済み：PP 、出荷完了済み：CL                                                       | 
| TotalBlockStatus                                                       | 受注ブロック状況。受注をブロックしたい場合は、”B”を入力する。                                                                                                                                                | 
| OverallOrdReltdBillgStatus                                             | 請求ステータス。以下のステータスが連携される。未請求：NP、一部請求完了済み：PP、請求完了済み：CL                                                       | 
| OverallSDDocReferenceStatus                                            | 受注参照状況。以下のステータスが連携される。見積参照受注：QT、引合参照受注：IN                                                       | 
| TransactionCurrency                                                    | 取引通貨。ビジネスパートナーマスタの販売エリアの通貨が提案される。変更の必要があれば、通貨コードマスタから選択して指定する。                                                                                   | 
| PricingDate                                                            | 価格設定日付。通常、発注日と同じ日付を入力する。価格設定日付を任意で決めたい場合、その日付を入力する。                                                                                                         | 
| SDDocumentReason                                                       | 受注理由。固定値""(ブランク)を指定する。                                         | 
| PriceDetnExchangeRate                                                  | 価格決定のための為替レート。必要な場合、為替レートを入力する。                                                                                                                                                 | 
| RequestedDeliveryDate                                                  | 希望納入日付を入力する。                                                                                                                                                                                       | 
| ShippingCondition                                                      | 出荷条件。ビジネスパートナーマスタの販売エリアの出荷条件データから固定値が連携される。入力不要。                                                                                                               | 
| CompleteDeliveryIsDefined                                              | 出荷完了ステータス。以下のステータスが連携される。true：出荷完了、false:出荷未完了。                                     | 
| ShippingType                                                           | 出荷タイプ。固定値"01"を指定する。                                                 | 
| HeaderBillingBlockReason                                               | 請求ブロック理由。請求をブロックしたい場合は、”B”を入力する。                                                                                                                                                | 
| DeliveryBlockReason                                                    | 出荷ブロック理由。出荷をブロックしたい場合は、”B”を入力する。                                                                                                                                                | 
| IncotermsClassification                                                | インコタームズ。貿易取引を行う際、輸送タイプに応じてインコタームズを指定する。ビジネスパートナーマスタの販売エリアのインコタームズが提案される。変更の必要があれば、インコタームズマスタから選択して指定する。 | 
| CustomerPriceGroup                                                     | 顧客価格グループ。固定値"01"を指定する。                                                 | 
| PriceListType                                                          | 価格リストタイプ。                                                                                                                                                                                             | 
| CustomerPaymentTerms                                                   | 顧客支払条件。ビジネスパートナーマスタの販売エリアの顧客支払条件が提案される。変更の必要があれば、支払条件マスタから選択して指定する。                                                                         | 
| PaymentMethod                                                          | 支払方法。支払条件マスタから提案される。変更不可。                                                                                                                                                             | 
| ReferenceSDDocument                                                    | 参照受注伝票。見積参照受注の場合は見積番号、引合参照受注の場合は引合番号を入力する。 | 
| CustomerAccountAssignmentGroup                                         | 得意先勘定設定グループ。ビジネスパートナーマスタの販売エリアの勘定設定グループが提案される。変更不要。                                                                                                         | 
| ReferenceSDDocumentCategory                                            | 参照受注伝票カテゴリ                                                                                                                                                                                           | 
| AccountingExchangeRate                                                 | 会計計上のための為替レート。必要な場合、為替レートを入力する。                                                                                                                                                 | 
| CustomerGroup                                                          | 産業コードを表す得意先グループ。ビジネスパートナーマスタの得意先グループが提案される。変更不可。                                                                                                               | 
| AdditionalCustomerGroup1                                               | 追加得意先グループ1。ビジネスパートナーマスタの得意先グループ1が提案される。変更不可。                                                                                                                         | 
| AdditionalCustomerGroup2                                               | 追加得意先グループ2。ビジネスパートナーマスタの得意先グループ2が提案される。変更不可。                                                                                                                         | 
| AdditionalCustomerGroup3                                               | 追加得意先グループ3。ビジネスパートナーマスタの得意先グループ3が提案される。変更不可。                                                                                                                         | 
| AdditionalCustomerGroup4                                               | 追加得意先グループ4。ビジネスパートナーマスタの得意先グループ4が提案される。変更不可。                                                                                                                         | 
| AdditionalCustomerGroup5                                               | 追加得意先グループ5。ビジネスパートナーマスタの得意先グループ5が提案される。変更不可。                                                                                                                         | 
| CustomerTaxClassification1                                             | 税分類。以下の値から選択して入力する。0:非課税、1:納税義務                                                             | 
| TotalCreditCheckStatus                                                 | 与信確認状況。項目としては存在するが不使用。                                                                                                                                                                   | 
| BillingDocumentDate                                                    | 請求書日付。支払条件の値により自動提案される。通常、請求書日付と同じ日付になる。任意の日付の場合は、その日付を入力する。                                                                                       | 
|                                                                        | 

## ＜Request Body JSONレイアウト＞
```
{
	"connection_key": "response",
	"result": true,
	"redis_key": "abcdefg",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"SalesOrder": {
		"SalesOrder": "",
		"SalesOrderType": "OR1",
		"SalesOrganization": "0001",
		"DistributionChannel": "01",
		"OrganizationDivision": "01",
		"SalesGroup": "",
		"SalesOffice": "",
		"SalesDistrict": "",
		"SoldToParty": "1",
		"CreationDate": null,
		"LastChangeDate": null,
		"ExternalDocumentID": "",
		"LastChangeDateTime": null,
		"PurchaseOrderByCustomer": "Test",
		"CustomerPurchaseOrderDate": "2022-09-15",
		"SalesOrderDate": null,
		"TotalNetAmount": null,
		"OverallDeliveryStatus": null,
		"TotalBlockStatus": "",
		"OverallOrdReltdBillgStatus": "",
		"OverallSDDocReferenceStatus": "",
		"TransactionCurrency": "",
		"SDDocumentReason": "",
		"PricingDate": null,
		"PriceDetnExchangeRate": null,
		"RequestedDeliveryDate": null,
		"ShippingCondition": "",
		"CompleteDeliveryIsDefined": false,
		"ShippingType": "",
		"HeaderBillingBlockReason": "",
		"DeliveryBlockReason": "",
		"IncotermsClassification": "",
		"CustomerPriceGroup": "",
		"PriceListType": "",
		"CustomerPaymentTerms": "",
		"PaymentMethod": "",
		"ReferenceSDDocument": "",
		"CustomerAccountAssignmentGroup": "",
		"ReferenceSDDocumentCategory": "",
		"AccountingExchangeRate": null,
		"CustomerGroup": "",
		"AdditionalCustomerGroup1": "",
		"AdditionalCustomerGroup2": "",
		"AdditionalCustomerGroup3": "",
		"AdditionalCustomerGroup4": "",
		"AdditionalCustomerGroup5": "",
		"CustomerTaxClassification1": "",
		"TotalCreditCheckStatus": "",
		"BillingDocumentDate": null,
		"HeaderPartner": {
			"PartnerFunction": "",
			"Customer": "",
			"Supplier": ""
		},
		"SalesOrderItem": [
			{
				"SalesOrderItem": "10",
				"SalesOrderItemCategory": null,
				"SalesOrderItemText": null,
				"PurchaseOrderByCustomer": null,
				"Material": "21",
				"MaterialByCustomer": null,
				"PricingDate": null,
				"BillingPlan": null,
				"RequestedQuantity": "1",
				"RequestedQuantityUnit": null,
				"OrderQuantityUnit": null,
				"ConfdDelivQtyInOrderQtyUnit": null,
				"ItemGrossWeight": "1",
				"ItemNetWeight": "1",
				"ItemWeightUnit": null,
				"ItemVolume": "1",
				"ItemVolumeUnit": null,
				"TransactionCurrency": null,
				"NetAmount": null,
				"MaterialGroup": null,
				"MaterialPricingGroup": null,
				"BillingDocumentDate": null,
				"Batch": null,
				"ProductionPlant": null,
				"StorageLocation": null,
				"DeliveryGroup": null,
				"ShippingPoint": null,
				"ShippingType": null,
				"DeliveryPriority": null,
				"IncotermsClassification": "",
				"TaxAmount": null,
				"ProductTaxClassification1": null,
				"MatlAccountAssignmentGroup": null,
				"CostAmount": null,
				"CustomerPaymentTerms": null,
				"CustomerGroup": null,
				"SalesDocumentRjcnReason": null,
				"ItemBillingBlockReason": null,
				"WBSElement": null,
				"ProfitCenter": null,
				"AccountingExchangeRate": null,
				"ReferenceSDDocument": null,
				"ReferenceSDDocumentItem": null,
				"SDProcessStatus": null,
				"DeliveryStatus": null,
				"OrderRelatedBillingStatus": null,
			"ItemPartner": {
				"PartnerFunction": "",
				"Customer": "",
				"Supplier": "",
				"Personnel": "",
				"ContactPerson": ""
				},
			"ItemPricingElement": {
				"PricingProcedureStep": "",
				"PricingProcedureCounter": "",
				"ConditionType": "",
				"PriceConditionDeterminationDte": "",
				"ConditionCalculationType": "",
				"ConditionBaseValue": "",
				"ConditionRateValue": "",
				"ConditionCurrency": "",
				"ConditionQuantity": "",
				"ConditionQuantityUnit": "",
				"ConditionCategory": "",
				"PricingScaleType": "",
				"ConditionRecord": "",
				"ConditionSequentialNumber": "",
				"TaxCode": "",
				"ConditionAmount": "",
				"TransactionCurrency": "",
				"PricingScaleBasis": "",
				"ConditionScaleBasisValue": "",
				"ConditionScaleBasisUnit": "",
				"ConditionScaleBasisCurrency": "",
				"ConditionIsManuallyChanged": false
			},
			"ItemScheduleLine": {
				"ScheduleLine": "",
				"RequestedDeliveryDate": "",
				"ConfirmedDeliveryDate": "",
				"OrderQuantityUnit": "",
				"ScheduleLineOrderQuantity": "",
				"ConfdOrderQtyByMatlAvailCheck": "",
				"DeliveredQtyInOrderQtyUnit": "",
				"OpenConfdDelivQtyInOrdQtyUnit": "",
				"CorrectedQtyInOrderQtyUnit": "",
				"DelivBlockReasonForSchedLine": ""
			}
		},
		{
			"SalesOrderItem": "20",
			"SalesOrderItemCategory": null,
			"SalesOrderItemText": null,
			"PurchaseOrderByCustomer": null,
			"Material": "21",
			"MaterialByCustomer": null,
			"PricingDate": null,
			"BillingPlan": null,
			"RequestedQuantity": "1",
			"RequestedQuantityUnit": null,
			"OrderQuantityUnit": null,
			"ConfdDelivQtyInOrderQtyUnit": null,
			"ItemGrossWeight": "1",
			"ItemNetWeight": "1",
			"ItemWeightUnit": null,
			"ItemVolume": "1",
			"ItemVolumeUnit": null,
			"TransactionCurrency": null,
			"NetAmount": null,
			"MaterialGroup": null,
			"MaterialPricingGroup": null,
			"BillingDocumentDate": null,
			"Batch": null,
			"ProductionPlant": null,
			"StorageLocation": null,
			"DeliveryGroup": null,
			"ShippingPoint": null,
			"ShippingType": null,
			"DeliveryPriority": null,
			"IncotermsClassification": "",
			"TaxAmount": null,
			"ProductTaxClassification1": null,
			"MatlAccountAssignmentGroup": null,
			"CostAmount": null,
			"CustomerPaymentTerms": null,
			"CustomerGroup": null,
			"SalesDocumentRjcnReason": null,
			"ItemBillingBlockReason": null,
			"WBSElement": null,
			"ProfitCenter": null,
			"AccountingExchangeRate": null,
			"ReferenceSDDocument": null,
			"ReferenceSDDocumentItem": null,
			"SDProcessStatus": null,
			"DeliveryStatus": null,
			"OrderRelatedBillingStatus": null,
			"ItemPartner": {
				"PartnerFunction": "",
				"Customer": "",
				"Supplier": "",
				"Personnel": "",
				"ContactPerson": ""
				},
			"ItemPricingElement": {
				"PricingProcedureStep": "",
				"PricingProcedureCounter": "",
				"ConditionType": "",
				"PriceConditionDeterminationDte": "",
				"ConditionCalculationType": "",
				"ConditionBaseValue": "",
				"ConditionRateValue": "",
				"ConditionCurrency": "",
				"ConditionQuantity": "",
				"ConditionQuantityUnit": "",
				"ConditionCategory": "",
				"PricingScaleType": "",
				"ConditionRecord": "",
				"ConditionSequentialNumber": "",
				"TaxCode": "",
				"ConditionAmount": "",
				"TransactionCurrency": "",
				"PricingScaleBasis": "",
				"ConditionScaleBasisValue": "",
				"ConditionScaleBasisUnit": "",
				"ConditionScaleBasisCurrency": "",
				"ConditionIsManuallyChanged": false
			},
			"ItemScheduleLine": {
				"ScheduleLine": "",
				"RequestedDeliveryDate": "",
				"ConfirmedDeliveryDate": "",
				"OrderQuantityUnit": "",
				"ScheduleLineOrderQuantity": "",
				"ConfdOrderQtyByMatlAvailCheck": "",
				"DeliveredQtyInOrderQtyUnit": "",
				"OpenConfdDelivQtyInOrdQtyUnit": "",
				"CorrectedQtyInOrderQtyUnit": "",
				"DelivBlockReasonForSchedLine": ""
			}
		}
	]
	},
	"api_schema": "SAPSalesOrderReads",
	"accepter": ["HeaderItem"],
	"sales_order": "",
	"deleted": false
}
```
