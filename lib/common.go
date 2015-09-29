package ospafLib

import (
	"bytes"
	"fmt"
	"os"
)

func ReadFile(file_url string) (valid bool, content string) {
	_, err := os.Stat(file_url)
	if err != nil {
		content = fmt.Sprintf("Cannot find the file %s.", file_url)
		return false, content
	}
	file, err := os.Open(file_url)
	defer file.Close()
	if err != nil {
		content = fmt.Sprintf("Cannot open the file %s.", file_url)
		return false, content
	}
	buf := bytes.NewBufferString("")
	buf.ReadFrom(file)
	content = buf.String()

	return true, content
}
