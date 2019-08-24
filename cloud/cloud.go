// Package p contains an ht Cloud Function.
package p

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/ando9527/poe-live-trader/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)
var conf env.Config
func init(){
	log.InitCloudLogger(false)
	conf = env.NewConfig()
}

//func basicAuth(username, password string) string {
//	auth := username + ":" + password
//	return base64.StdEncoding.EncodeToString([]byte(auth))
//}
//
//func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
//	req.Header.Add("Authorization","Basic " + basicAuth("admin","cycloneOP"))
//	return nil
//}
var mux = newMux()


func C(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w,r)
}


func newMux() *http.ServeMux{
	mux:=http.NewServeMux()
	mux.HandleFunc("/", handleAuth(handleSSID()))

	return mux
}


func updateInsert(ctx context.Context, client *firestore.Client, poessid string) error {
	// [START fs_update_create_if_missing]
	_, err := client.Collection("data").Doc("one").Set(ctx, map[string]interface{}{
		"poessid": poessid,
		"data":time.Now(),
	}, firestore.MergeAll)

	if err != nil {
		logrus.Errorf("An error has occurred: %s", err)
	}
	return err
}

func querySSID(ctx context.Context, client *firestore.Client)(ssid string, err error) {
	data := client.Collection("data")
	iter := data.OrderBy("poessid", firestore.Desc).Limit(1).Documents(ctx)

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

func handleSSID() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ctx := context.Background()
		client, err := firestore.NewClient(ctx, conf.GoogleProjectId)
		if err != nil {
			logrus.Fatalf("Failed to create client: %v", err)
		}
		defer client.Close()
		ssid, err := querySSID(ctx, client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, html.EscapeString(ssid))
	}
}

func handleAuth(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		user , pass, _ := r.BasicAuth()
		if user!=conf.User || pass !=conf.Pass{
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		h(w,r)
	}
}



