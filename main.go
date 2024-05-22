package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func String(s string) *string {
	return &s
}

// Структуры данных для XML
type ExpressCargoDeclaration struct {
	XMLName                          xml.Name                        `xml:"ExpressCargoDeclaration"`
	Xmlns                            string                          `xml:"xmlns,attr"`
	XmlnsNs2                         string                          `xml:"xmlns:ns2,attr"`
	XmlnsNs3                         string                          `xml:"xmlns:ns3,attr"`
	XmlnsNs4                         string                          `xml:"xmlns:ns4,attr"`
	XmlnsNs5                         string                          `xml:"xmlns:ns5,attr"`
	ExpressRegistryKindCode          string                          `xml:"ns3:ExpressRegistryKindCode"`
	DeclarationKindCode              string                          `xml:"ns3:DeclarationKindCode,omitempty"`
	EDocIndicatorCode                string                          `xml:"ns3:EDocIndicatorCode"`
	DeclarationFeatureCode           string                          `xml:"ns3:DeclarationFeatureCode,omitempty"`
	RegisterDocumentIdDetails        RegisterDocumentIdDetail        `xml:"ns4:RegisterDocumentIdDetails"`
	ExpressCargoDeclarationIdDetails ExpressCargoDeclarationIdDetail `xml:"ns4:ExpressCargoDeclarationIdDetails"`
	ECGoodsShipmentDetails           ECGoodsShipmentDetail           `xml:"ns4:ECGoodsShipmentDetails,omitempty"`
	SignatoryPersonV2Details         SignatoryPersonV2Detail         `xml:"ns4:SignatoryPersonV2Details"`
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
	ConsignmentItemOrdinal  string             `xml:"ns3:ConsignmentItemOrdinal"`       //14.3.6.1
	HMConsignmentItemNumber string             `xml:"ns3:HMConsignmentItemNumber"`      //14.3.6.7
	CommodityCode           string             `xml:"ns2:CommodityCode"`                //14.3.6.2
	UnifiedGrossMassMeasure string             `xml:"ns2:UnifiedGrossMassMeasure"`      //14.3.6.4
	GoodsMeasureDetails     GoodsMeasureDetail `xml:"ns4:GoodsMeasureDetails"`          //14.3.6.6
	GoodsDescriptionText    string             `xml:"ns3:GoodsDescriptionText"`         //14.3.6.3
	CAValueAmount           []CurrencyValue    `xml:"ns3:CAValueAmount"`                //14.3.6.6.11
	CustomsValueAmount      *CurrencyValue     `xml:"ns3:CustomsValueAmount,omitempty"` //14.3.6.12(11)
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

func downloadFile(url, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	isRecreate := flag.Bool("r", false, "recreate")

	flag.Parse()

	// Путь к файлу XLSX
	filePath := "samplefile.xlsx"
	var productOrder int = 0
	var declarationCost float32 = 0
	var declarationWeight float32 = 0
	ecd := ExpressCargoDeclaration{}
	ecd.Xmlns = "http://www.codecraft.kz/keden/ecd"
	ecd.XmlnsNs2 = "urn:EEC:M:SimpleDataObjects:v0.4.14"
	ecd.XmlnsNs3 = "urn:EEC:M:CA:SimpleDataObjects:v1.8.1"
	ecd.XmlnsNs4 = "urn:EEC:M:CA:ComplexDataObjects:v1.8.1"
	ecd.XmlnsNs5 = "urn:EEC:M:ComplexDataObjects:v0.4.14"
	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Файл отсутствует, начинаем загрузку...")
		if err := downloadFile("https://raw.githubusercontent.com/Gessar/express_xlsx_xml/main/samplefile.xlsx", filePath); err != nil {
			fmt.Println("Ошибка загрузки файла:", err)
			return
		}
		fmt.Println("Файл успешно загружен.")
	}

	// Открытие файла
	declaration_xlsx_file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer declaration_xlsx_file.Close()

	if *isRecreate {
		for i := 0; i <= 500; i++ {
			s := "Накладная " + fmt.Sprint(i)
			index, _ := declaration_xlsx_file.NewSheet(s)
			_ = declaration_xlsx_file.CopySheet(2, index)
		}

		declaration_xlsx_file.SetActiveSheet(0)
		// Сохраняем изменения обратно в файл XLSX
		if err := declaration_xlsx_file.SaveAs(filePath); err != nil {
			fmt.Println("Ошибка при сохранении файла:", err)
			return
		}
		fmt.Println("Листы успешно добавлены в файл samplefile.xlsx!")
		os.Exit(0)
	}

	for i := 0; i < 500; i++ {
		fmt.Println("Поиск накладной", i+1)
		cell, err := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "A28")
		if err != nil || cell == "" {
			fmt.Println("Накладных больше нет")
			break
		}
		fmt.Println("Работа с накладной", i+1)
		newHS := ECHouseShipmentDetail{}
		tdDocId, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B2")
		tdDocCreationDate, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B3")
		newHS.TransportDocumentDetails = TransportDocumentDetail{DocId: tdDocId, DocCreationDate: tdDocCreationDate}
		for j := 0; j < 999; j++ {
			fmt.Println("Поиск товара", j+1)
			jcell, err := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B"+strconv.FormatInt(int64(j+28), 10))
			if err != nil || jcell == "" {
				fmt.Println("Товаров больше нет")
				break
			}
			productOrder++
			fmt.Println("Работа с товаром", j+1)
			newGood := ECGoodsItemDetail{}
			commodityCode, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B"+strconv.FormatInt(int64(j+28), 10))
			newGood.CommodityCode = commodityCode
			newGood.HMConsignmentItemNumber = strconv.FormatInt(int64(j+1), 10)
			productDescription, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "C"+strconv.FormatInt(int64(j+28), 10))
			newGood.GoodsDescriptionText = productDescription
			grossWeight, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "D"+strconv.FormatInt(int64(j+28), 10))
			newGood.UnifiedGrossMassMeasure = grossWeight
			productCount, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "F"+strconv.FormatInt(int64(j+28), 10))
			newGood.GoodsMeasureDetails.GoodsMeasure.Value = productCount
			newGood.GoodsMeasureDetails.GoodsMeasure.MeasurementUnitCode = "796"
			newGood.GoodsMeasureDetails.GoodsMeasure.MeasurementUnitCodeListId = "2016"
			// newGood.CAValueAmount.Value, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "G"+strconv.FormatInt(int64(j+28), 10))
			fcostcurrency, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H"+strconv.FormatInt(int64(j+28), 10))
			fcostvalue, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "G"+strconv.FormatInt(int64(j+28), 10))
			// ccli := "2022"
			// kzttext := "KZT"
			fcost := CurrencyValue{Value: &fcostvalue, CurrencyCode: String(fcostcurrency), CurrencyCodeListId: String("2022")}
			kzcostvalue, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "I"+strconv.FormatInt(int64(j+28), 10))
			kzcost := CurrencyValue{Value: &kzcostvalue, CurrencyCode: String("KZT"), CurrencyCodeListId: String("2022")}
			newGood.CAValueAmount = append(newGood.CAValueAmount, fcost)
			newGood.CAValueAmount = append(newGood.CAValueAmount, kzcost)
			// newGood.CAValueAmount.CurrencyCode, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H"+strconv.FormatInt(int64(j+28), 10))
			newGood.ConsignmentItemOrdinal = strconv.FormatInt(int64(productOrder), 10)

			newHS.ECGoodsItemDetails = append(newHS.ECGoodsItemDetails, newGood)
		}
		newHS.ObjectOrdinal = strconv.FormatInt(int64(i+1), 10)
		cost, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "I26")
		weight, _ := declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "D26")
		costfloat, err := strconv.ParseFloat(cost, 32)
		weightfloat, err := strconv.ParseFloat(weight, 32)
		declarationCost += float32(costfloat)
		declarationWeight += float32(weightfloat)
		fmt.Println("cost is ", cost)
		newHS.CAValueAmount.Value = &cost
		// kzttext := "KZT"
		newHS.CAValueAmount.CurrencyCode = String("KZT") //&kzttext
		// ccli := "2022"
		newHS.CAValueAmount.CurrencyCodeListId = String("2022") //&ccli
		newHS.UnifiedGrossMassMeasure = strconv.FormatFloat(float64(declarationWeight), 'f', -1, 32)
		newHS.TransportDocumentDetails.DocId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "C2")
		newHS.TransportDocumentDetails.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "C3")
		newHS.HouseWaybillDetails.DocId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "C5")
		newHS.HouseWaybillDetails.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "C6")

		newHS.ConsignorDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B8")
		newHS.ConsignorDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B9")
		newHS.ConsignorDetails.TaxpayerId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B10")
		newHS.ConsignorDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B12")
		newHS.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B13")
		newHS.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
		newHS.ConsignorDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B14")
		newHS.ConsignorDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B15")
		newHS.ConsignorDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B16")
		newHS.ConsignorDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B17")
		newHS.ConsignorDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B18")
		newHS.ConsignorDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B19")
		newHS.ConsignorDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B20")
		newHS.ConsignorDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B22")
		newHS.ConsignorDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B23")
		newHS.ConsignorDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "B24")

		newHS.ConsigneeDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E8")
		newHS.ConsigneeDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E9")
		newHS.ConsigneeDetails.PersonId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E10")
		newHS.ConsigneeDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E12")
		newHS.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E13")
		newHS.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
		newHS.ConsigneeDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E14")
		newHS.ConsigneeDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E15")
		newHS.ConsigneeDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E16")
		newHS.ConsigneeDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E17")
		newHS.ConsigneeDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E18")
		newHS.ConsigneeDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E19")
		newHS.ConsigneeDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E20")
		newHS.ConsigneeDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E22")
		newHS.ConsigneeDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E23")
		newHS.ConsigneeDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "E24")
		newHS.ConsigneeDetails.IdentityDocV3Details.UnifiedCountryCode.CodeListId = "2021"
		newHS.ConsigneeDetails.IdentityDocV3Details.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H9")
		newHS.ConsigneeDetails.IdentityDocV3Details.IdentityDocKindCode.CodeListId = "2053"
		newHS.ConsigneeDetails.IdentityDocV3Details.IdentityDocKindCode.Value, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H10")
		newHS.ConsigneeDetails.IdentityDocV3Details.DocKindName, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H11")
		newHS.ConsigneeDetails.IdentityDocV3Details.DocSeriesId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H12")
		newHS.ConsigneeDetails.IdentityDocV3Details.DocId, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H13")
		newHS.ConsigneeDetails.IdentityDocV3Details.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Накладная "+strconv.FormatInt(int64(i+1), 10), "H14")

		ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails = append(ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails, newHS)
	}
	ecd.ExpressRegistryKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "C1")
	ecd.DeclarationKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B1")
	ecd.ExpressCargoDeclarationIdDetails.CustomsOfficeCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B3")
	ecd.DeclarationFeatureCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B5")
	ecd.RegisterDocumentIdDetails.RegistrationNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B10")
	ecd.RegisterDocumentIdDetails.DocKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B8")
	//declarationCost = float32(math.Round((float64(declarationCost * 100))) / 100)
	ecd.ECGoodsShipmentDetails.CAValueAmount.Value = String(strconv.FormatFloat(float64(declarationCost), 'f', -1, 32))
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCode = String("KZT")
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCodeListId = String("2022")
	ecd.ECGoodsShipmentDetails.UnifiedGrossMassMeasure = strconv.FormatFloat(float64(declarationWeight), 'f', -1, 32)
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.FirstName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B13")
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.LastName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B15")
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.MiddleName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B14")
	ecd.SignatoryPersonV2Details.SigningDetails.PositionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B16")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B17")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B18")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B19")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B22")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B23")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.CodeListId = "2021"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.CodeListId = "2053"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocKindName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B24")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocSeriesId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B25")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B26")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B27")
	ecd.EDocIndicatorCode = "ЭД" //, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B24")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.CodeListId = "2009"
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B29")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B30")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B31")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocStartDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B32")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocValidityDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B33")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B35")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B36")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.TaxpayerId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B37")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B39")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B40")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B41")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B42")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B43")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B44")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B45")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B46")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B47")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B49")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B50")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B51")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E35")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E36")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.TaxpayerId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E37")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E39")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E40")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E41")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E42")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E43")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E44")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E45")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E46")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E47")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E49")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E50")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E51")

	// xmlData, err := xml.MarshalIndent(expressCargoDeclaration, "", "  ")
	// if err != nil {
	// 	fmt.Println("Ошибка при сериализации в XML:", err)
	// 	return
	// }
	// outputstring := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n" + string(xmlData)

	// outputFile := "output.xml"
	// // copy(xmlData[:], "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n")

	// err = os.WriteFile(outputFile, []byte(outputstring), 0644)
	// if err != nil {
	// 	fmt.Println("Ошибка при записи XML файла:", err)
	// 	return
	// }

	// fmt.Printf("XML файл успешно создан: %s\n", outputFile)

	xmlData1, err := xml.MarshalIndent(ecd, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации в XML:", err)
		return
	}
	outputstring1 := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n" + string(xmlData1)

	outputFile1 := "output1.xml"
	// copy(xmlData[:], "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n")

	err = os.WriteFile(outputFile1, []byte(outputstring1), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи XML файла:", err)
		return
	}

	// fmt.Printf("XML файл успешно создан: %s\n", outputFile)
}
