package utils

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
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

// 判断远程文件是否存在
func SshFileIsExist(sftp *sftp.Client, path string) (err error) {
	_, err = sftp.Open(path)
	if err != nil {
		return err
	}
	return err
}

// 远程执行命令
func SshCmd(ssh *ssh.Client, cmd string) (err error) {
	defer ssh.Close()
	//获取session，这个session是用来远程执行操作的
	session, err := ssh.NewSession()
	if err != nil {
		return errors.New(fmt.Sprintf("SSH 创建 session 连接失败, 报错信息: %s", err))
	}
	//执行shell
	if _, err := session.CombinedOutput(cmd); err != nil {
		return errors.New(fmt.Sprintf("SSH 执行命令失败, 报错信息: %s", err))
	}
	return err
}

// 下载文件
func SftpDownload(sftp *sftp.Client, srcPath, dstPath string) (err error) {
	srcFile, err := sftp.Open(srcPath) //远程
	if err != nil {
		return errors.New(fmt.Sprintf("文件不存在, 报错信息: %s", err))
	}

	dstFile, err := os.Create(dstPath) //创建本地文件
	if err != nil {
		return errors.New(fmt.Sprintf("本地文件创建失败, 报错信息: %s", err))
	}
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return errors.New(fmt.Sprintf("文件下载失败, 报错信息: %s", err))
	}

	return err
}

// 上传文件
func SftpUpload(sftp *sftp.Client, srcPath, dstPath string) (err error) {
	srcFile, err := os.Open(srcPath) //本地
	if err != nil {
		return errors.New(fmt.Sprintf("读取平台公钥文件错误，报错信息: %s", err))
	}

	dstFile, err := sftp.Create(dstPath) //远程
	if err != nil {
		return errors.New(fmt.Sprintf("打开远程路径，错误，报错信息: %s", err))
	}

	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("上传文件错误, 报错信息:", err)
			} else {
				break
			}
		}
		_, _ = dstFile.Write(buf[:n])
	}
	return err
}
