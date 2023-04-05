package Kuesioner

import (
	"fmt"
	"github.com/stefanymelda/be_kuesioner/model"
	"github.com/stefanymelda/be_kuesioner/module"
	"testing"
)


func TestInsertKuesioner(t *testing.T) {
	long := 118.9808932
    lat := 32.2023727
    lokasi := "Jepang"
    email := "yukikato10@gmail.com"
    status := "Done"
    biodata := model.Responden{
        Nama : "Yuki Kato",
        Jenis_kelamin : "perempuan",
		Email : "yukikato10@gmail.com",
        Usia : 27,
        Phone_number : "081387230912",
		Jam_pengisian : model.JamPengisian{
			Durasi : 2,
			Jam_mulai : "12.00",
			Jam_selesai : "14.00",
			Deskripsi : "Telah mengisi selama 2 jam",
		},
        Hari_pengisian : "Rabu",

    }


	hasil:=module.InsertKuesioner(module.MongoConn, "kuesioner", long ,lat , lokasi , email , status , biodata )
	fmt.Println(hasil)
}

func TestInsertResponden(t *testing.T) {
	nama := "Stefany Melda"
	jenis_kelamin := "perempuan"
	usia := 19
	email := "stefanymelda01@gamil.com"
	phone_number := "081267902121"
	jam_pengisian := model.JamPengisian {
		Durasi : 2,
		Jam_mulai : "12.00",
		Jam_selesai : "14.00",
		Deskripsi : "Telah mengisi selama 2 jam",

	}
	hari_pengisian := "Senin"
	hasil := module.InsertResponden(module.MongoConn, "responden", nama , jenis_kelamin, usia, email, phone_number, jam_pengisian, hari_pengisian)
	fmt.Println(hasil)
}

func TestInsertJamPengisian(t *testing.T) {
	durasi := 2
	jam_mulai := "12 AM"
	jam_selesai := "14 AM"
	deskripsi := "telah mengisi selama 2 jam"
	hasil := module.InsertJamPengisian(module.MongoConn, "jampengisian", durasi, jam_mulai, jam_selesai, deskripsi)
	fmt.Println(hasil)
}

func TestInsertLokasi(t *testing.T) {
	nama := "ULBI"
	kategori := "Kampus"
	hasil := module.InsertLokasi(module.MongoConn, "lokasi", nama, kategori)
	fmt.Println(hasil)
}

func TestInsertSurvey(t *testing.T) {
	kode := 01
	title := "Boyband BTS"
	soal := model.Question {
		Nomor : 1,
		Text : "Sejak kapan BTS debut?",
		Options : "a.2016 b.2017",
	}
	hasil := module.InsertSurvey(module.MongoConn, "survey", kode , title, soal)
	fmt.Println(hasil)
}

func TestGetKuesionerFromStatus(t *testing.T) {
	status := "selesai"
	biodata := module.GetKuesionerFromStatus(status, module.MongoConn, "kuesioner")
	fmt.Println(biodata)
}

func TestGetRespondenFromUsia(t *testing.T) {
	usia := 19
	data := module.GetRespondenFromUsia(usia, module.MongoConn, "responden")
	fmt.Println(data)
}

func TestGetJamPengisianFromDurasi(t *testing.T) {
	durasi := 2
	data := module.GetJamPengisianFromDurasi(durasi, module.MongoConn, "jampengisian")
	fmt.Println(data)
}

func TestGetLokasiFromNama(t *testing.T) {
	nama := "ULBI"
	data := module.GetLokasiFromNama(nama, module.MongoConn, "lokasi")
	fmt.Println(data)
}

func TestGetAllKuesionerFromEmail(t *testing.T) {
	email := "stefanymelda01@gmail.com"
	data := module.GetAllKuesionerFromEmail(email, module.MongoConn, "kuesioner")
	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllKuesioner(module.MongoConn, "kuesioner")
	fmt.Println(data)
}