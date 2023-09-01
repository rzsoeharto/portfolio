package jwt_utils

import (
	"fmt"
	"io"
	"os"
)

func LoadPEMFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	pem, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return pem, nil
}
