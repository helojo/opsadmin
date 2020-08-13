package utils

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
)

// 拉取代码
func Gitpull(tag string, projectUrl string, filepath string) (localpath string, err error) {
	var token string = global.GVA_CONFIG.Gitlab.Token
	var path string = global.GVA_CONFIG.Gitpull.Path + filepath + "/"

	//判断如果目录不存在则创建，存在则删除重建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	} else {
		os.RemoveAll(path)
		os.Mkdir(path, 0755)
	}

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      projectUrl,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "",
			Password: token,
		},
		ReferenceName: plumbing.NewTagReferenceName(tag),
		SingleBranch:  true,
	})

	if err != nil {
		return path, errors.New(fmt.Sprintf("克隆项目报错:%s", err))
	}
	return path, err
}
