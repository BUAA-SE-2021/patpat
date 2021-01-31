package initialize

import (
	"io/ioutil"
	"os"
	"patpat/global"
	"patpat/model"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func FetchMySQLConfig() (host string, port string, username string, password string, db string) {
	t := MySQLConfig{}
	fin, err := os.Open("mysql.yaml")
	if err != nil {
		panic(err)
	}
	cin, _ := ioutil.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}
	return t.Host, t.Port, t.Username, t.Password, t.Database
}

func InitMySQL() {
	host, port, username, password, database := FetchMySQLConfig()
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.JudgeResultUsual{},
		&model.JudgeResultFormal{},
	)
}
