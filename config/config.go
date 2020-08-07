package config

type Server struct {
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mongo     Mongo     `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Qiniu     Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
	ZapConfig ZapConfig `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	GinMode       string `mapstructure:"gin_mode" json:"gin_mode" yaml:"gin_mode"`
	Domain        string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Ishttps       bool   `mapstructure:"ishttps" json:"ishttps" yaml:"ishttps"`
	SslCertFile   string   `mapstructure:"sslcertfile" json:"sslcertfile" yaml:"sslcertfile"`
	SslKeyFile    string   `mapstructure:"sslkeyfile" json:"sslkeyfile" yaml:"sslkeyfile"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password      string `mapstructure:"password" json:"password" yaml:"password"`
	DB            int    `mapstructure:"db" json:"db" yaml:"db"`
	RedisPoolSize int    `mapstructure:"poolsize" json:"pool_size" yaml:"pool_size"`
}

type Mongo struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DbName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
}

type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath   string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

// Config is the struct for logger information
type ZapConfig struct {
	Writers          string `mapstructure:"writers" json:"writers" yaml:"writers"`
	LoggerLevel      string `mapstructure:"logger_level" json:"logger_level" yaml:"logger_level"`
	LoggerFile       string `mapstructure:"logger_file" json:"logger_file" yaml:"logger_file"`
	LoggerWarnFile   string `mapstructure:"logger_warn_file" json:"logger_warn_file" yaml:"logger_warn_file"`
	LoggerErrorFile  string `mapstructure:"logger_error_file" json:"logger_error_file" yaml:"logger_error_file"`
	LogFormatText    bool   `mapstructure:"log_format_text" json:"log_format_text" yaml:"log_format_text"`
	LogRollingPolicy string `mapstructure:"log_rolling_policy" json:"log_rolling_policy" yaml:"log_rolling_policy"`
	LogRotateDate    int    `mapstructure:"log_rotate_date" json:"log_rotate_date" yaml:"log_rotate_date"`
	LogRotateSize    int    `mapstructure:"log_rotate_size" json:"log_rotate_size" yaml:"log_rotate_size"`
	LogBackupCount   int    `mapstructure:"log_backup_count" json:"log_backup_count" yaml:"log_backup_count"`
	LogRollingType   string `mapstructure:"log_rolling_type" json:"log_rolling_type" yaml:"log_rolling_type"`
}
