package main
/**
 author j.yao
 tags: 最烂的 go 代码


 */
import (
	"fmt"
	"github.com/opesun/goquery"
	"strings"
	"bytes"

	"encoding/json"

	"io/ioutil"
	"strconv"
/*	"math"
	"os"

	"bufio"
	"io"*/
	"runtime"
	"os"
)

type Config struct {
	Server  string
	Server_port int64
	Local_address  string
	Local_port int64
	Password string
	Timeout int64
	Method string
	Fast_open bool
	Workers  int64
}
/*var FdMap = map[string]string{}  // {}为初始化成空
func readFile(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &FdMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}
	return FdMap, nil
}*/

func main() {
	var url  = "http://www.ishadowsocks.com/#free"

	section,err := goquery.ParseUrl(url)

	Host0:=Config{Workers:1,Fast_open:false,Local_address:"127.0.0.1",Local_port:1080,Timeout:300}
	//Host1:=Config{Workers:1,Fast_open:false,Local_address:"127.0.0.1",Local_port:1080,Timeout:300}
	//Host2:=Config{Workers:1,Fast_open:false,Local_address:"127.0.0.1",Local_port:1080,Timeout:300}


	if err != nil {
		panic(err)
	} else {
		//section_info := section.Find("#free .container .row .col-lg-4 ").Each(func(i int, contentSelection *goquery.Selection) {
           	//
		//
		//})// 获取 section id=free
		section_info := section.Find("#free .container .row .col-lg-4 h4").HtmlAll()
		//section_info1 := section.Find("#free h4").HtmlAll()// 获取 section id=free

	/*	for i:=0;i<len(section_info)/6;i++ {


			for j := 0; j < len(section_info)/3; j++ {
				//fmt.Println(i,j+6*i)
				tmp:=strings.Split(section_info[j+6*i]," ")[0]

				tmp1:= strings.Replace(tmp,"<font","",-1)
 				list_.PushBack(strings.Split(tmp1,":"))
			}




		}*/
  		strs:=bytes.Buffer{}
		for i:=0;i<len(section_info);i++ {
			tmp:= strings.Replace(strings.Split(section_info[i]," ")[0],"<font","",-1)
			tmp1:=strings.Split(tmp,":")

			if len(tmp1) ==2 {
				strs.WriteString((tmp1[1]))
				strs.WriteString(" ")
			}else {
				strs.WriteString(tmp1[0])
				strs.WriteString(" ")
			}
			fmt.Print(string("处理得到的数据。。。。。。"))
			fmt.Println(tmp1)
		}
		fmt.Println(strs.String())
		 datas:=strs.String()
		 data_tmp:= strings.Split(datas," ")
		// 得到 账号 端口 及 密码数据
		fmt.Println("-----得到 账号 端口- data_tmp----------------")
		fmt.Println(data_tmp)


         var  config Config
		data,err:=ioutil.ReadFile("/etc/shadowsocks/config.json")
		if err != nil {
			fmt.Println("ReadERr",err.Error())
		}
		json.Unmarshal(data,&config)
		fmt.Println("------获取到系统路径下/etc/shadowsocks/config.json 的配置内容为： --------")

		fmt.Println(config)
		fmt.Println("----------------------")
		//
		Host0.Server=data_tmp[0]

		Host0.Server_port,err=strconv.ParseInt(data_tmp[1],10,64)

		if err != nil {
			panic(err)
		}
		Host0.Password=data_tmp[2]
		Host0.Method=data_tmp[3]

		fmt.Println("----------------------")
		fmt.Println()
		fmt.Println(Host0)
		fmt.Println("----------------------")


		//
		//Host1.Server=data_tmp[4]
		//fmt.Println(data_tmp[4] + "////")
		//fmt.Println(len(data_tmp))
		//fmt.Println("----------------------")
		//Host1.Server_port,err=strconv.ParseInt(data_tmp[5],10,64)
		//
		//if err != nil {
		//	panic(err)
		//}
		//Host1.Password=data_tmp[6]
		//Host1.Method=data_tmp[7]
		//
		//fmt.Println("----------------------")
		//fmt.Println()
		//fmt.Println(Host1)
		//fmt.Println("----------------------")
		//
		//Host2.Server=data_tmp[8]
		//
		//Host2.Server_port,err=strconv.ParseInt(data_tmp[9],10,64)
		//
		//if err != nil {
		//	panic(err)
		//}
		//Host2.Password=data_tmp[10]
		//Host2.Method=data_tmp[11]
		//
		//fmt.Println("----------------------")
		//fmt.Println()
		//fmt.Println(Host2)
		//fmt.Println("----------------------")

		data1, err:=json.Marshal(Host0)
		if err != nil {
			panic(err)
		}
		//data2, err:=json.Marshal(Host1)
		//if err != nil {
		//	panic(err)
		//}
		//data3, err:=json.Marshal(Host2)
		//if err != nil {
		//	panic(err)
		//}
		//
		fmt.Println("----------------------")
		fmt.Println(string(data1))
		//fmt.Println(string(data2))
		//fmt.Println(string(data3))
		fmt.Println("----------------------")
		//  get  姿势 获取当前运行文件
		_, filename, line, _ := runtime.Caller(0)
		fmt.Println("你正在执行的程序路径是："+filename)
		fmt.Print("当前执行到第：")
		fmt.Print(line)
		fmt.Println(" 行！")
		currentDir,err:=os.Getwd()
		fmt.Println("---配置文件将存放到以下目录：-----")
		fmt.Println(currentDir)
		fmt.Println("----------------------")
 		data_to_lower:= strings.ToLower(string(data1))
		fmt.Println("----------------------")
		fmt.Println("生成配置文件---->>  匹配第二个和第三个服务器数据发送错乱，暂时不对其处理")
		fmt.Println(data_to_lower)
		ioutil.WriteFile(currentDir+"/config.json",[]byte (data_to_lower),os.ModePerm)
	}
}