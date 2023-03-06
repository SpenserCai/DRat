/*
 * @Author: SpenserCai
 * @Date: 2023-03-05 21:35:06
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-05 21:37:40
 * @Description: file content
 */
package message

import (
	DRatConfig "DRat/config"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

func InitBot() {
	// 机器人被@时的回复
	teleBotSettings := tele.Settings{
		Token:  DRatConfig.TELBOT_TOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	if DRatConfig.LOCAL_PROXY_PORT != 0 {
		teleBotSettings.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(&url.URL{
					Scheme: "socks5",
					Host:   "127.0.0.1:" + strconv.Itoa(DRatConfig.LOCAL_PROXY_PORT),
				}),
			},
		}
	}
	tempBot, err := tele.NewBot(teleBotSettings)
	if err != nil {
		fmt.Println(err)
		return
	}
	DRatConfig.TelBot = tempBot
	TeleBotCommand()
}

// 发送纯文本消息
func TeleSendMessge(message string) {
	DRatConfig.TelBot.Send(tele.ChatID(-DRatConfig.TELBOT_CHAT_ID), message)
}

// 发送支持文件和markdown的消息
func TeleSendMarkDownMessage(message string) {
	// 发送markdown
	DRatConfig.TelBot.Send(tele.ChatID(-DRatConfig.TELBOT_CHAT_ID), message, &tele.SendOptions{
		ParseMode: tele.ModeMarkdown,
	})

}
