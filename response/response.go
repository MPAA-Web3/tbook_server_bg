package response

// Task struct to represent a user's task.
type Task struct {
	Name   string `json:"name"`   // 任务名称
	Status bool   `json:"status"` // 是否完成任务
}

// UserTasks struct to represent the tasks of a user.
type UserTasks struct {
	UserID string `json:"user_id"` // 用户的唯一标识符
	Task   []Task `json:"task"`    // 用户的任务列表
}

type Rewards struct {
	UserID uint   `json:"user_id"` // 用户 ID
	Name   string `json:"name"`    // 奖励名称
	Type   string `json:"type"`    // 奖励类型
	Value  string `json:"value"`   // 奖励值
	Game   string `json:"game"`    // 所属游戏
	Time   string `json:"time"`    // 获得奖励的时间
}

type InviteeTasks struct {
	UserID   string `json:"user_id"`  // 用户ID
	Num      int    `json:"num"`      // 邀请人数
	Wallet   string `json:"wallet"`   // 邀请绑定钱包的人数
	Twitter  string `json:"twitter"`  // 邀请关注Twitter 数量
	Telegram string `json:"telegram"` // 邀请关注Telegram 账号
}

type Purchase struct {
	UserID string  `json:"user_id"` // 用户ID
	Time   string  `json:"time"`    // 购买时间
	Points float64 `json:"points"`  // 购买所获得的积分
	Level  string  `json:"level"`   // 用户当前的等级
}

type Order struct {
	UserID string  `json:"user_id"`
	Time   string  `json:"time"`
	Ton    float64 `json:"points"`
	Value  float64 `json:"level"`
}

// ImgList
type ImgList struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
