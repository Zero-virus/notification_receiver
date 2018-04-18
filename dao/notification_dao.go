package dao

import (
	"log"

	//"github.com/mlabouardy/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NotificationsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "notification"
)

func (m *NotificationsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *NotificationsDAO) FindAll() ([]paymentResponse, error) {
	var notifications []paymentResponse
	err := db.C(COLLECTION).Find(bson.M{}).All(&notification)
	return notifications, err
}

func (m *NotificationsDAO) FindById(id string) (paymentResponse, error) {
	var notification paymentResponse
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&notification)
	return notification, err
}

func (m *NotificationsDAO) Insert(notification paymentResponse) error {
	err := db.C(COLLECTION).Insert(&notification)
	return err
}

func (m *NotificationDAO) Delete(notification paymentResponse) error {
	err := db.C(COLLECTION).Remove(&notification)
	return err
}

func (m *NotificationsDAO) Update(notification paymentResponse) error {
	err := db.C(COLLECTION).UpdateId(notification.TransactionID, &notification)
	return err
}
