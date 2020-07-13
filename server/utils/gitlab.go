package utils

import (
	"fmt"
	"gin-vue-admin/global"
	"github.com/xanzy/go-gitlab"
	"strconv"
)

func GetGitlabprojects() (projecs map[int]map[string]string, err error) {
	// 获取项目列表
	var url string = global.GVA_CONFIG.Gitlab.Url
	var token string = global.GVA_CONFIG.Gitlab.Token
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		panic(err)
	}
	page := 0
	var project_list = make(map[int]map[string]string)
	for i := 1; i <= 100000; i++ {
		page = i
		opt := &gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		}
		projects, _, err := git.Projects.ListProjects(opt)
		if len(projects) == 0 && err == nil {
			break
		}
		for _, pt := range projects {
			var project = make(map[string]string)
			project["id"] = strconv.Itoa(pt.ID)
			project["name"] = pt.Name
			project["url"] = pt.HTTPURLToRepo
			project_list[pt.ID] = project
		}
	}
	return project_list, err
}

func ProjectBranches(id int) (taglist []interface{}, err error) {
	// 获取项目列表
	var url string = global.GVA_CONFIG.Gitlab.Url
	var token string = global.GVA_CONFIG.Gitlab.Token
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return taglist, err
	}
	//获取分支tag
	opttag := &gitlab.ListTagsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	tags, _, err := git.Tags.ListTags(id, opttag)
	var tagList []interface{}
	for _, tag := range tags {
		tagMap := make(map[string]string)
		tagMap["id"] = tag.Name
		tagMap["name"] = tag.Name
		tagList = append(tagList, tagMap)
	}
	fmt.Println(tagList, err)
	return tagList, err
}

func GetProjectBranches(id int) (branchelist []interface{}, taglist []interface{}, err error) {
	// 获取项目列表
	var url string = global.GVA_CONFIG.Gitlab.Url
	var token string = global.GVA_CONFIG.Gitlab.Token
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		panic(err)
	}
	// 获取项目分支
	opt := &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	branches, _, err := git.Branches.ListBranches(id, opt)
	var brancheList []interface{}
	for _, branche := range branches {
		brancheMap := make(map[string]string)
		brancheMap["id"] = branche.Name
		brancheMap["name"] = branche.Name
		brancheList = append(brancheList, brancheMap)
	}

	//获取分支tag
	opttag := &gitlab.ListTagsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	tags, _, err := git.Tags.ListTags(id, opttag)
	var tagList []interface{}
	for _, tag := range tags {
		tagMap := make(map[string]string)
		tagMap["id"] = tag.Name
		tagMap["name"] = tag.Name
		tagList = append(tagList, tagMap)
	}
	return brancheList, tagList, err
}
