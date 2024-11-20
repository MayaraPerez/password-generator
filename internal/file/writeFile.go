package file

import (
	"fmt"
	"os"
)

func GeneratedFile(fileName string, content string) (string, error) {
	fileData, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	_, err = fileData.WriteString(content + "\n")
	if err != nil {
		return "", fmt.Errorf("erro ao escrever no arquivo: %v", err)
		
	}
	return fileName, nil
}