package main

import (
	sap_api_caller "sap-api-integrations-sales-order-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-sales-order-reads/SAP_API_Input_Reader"
	"sap-api-integrations-sales-order-reads/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_get_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	gc := sap_api_get_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		gc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Sales_Order_Item_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)
	accepter := inputSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}

	caller.AsyncGetSalesOrder(
		inputSDC.SalesOrder.SalesOrder,
		inputSDC.SalesOrder.SalesOrderItem.SalesOrderItem,
		accepter,
	)
}
