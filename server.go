package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Agent struct {
	Name  string `json:"agentName"`
	Image string `json:"agentImage"`
}

type AgentsResponse struct {
	Agents []Agent `json:"agents"`
}

var agents = map[int]string{
	1:   "荒芜拉普兰德",
	2:   "忍冬",
	3:   "弑君者",
	4:   "维娜·维多利亚",
	5:   "玛露西尔",
	6:   "佩佩",
	7:   "娜仁图亚",
	8:   "妮芙",
	9:   "乌尔比安",
	10:  "维什戴尔",
	11:  "逻各斯",
	12:  "魔王",
	13:  "阿斯卡纶",
	14:  "艾拉",
	15:  "黍",
	16:  "左乐",
	17:  "莱伊",
	18:  "锏",
	19:  "塑心",
	20:  "薇薇安娜",
	21:  "止颂",
	22:  "赫德雷",
	23:  "涤火杰西卡",
	24:  "纯烬艾雅法拉",
	25:  "琳琅诗怀雅",
	26:  "提丰",
	27:  "圣约送葬人",
	28:  "缪尔赛思",
	29:  "霍尔海雅",
	30:  "淬羽赫默",
	31:  "伊内丝",
	32:  "麒麟R夜刀",
	33:  "仇白",
	34:  "重岳",
	35:  "林",
	36:  "焰影苇草",
	37:  "缄默德克萨斯",
	38:  "斥罪",
	39:  "伺夜",
	40:  "白铁",
	41:  "玛恩纳",
	42:  "百炼嘉维尔",
	43:  "鸿雪",
	44:  "多萝西",
	45:  "黑键",
	46:  "归溟幽灵鲨",
	47:  "艾丽妮",
	48:  "流明",
	49:  "号角",
	50:  "菲亚梅塔",
	51:  "澄闪",
	52:  "令",
	53:  "老鲤",
	54:  "灵知",
	55:  "耀骑士临光",
	56:  "焰尾",
	57:  "远牙",
	58:  "琴柳",
	59:  "假日威龙陈",
	60:  "水月",
	61:  "帕拉斯",
	62:  "卡涅利安",
	63:  "浊心斯卡蒂",
	64:  "凯尔希",
	65:  "歌蕾蒂娅",
	66:  "异客",
	67:  "灰烬",
	68:  "夕",
	69:  "嵯峨",
	70:  "空弦",
	71:  "山",
	72:  "迷迭香",
	73:  "泥岩",
	74:  "瑕光",
	75:  "史尔特尔",
	76:  "森蚺",
	77:  "棘刺",
	78:  "铃兰",
	79:  "早露",
	80:  "W",
	81:  "温蒂",
	82:  "傀影",
	83:  "风笛",
	84:  "刻俄柏",
	85:  "年",
	86:  "阿",
	87:  "煌",
	88:  "莫斯提马",
	89:  "麦哲伦",
	90:  "赫拉格",
	91:  "黑",
	92:  "陈",
	93:  "斯卡蒂",
	94:  "银灰",
	95:  "塞雷娅",
	96:  "星熊",
	97:  "夜莺",
	98:  "闪灵",
	99:  "安洁莉娜",
	100: "艾雅法拉",
	101: "伊芙利特",
	102: "推进之王",
	103: "能天使",
}
var agentImages = make(map[int]string)

func init() {
	for i := 1; i <= len(agents); i++ {
		agentImages[i] = fmt.Sprintf("assests\\arknight_image\\%d.png", i)
	}
}

func randomAgentHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	agentsResponse := AgentsResponse{}

	// 从前端获取要返回的干员数量
	query := r.URL.Query()
	countStr := query.Get("count")
	var count int
	if countStr != "" {
		// 尝试将字符串转换为整数
		if tmpCount, err := strconv.Atoi(countStr); err == nil {
			count = tmpCount
		} else {
			// 如果转换失败，返回错误信息
			http.Error(w, "Invalid count parameter", http.StatusBadRequest)
			return
		}
	} else {
		// 如果没有提供count参数，默认为12
		count = 12
	}

	// 创建一个包含所有干员索引的切片
	var availableIndices []int
	for index := range agents {
		availableIndices = append(availableIndices, index)
	}

	// 确保只选取12个不重复的干员
	for len(agentsResponse.Agents) < count {
		randIndex := rand.Intn(len(availableIndices))
		agentIndex := availableIndices[randIndex]

		// 将随机选择的干员添加到响应中
		agentName := agents[agentIndex]
		agentImage := agentImages[agentIndex]
		agentResponse := Agent{
			Name:  agentName,
			Image: agentImage,
		}
		agentsResponse.Agents = append(agentsResponse.Agents, agentResponse)

		// 从候选列表中移除已选择的干员
		availableIndices[randIndex] = availableIndices[len(availableIndices)-1]
		availableIndices = availableIndices[:len(availableIndices)-1]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agentsResponse)
}

func main() {
	http.Handle("/assests/", http.StripPrefix("/assests/", http.FileServer(http.Dir("./assests/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	})
	http.HandleFunc("/random", randomAgentHandler)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
