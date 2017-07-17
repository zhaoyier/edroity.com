package static

import (
	"edroity.com/server/common/log"
	"edroity.com/server/common/conf/json"
	"fmt"
)

var ItemData map[string]Item

func init()  {
	if err := json.Parse(ItemJson, &ItemData); err != nil {
		fmt.Println("[init] error:", err)
	}

	log.Debug("[init]数据初始化完毕.")
}

func GetItem(id string) Item {
	return ItemData[id]
}
