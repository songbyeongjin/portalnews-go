package db
import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Connector struct{
	Dialect string `yaml:"dialect"`
	Host string `yaml:"host"`
	Dbname string `yaml:"db_name"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

var Instance *gorm.DB = nil


func (d Connector) GetConnectString() string{
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",d.Host, d.Port, d.User, d.Dbname, d.Password)
}

func (d Connector) SetDbInstance() error{
	var err error
	Instance, err = gorm.Open(d.Dialect, d.GetConnectString())
	return err
}