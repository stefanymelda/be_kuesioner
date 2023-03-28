package Kuesioner

import (
	"fmt"
	"github.com/stefanymelda/be_kuesioner/model"
	"github.com/stefanymelda/be_kuesioner/module"
	"testing"
)

func TestInsertKuesioner(t *testing.T) {
	var jamIsi1 = model.JamPengisian{
		Durasi:     2,
		Jam_mulai:  "10:00",
		Jam_selesai: "12:00",
		Gmt:        7,
		Hari:       []string{"Rabu", "Jumat", "Sabtu"},
	}
	var jamIsi2 = model.JamPengisian{
		Durasi:     4,
		Jam_mulai:  "15:00",
		Jam_selesai: "19:00",
		Gmt:        7,
		Hari:       []string{"Senin"},
	}

	long := -6.873165194466392
	lat := 107.57589801606773
	lokasi := "ULBI"
	email := "juwitastefany13@gmail.com"
	status := "selesai"
	biodata := model.Responden{
		Nama : "Juwita Stefany",
		Email : "juwitastefany13@gmail.com",
		Jenis_kelamin : "perempuan",
		Jam_pengisian:   []model.JamPengisian{jamIsi1, jamIsi2},
		Hari_pengisian:  []string{"Rabu", "Kamis"},
	}
	hasil:= module.InsertKuesioner(module.MongoConn, "kuesioner", long ,lat , lokasi , email , status , biodata )
	fmt.Println(hasil)

}


func TestGetRespondenFromEmail(t *testing.T) {
	email := "juwitastefany13@gmail.com"
	biodata := module.GetRespondenFromEmail(email, module.MongoConn, "kuesioner")
	fmt.Println(biodata)
}

func TestGetRespondenFromStatus(t *testing.T) {
	status := "selesai"
	biodata := module.GetKuesionerFromStatus(status, module.MongoConn, "kuesioner")
	fmt.Println(biodata)
}

func TestGetAllRespondenFromStatus(t *testing.T) {
	status := "selesai"
	data := module.GetAllKuesionerFromStatus(status, module.MongoConn, "kuesioner")
	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllKuesioner(module.MongoConn, "kuesioner")
	fmt.Println(data)
}
