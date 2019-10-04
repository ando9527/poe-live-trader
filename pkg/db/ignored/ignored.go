package ignored

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

func (c *Client)Connect(name string) (e error) {

	db, err := gorm.Open("sqlite3", name)
	if err != nil {
		return err
	}
	c.db =db
	return nil
}

func (c *Client)Migration() (e error){
	return c.db.AutoMigrate(&Ignored{}).Error
}

func (c *Client)GetIgnoreMap()(m map[string]bool,err error){
	users,err:=c.GetAll()
	if err != nil {
		return nil, err
	}
	m = map[string]bool{}
	for _,v :=range users{
		m[v.Name]=true
	}

	return m, nil
}

func (c *Client)Add(name string) (err error){
	 err = c.db.Create(&Ignored{
		Name: name,
	}).Error
	return err
}

func (c *Client)Remove(name string)(err error){
	ig:=&Ignored{}
	err = c.db.Where("name = ?", name).Delete(ig).Error
	if err != nil {
		return err
	}

	if ig.Name==""{
		logrus.Warnf("%s not exist", name)
	}

	return nil
}

func (c *Client)GetAll()(users []Ignored, err error){
	users=[]Ignored{}
	err= c.db.Find(&users).Error
	return users, err
}