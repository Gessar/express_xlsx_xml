package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Открываем файл XLSX для редактирования
	f, err := excelize.OpenFile("samplefile.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создаем новый лист с именем "Новый лист"
	index, err := f.NewSheet("Новый лист")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Устанавливаем активным листом только что созданный
	f.SetActiveSheet(index)

	// Сохраняем изменения обратно в файл XLSX
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Лист успешно добавлен в файл samplefile.xlsx!")
}
