// utils/file.go
package utils

import (
	"io/ioutil"
)

func SaveFile(content, path string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}
