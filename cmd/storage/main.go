package main

// При импорте самописного пакета нужно указывать полный путь
// Например: Объявили пакет "github.com/d-jackalope/storage"
// к нему далее путь "*пакет*/internal/storage"

import (
	"fmt"
	"log"

	"github.com/d-jackalope/storage/internal/storage"
)

// В данном контексте название директории ./storage - это
// название приложения, в ./cmd/ хранятся исходники

// Это приложение является некоторым "Hello World!" c модулями
// У нас есть storage, который внутри себя (памяти) хранит файлы
// и этот файл можно будет получить по ID

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello, newFile!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It's uploaded", file)

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It's restored", restoredFile)

}
