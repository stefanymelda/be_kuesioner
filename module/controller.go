package module

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atdb"
	"github.com/stefanymelda/be_kuesioner/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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


// func InsertKuesioner(db *mongo.Database, col string, lat float64, long float64, lokasi string, email string, status string, biodata model.Responden) (InsertedID interface{}) {
// 	var kuesioner model.Kuesioner
// 	kuesioner.Latitude = lat
// 	kuesioner.Longitude = long
// 	kuesioner.Location = lokasi
// 	kuesioner.Email = email
// 	kuesioner.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
// 	kuesioner.Status = status
// 	kuesioner.Biodata = biodata
// 	return InsertOneDoc(db, col, kuesioner)
// }

func InsertResponden(db *mongo.Database, col string, nama string,  jenis_kelamin string, usia int, email string, phone_number string) (InsertedID interface{}) {
	var responden model.Responden
	responden.Nama = nama
    responden.Jenis_kelamin = jenis_kelamin
    responden.Usia = usia
    responden.Email = email
    responden.Phone_number = phone_number
	// responden.Jam_pengisian = jam_pengisian
	// responden.Hari_pengisian = hari_pengisian
	return InsertOneDoc(db, col, responden)
}

// func InsertJamPengisian(db *mongo.Database, col string, durasi int, jam_mulai string, jam_selesai string, deskripsi string) (InsertedID interface{}) {
// 	var jampengisian model.JamPengisian
// 	jampengisian.Durasi = durasi
// 	jampengisian.Jam_mulai = jam_mulai
// 	jampengisian.Jam_selesai = jam_selesai
// 	jampengisian.Deskripsi = deskripsi
// 	return InsertOneDoc(db, col, jampengisian)
// }

func InsertLokasi(db *mongo.Database, col string, nama string, kategori string) (InsertedID interface{}) {
	var lokasi model.Lokasi
	lokasi.Nama = nama
	lokasi.Kategori = kategori
	return InsertOneDoc(db, col, lokasi)
}

func InsertQuestion(db *mongo.Database, col string, nomor int, text string, options string) (InsertedID interface{}) {
	var question model.Question
	question.Nomor = nomor
	question.Text = text
	question.Options = options
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

func GetAllSurvey(db *mongo.Database, col string) (data []model.Survey) {
	data_survey := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_survey.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
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

// func GetJamPengisianFromDurasi(durasi int, db *mongo.Database, col string) (jp model.JamPengisian) {
// 	jampengisian := db.Collection(col)
// 	filter := bson.M{"durasi": durasi}
// 	err := jampengisian.FindOne(context.TODO(), filter).Decode(&jp)
// 	if err != nil {
// 		fmt.Printf("getJamPengisianFromDurasi: %v\n", err)
// 	}
// 	return jp
// }

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

func GetAllKuesioner(db *mongo.Database, col string) (data []model.Kuesioner) {
	data_kuesioner := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_kuesioner.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//TUGASBESAR

//KUESIONER
func InsertKuesioner(db *mongo.Database, col string, lat float64, long float64, lokasi string, email string, status string, biodata model.Responden) (insertedID primitive.ObjectID, err error) {
	kuesioner := bson.M{
		"longitude":    long,
		"latitude":     lat,
		"location":     lokasi,
		"email":     	email,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"status"  :     status,
		"biodata":      biodata,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), kuesioner)
	if err != nil {
		fmt.Printf("InsertKuesioner: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateKuesioner(db *mongo.Database, col string, id primitive.ObjectID, lat float64, long float64, lokasi string, email string, status string, biodata model.Responden) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"longitude":    long,
			"latitude":     lat,
			"location":     lokasi,
			"email":     	email,
			"status":      	status,
			"biodata":      biodata,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateKuesioner: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func GetKuesionerFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ksr model.Kuesioner, errs error) {
	responden := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := responden.FindOne(context.TODO(), filter).Decode(&ksr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ksr, fmt.Errorf("no data found for ID %s", _id)
		}
		return ksr, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ksr, nil
}

func DeleteKuesionerByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	responden := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := responden.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
//ENDKUESIONER

//SURVEY

func InsertSurvey1(db *mongo.Database, col string, kode int, title string, soal model.Question) (insertedID primitive.ObjectID, err error) {
	survey := bson.M{
		"kode":    kode,
		"title":   title,
		"soal":    soal,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), survey)
	if err != nil {
		fmt.Printf("InsertSurvey1: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateSurvey(db *mongo.Database, id primitive.ObjectID, col string, kode int, title string, soal model.Question) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kode":    kode,
			"title":   title,
			"soal":    soal,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateSurvey: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func GetSurveyFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ksr model.Survey, errs error) {
	survey := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := survey.FindOne(context.TODO(), filter).Decode(&ksr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ksr, fmt.Errorf("no data found for ID %s", _id)
		}
		return ksr, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ksr, nil
}

func DeleteSurveyByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	question := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := question.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//ENDSURVEY

//ENDTUGASBESAR

func InsertAdmin(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
		"username": username,
		"password": password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

//LogAdmin

func LogAdmin(db *mongo.Database, col string, username string, password string) (authenticated bool, err error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}

	result, err := db.Collection(col).CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("LogAdmin: %v\n", err)
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

//ENDLogAdmin