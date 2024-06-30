// utils/file.go
package utils

import "os"

func SaveFile(content, path string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
