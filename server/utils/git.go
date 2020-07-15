package utils

import (
	"gin-vue-admin/global"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"net/http"
	"os"
)

// 拉取代码
func Gitpull(path string, projectUrl string) (err error) {
	var url string = global.GVA_CONFIG.Gitlab.Url
	var token string = global.GVA_CONFIG.Gitlab.Token
	_, err := git.PlainClone("temp/foo", false, &git.CloneOptions{
		URL:      "http://git.wjh.com/devops/helloword.git",
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "",
			Password: "cNwcPatbFYxc85DHndtc",
		},
		ReferenceName: plumbing.NewTagReferenceName("v1.0.0"),
		//ReferenceName: plumbing.NewBranchReferenceName("ci_dev"),
		SingleBranch: true,
	})
	if err != nil {
		fmt.Println("克隆项目报错:", err)
	}
}
