package infra_clean

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DbConnector struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"db_name"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (d *DbConnector) GetConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", d.Host, d.Port, d.User, d.Dbname, d.Password)
}

//Set Db information from yaml
func getDbConnector() (*DbConnector, error) {
	//temporary path for debug mode
	buf, err := ioutil.ReadFile(`C:\Users\SONG\Documents\study\go\src\portal_news\db_info.yaml`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var dbConnector DbConnector

	err = yaml.Unmarshal(buf, &dbConnector)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dbConnector, nil
}



type DbHandler struct {
	Conn *gorm.DB
}

func NewDbHandler() *DbHandler {
	dbHandler := &DbHandler{}
	DbConnector, err := getDbConnector()

	if err != nil{
		log.Print(err)
		return nil
	}
	dbHandler.Conn, err = gorm.Open(DbConnector.Dialect, DbConnector.GetConnectString())
	if err != nil{
		log.Print(err)
		return nil
	}

	return dbHandler
}
