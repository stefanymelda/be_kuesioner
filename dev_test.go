package Kuesioner

import (
	"fmt"
	"github.com/stefanymelda/be_kuesioner/model"
	"github.com/stefanymelda/be_kuesioner/module"
	"testing"
)


func TestInsertKuesioner(t *testing.T) {
	long := 103.5394871
    lat := -1.6102634
    lokasi := "Jambi"
    email := "rinaldibarimbing@gmail.com"
    status := "Done"
    biodata := model.Responden{
        Nama : "Rinaldi Barimbing",
        Jenis_kelamin : "laki-laki",
		Email : "rinaldibarimbing@gmail.com",
        Usia : 25,
        Phone_number : "082264531232",

    }


	hasil:=module.InsertKuesioner(module.MongoConn, "kuesioner", long ,lat , lokasi , email , status , biodata )
	fmt.Println(hasil)
}

func TestInsertResponden(t *testing.T) {
	nama := "Wina Sibuea"
	jenis_kelamin := "perempuan"
	usia := 19
	email := "winasibuea22@gmail.com"
	phone_number := "082174302874"
	hasil := module.InsertResponden(module.MongoConn, "responden", nama , jenis_kelamin, usia, email, phone_number)
	fmt.Println(hasil)
}

// func TestInsertJamPengisian(t *testing.T) {
// 	durasi := 2
// 	jam_mulai := "12 AM"
// 	jam_selesai := "14 AM"
// 	deskripsi := "telah mengisi selama 2 jam"
// 	hasil := module.InsertJamPengisian(module.MongoConn, "jampengisian", durasi, jam_mulai, jam_selesai, deskripsi)
// 	fmt.Println(hasil)
// }

func TestInsertLokasi(t *testing.T) {
	nama := "Bali"
	kategori := "Kota"
	hasil := module.InsertLokasi(module.MongoConn, "lokasi", nama, kategori)
	fmt.Println(hasil)
}

func TestInsertSurvey(t *testing.T) {
	kode := 0141
	title := "Boyband BTS"
	soal := model.Question {
		Nomor : 15,
		Text : "Siapa anjing Taehyung?",
		Options : "a.Yeontan b.Jolly",
	}
	hasil := module.InsertSurvey(module.MongoConn, "survey", kode , title, soal)
	fmt.Println(hasil)
}

func TestGetKuesionerFromStatus(t *testing.T) {
	status := "Done"
	biodata := module.GetKuesionerFromStatus(status, module.MongoConn, "kuesioner")
	fmt.Println(biodata)
}

func TestGetRespondenFromUsia(t *testing.T) {
	usia := 19
	data := module.GetRespondenFromUsia(usia, module.MongoConn, "responden")
	fmt.Println(data)
}

// func TestGetJamPengisianFromDurasi(t *testing.T) {
// 	durasi := 2
// 	data := module.GetJamPengisianFromDurasi(durasi, module.MongoConn, "jampengisian")
// 	fmt.Println(data)
// }

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