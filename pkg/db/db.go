package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)
type Ignored struct {
	ID   int64
	Name string
}

type Client struct{
	db *gorm.DB
}

func NewClient()(c *Client){
	return &Client{
		db: nil,
	}
}

func (c *Client) IsIgnored(id string) bool {
	logrus.Debug("checking ignored in database")
	var out []Ignored
	err := c.db.Where("name = ?", id).Find(&out).Error
	if err != nil {
		logrus.Error("database error", err)
	}
	if len(out)>=1{
		logrus.Debug("user ID is ignored in database, ", id)
		return true
	}
	return false
}




func (c *Client)Connect(name string) {

	db, err := gorm.Open("sqlite3", name)
	if err != nil {
		logrus.Fatal(err)
	}
	c.db =db
}

func (c *Client)Migration() {
	err:= c.db.AutoMigrate(&Ignored{}).Error
	if err != nil {
		logrus.Fatal(err)
	}
}

func (c *Client)GetIgnoreMap()(m map[string]bool,err error){
	users,err:=c.GetIgnoredList()
	if err != nil {
		return nil, err
	}
	m = map[string]bool{}
	for _,v :=range users{
		m[v.Name]=true
	}

	return m, nil
}

func (c *Client) AddIgnored(name string) (err error){
	 err = c.db.Create(&Ignored{
		Name: name,
	}).Error
	return err
}

func (c *Client) RemoveIgnored(name string)(err error){
	err = c.db.Where("name = ?", name).Delete(Ignored{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetIgnoredList()(users []Ignored, err error){
	users=[]Ignored{}
	err= c.db.Find(&users).Error
	return users, err
}