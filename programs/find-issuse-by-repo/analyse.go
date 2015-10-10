package main

import (
	github "../../github"
	ospaf "../../lib"
	"encoding/json"
	"fmt"
)

type Config struct {
	Owner string
	Repo  string
}

func openIssueReport(issueList []github.Issue) {
}

func main() {
	content, err := ospaf.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config
	json.Unmarshal([]byte(content), &config)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", config.Owner, config.Repo)
	fileUrl := fmt.Sprintf("data/issue-of-repo-%s", ospaf.MD5(url))

	content, err = ospaf.ReadFile(fileUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	var issueList []github.Issue
	json.Unmarshal([]byte(content), &issueList)

	for index := 0; index < len(issueList); index++ {
		fmt.Println("\n\n")
		fmt.Println(issueList[index])
	}

	ospaf.PreparePath("report", "")

}
