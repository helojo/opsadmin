package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"io/ioutil"
	"os"
)

// @title    ServerList
// @description   get server list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func ServerList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ResourceServer{})
	var serverList []model.ResourceServer
	err = db.Count(&total).Error
	err = db.Preload("ResourceEnv").Limit(limit).Offset(offset).Find(&serverList).Error
	return err, serverList, total
}

// @title    ServerCreate
// @description   create base server, 新增主机
// @auth                     （2020/07/07  09:45）
// @param     api             model.ResourceServer
// @return                    error

func ServerCreate(server model.ResourceServer) (err error) {
	findOne := global.GVA_DB.Where("host = ? and resource_env_id = ?", server.Host, server.ResourceEnvId).Find(&model.ResourceServer{}).Error
	if findOne == nil {
		return errors.New("存在相同主机地址")
	} else {
		server.Status = 1
		server.Pwd, err = utils.EnPwdCode([]byte(server.Pwd))
		if err != nil {
			return err
		}
		err = global.GVA_DB.Create(&server).Error
	}
	return err
}

// @title    ServerMsgUpdate
// @description   更新主机信息
// @auth                     （2020/04/05  20:22）
// @param     env             model.ResourceServer
// @return                    error

func ServerMsgUpdate(server model.ResourceServer) (err error) {
	err = global.GVA_DB.Where("id = ?", server.ID).First(&model.ResourceServer{}).Updates(&server).Error
	return err
}

// @title    ServerUpdate
// @description   更新主机信息
// @auth                     （2020/07/07  11:36）
// @param     env             model.ResourceServer
// @return                    error

func ServerUpdate(server model.ResourceServer) (err error) {
	var serverOld model.ResourceServer
	findOne := global.GVA_DB.Where("id = ?", server.ID).Find(&serverOld).Error
	if findOne != nil {
		return errors.New("该主机不存在!")
	}
	if server.Pwd != serverOld.Pwd {
		newpassword, _ := utils.EnPwdCode([]byte(server.Pwd))
		server.Status = 1
		server.Pwd = newpassword
		return ServerMsgUpdate(server)
	}
	server.Status = 1
	return ServerMsgUpdate(server)
}

// @title    ServerDelete
// @description    删除主机信息
// @auth                     （2020/07/07  11:52）
// @param     id
// @return    err             error

func ServerDelete(id float64) (err error) {
	var server model.ResourceServer
	err = global.GVA_DB.Where("id = ?", id).Delete(&server).Error
	return err
}

// @title    isExist
// @description  判断文件是否存在
// @auth                     （2020/07/08  11:11)
// @return    bool             true|flase
func isExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil || os.IsExist(err)
}

// @title    ServerCreateKey
// @description    创建平台密钥对
// @auth                     （2020/07/08  9:51）
// @param     id
// @return    err             error

func PlatformCreateKey() (err error) {
	id_rsa := global.GVA_CONFIG.Platformkey.Path + "id_rsa"
	id_rsa_pub := global.GVA_CONFIG.Platformkey.Path + "id_rsa.pub"

	if isExist(id_rsa) || isExist(id_rsa_pub) {
		return errors.New("平台密钥对文件, 已经存在!")
	}

	pkey, pubkey, err := utils.MakeSSHKeyPair()
	if err != nil {
		return errors.New(fmt.Sprintf("平台生成密钥对失败, 报错信息: %s", err))
	}

	err = ioutil.WriteFile(id_rsa, []byte(pkey), 0600|os.ModeAppend)
	if err == nil {
		err = ioutil.WriteFile(id_rsa_pub, []byte(pubkey), 0600|os.ModePerm)
	}
	return err
}

// @title    ServerConnect
// @description    测试连接主机
// @auth                     （2020/07/08  15:12）
// @param     id
// @return    err             error

func ServerConnect(id float64) (err error) {
	var server model.ResourceServer
	findOne := global.GVA_DB.Where("id = ?", id).Find(&server).Error
	if findOne != nil {
		return errors.New("该主机不存在!")
	}

	password, err := utils.DePwdCode(server.Pwd)
	if err != nil {
		return errors.New(fmt.Sprintf("服务器密码解密错误, 报错信息: %s", err))
	}

	sshClient, sftpClient, err := utils.SshClient(server.Host, server.Port, server.User, string(password))
	if err != nil {
		server.Status = 4
		_ = ServerMsgUpdate(server)
		return err
	}

	var dstDir string
	if server.User != "root" {
		dstDir = "/home/" + server.User + "/.ssh"
	} else {
		dstDir = "/root/.ssh"
	}

	if err = utils.SshDirIsExist(sftpClient, dstDir); err != nil {
		cmd := fmt.Sprintf("mkdir -p %s && chmod  0700 %s -R", dstDir, dstDir)
		if err = utils.SshCmd(sshClient, cmd); err != nil {
			server.Status = 4
			_ = ServerMsgUpdate(server)
			return errors.New(fmt.Sprintf("%s 该主机上, %s 目录不存在, 且创建失败，报错信息：%s", server.Host, dstDir, err))
		}
	}
	server.Status = 3
	_ = ServerMsgUpdate(server)

	return err
}

// @title    ServerPushKey
// @description    推送公钥
// @auth                     （2020/07/09  11:52）
// @param     id
// @return    err             error

func ServerPushKey(id float64) (err error) {
	var server model.ResourceServer
	findOne := global.GVA_DB.Where("id = ?", id).Find(&server).Error
	if findOne != nil {
		return errors.New("该主机不存在!")
	}

	if server.Status != 3 {
		return errors.New("请测试连接，再进行密钥推送!")
	}

	password, err := utils.DePwdCode(server.Pwd)
	if err != nil {
		return errors.New(fmt.Sprintf("服务器密码解密错误, 报错信息: %s", err))
	}

	sshClient, sftpClient, err := utils.SshClient(server.Host, server.Port, server.User, string(password))
	if err != nil {
		return errors.New(fmt.Sprintf("连接远程主机失败, 报错信息: %s", err))
	}

	var dstFile string
	if server.User != "root" {
		dstFile = "/home/" + server.User + "/.ssh/authorized_keys"
	} else {
		dstFile = "/root/.ssh/authorized_keys"
	}
	//判断远程文件是否存在
	err = utils.SshFileIsExist(sftpClient, dstFile)
	if err != nil {
		// 文件不存在，则直接推送文件到机器
		id_rsa_pub := global.GVA_CONFIG.Platformkey.Path + "id_rsa.pub"
		err := utils.SftpUpload(sftpClient, id_rsa_pub, dstFile)
		if err == nil {
			cmd := fmt.Sprintf("chmod  0600 %s", dstFile)
			err := utils.SshCmd(sshClient, cmd)
			if err != nil {
				server.Status = 6
				return errors.New(fmt.Sprintf("密钥推送成功，文件授权失败, 报错信息: %s", err))
			}
			server.Status = 5
			_ = ServerMsgUpdate(server)
			return err
		} else {
			return errors.New(fmt.Sprintf("密钥文件上传失败, 报错信息: %s", err))
		}

	}
	return err
}
