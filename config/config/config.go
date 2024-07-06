package config

type Config struct {
	IsDebug   *bool         `yaml:"is_debug" env-required:"true"`
	Listen    ListenConfig  `yaml:"listen"`
	Storage   StorageConfig `yaml:"storage"`
	SecretKey JwtConfig     `yaml:"jwt"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
}

type ListenConfig struct {
	Type   string `yaml:"type" env-default:"port"`
	BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
