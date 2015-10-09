package main

import (
	github "../../github"
	ospaf "../../lib"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func loadComments(url string, pool ospaf.Pool) {
	fmt.Println("Start to load: ", url)

	var commentList []github.Comment
	paras := make(map[string]string)
	for page := 1; page != -1; {
		paras["page"] = strconv.Itoa(page)
		value, code, nextPage, _ := pool.ReadPage(url, paras)
		if code != 200 {
			break
		}
		page = nextPage

		var comments []github.Comment
		json.Unmarshal([]byte(value), &comments)
		for index := 0; index < len(comments); index++ {
			commentList = append(commentList, comments[index])
		}
	}

	fileUrl := fmt.Sprintf("data/comment-of-issue-%s", ospaf.MD5(url))
	content, _ := json.MarshalIndent(commentList, "", "  ")
	fout, err := os.Create(fileUrl)
	if err != nil {
		fmt.Println(fileUrl, err)
	} else {
		fout.WriteString(string(content))
		fmt.Println("Save ", url, " to ", fileUrl)
		fout.Close()
	}
}

func main() {
	pool, err := ospaf.InitPool()
	if err != nil {
		fmt.Println(err)
		return
	}

	ospaf.PreparePath("data", "")

	owner := "fakeowner"
	repo := "fakerepo"
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	paras := make(map[string]string)
	paras["state"] = "all"

	var issueList []github.Issue
	for page := 1; page != -1; {
		paras["page"] = strconv.Itoa(page)
		value, code, nextPage, _ := pool.ReadPage(url, paras)
		if code != 200 {
			break
		}
		page = nextPage

		var issues []github.Issue
		json.Unmarshal([]byte(value), &issues)

		for index := 0; index < len(issues); index++ {
			if issues[index].Pull_request.Url == "" {
				fmt.Println("Pull request, drop")
				continue
			}
			issueList = append(issueList, issues[index])
			commentUrl := issues[index].Comments_url
			loadComments(commentUrl, pool)
		}
	}

	fileUrl := fmt.Sprintf("data/issue-of-repo-%s", ospaf.MD5(url))
	content, _ := json.MarshalIndent(issueList, "", "  ")
	fout, err := os.Create(fileUrl)
	if err != nil {
		fmt.Println(fileUrl, err)
	} else {
		fout.WriteString(string(content))
		fmt.Println("Save ", url, " to ", fileUrl)
		fout.Close()
	}

}
