package declaration

import (
	"fmt"
	"os"
	"strconv"

	// "github.com/Gessar/express_xlsx_xml/internal/declaration"
	"github.com/xuri/excelize/v2"
)

func String(s string) *string {
	return &s
}

// FormatCurrencyValue форматирует значение валюты с двумя десятичными знаками
func FormatCurrencyValue(value float64) *string {
	formatted := strconv.FormatFloat(value, 'f', 2, 64)
	return &formatted
}

func FormatGrossMassValue(value float64) *string {
	formatted := strconv.FormatFloat(value, 'f', 3, 64)
	return &formatted
}

func AddLists(file *excelize.File, path string, count int) {
	fmt.Println("start")
	for i := 2; i <= count; i++ {
		s := "Накладная " + fmt.Sprint(i)
		index, _ := file.NewSheet(s)
		_ = file.CopySheet(2, index)
	}

	file.SetActiveSheet(0)
	// Сохраняем изменения обратно в файл XLSX
	if err := file.SaveAs(path); err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return
	}
	fmt.Println("Листы успешно добавлены в файл samplefile.xlsx!")
}

func WriteToFile(path string, data string) {
	err := os.WriteFile(path, []byte(data), 0064)
	if err != nil {
		fmt.Println("Ошибка при записи XML файла:", err)
		return
	}
}

func GetCellValue(sheet string, cell string, xlsx *excelize.File) string {
	cellValue, _ := xlsx.GetCellValue(sheet, cell)
	return cellValue
}

func ReadGoodsItemDetail(itemCount int, xlsx *excelize.File) ECGoodsItemDetail {
	good := ECGoodsItemDetail{}
	good.ConsignmentItemOrdinal = GetCellValue("Накладная", "AV"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.CommodityCode = GetCellValue("Накладная", "AY"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.GoodsDescriptionText = GetCellValue("Накладная", "AX"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.UnifiedGrossMassMeasure = GetCellValue("Накладная", "BA"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.GoodsMeasureDetails.GoodsMeasure.Value = GetCellValue("Накладная", "AZ"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.GoodsMeasureDetails.GoodsMeasure.MeasurementUnitCode = "796"
	good.HMConsignmentItemNumber = GetCellValue("Накладная", "AW"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	fcostcurrency := GetCellValue("Накладная", "BB"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	fcostvalue := GetCellValue("Накладная", "BC"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	fcost := CurrencyValue{Value: &fcostvalue, CurrencyCode: String(fcostcurrency), CurrencyCodeListId: String("2022")}
	kzcostvalue := GetCellValue("Накладная", "BD"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	kzcost := CurrencyValue{Value: &kzcostvalue, CurrencyCode: String("KZT"), CurrencyCodeListId: String("2022")}
	good.CAValueAmount = append(good.CAValueAmount, fcost)
	good.CAValueAmount = append(good.CAValueAmount, kzcost)
	// fcost := CurrencyValue{Value: &fcostvalue, CurrencyCode: String(fcostcurrency), CurrencyCodeListId: String("2022")}
	// kzcost := CurrencyValue{Value: &kzcostvalue, CurrencyCode: String("KZT"), CurrencyCodeListId: String("2022")}

	good.ECPresentedDocDetails.DocKindCode.Value = GetCellValue("Накладная", "BE"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.DocKindCode.CodeListId = "2009"
	good.ECPresentedDocDetails.DocId = GetCellValue("Накладная", "BG"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.DocCreationDate = GetCellValue("Накладная", "BH"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.DocStartDate = GetCellValue("Накладная", "BI"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.DocValidityDate = GetCellValue("Накладная", "BJ"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.UnifiedCountryCode.Value = GetCellValue("Накладная", "BK"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.UnifiedCountryCode.CodeListId = "2021"
	good.ECPresentedDocDetails.AuthorityName = GetCellValue("Накладная", "BL"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.AuthorityId = GetCellValue("Накладная", "BM"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	good.ECPresentedDocDetails.DocumentPresentingDetails.DocPresentKindCode = GetCellValue("Накладная", "BF"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)

	return good
}

func ReadECHouseShipmentDetail(itemCount int, xlsx *excelize.File) ECHouseShipmentDetail {
	hsd := ECHouseShipmentDetail{}
	hsd.ObjectOrdinal = GetCellValue("Накладная", "A"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.TransportDocumentDetails.DocId = GetCellValue("Накладная", "B"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.TransportDocumentDetails.DocCreationDate = GetCellValue("Накладная", "C"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.HouseWaybillDetails.DocId = GetCellValue("Накладная", "D"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.HouseWaybillDetails.DocCreationDate = GetCellValue("Накладная", "E"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectName = GetCellValue("Накладная", "F"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectBriefName = GetCellValue("Накладная", "G"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.TaxpayerId = GetCellValue("Накладная", "H"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.AddressKindCode = GetCellValue("Накладная", "O"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.Value = GetCellValue("Накладная", "P"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	hsd.ConsignorDetails.SubjectAddressDetails.RegionName = GetCellValue("Накладная", "Q"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.DistrictName = GetCellValue("Накладная", "R"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.CityName = GetCellValue("Накладная", "S"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.SettlementName = GetCellValue("Накладная", "T"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.StreetName = GetCellValue("Накладная", "U"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.BuildingNumberId = GetCellValue("Накладная", "V"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.SubjectAddressDetails.RoomNumberId = GetCellValue("Накладная", "W"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.CommunicationDetails.CommunicationChannelCode = GetCellValue("Накладная", "X"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.CommunicationDetails.CommunicationChannelName = GetCellValue("Накладная", "Y"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsignorDetails.CommunicationDetails.CommunicationChannelId = GetCellValue("Накладная", "Z"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectName = GetCellValue("Накладная", "AA"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectBriefName = GetCellValue("Накладная", "AB"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.PersonId = GetCellValue("Накладная", "AC"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.AddressKindCode = GetCellValue("Накладная", "AJ"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.Value = GetCellValue("Накладная", "AK"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	hsd.ConsigneeDetails.SubjectAddressDetails.RegionName = GetCellValue("Накладная", "AL"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.DistrictName = GetCellValue("Накладная", "AM"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.CityName = GetCellValue("Накладная", "AN"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.SettlementName = GetCellValue("Накладная", "AO"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.StreetName = GetCellValue("Накладная", "AP"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.BuildingNumberId = GetCellValue("Накладная", "AQ"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.SubjectAddressDetails.RoomNumberId = GetCellValue("Накладная", "AR"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.CommunicationDetails.CommunicationChannelCode = GetCellValue("Накладная", "AS"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.CommunicationDetails.CommunicationChannelName = GetCellValue("Накладная", "AT"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.CommunicationDetails.CommunicationChannelId = GetCellValue("Накладная", "AU"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.UnifiedCountryCode.CodeListId = "2021"
	hsd.ConsigneeDetails.IdentityDocV3Details.UnifiedCountryCode.Value = GetCellValue("Накладная", "AD"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.IdentityDocKindCode.CodeListId = "2053"
	hsd.ConsigneeDetails.IdentityDocV3Details.IdentityDocKindCode.Value = GetCellValue("Накладная", "AE"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.DocKindName = GetCellValue("Накладная", "AF"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.DocSeriesId = GetCellValue("Накладная", "AG"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.DocId = GetCellValue("Накладная", "AH"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	hsd.ConsigneeDetails.IdentityDocV3Details.DocCreationDate = GetCellValue("Накладная", "AI"+strconv.FormatInt(int64(itemCount+5), 10), xlsx)
	return hsd
}

func ReadECGoodsShipmentDetails(ecd *ExpressCargoDeclaration, xlsx *excelize.File) {
	ecd.ExpressRegistryKindCode = GetCellValue("Общие сведения", "E1", xlsx)
	// Тип декларации
	ecd.DeclarationKindCode = GetCellValue("Общие сведения", "D1", xlsx)
	// Код таможенного органа
	ecd.ExpressCargoDeclarationIdDetails.CustomsOfficeCode = GetCellValue("Общие сведения", "H1", xlsx)
	ecd.DeclarationFeatureCode = GetCellValue("Общие сведения", "W1", xlsx)
	// Регистрационный номер юридического лица при включении в реестр
	ecd.RegisterDocumentIdDetails.RegistrationNumberId = GetCellValue("Общие сведения", "AC15", xlsx)
	// Код вида документа
	ecd.RegisterDocumentIdDetails.DocKindCode = GetCellValue("Общие сведения", "AA15", xlsx)
	ecd.ECGoodsShipmentDetails.CAValueAmount.Value = func() *string {
		str := "0.00"
		return &str
	}()
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCode = String("KZT")
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCodeListId = String("2022")
	ecd.ECGoodsShipmentDetails.UnifiedGrossMassMeasure = "0.000"
	ecd.ExpressRegistryKindCode = GetCellValue("Общие сведения", "E1", xlsx)
	ecd.DeclarationKindCode = GetCellValue("Общие сведения", "D1", xlsx)
	ecd.ExpressCargoDeclarationIdDetails.CustomsOfficeCode = GetCellValue("Общие сведения", "H1", xlsx)
	ecd.DeclarationFeatureCode = GetCellValue("Общие сведения", "W1", xlsx)
	ecd.RegisterDocumentIdDetails.RegistrationNumberId = GetCellValue("Общие сведения", "AC15", xlsx)
	ecd.RegisterDocumentIdDetails.DocKindCode = GetCellValue("Общие сведения", "AA15", xlsx)
	ecd.ECGoodsShipmentDetails.CAValueAmount.Value = func() *string {
		str := "0.00"
		return &str
	}()
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCode = func() *string {
		str := "KZT"
		return &str
	}()
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCodeListId = func() *string {
		str := "2022"
		return &str
	}()
	ecd.ECGoodsShipmentDetails.UnifiedGrossMassMeasure = func() string {
		str := "0.000"
		return str
	}()
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.FirstName = GetCellValue("Общие сведения", "A15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.LastName = GetCellValue("Общие сведения", "E15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.MiddleName = GetCellValue("Общие сведения", "C15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.PositionName = GetCellValue("Общие сведения", "G15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelCode = GetCellValue("Общие сведения", "J15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelName = GetCellValue("Общие сведения", "K15", xlsx)
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelId = GetCellValue("Общие сведения", "M15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.Value = GetCellValue("Общие сведения", "O15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.Value = GetCellValue("Общие сведения", "P15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.CodeListId = "2021"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.CodeListId = "2053"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocKindName = GetCellValue("Общие сведения", "R15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocSeriesId = GetCellValue("Общие сведения", "S15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocId = GetCellValue("Общие сведения", "T15", xlsx)
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocCreationDate = GetCellValue("Общие сведения", "U15", xlsx)
	ecd.EDocIndicatorCode = "ЭД"
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.CodeListId = "2009"
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.Value = GetCellValue("Общие сведения", "V15", xlsx)
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocId = GetCellValue("Общие сведения", "W15", xlsx)
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocCreationDate = GetCellValue("Общие сведения", "X15", xlsx)
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocStartDate = GetCellValue("Общие сведения", "Y15", xlsx)
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocValidityDate = GetCellValue("Общие сведения", "Z15", xlsx)
	//
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectName = GetCellValue("Общие сведения", "A6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectBriefName = GetCellValue("Общие сведения", "B6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.TaxpayerId = GetCellValue("Общие сведения", "C6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.AddressKindCode = GetCellValue("Общие сведения", "E6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.Value = GetCellValue("Общие сведения", "F6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RegionName = GetCellValue("Общие сведения", "G6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.DistrictName = GetCellValue("Общие сведения", "H6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.CityName = GetCellValue("Общие сведения", "I6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.SettlementName = GetCellValue("Общие сведения", "J6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.StreetName = GetCellValue("Общие сведения", "K6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.BuildingNumberId = GetCellValue("Общие сведения", "L6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RoomNumberId = GetCellValue("Общие сведения", "M6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelCode = GetCellValue("Общие сведения", "N6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelName = GetCellValue("Общие сведения", "O6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelId = GetCellValue("Общие сведения", "P6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectName = GetCellValue("Общие сведения", "Q6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectBriefName = GetCellValue("Общие сведения", "R6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.TaxpayerId = GetCellValue("Общие сведения", "S6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.AddressKindCode = GetCellValue("Общие сведения", "U6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.Value = GetCellValue("Общие сведения", "V6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RegionName = GetCellValue("Общие сведения", "W6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.DistrictName = GetCellValue("Общие сведения", "X6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.CityName = GetCellValue("Общие сведения", "Y6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.SettlementName = GetCellValue("Общие сведения", "Z6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.StreetName = GetCellValue("Общие сведения", "AA6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.BuildingNumberId = GetCellValue("Общие сведения", "AB6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RoomNumberId = GetCellValue("Общие сведения", "AC6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelCode = GetCellValue("Общие сведения", "AD6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelName = GetCellValue("Общие сведения", "AE6", xlsx)
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelId = GetCellValue("Общие сведения", "AF6", xlsx)
}

func SumHsdCAValueAmount(ehsp *ECHouseShipmentDetail) {
	var ehspCAValueAmount float64                   //Итого стоимость по декларации
	for _, goods := range ehsp.ECGoodsItemDetails { //Проходим по товарам
		fmt.Println("CAValueAmount is ", *goods.CAValueAmount[1].Value)
		value, _ := strconv.ParseFloat(*goods.CAValueAmount[1].Value, 64) //Стоимость товара в тенге
		ehspCAValueAmount += value                                        //Суммируем
	}

	strValue := fmt.Sprintf("%.2f", ehspCAValueAmount)
	ehsp.CAValueAmount.Value = &strValue

	// return *ehsp.CAValueAmount.Value
}

func SumHsdUnifiedGrossMassMeasure(ehsp *ECHouseShipmentDetail) {
	var ehspUnifiedGrossMassMeasure float64
	for _, goods := range ehsp.ECGoodsItemDetails { //Проходим по товарам
		fmt.Println("UnifiedGrossMassMeasure is ", goods.UnifiedGrossMassMeasure)
		value, _ := strconv.ParseFloat(goods.UnifiedGrossMassMeasure, 64) //Стоимость товара в тенге
		ehspUnifiedGrossMassMeasure += value                              //Суммируем
	}

	if ehspUnifiedGrossMassMeasure < 0.001 {
		strValue := fmt.Sprintf("%.6f", ehspUnifiedGrossMassMeasure)
		ehsp.UnifiedGrossMassMeasure = strValue
	} else {
		strValue := fmt.Sprintf("%.3f", ehspUnifiedGrossMassMeasure)
		ehsp.UnifiedGrossMassMeasure = strValue
	}
}

func SumEcdCAValueAmount(ecd *ECGoodsShipmentDetail) {
	var ehspCAValueAmount float64
	var ecdUnifiedGrossMassMeasure float64
	for i, hsd := range ecd.ECHouseShipmentDetails {
		CAValueAmount, _ := strconv.ParseFloat(*hsd.CAValueAmount.Value, 64)
		fmt.Println("CAValueAmount of ", i, "is ", CAValueAmount)
		ehspCAValueAmount += CAValueAmount
		UnifiedGrossMassMeasure, _ := strconv.ParseFloat(hsd.UnifiedGrossMassMeasure, 64)
		fmt.Println("ecdUnifiedGrossMassMeasure of ", i, "is ", ecdUnifiedGrossMassMeasure)
		ecdUnifiedGrossMassMeasure += UnifiedGrossMassMeasure
	}
	strCAValueAmount := fmt.Sprintf("%.2f", ehspCAValueAmount)
	ecd.CAValueAmount.Value = &strCAValueAmount
	if ecdUnifiedGrossMassMeasure < 0.001 {
		strUnifiedGrossMassMeasure := fmt.Sprintf("%.6f", ecdUnifiedGrossMassMeasure)
		ecd.UnifiedGrossMassMeasure = strUnifiedGrossMassMeasure
	} else {
		strUnifiedGrossMassMeasure := fmt.Sprintf("%.3f", ecdUnifiedGrossMassMeasure)
		ecd.UnifiedGrossMassMeasure = strUnifiedGrossMassMeasure
	}
}

// func SetECHouseShipmentDetail() declaration.ECHouseShipmentDetail {
// 	nak := ECHouseShipmentDetail{}
// 	nak.ObjectOrdinal := GetCellValue()
// 	return nak
// }
