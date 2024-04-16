package config

type Config struct {
	Email struct {
		Server    string `yaml:"server"`
		Port      int    `yaml:"port"`
		User      string `yaml:"account"`
		Password  string `yaml:"password"`
		AdminMail string `yaml:"admin_email"`
	} `yaml:"email"`
	Auth struct {
		Key string `yaml:"key"`
	} `yaml:"secret"`
	Domain struct {
		Host string `yaml:"host"`
		Api  string `yaml:"api"`
	} `yaml:"domain"`
}
