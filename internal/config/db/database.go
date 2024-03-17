package db

type DatabaseList struct {
	Movie Database `yaml:"movie" json:"movie"`
}
type Database struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Dbname   string `yaml:"dbname" json:"dbname"`
}
