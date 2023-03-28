package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Responden struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama             string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Jenis_kelamin    string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
	Usia             int      			 `bson:"usia,omitempty" json:"usia,omitempty"`
	Email      	     string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number     string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jam_pengisian    []JamPengisian     `bson:"jam_pengisian,omitempty" json:"jam_pengisian,omitempty"`
	Hari_pengisian   []string           `bson:"hari_pengisian,omitempty" json:"hari_pengisian,omitempty"`
}

type JamPengisian struct {
	Durasi      int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_mulai   string   `bson:"jam_mulai,omitempty" json:"jam_mulai,omitempty"`
	Jam_selesai string   `bson:"jam_selesai,omitempty" json:"jam_selesai,omitempty"`
	Gmt         int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari        []string `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Kuesioner struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	Biodata      Responden           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
