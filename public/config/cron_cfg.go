package config

type CronCfg struct {
	Name     string     `json:"Name"`
	QuitWait int        `json:"Wait,default=25"`
	Cron     [][]string `json:"Cron"`
}
