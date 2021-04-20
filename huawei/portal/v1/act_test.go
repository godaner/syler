package v1

import (
	"testing"
	"net"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

//func TestChallange(t *testing.T) {
//	ip := net.IPv4(192, 168, 56, 2)
//	fmt.Println(NewChallenge(ip, "it is a secret"))
//}
//
//func TestAuth(t *testing.T) {
//	ip := net.IPv4(192, 168, 56, 2)
//	fmt.Println(NewAuth(ip, "it is a secret", []byte("刘铭"), []byte("456"), uint16(1234)))
//}
//
//func TestRunChallange(t *testing.T) {
//	Challenge(net.IPv4(192, 168, 1, 1), "it is a secret", net.IPv4(192, 168, 56, 2))
//}
//
//func TestUnmarshal(t *testing.T) {
//	msg := Unmarshall([]byte{0x01, 0x02, 0x00, 0x00, 0x6f, 0x3c, 0x00, 0x06, 0xc0, 0xa8, 0x0a, 0xfe, 0x00, 0x00, 0x00, 0x01, 0x03, 0x12, 0xef, 0x47, 0x25, 0x3d, 0xc5, 0x19, 0x41, 0xb7, 0x63, 0x97, 0x35, 0x07, 0x75, 0xe7, 0x3d, 0x95})
//	fmt.Println(msg)
//}



func Test_int(t *testing.T) {

	bytes := []byte{0x43,0x38,0x3a,0x32,0x31, 0x3a,0x35, 0x38,0x3a, 0x33,0x43, 0x3a,0x41,0x37,0x3a, 0x42,0x37}
	mac,err := net.ParseMAC(string(bytes))
	if err==nil{
		t.Log("---------")
		t.Log(mac)
	}else{
		t.Log(err)
	}

	s:="00:13:32:10:32:40:i-town"
	t.Log(parse(s))



}

func parse(s string )string{
	 d := strings.Split(s,":")
	mac:=""
	for i:=0;i<6;i++{
		if i==5{
			mac += d[i]
		}else{
			mac += d[i]+":"
		}

	}
	return mac
}

func TestT_Attr_Byte(t *testing.T) {
	JsonParse := NewJsonStruct()
	v := make(map[string]interface{})
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("/version.txt", &v)
	fmt.Println(v["versions"])
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func TestT_Attr_Length(t *testing.T) {
	coremap := sync.Map{}

	coremap.Store("123",make(chan map[string]interface{}))
	mm := make(map[string]string)
	mm["12:12:12:12:12:12"]="success"
	coremap.Store("124",logcache{mm})
	tt,_:=coremap.Load("124")


	switch tt.(type){
	case chan map[string]interface{} :
		t.Log("chan ")
	case logcache :
		t.Log("struct ")
		for k,v:=range tt.(logcache).m{
			t.Log("k:",k,"v:",v)
		}


	}
}

type logcache struct {
	m   map[string]string
}
