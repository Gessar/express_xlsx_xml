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
	tea "github.com/charmbracelet/bubbletea"
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

// menu
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "spacebar", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n↑[j] for up, ↓[k] for down, spacebar for select\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

// menu_end
func main() {
	// var isX, isY, isZ, isMenu bool //debug
	// var isRecreate bool
	// var isCopy bool
	// // var listsCount int
	// var rootCmd = &cobra.Command{
	// 	Use:   "app",
	// 	Short: "A brief description of your application",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("x is ", isX)
	// 		fmt.Println("y is ", isY)
	// 		fmt.Println("z is ", isZ)
	// 		fmt.Println("for create lists", isRecreate)
	// 		fmt.Println("for local copy template", isCopy)
	// 		// fmt.Println("r is ", isRecreate)
	// 		// fmt.Println("m is ", isMenuDebug)
	// 		// fmt.Println("n is ", name)
	// 	},
	// }
	// rootCmd.Flags().BoolVarP(&isX, "x", "x", false, "is x")
	// rootCmd.Flags().BoolVarP(&isY, "y", "y", false, "is y")
	// rootCmd.Flags().BoolVarP(&isZ, "z", "z", false, "is z")
	// rootCmd.Flags().BoolVarP(&isMenu, "menu", "m", false, "is menu")
	// rootCmd.Flags().BoolVarP(&isRecreate, "recreate", "r", false, "for create lists")
	// rootCmd.Flags().BoolVarP(&isCopy, "copy", "c", false, "for local copy")

	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// is_x := flag.Bool("x", false, "is x")
	// is_y := flag.Bool("y", false, "is y")
	// is_z := flag.Bool("z", false, "is z")
	// isRecreate := flag.Bool("r", false, "recreate")
	// isMenuDebug := flag.Bool("m", false, "menudebug")

	// // flag.Parse()
	// flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	// err := flag.CommandLine.Parse(filterCombinedFlags(os.Args[1:]))
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println("x is ", *is_x)
	// fmt.Println("y is ", *is_y)
	// fmt.Println("z is ", *is_z)

	// if isMenu {
	// 	fmt.Println("Enter menu debug")
	// 	p := tea.NewProgram(initialModel())
	// 	if _, err := p.Run(); err != nil {
	// 		fmt.Printf("Alas, there's been an error: %v", err)
	// 		os.Exit(1)
	// 	}
	// 	os.Exit(0)
	// }

	// // os.Exit(0)
	// // Путь к файлу XLSX
	// if !isX {
	// 	log.SetOutput(io.Discard)
	// }

	filePathOld := "samplefile.xlsx"
	filePathNew := "template.xlsx"
	var declaration_xlsx_file *excelize.File
	// var productOrder int = 0
	var declarationCost float32 = 0
	var declarationWeight float32 = 0
	ecd := declaration.ExpressCargoDeclaration{}
	ecd.Xmlns = "http://www.codecraft.kz/keden/ecd"
	ecd.XmlnsNs2 = "urn:EEC:M:SimpleDataObjects:v0.4.14"
	ecd.XmlnsNs3 = "urn:EEC:M:CA:SimpleDataObjects:v1.8.1"
	ecd.XmlnsNs4 = "urn:EEC:M:CA:ComplexDataObjects:v1.8.1"
	ecd.XmlnsNs5 = "urn:EEC:M:ComplexDataObjects:v0.4.14"

	// if isCopy {
	// 	err := os.Remove("samplefile.xlsx")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	data, err := os.ReadFile("samplefile_copy.xlsx")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = os.WriteFile("samplefile.xlsx", data, 0644)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	os.Exit(0)
	// }

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
	// if isRecreate {
	// 	declaration.AddLists(declaration_xlsx_file, filePathOld, 5)
	// 	os.Exit(1)
	// }

	// Вид декларации для экспресс-грузов
	// ecd.ExpressRegistryKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E1")
	ecd.ExpressRegistryKindCode = declaration.GetCellValue("Общие сведения", "E1", declaration_xlsx_file)
	if ecd.ExpressRegistryKindCode == "ПТДЭГ" {
		log.Println("this is PTDEG")
		// fmt.Println()
	} else if ecd.ExpressRegistryKindCode == "ДТЭГ" {
		fmt.Println("this is DTEG")
	} else {
		fmt.Println("use \"ПТДЭГ\" or \"ДТЭГ\" in cell \"C1\" list \"Общие сведения\"")
		os.Exit(0)
	}
	// Тип декларации
	ecd.DeclarationKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "D1")
	// Код таможенного органа
	ecd.ExpressCargoDeclarationIdDetails.CustomsOfficeCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "H1")
	ecd.DeclarationFeatureCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "W1")
	// Регистрационный номер юридического лица при включении в реестр
	ecd.RegisterDocumentIdDetails.RegistrationNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AC15")
	// Код вида документа
	ecd.RegisterDocumentIdDetails.DocKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AA15")
	formattedValue := declaration.FormatCurrencyValue(float64(declarationCost))

	ecd.ECGoodsShipmentDetails.CAValueAmount.Value = declaration.String(*formattedValue)
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCode = declaration.String("KZT")
	ecd.ECGoodsShipmentDetails.CAValueAmount.CurrencyCodeListId = declaration.String("2022")
	ecd.ECGoodsShipmentDetails.UnifiedGrossMassMeasure = *declaration.FormatGrossMassValue(float64(declarationWeight))
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.FirstName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "A15")
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.LastName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E15")
	ecd.SignatoryPersonV2Details.SigningDetails.FullNameDetails.MiddleName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "C15")
	ecd.SignatoryPersonV2Details.SigningDetails.PositionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "G15")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "J15")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "K15")
	ecd.SignatoryPersonV2Details.SigningDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "M15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "O15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "P15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.UnifiedCountryCode.CodeListId = "2021"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.IdentityDocKindCode.CodeListId = "2053"
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocKindName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "R15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocSeriesId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "S15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "T15")
	ecd.SignatoryPersonV2Details.IdentityDocV3Details.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "U15")
	ecd.EDocIndicatorCode = "ЭД"
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.CodeListId = "2009"
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocKindCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "V15")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "W15")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocCreationDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "X15")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocStartDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "Y15")
	ecd.SignatoryPersonV2Details.PowerOfAttorneyDetails.DocValidityDate, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "Z15")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "A6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "B6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.TaxpayerId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "C6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "E6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "F6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "G6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "H6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "I6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "J6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "K6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "L6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "M6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "N6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "O6")
	ecd.ECGoodsShipmentDetails.ConsignorDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "P6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "Q6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectBriefName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "R6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.TaxpayerId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "S6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.AddressKindCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "U6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.Value, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "V6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.UnifiedCountryCode.CodeListId = "2021"
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RegionName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "W6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.DistrictName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "X6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.CityName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "Y6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.SettlementName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "Z6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.StreetName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AA6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.BuildingNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AB6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.SubjectAddressDetails.RoomNumberId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AC6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelCode, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AD6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelName, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AE6")
	ecd.ECGoodsShipmentDetails.ConsigneeDetails.CommunicationDetails.CommunicationChannelId, _ = declaration_xlsx_file.GetCellValue("Общие сведения", "AF6")

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

	// err = os.WriteFile(outputFile, []byte(outputstring), 0644)
	// if err != nil {
	// 	fmt.Println("Ошибка при записи XML файла:", err)
	// 	return
	// }

	fmt.Printf("XML файл успешно создан: %s\n", outputFile)
}
