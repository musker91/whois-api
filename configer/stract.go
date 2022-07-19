package configer

type ConfigerStruct struct {
	AppMode string `default:"development"`
	Serve   ServeConfiger
	Redis   RedisConfiger
}

type ServeConfiger struct {
	Port         string `default:"8091"`
	LogLevel     string `default:"debug"`
	LogType      string `dedfault:"json"`
	LogOutPath   string `default:"console"`
	LogSaveDays  int    `default:"7"`
	LogSplitTime int    `default:"24"`
}

type RedisConfiger struct {
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"6379"`
	Password string
}
