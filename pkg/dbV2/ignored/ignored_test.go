package ignored

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
)

func fakeClient()(c *Client){
	c=NewClient()
	e := c.Connect(filepath.Join(os.TempDir(), "gorm.db"))
	if e != nil {
		panic(e)
	}
	c.db.DropTableIfExists(&Ignored{})
	e = c.Migration()
	if e != nil {
		panic(e)
	}
	err := c.Add("yolo")
	if err != nil {
		panic(err)
	}
	return c
}

func TestClient_Query(t *testing.T) {
	c:=fakeClient()
	defer c.db.Close()

	users, e := c.GetAll()
	if e != nil {
		t.Fatal(e)
	}
	is:=is.New(t)
	is.Equal(len(users), 1)
}

func TestClient_Remove(t *testing.T) {
	c:=fakeClient()
	defer c.db.Close()

	err := c.Remove("yolo")
	if err != nil {
		t.Fatal(err)
	}
}