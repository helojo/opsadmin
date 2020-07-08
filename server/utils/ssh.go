package utils

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

// 连接ssh 返回对象
func SshClient(host string, port int64, username, password string) (sshClient *ssh.Client, sftpClient *sftp.Client, err error) {
	config := ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", host, port)

	sshClient, err = ssh.Dial("tcp", addr, &config)
	if err != nil {
		return sshClient, sftpClient, err
	}

	//此时获取了sshClient，下面使用sshClient构建sftpClient
	sftpClient, err = sftp.NewClient(sshClient)
	if err != nil {
		return sshClient, sftpClient, err
	}

	return sshClient, sftpClient, err
}

// 判断远程目录是否存在
func SshDirIsExist(sftp *sftp.Client, path string) (err error) {
	_, err = sftp.ReadDir(path)
	if err != nil {
		return err
	}
	return err
}

// 为用户创建.ssh 目录
func SshCreateDir(ssh *ssh.Client, path string) (err error) {
	//获取session，这个session是用来远程执行操作的
	session, err := ssh.NewSession()
	if err != nil {
		return errors.New(fmt.Sprintf("SSH 创建 session 连接失败, 报错信息: %s", err))
	}
	//执行shell
	if _, err := session.CombinedOutput(fmt.Sprintf("mkdir -p %s", path)); err != nil {
		return errors.New(fmt.Sprintf("SSH 创建目录失败, 报错信息: %s", err))
	}
	return err
}
