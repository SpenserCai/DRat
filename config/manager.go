/*
 * @Author: SpenserCai
 * @Date: 2023-03-05 22:09:44
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-05 22:10:58
 * @Description: file content
 */
package config

import (
	DRatWeb3 "DRat/web3"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type CfgManager struct{}

func (cfg *CfgManager) DecryptEnsConfig(ensdomain string, data string) (map[string]interface{}, error) {
	// 将endsdomain的最后一个点以及后面的内容去掉
	ensdomain = ensdomain[:strings.LastIndex(ensdomain, ".")]
	pwd1 := []byte(ensdomain)
	pwd2 := []byte{}
	for _, v := range pwd1 {
		pwd2 = append(pwd2, v+1)
	}
	pwd := append(pwd1, pwd2...)
	iv := []byte{}
	for i := len(pwd) - 1; i >= 0; i-- {
		iv = append(iv, pwd[i])
	}
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	t, err := aes.NewCipher(pwd)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(t, iv)
	decrypted := make([]byte, len(dataByte))
	blockMode.CryptBlocks(decrypted, dataByte)
	decrypted = decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]
	// 解密结果为json字符串
	var result map[string]interface{}
	err = json.Unmarshal(decrypted, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (cfg *CfgManager) GetConfigFromEns(ethdomain string) (map[string]interface{}, error) {
	ensClient := &DRatWeb3.EnsClient{}
	dConfig, err := ensClient.GetTextRecordByDomain(ethdomain, "description")
	if err != nil {
		return nil, err
	}
	// aes解密数据
	eConfig, err := cfg.DecryptEnsConfig(ethdomain, dConfig)
	if err != nil {
		return nil, err
	}
	return eConfig, nil

}
