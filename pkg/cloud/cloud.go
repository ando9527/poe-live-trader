package cloud

import (
	"context"
	"errors"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

type Client struct {
	*firestore.Client
	ctx context.Context
}

func NewClient(ctx context.Context)(c *Client, err error){
	f, err := firestore.NewClient(ctx, os.Getenv("APP_GOOGLE_PROJECT_ID"))
	if err != nil {
		return nil, err
	}
	c = &Client{
		Client: f,
		ctx:    ctx,
	}
	return c,nil
}
func (c *Client)UpdateInsert(poessid string) (err error) {
	// [START fs_update_create_if_missing]
	_, err = c.Collection("data").Doc("one").Set(c.ctx, map[string]interface{}{
		"poessid": poessid,
		"data":    time.Now(),
	}, firestore.MergeAll)

	if err != nil {
		logrus.Errorf("An error has occurred: %s", err)
	}
	return err
}

func (c *Client) QuerySSID() (ssid string, err error) {
	data := c.Collection("data")
	iter := data.OrderBy("poessid", firestore.Desc).Limit(1).Documents(c.ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logrus.Error(err)
		}
		ssid = doc.Data()["poessid"].(string)
		return ssid, nil
	}
	return "", errors.New("empty query")

}



