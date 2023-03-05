/*
 * @Author: SpenserCai
 * @Date: 2023-03-02 23:17:29
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-05 21:27:06
 * @Description: file content
 */
package web3

import (
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

type EnsClient struct{}

// 通过域名获取RECORDS
func (e *EnsClient) GetTextRecordByDomain(domain string, textName string) (string, error) {
	// 通过以太坊域名获取records
	// 1. 连接以太坊goerli测试网络
	client, err := ethclient.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		return "", err
	}
	// 2. 获取域名解析器
	resolver, err := ens.NewResolver(client, domain)
	if err != nil {
		return "", err
	}
	// 3. 获取域名解析器的内容
	textRecord, err := resolver.Text(textName)
	if err != nil {
		return "", err
	}
	return textRecord, nil
}
