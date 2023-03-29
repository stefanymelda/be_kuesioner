package Kuesioner

import (
	"fmt"
	"github.com/stefanymelda/be_kuesioner/model"
	"github.com/stefanymelda/be_kuesioner/module"
	"testing"
)


func TestInsertKuesioner(t *testing.T) {
	long := -6.873165194466392
	lat := 107.57589801606773
	lokasi := "ULBI"
	email := "stefanymelda01@gmail.com"
	status := "selesai"
	biodata := model.Responden{
		Nama : "Stefany Melda",
		Email : "stefanymelda01@gmail.com",
		Jenis_kelamin : "perempuan",
	}
	hasil:=module.InsertKuesioner(module.MongoConn, "kuesioner", long ,lat , lokasi , email , status , biodata )
	fmt.Println(hasil)
}

func TestInsertResponden(t *testing.T) {
	nama := "Stefany Melda"
	jenis_kelamin := "perempuan"
	usia := 19
	phone_number := "081267902121"
	hari_pengisian := "Senin"
	hasil := module.InsertResponden(module.MongoConn, "responden", nama , jenis_kelamin, usia, phone_number, hari_pengisian)
	fmt.Println(hasil)
}

func TestInsertJamPengisian(t *testing.T) {
	durasi := 2
	jam_mulai := "10 AM"
	jam_selesai := "11 AM"
	gmt := 7
	hasil := module.InsertJamPengisian(module.MongoConn, "jampengisian", durasi, jam_mulai, jam_selesai, gmt)
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
		Text : "Sejak kapan BTS debut?",
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
