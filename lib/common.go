package ospafLib

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

//testStr := "https://api.github.com/repositories/36960293/issues/20/comments?page=2>; rel=\"last\", <https://api.github.com/repositories/36960293/issues/20/comments?page=1>; rel=\"first\", <https://api.github.com/repositories/36960293/issues/20/comments?page=23>; rel=\"prev\""
func GetPageMap(link string) (pageMap map[string]int) {
	pageMap = make(map[string]int)
	strSet := strings.Split(link, ",")
	for index := 0; index < len(strSet); index++ {
		re, _ := regexp.Compile("page=(\\d+)>; rel=\"(last|first|prev|next)\"")
		result := re.FindStringSubmatch(strSet[index])
		if len(result) == 3 {
			pageMap[result[2]], _ = strconv.Atoi(result[1])
		}
	}

	return pageMap
}
