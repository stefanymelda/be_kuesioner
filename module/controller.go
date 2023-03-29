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

func InsertResponden(db *mongo.Database, col string, nama string,  jenis_kelamin string, usia int, email string, phone_number string) (InsertedID interface{}) {
	var responden model.Responden
	responden.Nama = nama
    responden.Jenis_kelamin = jenis_kelamin
    responden.Usia = usia
    responden.Email = email
    responden.Phone_number = phone_number
	return InsertOneDoc(db, col, responden)
}

func InsertJamPengisian(db *mongo.Database, col string, durasi int, jam_mulai string, jam_selesai string, gmt int) (InsertedID interface{}) {
	var jampengisian model.JamPengisian
	jampengisian.Durasi = durasi
	jampengisian.Jam_mulai = jam_mulai
	jampengisian.Jam_selesai = jam_selesai
	jampengisian.Gmt = gmt
	return InsertOneDoc(db, col, jampengisian)
}

func InsertLokasi(db *mongo.Database, col string, nama string, kategori string) (InsertedID interface{}) {
	var lokasi model.Lokasi
	lokasi.Nama = nama
	lokasi.Kategori = kategori
	return InsertOneDoc(db, col, lokasi)
}

func InsertQuestion(db *mongo.Database, col string, nomor int, text string) (InsertedID interface{}) {
	var question model.Question
	question.Nomor = nomor
	question.Text = text
	return InsertOneDoc(db, col, question)
}

func InsertAnswer(db *mongo.Database, col string, question_nomor int, text2 string) (InsertedID interface{}) {
	var answer model.Answer
	answer.Question_Nomor = question_nomor
	answer.Text2 = text2
	return InsertOneDoc(db, col, answer)
}

func InsertSurvey(db *mongo.Database, col string, kode int, title string, soal model.Question) (InsertedID interface{}) {
	var survey model.Survey
	survey.Kode = kode
	survey.Title = title
	survey.Soal = soal
	return InsertOneDoc(db, col, survey)
}

func GetKuesionerFromStatus(status string, db *mongo.Database, col string) (ksr model.Kuesioner) {
	kuesioner := db.Collection(col)
	filter := bson.M{"status": status}
	err := kuesioner.FindOne(context.TODO(), filter).Decode(&ksr)
	if err != nil {
		fmt.Printf("getKuesionerFromStatus: %v\n", err)
	}
	return ksr
}

func GetRespondenFromUsia(usia int, db *mongo.Database, col string) (rsp model.Responden) {
	responden := db.Collection(col)
	filter := bson.M{"usia": usia}
	err := responden.FindOne(context.TODO(), filter).Decode(&rsp)
	if err != nil {
		fmt.Printf("getRespondenFromUsia: %v\n", err)
	}
	return rsp
}

func GetJamPengisianFromDurasi(durasi int, db *mongo.Database, col string) (jp model.JamPengisian) {
	jampengisian := db.Collection(col)
	filter := bson.M{"durasi": durasi}
	err := jampengisian.FindOne(context.TODO(), filter).Decode(&jp)
	if err != nil {
		fmt.Printf("getJamPengisianFromDurasi: %v\n", err)
	}
	return jp
}

func GetLokasiFromNama(nama string, db *mongo.Database, col string) (lks model.Lokasi) {
	lokasi := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := lokasi.FindOne(context.TODO(), filter).Decode(&lks)
	if err != nil {
		fmt.Printf("getLokasiFromNama: %v\n", err)
	}
	return lks
}

func GetAllKuesionerFromEmail(email string, db *mongo.Database, col string) (aksr []model.Kuesioner) {
	kuesioner := db.Collection(col)
	filter := bson.M{"email": email}
	cursor, err := kuesioner.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllKuesionerFromEmail: %v\n", err)
	}
	err = cursor.All(context.TODO(), &aksr)
	if err != nil{
		fmt.Println(err)
	}
	return 
}