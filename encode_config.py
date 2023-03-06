'''
Author: SpenserCai
Date: 2023-03-03 13:07:44
version: 
LastEditors: SpenserCai
LastEditTime: 2023-03-06 10:07:11
Description: file content
'''
import Crypto
from Crypto.Cipher import AES
import base64
import sys
import json

# 此文件用于加密存放在ens上的配置文件
# 传入的参数为文件路径，密码
# 传入的文件内容进行aes加密，CBC模式
def aes_encrypt(data):
    # 密码，从第二个参数获得
    pwd1 = sys.argv[2]
    # pwd2，将pwd1每位加上1
    pwd2 = ""
    for i in pwd1:
        pwd2 += chr(ord(i)+1)
    pwd = (pwd1 + pwd2).encode()
    # 向量，密码的倒序
    iv = pwd[::-1]
    # 加密
    cipher = AES.new(pwd, AES.MODE_CBC, iv)
    # 补位
    length = 16
    count = len(data)
    add = length - (count % length)
    data = data + (chr(add) * add).encode()
    # 加密
    encrypted_data = cipher.encrypt(data)
    # base64编码
    return base64.b64encode(encrypted_data)

def aes_decrypt(data):
    # 密码，从第二个参数获得
    pwd1 = sys.argv[2]
    # pwd2，将pwd1每位加上1
    pwd2 = ""
    for i in pwd1:
        pwd2 += chr(ord(i)+1)
    pwd = (pwd1 + pwd2).encode()
    # 向量，密码的倒序
    iv = pwd[::-1]
    # base64解码
    data = base64.b64decode(data)
    # 解密
    cipher = AES.new(pwd, AES.MODE_CBC, iv)
    decrypted_data = cipher.decrypt(data)
    # 去掉补位
    decrypted_data = decrypted_data[:-decrypted_data[-1]]
    return decrypted_data

if __name__ == '__main__':
    # 读取文件内容
    with open(sys.argv[1], 'r') as f:
        data = f.read()
    # 加密
    encrypted_data = aes_encrypt(data.encode("utf-8"))
    # 写入到文件中
    print(encrypted_data.decode())
    # 解密
    decrypted_data = aes_decrypt(encrypted_data)
    print(decrypted_data.decode("utf-8"))
    

