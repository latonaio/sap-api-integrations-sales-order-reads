package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			SalesOrder:                     data.SalesOrder,
			SalesOrderType:                 data.SalesOrderType,
			SalesOrganization:              data.SalesOrganization,
			DistributionChannel:            data.DistributionChannel,
			OrganizationDivision:           data.OrganizationDivision,
			SalesGroup:                     data.SalesGroup,
			SalesOffice:                    data.SalesOffice,
			SalesDistrict:                  data.SalesDistrict,
			SoldToParty:                    data.SoldToParty,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			ExternalDocumentID:             data.ExternalDocumentID,
			LastChangeDateTime:             data.LastChangeDateTime,
			PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
			CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
			SalesOrderDate:                 data.SalesOrderDate,
			TotalNetAmount:                 data.TotalNetAmount,
			OverallDeliveryStatus:          data.OverallDeliveryStatus,
			TotalBlockStatus:               data.TotalBlockStatus,
			OverallOrdReltdBillgStatus:     data.OverallOrdReltdBillgStatus,
			OverallSDDocReferenceStatus:    data.OverallSDDocReferenceStatus,
			TransactionCurrency:            data.TransactionCurrency,
			SDDocumentReason:               data.SDDocumentReason,
			PricingDate:                    data.PricingDate,
			PriceDetnExchangeRate:          data.PriceDetnExchangeRate,
			RequestedDeliveryDate:          data.RequestedDeliveryDate,
			ShippingCondition:              data.ShippingCondition,
			CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
			ShippingType:                   data.ShippingType,
			HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
			DeliveryBlockReason:            data.DeliveryBlockReason,
			IncotermsClassification:        data.IncotermsClassification,
			CustomerPriceGroup:             data.CustomerPriceGroup,
			PriceListType:                  data.PriceListType,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			PaymentMethod:                  data.PaymentMethod,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
			CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
			AccountingExchangeRate:         data.AccountingExchangeRate,
			CustomerGroup:                  data.CustomerGroup,
			AdditionalCustomerGroup1:       data.AdditionalCustomerGroup1,
			AdditionalCustomerGroup2:       data.AdditionalCustomerGroup2,
			AdditionalCustomerGroup3:       data.AdditionalCustomerGroup3,
			AdditionalCustomerGroup4:       data.AdditionalCustomerGroup4,
			AdditionalCustomerGroup5:       data.AdditionalCustomerGroup5,
			CustomerTaxClassification1:     data.CustomerTaxClassification1,
			TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
			BillingDocumentDate:            data.BillingDocumentDate,
			ToHeaderPartner:                data.ToHeaderPartner.Deferred.URI,
			ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			SalesOrder:                  data.SalesOrder,
			SalesOrderItem:              data.SalesOrderItem,
			SalesOrderItemCategory:      data.SalesOrderItemCategory,
			SalesOrderItemText:          data.SalesOrderItemText,
			PurchaseOrderByCustomer:     data.PurchaseOrderByCustomer,
			Material:                    data.Material,
			MaterialByCustomer:          data.MaterialByCustomer,
			PricingDate:                 data.PricingDate,
			BillingPlan:                 data.BillingPlan,
			RequestedQuantity:           data.RequestedQuantity,
			RequestedQuantityUnit:       data.RequestedQuantityUnit,
			OrderQuantityUnit:           data.OrderQuantityUnit,
			ConfdDelivQtyInOrderQtyUnit: data.ConfdDelivQtyInOrderQtyUnit,
			ItemGrossWeight:             data.ItemGrossWeight,
			ItemNetWeight:               data.ItemNetWeight,
			ItemWeightUnit:              data.ItemWeightUnit,
			ItemVolume:                  data.ItemVolume,
			ItemVolumeUnit:              data.ItemVolumeUnit,
			TransactionCurrency:         data.TransactionCurrency,
			NetAmount:                   data.NetAmount,
			MaterialGroup:               data.MaterialGroup,
			MaterialPricingGroup:        data.MaterialPricingGroup,
			BillingDocumentDate:         data.BillingDocumentDate,
			Batch:                       data.Batch,
			ProductionPlant:             data.ProductionPlant,
			StorageLocation:             data.StorageLocation,
			DeliveryGroup:               data.DeliveryGroup,
			ShippingPoint:               data.ShippingPoint,
			ShippingType:                data.ShippingType,
			DeliveryPriority:            data.DeliveryPriority,
			IncotermsClassification:     data.IncotermsClassification,
			TaxAmount:                   data.TaxAmount,
			ProductTaxClassification1:   data.ProductTaxClassification1,
			MatlAccountAssignmentGroup:  data.MatlAccountAssignmentGroup,
			CostAmount:                  data.CostAmount,
			CustomerPaymentTerms:        data.CustomerPaymentTerms,
			CustomerGroup:               data.CustomerGroup,
			SalesDocumentRjcnReason:     data.SalesDocumentRjcnReason,
			ItemBillingBlockReason:      data.ItemBillingBlockReason,
			WBSElement:                  data.WBSElement,
			ProfitCenter:                data.ProfitCenter,
			AccountingExchangeRate:      data.AccountingExchangeRate,
			ReferenceSDDocument:         data.ReferenceSDDocument,
			ReferenceSDDocumentItem:     data.ReferenceSDDocumentItem,
			SDProcessStatus:             data.SDProcessStatus,
			DeliveryStatus:              data.DeliveryStatus,
			OrderRelatedBillingStatus:   data.OrderRelatedBillingStatus,
			ToItemPricingElement:        data.ToItemPricingElement.Deferred.URI,
			ToItemScheduleLine:          data.ToItemScheduleLine.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
			SalesOrder:      data.SalesOrder,
			PartnerFunction: data.PartnerFunction,
			Customer:        data.Customer,
			Supplier:        data.Supplier,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			SalesOrder:                  data.SalesOrder,
			SalesOrderItem:              data.SalesOrderItem,
			SalesOrderItemCategory:      data.SalesOrderItemCategory,
			SalesOrderItemText:          data.SalesOrderItemText,
			PurchaseOrderByCustomer:     data.PurchaseOrderByCustomer,
			Material:                    data.Material,
			MaterialByCustomer:          data.MaterialByCustomer,
			PricingDate:                 data.PricingDate,
			BillingPlan:                 data.BillingPlan,
			RequestedQuantity:           data.RequestedQuantity,
			RequestedQuantityUnit:       data.RequestedQuantityUnit,
			OrderQuantityUnit:           data.OrderQuantityUnit,
			ConfdDelivQtyInOrderQtyUnit: data.ConfdDelivQtyInOrderQtyUnit,
			ItemGrossWeight:             data.ItemGrossWeight,
			ItemNetWeight:               data.ItemNetWeight,
			ItemWeightUnit:              data.ItemWeightUnit,
			ItemVolume:                  data.ItemVolume,
			ItemVolumeUnit:              data.ItemVolumeUnit,
			TransactionCurrency:         data.TransactionCurrency,
			NetAmount:                   data.NetAmount,
			MaterialGroup:               data.MaterialGroup,
			MaterialPricingGroup:        data.MaterialPricingGroup,
			BillingDocumentDate:         data.BillingDocumentDate,
			Batch:                       data.Batch,
			ProductionPlant:             data.ProductionPlant,
			StorageLocation:             data.StorageLocation,
			DeliveryGroup:               data.DeliveryGroup,
			ShippingPoint:               data.ShippingPoint,
			ShippingType:                data.ShippingType,
			DeliveryPriority:            data.DeliveryPriority,
			IncotermsClassification:     data.IncotermsClassification,
			TaxAmount:                   data.TaxAmount,
			ProductTaxClassification1:   data.ProductTaxClassification1,
			MatlAccountAssignmentGroup:  data.MatlAccountAssignmentGroup,
			CostAmount:                  data.CostAmount,
			CustomerPaymentTerms:        data.CustomerPaymentTerms,
			CustomerGroup:               data.CustomerGroup,
			SalesDocumentRjcnReason:     data.SalesDocumentRjcnReason,
			ItemBillingBlockReason:      data.ItemBillingBlockReason,
			WBSElement:                  data.WBSElement,
			ProfitCenter:                data.ProfitCenter,
			AccountingExchangeRate:      data.AccountingExchangeRate,
			ReferenceSDDocument:         data.ReferenceSDDocument,
			ReferenceSDDocumentItem:     data.ReferenceSDDocumentItem,
			SDProcessStatus:             data.SDProcessStatus,
			DeliveryStatus:              data.DeliveryStatus,
			OrderRelatedBillingStatus:   data.OrderRelatedBillingStatus,
			ToItemPricingElement:        data.ToItemPricingElement.Deferred.URI,
			ToItemScheduleLine:          data.ToItemScheduleLine.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			PricingProcedureStep:           data.PricingProcedureStep,
			PricingProcedureCounter:        data.PricingProcedureCounter,
			ConditionType:                  data.ConditionType,
			PriceConditionDeterminationDte: data.PriceConditionDeterminationDte,
			ConditionCalculationType:       data.ConditionCalculationType,
			ConditionBaseValue:             data.ConditionBaseValue,
			ConditionRateValue:             data.ConditionRateValue,
			ConditionCurrency:              data.ConditionCurrency,
			ConditionQuantity:              data.ConditionQuantity,
			ConditionQuantityUnit:          data.ConditionQuantityUnit,
			ConditionCategory:              data.ConditionCategory,
			PricingScaleType:               data.PricingScaleType,
			ConditionRecord:                data.ConditionRecord,
			ConditionSequentialNumber:      data.ConditionSequentialNumber,
			TaxCode:                        data.TaxCode,
			ConditionAmount:                data.ConditionAmount,
			TransactionCurrency:            data.TransactionCurrency,
			PricingScaleBasis:              data.PricingScaleBasis,
			ConditionScaleBasisValue:       data.ConditionScaleBasisValue,
			ConditionScaleBasisUnit:        data.ConditionScaleBasisUnit,
			ConditionScaleBasisCurrency:    data.ConditionScaleBasisCurrency,
			ConditionIsManuallyChanged:     data.ConditionIsManuallyChanged})
	}

	return toItemPricingElement, nil
}

func ConvertToToItemScheduleLine(raw []byte, l *logger.Logger) ([]ToItemScheduleLine, error) {
	pm := &responses.ToItemScheduleLine{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toItemScheduleLine := make([]ToItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemScheduleLine = append(toItemScheduleLine, ToItemScheduleLine{
			SalesOrder:                    data.SalesOrder,
			SalesOrderItem:                data.SalesOrderItem,
			ScheduleLine:                  data.ScheduleLine,
			RequestedDeliveryDate:         data.RequestedDeliveryDate,
			ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
			OrderQuantityUnit:             data.OrderQuantityUnit,
			ScheduleLineOrderQuantity:     data.ScheduleLineOrderQuantity,
			ConfdOrderQtyByMatlAvailCheck: data.ConfdOrderQtyByMatlAvailCheck,
			DeliveredQtyInOrderQtyUnit:    data.DeliveredQtyInOrderQtyUnit,
			OpenConfdDelivQtyInOrdQtyUnit: data.OpenConfdDelivQtyInOrdQtyUnit,
			CorrectedQtyInOrderQtyUnit:    data.CorrectedQtyInOrderQtyUnit,
			DelivBlockReasonForSchedLine:  data.DelivBlockReasonForSchedLine,
		})
	}

	return toItemScheduleLine, nil
}
