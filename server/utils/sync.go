package utils

import (
	"errors"
	"fmt"
	"gin-vue-admin/utils/grsync"
	"strings"
)

func FileContrast(sourceDir string, user string, tagetServer string, tagetDir string, exclude []string) (err error, list interface{}) {
	destination := fmt.Sprintf("%s@%s:%s", user, tagetServer, tagetDir)
	task := grsync.NewTask(
		sourceDir,
		destination,
		grsync.RsyncOptions{
			Archive:   true,                                   //归档模式，表示以递归方式传输文件
			Compress:  true,                                   //压缩处理
			Recursive: true,                                   //对子目录以递归模式处理
			Progress:  true,                                   //显示过程
			Delete:    true,                                   //删除那些DST中SRC没有的文件
			Force:     true,                                   //强制删除目录，即使不为空
			DryRun:    true,                                   //显示哪些文件将被传输
			Checksum:  true,                                   //打开校验开关，强制对文件传输进行校验
			OutFormat: true,                                   //输出日志格式化
			Exclude:   exclude,                                //过滤文件
			Rsh:       "ssh -i ./resource/platformkey/id_rsa", //指定平台私钥
		},
	)
	if err := task.Run(); err != nil {
		return errors.New(fmt.Sprint("Git拉取项目报错, 报错信息: %s", err)), list
	}
	result := strings.Split(task.Log().Stdout, "\n")
	result_list := make([]map[string]string, 0)
	for _, v := range result {
		file := map[string]string{}
		if !strings.HasSuffix(v, "/\"") {
			if !strings.HasPrefix(v, "deleting") && v != "" {
				v = strings.ReplaceAll(v, "\"", "")
				if v != "" {
					file["key"] = v
					result_list = append(result_list, file)
				}
			} else {
				if v != "" {
					file["key"] = v
					result_list = append(result_list, file)
				}

			}
		}
	}
	return err, result_list
}

func FileSync(sourceDir string, user string, tagetServer string, tagetDir string, exclude []string) (err error, result string) {
	destination := fmt.Sprintf("%s@%s:%s", user, tagetServer, tagetDir)
	task := grsync.NewTask(
		sourceDir,
		destination,
		grsync.RsyncOptions{
			Archive:   true, //归档模式，表示以递归方式传输文件
			Compress:  true, //压缩处理
			Recursive: true, //对子目录以递归模式处理
			Progress:  true, //显示过程
			Delete:    true, //删除那些DST中SRC没有的文件
			Force:     true, //强制删除目录，即使不为空
			//DryRun:    true,                                 //显示哪些文件将被传输
			Checksum:  true,                                   //打开校验开关，强制对文件传输进行校验
			OutFormat: true,                                   //输出日志格式化
			Exclude:   exclude,                                //过滤文件
			Rsh:       "ssh -i ./resource/platformkey/id_rsa", //指定平台私钥
		},
	)
	if err := task.Run(); err != nil {
		return errors.New(fmt.Sprint("同步报错, 报错信息: %s", err)), result
	}
	syncresult := task.Log().Stdout
	return err, syncresult
}
