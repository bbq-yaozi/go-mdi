package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/bbq-yaozi/go-mdi/internal/generate"
)

func main() {
	metaList, err := generate.LoadMetaList()
	if err != nil {
		panic(err)
	}

	typesTempl, err := generate.LoadTemplate(generate.TypesGoFile)
	if err != nil {
		panic(err)
	}

	funcsTempl, err := generate.LoadTemplate(generate.FuncsGoFile)
	if err != nil {
		panic(err)
	}

	iconTempl, err := generate.LoadTemplate(generate.IconGoFile)
	if err != nil {
		panic(err)
	}

	_ = os.RemoveAll("mdi")
	_ = os.MkdirAll("mdi", os.ModePerm)

	metaList = metaList[:1]

	// create mdi/types.go
	var typesContent bytes.Buffer
	if err = typesTempl.Execute(&typesContent, map[string]any{
		"version": "v7.4.47",
	}); err != nil {
		panic(err)
	}

	if err = saveToFile("mdi/types.go", typesContent.Bytes()); err != nil {
		panic(err)
	}

	// create mdi/funcs.go
	var funcsContent bytes.Buffer
	if err = funcsTempl.Execute(&funcsContent, map[string]any{
		"metaList": metaList,
	}); err != nil {
		panic(err)
	}

	fmt.Println(string(funcsContent.Bytes()))
	if err = saveToFile("mdi/funcs.go", funcsContent.Bytes()); err != nil {
		panic(err)
	}

	// create mdi/[IconName].go
	for _, meta := range metaList {
		var iconContent bytes.Buffer
		if err = iconTempl.Execute(&iconContent, map[string]any{
			"meta": meta,
		}); err != nil {
			panic(err)
		}

		if err = saveToFile(fmt.Sprintf("mdi/%s.go", meta.Name), iconContent.Bytes()); err != nil {
			panic(err)
		}
	}
}

func saveToFile(fileName string, data []byte) error {
	if _, err := os.Stat(fileName); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat file[%s] error:%s", fileName, err)
		}

	} else {
		if err := os.Remove(fileName); err != nil {
			return fmt.Errorf("remove file[%s] error:%s", fileName, err)
		}
	}

	fs, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fs.Close()

	_, err = fs.Write(data)
	return err
}
