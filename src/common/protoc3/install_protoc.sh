#!/bin/bash
set -e
#下载protoc3
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip -c -P /opt

if [! -d "/opt/protoc" ];then
    mkdir /opt/protoc
fi
#下载zip
apt-get -y install zip

#解压zip包
unzip /opt/protoc-3.0.0-linux-x86_64.zip -d /opt/protoc
chmod +x /opt/protoc/ -R

#配置环境变量
echo export PATH=$PATH:/opt/protoc/bin >>/etc/profile

source /etc/profile

exit 0