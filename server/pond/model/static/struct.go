package static

type Item struct {
	SystemOpenID int64  `json:"SystemOpenID"`
	Desc         string `json:"desc"`
	ExtraInfo    string `json:"extra_info"`
	GetCondition []struct {
		Data string `json:"data"`
		ID   int64  `json:"id"`
	} `json:"get_condition"`
	Icon     string `json:"icon"`
	ID       int64  `json:"id"`
	Info     string `json:"info"`
	Key      string `json:"key"`
	Level    int64  `json:"level"`
	Maxnum   int64  `json:"maxnum"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quality  int64  `json:"quality"`
	StackNum int64  `json:"stack_num"`
	Type     int64  `json:"type"`
	UseMaxLv int64  `json:"useMaxLv"`
	UseMinLv int64  `json:"useMinLv"`
	UseSub   []struct {
		ID  int64 `json:"id"`
		Num int64 `json:"num"`
	} `json:"useSub"`
	UseState int64 `json:"use_state"`
}