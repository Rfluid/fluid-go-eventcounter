package common_service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CreateAndWriteFile(path, name string, content map[string]int) error {
	file, err := os.Create(fmt.Sprintf("%s/%s.json", path, name))
	if err != nil {
		log.Printf("can't write file %s.json, err: %s", name, err)
		return err
	}
	defer file.Close()

	b, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		log.Printf("can't marshal data for file %s.json, err: %s", name, err)
		return err
	}

	if _, err := file.Write(b); err != nil {
		log.Printf("can't write data for file %s.json, err: %s", name, err)
		return err
	}

	return nil
}
