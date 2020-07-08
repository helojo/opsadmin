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
		err = global.GVA_DB.Create(&server).Error
	}
	return err
}

// @title    ServerUpdate
// @description   更新主机信息
// @auth                     （2020/07/07  11:36）
// @param     env             model.ResourceServer
// @return                    error

func ServerUpdate(server model.ResourceServer) (err error) {
	err = global.GVA_DB.Where("id = ?", server.ID).First(&model.ResourceServer{}).Updates(&server).Error
	return err
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
// @param     string  p
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

	if isExist(id_rsa) && isExist(id_rsa_pub) {
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
