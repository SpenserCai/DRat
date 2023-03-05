/*
 * @Author: SpenserCai
 * @Date: 2023-03-01 12:57:08
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-03-05 21:31:26
 * @Description: file content
 */
package util

import "golang.org/x/text/encoding/simplifiedchinese"

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
