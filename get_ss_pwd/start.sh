#! /bin/sh

echo "正在下载所需软件 shadowsocks"
sleep 3
sudo apt-get install python-pip &&\
sudo apt-get install shadowsocks -y && \
sudo  pip install  request &&\
sudo apt-get install python-m2crypto

path_Dir=$PWD
arch=`uname --machine`

echo $path_Dir
echo 系统架构 $arch

if  [  `echo $arch | grep -e arm`  ]; then
    echo "执行 for_arm 版本"
	sleep 1
	sudo $PWD/get_ss_pwd_arm
 else
   echo "执行 for_pc 版本"
   sleep 1
    sudo $PWD/get_ss_pwd_for_pc
fi

sudo  /usr/bin/ssserver -c $path_Dir/config.json