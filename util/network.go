/*
 * @Author: SpenserCai
 * @Date: 2023-03-12 21:09:45
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-12 21:12:36
 * @Description: file content
 */
package util

import (
	"io"
	"net/http"
)

// 获取公网ip
func GetPublicIp() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
