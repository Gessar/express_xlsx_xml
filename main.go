package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xuri/excelize/v2"
)

func downloadFile(url, filepath string) error {
	// Создаем файл для записи
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Получаем содержимое файла
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Путь к файлу XLSX
	filePath := "samplefile.xlsx"

	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Если файл отсутствует, скачиваем его
		fmt.Println("Файл отсутствует, начинаем загрузку...")
		if err := downloadFile("https://raw.githubusercontent.com/Gessar/express_xlsx_xml/main/samplefile.xlsx", filePath); err != nil {
			fmt.Println("Ошибка загрузки файла:", err)
			return
		}
		fmt.Println("Файл успешно загружен.")
	}

	// Открываем файл XLSX для редактирования
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	// Создаем новый лист с именем "Новый лист"
	index, err := f.NewSheet("Новый лист")

	// Устанавливаем активным листом только что созданный
	f.SetActiveSheet(index)

	// Сохраняем изменения обратно в файл XLSX
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return
	}

	fmt.Println("Лист успешно добавлен в файл samplefile.xlsx!")
}
