/*
 * @Author: SpenserCai
 * @Date: 2023-03-12 21:15:13
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-12 21:18:55
 * @Description: file content
 */
package util

import (
	"os"
	"os/exec"
	"os/user"
)

func GetCurrentUser() (*user.User, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}

func StartDrat() error {
	// 获取当前程序路径
	filePath, err := os.Executable()
	if err != nil {
		return err
	}
	// 异步执行程序
	err = exec.Command(filePath).Start()
	if err != nil {
		return err
	}
	return nil
}

func ExitProcess() {
	os.Exit(0)
}
