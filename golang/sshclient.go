package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"os/exec"
	"strings"
	"time"
)

type SSHClintHandler struct {
	Host     string
	User     string
	Password string
	Port     int
	ShType   string
	Shell    string
}

// SSHClient 执行远程命令
func SSHClient(shCli *SSHClintHandler) {
	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            shCli.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if shCli.ShType == "" || shCli.ShType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(shCli.Password)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", shCli.Host, shCli.Port)
	fmt.Println("addr: ", addr)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		logrus.Error("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		logrus.Error("创建ssh session 失败", err)
	}
	defer session.Close()

	stdout, _ := session.StdoutPipe()
	stderr, _ := session.StderrPipe()
	go asyncLog(io.NopCloser(stdout))
	go asyncLog(io.NopCloser(stderr))

	//执行远程命令
	err = session.Start(shCli.Shell)
	if err != nil {
		logrus.Error("远程执行cmd 失败", err)
	}
	if err := session.Wait(); err != nil {
		logrus.Errorf("wait run shell error: %v", err)
		return
	}
	//logrus.Info("命令输出:", string(combo))
	return
}

// ExecuteShell 运行本地shell命令
func ExecuteShell(shell string) error {
	fmt.Println("----shell------", shell)
	cmd := exec.Command("bash", "-c", shell)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		logrus.Errorf("run %v error: %v", shell, err)
		return err
	}
	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		logrus.Errorf("wait shell run error: %v", err)
		return err
	}
	return nil
}

func asyncLog(reader io.ReadCloser) error {
	bucket := make([]byte, 0)
	buffer := make([]byte, 100)
	for {
		num, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				err = nil
			}
			return err
		}
		if num > 0 {
			line := ""
			bucket = append(bucket, buffer[:num]...)
			tmp := string(bucket)
			if strings.Contains(tmp, "\n") {
				ts := strings.Split(tmp, "\n")
				if len(ts) > 1 {
					line = strings.Join(ts[:len(ts)-1], "\n")
					bucket = []byte(ts[len(ts)-1]) //不够整行的以后再处理
				} else {
					line = ts[0]
					bucket = bucket[:0]
				}
				fmt.Printf("%s\n", line)
			}

		}
	}
	return nil
}
