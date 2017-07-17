package json

import (
	"encoding/json"
	"io/ioutil"
	"os"

	//"edroity.com/server/common/log"
	"errors"
	//"fmt"
)

//将json config文件转换成json对象
func Parse(filename string, structure interface{}) error {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		return errors.New("加载" + filename + "配置文件出错")
		//return log.Error("加载" + filename + "配置文件出错")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	//fmt.Println("[Parse] data is:", data, "|", string(data))
	if err != nil {
		//return log.Error("加载" + filename + "配置文件出错")
		return errors.New("加载" + filename + "配置文件出错")
	}
	return json.Unmarshal(data, &structure)
}

//将json 字符串转换成json对象
func ParseJsonString(jsonstr string, structure interface{}) error {
	return json.Unmarshal([]byte(jsonstr), &structure)
}

//将json byte转换成json对象
func ParseJsonByte(jsonstr []byte, structure interface{}) error {
	return json.Unmarshal(jsonstr, &structure)
}
