package utils

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
	"strings"
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
	defer sftp.Close()
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

//  对比目标文件是否已经存在平台公钥
func ContainPubkey(keyPath string, downkeyPath string) (err error, diff bool) {
	var source_data string
	keydata, err := os.Open(keyPath)
	defer keydata.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("平台公钥读取错误，请检查是否生成平台密钥对, 报错信息: %s", err)), false
	} else {
		scanner := bufio.NewScanner(keydata)
		for scanner.Scan() {
			if scanner.Text() != "" {
				source_data = scanner.Text()
			}
		}
	}

	file, err := os.Open(downkeyPath)
	defer file.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("远程公钥文件本地读取错误, 报错信息: %s", err)), false
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.EqualFold(scanner.Text(), source_data) {
				return err, true
			}
		}
		return err, false
	}
}

// 合并公钥文件
func MergePublicKey(keyPath string, downkeyPath string) (err error) {
	var source_data string
	keydata, err := os.Open(keyPath)
	defer keydata.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("平台公钥读取错误，请检查是否生成平台密钥对, 报错信息: %s", err))
	} else {
		scanner := bufio.NewScanner(keydata)
		for scanner.Scan() {
			if scanner.Text() != "" {
				source_data = scanner.Text()
			}
		}
	}

	file, err := os.OpenFile(downkeyPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("本地远程公钥文件读取错误, 报错信息: %s", err))
	} else {
		file.Write([]byte(source_data))
	}
	return err
}
