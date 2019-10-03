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
	DB *gorm.DB
}

func NewClient()( c *Client){

	db, err := gorm.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	c= &Client{
		DB: db,
	}
	c.DB.AutoMigrate(&Ignored{})
	return c
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
	 err = c.DB.Create(&Ignored{
		Name: name,
	}).Error
	return err
}

func (c *Client)Remove(name string)(err error){
	ig:=&Ignored{}
	err = c.DB.Where("name = ?", name).Delete(ig).Error
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
	err= c.DB.Find(&users).Error
	return users, err
}