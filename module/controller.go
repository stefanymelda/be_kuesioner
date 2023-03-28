package module

import (
	"context"
	"fmt"
	"github.com/stefanymelda/be_kuesioner/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "data_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKuesionerOld(db *mongo.Database, long float64, lat float64, lokasi string, status string, biodata model.Responden) (InsertedID interface{}) {
	var kuesioner model.Kuesioner
	kuesioner.Latitude = long
	kuesioner.Longitude = lat
	kuesioner.Location = lokasi
	kuesioner.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	kuesioner.Status = status
	kuesioner.Biodata = biodata
	return InsertOneDoc(db, "kuesioner", kuesioner)
}

func InsertKuesioner(db *mongo.Database, col string, long float64, lat float64, lokasi string, email string, status string, biodata model.Responden) (InsertedID interface{}) {
	var kuesioner model.Kuesioner
	kuesioner.Latitude = long
	kuesioner.Longitude = lat
	kuesioner.Location = lokasi
	kuesioner.Email = email
	kuesioner.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	kuesioner.Status = status
	kuesioner.Biodata = biodata
	return InsertOneDoc(db, col, kuesioner)
}

func InsertResponden(db *mongo.Database, nama string, email string, jenis_kelamin string, jam_pengisian []model.JamPengisian, hari_pengisian []string) (InsertedID interface{}) {
	var responden model.Responden
	responden.Nama = nama
	responden.Email = email
	responden.Jenis_kelamin = jenis_kelamin
	responden.Jam_pengisian = jam_pengisian
	responden.Hari_pengisian = hari_pengisian
	return InsertOneDoc(db, "responden", responden)
}

func GetRespondenFromEmail(email string, db *mongo.Database, col string) (rsp model.Kuesioner) {
	responden := db.Collection(col)
	filter := bson.M{"email": email}
	err := responden.FindOne(context.TODO(), filter).Decode(&rsp)
	if err != nil {
		fmt.Printf("getRespondenFromEmail: %v\n", err)
	}
	return rsp
}

func GetRespondenFromNama(usia string, db *mongo.Database, col string) (rsp model.Kuesioner) {
	nm_responden := db.Collection(col)
	filter := bson.M{"biodata.usia": usia}
	err := nm_responden.FindOne(context.TODO(), filter).Decode(&rsp)
	if err != nil {
		fmt.Printf("getRespondenFromNama: %v\n", err)
	}
	return rsp
}

func GetKuesionerFromStatus(status string, db *mongo.Database, col string) (data model.Kuesioner) {
	responden := db.Collection(col)
	filter := bson.M{"status": status}
	err := responden.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getRespondenFromEmail: %v\n", err)
	}
	return data
}

func GetAllKuesionerFromStatus(status string, db *mongo.Database, col string) (data []model.Kuesioner) {
	responden := db.Collection(col)
	filter := bson.M{"status": status}
	cursor, err := responden.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllKuesioner(db *mongo.Database, col string) (data []model.Kuesioner) {
	responden := db.Collection(col)
	filter := bson.M{}
	cursor, err := responden.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
