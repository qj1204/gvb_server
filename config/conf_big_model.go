package config

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disabled" json:"disabled"`
}

type Setting struct {
	Name      string `yaml:"name" json:"name"`
	Enable    bool   `yaml:"enable" json:"enable"`
	Order     int    `json:"order" yaml:"order"` // 菜单的序号
	ApiKey    string `yaml:"api-key" json:"api_key"`
	ApiSecret string `yaml:"api-secret" json:"api_secret"`
	Title     string `yaml:"title" json:"title"`
	Prompt    string `yaml:"prompt" json:"prompt"` // 加强提示词
	Slogan    string `yaml:"slogan" json:"slogan"` // slogan
}

type SessionSetting struct {
	ChatScope    int `yaml:"chat-scope" json:"chat_scope"`       // 对话的积分消耗
	SessionScope int `yaml:"session-scope" json:"session_scope"` // 会话的积分消耗
	DayScope     int `yaml:"day-scope" json:"day_scope"`         // 每日可以领取的积分
}

type BigModel struct {
	Setting        Setting        `yaml:"setting"`
	ModelList      []ModelOption  `yaml:"model-list"`
	SessionSetting SessionSetting `yaml:"session-setting"`
}
