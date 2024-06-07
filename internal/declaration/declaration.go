package declaration

import "encoding/xml"

type ExpressCargoDeclaration struct {
	XMLName                          xml.Name                          `xml:"ExpressCargoDeclaration"`
	Xmlns                            string                            `xml:"xmlns,attr"`
	XmlnsNs2                         string                            `xml:"xmlns:ns2,attr"`
	XmlnsNs3                         string                            `xml:"xmlns:ns3,attr"`
	XmlnsNs4                         string                            `xml:"xmlns:ns4,attr"`
	XmlnsNs5                         string                            `xml:"xmlns:ns5,attr"`
	ExpressRegistryKindCode          string                            `xml:"ns3:ExpressRegistryKindCode"`
	DeclarationKindCode              string                            `xml:"ns3:DeclarationKindCode,omitempty"`
	CustomsProcedureCode             *CustomsProcedureCode             `xml:"ns3:CustomsProcedureCode,omitempty"`
	PreviousCustomsProcedureModeCode *PreviousCustomsProcedureModeCode `xml:"ns3:PreviousCustomsProcedureModeCode,omitempty"`
	EDocIndicatorCode                string                            `xml:"ns3:EDocIndicatorCode"`
	DeclarationFeatureCode           string                            `xml:"ns3:DeclarationFeatureCode,omitempty"`
	RegisterDocumentIdDetails        RegisterDocumentIdDetail          `xml:"ns4:RegisterDocumentIdDetails"`
	ExpressCargoDeclarationIdDetails ExpressCargoDeclarationIdDetail   `xml:"ns4:ExpressCargoDeclarationIdDetails"`
	ECGoodsShipmentDetails           ECGoodsShipmentDetail             `xml:"ns4:ECGoodsShipmentDetails,omitempty"`
	SignatoryPersonV2Details         SignatoryPersonV2Detail           `xml:"ns4:SignatoryPersonV2Details"`
}

type CustomsProcedureCode struct {
	CodeListId string `xml:"codeListId,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type PreviousCustomsProcedureModeCode struct {
	CodeListId string `xml:"codeListId,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type RegisterDocumentIdDetail struct { //15
	DocKindCode          string `xml:"ns2:DocKindCode,omitempty"`          //15.1
	UnifiedCountryCode   string `xml:"ns2:UnifiedCountryCode,omitempty"`   //15.2
	RegistrationNumberId string `xml:"ns3:RegistrationNumberId,omitempty"` //15.3
	ReregistrationCode   string `xml:"ns3:ReregistrationCode,omitempty"`   //15.4
}

type ExpressCargoDeclarationIdDetail struct {
	CustomsOfficeCode string `xml:"ns2:CustomsOfficeCode,omitempty"` //7.1
}

type ECGoodsShipmentDetail struct {
	ConsignorDetails        ConsignorDetail         `xml:"ns4:ConsignorDetails"`                 //14.1
	ConsigneeDetails        ConsigneeDetail         `xml:"ns4:ConsigneeDetails"`                 //14.2
	UnifiedGrossMassMeasure string                  `xml:"ns2:UnifiedGrossMassMeasure"`          //14.4
	CAValueAmount           CurrencyValue           `xml:"ns3:CAValueAmount"`                    //14.5
	ECHouseShipmentDetails  []ECHouseShipmentDetail `xml:"ns4:ECHouseShipmentDetails,omitempty"` //14.3
}

type ECHouseShipmentDetail struct {
	ObjectOrdinal            string                  `xml:"ns2:ObjectOrdinal"`                //14.3.1
	TransportDocumentDetails TransportDocumentDetail `xml:"ns4:TransportDocumentDetails"`     //14.3.2
	HouseWaybillDetails      HouseWaybillDetail      `xml:"ns4:HouseWaybillDetails"`          //14.3.3
	ConsignorDetails         ConsignorDetail         `xml:"ns4:ConsignorDetails"`             //14.3.4
	ConsigneeDetails         ConsigneeDetail         `xml:"ns4:ConsigneeDetails"`             //14.3.5
	ECGoodsItemDetails       []ECGoodsItemDetail     `xml:"ns4:ECGoodsItemDetails,omitempty"` //14.3.6
	UnifiedGrossMassMeasure  string                  `xml:"ns2:UnifiedGrossMassMeasure"`
	CAValueAmount            CurrencyValue           `xml:"ns3:CAValueAmount"`
}

type TransportDocumentDetail struct {
	DocId           string `xml:"ns2:DocId"`           //14.3.2.3
	DocCreationDate string `xml:"ns2:DocCreationDate"` //14.3.2.4
}

type HouseWaybillDetail struct {
	DocId           string `xml:"ns2:DocId"`           //14.3.3.3
	DocCreationDate string `xml:"ns2:DocCreationDate"` //14.3.3.4
}

type ConsignorDetail struct {
	SubjectName           string              `xml:"ns2:SubjectName,omitempty"`           //*.1.2
	SubjectBriefName      string              `xml:"ns2:SubjectBriefName,omitempty"`      //*.1.3
	SubjectAddressDetails AddressDetail       `xml:"ns5:SubjectAddressDetails,omitempty"` //*.1.12
	CommunicationDetails  CommunicationDetail `xml:"ns5:CommunicationDetails,omitempty"`  //*.1.13
	TaxpayerId            string              `xml:"ns2:TaxpayerId,omitempty"`            //*.1.8
}

type ConsigneeDetail struct {
	SubjectName           string              `xml:"ns2:SubjectName"`
	SubjectBriefName      string              `xml:"ns2:SubjectBriefName"`
	TaxpayerId            string              `xml:"ns2:TaxpayerId,omitempty"`
	IdentityDocV3Details  IdentityDocV3Detail `xml:"ns5:IdentityDocV3Details,omitempty"`
	SubjectAddressDetails AddressDetail       `xml:"ns5:SubjectAddressDetails"`
	CommunicationDetails  CommunicationDetail `xml:"ns5:CommunicationDetails"`
	PersonId              string              `xml:"ns3:PersonId,omitempty"`
}

type IdentityDocV3Detail struct {
	UnifiedCountryCode  UnifiedCountryCode  `xml:"ns2:UnifiedCountryCode,omitempty"`
	IdentityDocKindCode IdentityDocKindCode `xml:"ns2:IdentityDocKindCode,omitempty"`
	DocKindName         string              `xml:"ns2:DocKindName,omitempty"`
	DocSeriesId         string              `xml:"ns2:DocSeriesId,omitempty"`
	DocId               string              `xml:"ns2:DocId,omitempty"`
	DocCreationDate     string              `xml:"ns2:DocCreationDate,omitempty"`
	// AuthorityName       string              `xml:"ns2:AuthorityName"`
}

type UnifiedCountryCode struct {
	CodeListId string `xml:"codeListId,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type IdentityDocKindCode struct {
	CodeListId string `xml:"codeListId,attr"`
	Value      string `xml:",chardata"`
}

type AddressDetail struct {
	AddressKindCode    string             `xml:"ns2:AddressKindCode"`            //*.1
	UnifiedCountryCode UnifiedCountryCode `xml:"ns2:UnifiedCountryCode"`         //*.2
	RegionName         string             `xml:"ns2:RegionName,omitempty"`       //*.4
	DistrictName       string             `xml:"ns2:DistrictName,omitempty"`     //*.5
	CityName           string             `xml:"ns2:CityName,omitempty"`         //*.6
	SettlementName     string             `xml:"ns2:SettlementName,omitempty"`   //*.7
	StreetName         string             `xml:"ns2:StreetName,omitempty"`       //*.8
	BuildingNumberId   string             `xml:"ns2:BuildingNumberId,omitempty"` //*.9
	RoomNumberId       string             `xml:"ns2:RoomNumberId,omitempty"`     //*.10
}

type CommunicationDetail struct {
	CommunicationChannelCode string `xml:"ns2:CommunicationChannelCode"`
	CommunicationChannelName string `xml:"ns2:CommunicationChannelName"`
	CommunicationChannelId   string `xml:"ns2:CommunicationChannelId"`
}

type ECGoodsItemDetail struct {
	ConsignmentItemOrdinal  string                `xml:"ns3:ConsignmentItemOrdinal"`          //14.3.6.1
	HMConsignmentItemNumber string                `xml:"ns3:HMConsignmentItemNumber"`         //14.3.6.7
	CommodityCode           string                `xml:"ns2:CommodityCode"`                   //14.3.6.2
	UnifiedGrossMassMeasure string                `xml:"ns2:UnifiedGrossMassMeasure"`         //14.3.6.4
	GoodsMeasureDetails     GoodsMeasureDetail    `xml:"ns4:GoodsMeasureDetails"`             //14.3.6.6
	GoodsDescriptionText    string                `xml:"ns3:GoodsDescriptionText"`            //14.3.6.3
	CAValueAmount           []CurrencyValue       `xml:"ns3:CAValueAmount"`                   //14.3.6.6.11
	CustomsValueAmount      *CurrencyValue        `xml:"ns3:CustomsValueAmount,omitempty"`    //14.3.6.12(11)
	ECPresentedDocDetails   ECPresentedDocDetails `xml:"ns4:ECPresentedDocDetails,omitempty"` //14.3.6.14
}

type ECPresentedDocDetails struct {
	DocKindCode               DocKindCode               `xml:"ns2:DocKindCode"`                  //14.3.6.14.1
	DocId                     string                    `xml:"ns2:DocId,omitempty"`              //14.3.6.14.3
	DocCreationDate           string                    `xml:"ns2:DocCreationDate,omitempty"`    //14.3.6.14.4
	DocStartDate              string                    `xml:"ns2:DocStartDate,omitempty"`       //14.3.6.14.5
	DocValidityDate           string                    `xml:"ns2:DocValidityDate,omitempty"`    //14.3.6.14.6
	UnifiedCountryCode        UnifiedCountryCode        `xml:"ns2:UnifiedCountryCode,omitempty"` //14.3.6.14.7
	AuthorityName             string                    `xml:"ns2:AuthorityName,omitempty"`      //14.3.6.14.8
	AuthorityId               string                    `xml:"ns2:AuthorityId,omitempty"`        //14.3.6.14.9
	DocumentPresentingDetails DocumentPresentingDetails `xml:"ns4:DocumentPresentingDetails"`    //14.3.6.14.14
}

type DocumentPresentingDetails struct {
	DocPresentKindCode string `xml:"ns3:DocPresentKindCode"` //14.3.6.14.14.1
}

type GoodsMeasureDetail struct {
	GoodsMeasure GoodsMeasure `xml:"ns3:GoodsMeasure"` //14.3.6.6.1
}

type GoodsMeasure struct {
	MeasurementUnitCode       string `xml:"measurementUnitCode,attr"`
	MeasurementUnitCodeListId string `xml:"measurementUnitCodeListId,attr"`
	Value                     string `xml:",chardata"`
}

type CurrencyValue struct {
	CurrencyCode       *string `xml:"currencyCode,attr,omitempty"`
	CurrencyCodeListId *string `xml:"currencyCodeListId,attr,omitempty"`
	Value              *string `xml:",chardata"`
}

type DocKindCode struct {
	CodeListId string `xml:"codeListId,attr"`
	Value      string `xml:",chardata"`
}

type SignatoryPersonV2Detail struct {
	SigningDetails         SigningDetail         `xml:"ns4:SigningDetails"`
	IdentityDocV3Details   IdentityDocV3Detail   `xml:"ns5:IdentityDocV3Details"`
	PowerOfAttorneyDetails PowerOfAttorneyDetail `xml:"ns4:PowerOfAttorneyDetails"`
}

type SigningDetail struct {
	FullNameDetails      FullNameDetail      `xml:"ns5:FullNameDetails"`
	PositionName         string              `xml:"ns2:PositionName"`
	CommunicationDetails CommunicationDetail `xml:"ns5:CommunicationDetails"`
}

type FullNameDetail struct {
	FirstName  string `xml:"ns2:FirstName"`
	MiddleName string `xml:"ns2:MiddleName"`
	LastName   string `xml:"ns2:LastName"`
}

type PowerOfAttorneyDetail struct { //16.4
	DocKindCode     PoadDocKindCode `xml:"ns2:DocKindCode"`               //16.4.1
	DocId           string          `xml:"ns2:DocId"`                     //16.4.3
	DocCreationDate string          `xml:"ns2:DocCreationDate"`           //16.4.4
	DocStartDate    string          `xml:"ns2:DocStartDate,omitempty"`    //16.4.5
	DocValidityDate string          `xml:"ns2:DocValidityDate,omitempty"` //16.4.6
}

type PoadDocKindCode struct {
	CodeListId string `xml:"codeListId,attr"`
	Value      string `xml:",chardata"`
}

//cacdo - ns4
//casdo - ns3
//csdo - ns2
//ccdo - ns5
