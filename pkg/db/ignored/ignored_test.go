package ignored

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
)

func fakeClient()(c *Client){
	c=NewClient()
	c.Connect(filepath.Join(os.TempDir(), "gorm.db"))

	c.db.DropTableIfExists(&Ignored{})
	c.Migration()

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

func TestClient_IsIgnored(t *testing.T) {
	c:=fakeClient()
	defer c.db.Close()
	ignored := c.IsIgnored("yolo")
	is:=is.New(t)
	is.Equal(ignored, true)

}