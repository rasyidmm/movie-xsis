package server

type ServerList struct {
	Rest Server `yaml:"rest" json:"rest"`
}

type Server struct {
	Name string `yaml:"name" json:"name"`
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}
