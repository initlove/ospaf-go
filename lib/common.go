package ospafLib

import (
	"bytes"
	"fmt"
	"os"
)

func ReadFile(file_url string) (content string, err error) {
	_, err = os.Stat(file_url)
	if err != nil {
		content = fmt.Sprintf("Cannot find the file %s.", file_url)
		return content, err
	}
	file, err := os.Open(file_url)
	defer file.Close()
	if err != nil {
		content = fmt.Sprintf("Cannot open the file %s.", file_url)
		return content, err
	}
	buf := bytes.NewBufferString("")
	buf.ReadFrom(file)
	content = buf.String()

	return content, nil
}
