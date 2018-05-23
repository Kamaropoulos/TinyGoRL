package models

import (
	"log"

	goctapus "github.com/Kamaropoulos/goctapus-mongo"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"

	"github.com/zemirco/uid"

	Log "github.com/sirupsen/logrus"
)

type URL struct {
	ID   	 bson.ObjectId `json:"id" bson:"_id"`
	ShortURL string		   `json:"shorturl" bson:"shorturl"`
	URL		 string        `json:"url" bson:"url"`
}

type URLCollection struct {
	URLs []URL `json:"items" bson:"items"`
}

// GetURLs from the DB
func GetURLs() URLCollection {
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("tinygorl").C("urls")
	result := URLCollection{}
	err := c.Find(nil).All(&result.URLs)

	if err != nil {
		panic(err)
	}

	return result
}

// GetURL gets a specific url from the DB
func GetURL(shortURL string) (URL, error) {
	Log.Debug("getting: ", shortURL)
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("tinygorl").C("urls")
	url := URL{}
	err := c.Find(bson.M{"shorturl": shortURL}).One(&url)
	Log.WithField("url", url).Debug("result:")

	if err != nil {
		Log.Warn(err)
		return URL{}, err
	}

	return url, nil
}

// PutURL into DB
func PutURL(longURL string) (string, error) {
	Log.Debug(longURL)
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("tinygorl").C("urls")
	url := URL{ID: bson.NewObjectId(), ShortURL: uid.New(10), URL: longURL}
	err := c.Insert(&url)
	if err != nil {
		log.Fatal(err)
	}
	return url.ShortURL, nil
}

// DeleteTask from DB
func DeleteURL(url string) (int64, error) {
	database := goctapus.ConnectDB(goctapus.Config)
	defer database.Close()
	c := database.DB("tinygorl").C("urls")
	err := c.Remove(bson.M{"shorturl": url})

	if err != nil {
		Log.Warn(err)
	}

	return 0, nil
}
