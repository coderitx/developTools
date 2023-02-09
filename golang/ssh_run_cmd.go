package tools


import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

// RunCmd 远端执行命令
// cfg: 远端ssh相关配置
// sshAddr: ssh地址
// cmdLine: 执行的命令 
func RunCmd(cfg *ssh.ClientConfig, sshAddr, cmdLine string) error {
	client, err := ssh.Dial("tcp", sshAddr, cfg)
	if err != nil {
		return fmt.Errorf("ssh连接目标%s失败:%v", sshAddr, err)
	}
	defer client.Close()


	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("开启session addr:%s失败:%v", sshAddr, err)
	}
	// 设置 session的 tty 配置
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	err = session.RequestPty("xterm", 24, 80, modes)
	if err != nil {
		return fmt.Errorf("设置TTY:%s失败:%v", sshAddr, err)
	}
	defer session.Close()

	// stdout stderr
	var b, eb bytes.Buffer
	session.Stdout = &b
	session.Stderr = &eb
	err = session.Run(cmdLine)
	log.Printf("HOST:[%s]  CMD:[%s] Err:[%v] OUT:[%s] OUT_Err:[%s]\n", sshAddr, cmdLine, err, b.String(), eb.String())
	if err != nil {
		return fmt.Errorf("ssh执行cmd:[ %s ]失败:%v", cmdLine, err)
	}
	return nil
}