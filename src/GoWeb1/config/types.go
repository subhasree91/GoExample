package config

type DBconfig struct {
	DBname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type Appconfig struct {
	PortNo string   `json:"portNo"`
	DB     DBconfig `json:"db"`
}
