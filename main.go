package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Gessar/express_xlsx_xml/internal/declaration"
	"github.com/xuri/excelize/v2"
)

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

func counter() {
	i := 0
	for {
		// fmt.Println(i)
		time.Sleep(time.Second * 1)
		i++
	}
}

func main() {

	filePathOld := "samplefile.xlsx"
	filePathNew := "template.xlsx"
	var declaration_xlsx_file *excelize.File

	// var declarationWeight float32 = 0
	ecd := declaration.ExpressCargoDeclaration{}
	ecd.Xmlns = "http://www.codecraft.kz/keden/ecd"
	ecd.XmlnsNs2 = "urn:EEC:M:SimpleDataObjects:v0.4.14"
	ecd.XmlnsNs3 = "urn:EEC:M:CA:SimpleDataObjects:v1.8.1"
	ecd.XmlnsNs4 = "urn:EEC:M:CA:ComplexDataObjects:v1.8.1"
	ecd.XmlnsNs5 = "urn:EEC:M:ComplexDataObjects:v0.4.14"

	// Проверяем, существует ли файл шаблона
	if _, err := os.Stat(filePathOld); os.IsNotExist(err) {
		// Если старый файл не существует, проверяем новый файл
		if _, err := os.Stat(filePathNew); os.IsNotExist(err) {
			// Если новый файл тоже не существует, скачиваем его
			fmt.Println("Файл отсутствует, начинаем загрузку...")
			if err := downloadFile("https://github.com/Gessar/express_xlsx_xml/raw/main/template.xlsx", filePathNew); err != nil {
				fmt.Println("Ошибка загрузки файла:", err)
				return
			}
		}
		// Открываем новый файл
		declaration_xlsx_file, err = excelize.OpenFile(filePathNew)
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			return
		}
		defer declaration_xlsx_file.Close()
	} else {
		// Если старый файл существует, говорим, что-бы удалили его

		fmt.Println("Удалите файл \"samplefile.xlsx\" и запустите снова", err)
		return

	}

	declaration.ReadECGoodsShipmentDetails(&ecd, declaration_xlsx_file)
	if ecd.ExpressRegistryKindCode == "ПТДЭГ" {
		log.Println("this is PTDEG")
		// fmt.Println()
	} else if ecd.ExpressRegistryKindCode == "ДТЭГ" {
		fmt.Println("this is DTEG")
	} else {
		fmt.Println("use \"ПТДЭГ\" or \"ДТЭГ\" in cell \"C1\" list \"Общие сведения\"")
		os.Exit(0)
	}
	ehsp := declaration.ECHouseShipmentDetail{} //Создаем накладную

	for i := 1; i <= 999; i++ { //От 1 до 999 товаров
		git := declaration.ECGoodsItemDetail{} //Создаем товар

		git = declaration.ReadGoodsItemDetail(i, declaration_xlsx_file) //Считываем товар

		if git.ConsignmentItemOrdinal == "" { //Если товаров больше нет, то добавить накладную к декларации

			declaration.SumHsdCAValueAmount(&ehsp)

			declaration.SumHsdUnifiedGrossMassMeasure(&ehsp)
			ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails = append(ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails, ehsp)
			break
		}

		hsdTrasportDocId, _ := declaration_xlsx_file.GetCellValue("Накладная", "B"+strconv.FormatInt(int64(i+5), 10)) //Смотрим есть ли на новой строке объявление накладной

		if hsdTrasportDocId != "" && i != 1 { //Если объявление есть, то добавляем старую накладную в декларацию, считываем новые данные по декларации

			declaration.SumHsdCAValueAmount(&ehsp)

			declaration.SumHsdUnifiedGrossMassMeasure(&ehsp)
			ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails = append(ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails, ehsp) //Добавляем накладную к декларации
			ehsp = declaration.ReadECHouseShipmentDetail(i, declaration_xlsx_file)                                              //Считаем новые данные накладной
		} else if i == 1 {
			ehsp = declaration.ReadECHouseShipmentDetail(i, declaration_xlsx_file)
		}

		ehsp.ECGoodsItemDetails = append(ehsp.ECGoodsItemDetails, git) //Добавляем товар в накладную

		if i == 999 { //Если последний товар, то добавить накладную к декларации и закончить

			declaration.SumHsdCAValueAmount(&ehsp)

			declaration.SumHsdUnifiedGrossMassMeasure(&ehsp)
			ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails = append(ecd.ECGoodsShipmentDetails.ECHouseShipmentDetails, ehsp)
			break
		}

		fmt.Println("tovar is number ", git.ConsignmentItemOrdinal, "\\", git.HMConsignmentItemNumber)
	}
	declaration.SumEcdCAValueAmount(&ecd.ECGoodsShipmentDetails)

	xmlData, err := xml.MarshalIndent(ecd, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации в XML:", err)
		return
	}
	outputstring := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n" + string(xmlData)

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02_150405")
	outputFile := fmt.Sprintf("express_%s.xml", formattedTime)

	declaration.WriteToFile(outputFile, outputstring)

	fmt.Printf("XML файл успешно создан: %s\n", outputFile)
	fmt.Println("Created by Gessar")
	fmt.Println("https://github.com/Gessar/express_xlsx_xml")
	go counter()
	fmt.Print("Press 'Enter' to exit...")
	fmt.Scanln()
}
