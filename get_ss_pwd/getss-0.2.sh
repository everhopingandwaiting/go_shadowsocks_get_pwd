#!/bin/bash
############################
#
############################

unset page_url
page_url='http://www.ishadowsocks.com/'
echo "from ${page_url}"

#配置文件目录
config_dir=$HOME/.ishadowsocks

example_config=${config_dir}/ishadowsocks.json.example
config_A=${config_dir}/ishadowsocks-A.json
config_B=${config_dir}/ishadowsocks-B.json
config_C=${config_dir}/ishadowsocks-C.json

#数据
set data

############################
#
#
#############################
arr_config=(
$config_A
$config_B
$config_C
)

arr_keywords=(
"server"
"server_port"
"password"
"method"
)

############################
#creat config dir
#
############################
function _creat_dir(){
    if [ ! -d "$config_dir" ];then
        mkdir -pv "${config_dir}"
    fi
}

############################
#creat example config
#
############################
function _creat_example_config(){
    if [ ! -f "${example_config}" ];then
        $(echo -e '{
            "server":"8.8.8.8",
            "server_port":8388,
            "local_port":1080,
            "password":"123456",
            "timeout":300,
            "method":"aes-256-cfb"
        }' > ${example_config})
    fi
}

############################
#check file
#
############################
function _check_file(){

    _max=${#arr_config[@]}

    for((_tmp=0;_tmp<$_max;_tmp++))
    do
        if [ ! -f "${arr_config[$_tmp]}" ];then

            `cp ${example_config} ${arr_config[$_tmp]}`
        fi
    done
}

############################
#get data
#
############################
function _get_data(){
    echo "downloading……"

    data=`curl -s ${page_url} | sed -n 's/.*<h4>\([^<]*\)<\/h4>.*/\1/p' | sed -n '1,12p'`

    `echo ${data} > ${config_dir}/ishadowsocks.log`
}

############################
#update time
#
############################
update_time=`date`


############################
#
#
############################
function _replace()
{
    tmp=0;
    max=3;
    while [ $tmp -lt $max ]
    do
        for((j=0;j<4;j++))
        do
            _ip=$(($tmp*4+$j+1))

            new=`cat ${config_dir}/ishadowsocks.log | awk '{print $+"'$_ip'"}' | awk -F : '{print $2}'`

            if [ $j -ne 1 ];then
                old=`cat ${arr_config[$tmp]} | sed -n '/'${arr_keywords[$j]}'/p' | awk -F \" '{print $4}'`
            else
                old=`cat ${arr_config[$temp]} | sed -n '/'${arr_keywords[$j]}'/p' | awk -F : '{print $2}' | awk -F , '{print $1}'`
            fi

            `sed -i "s/$old/$new/g" ${arr_config[$tmp]}`

            if [ $j -eq 0 ];then
                echo "replacing ${arr_config[$tmp]}"
            fi
        done
        tmp=$(($tmp+1))
    done

}

function _del_log(){
    if [ -f ${config_dir}/ishadowsocks.json.example ];then
        `rm -f ${config_dir}/ishadowsocks.json.example`
    fi
    if [ -f ${config_dir}/ishadowsocks.log ];then
        `rm -f ${config_dir}/ishadowsocks.log`
    fi
}
############################
#
#
############################
function _main_(){
    _creat_dir
    _creat_example_config
    _check_file
    _get_data
    _replace
    _del_log
    echo "Done!"
}

############################
#
#
############################
_main_
