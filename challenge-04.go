package main

import (
	"fmt"
	"os"
	"sync"
)

var ai autoInc // global instance

type Biodata struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id + 1
	a.id++
	return
}

func GetBiodata(absent int) Biodata {
	biodataList := []Biodata{
		{
			ID:      ai.ID(),
			Name:    "Arnawa",
			Address: "Kulon Progo",
			Job:     "Jobless",
			Reason:  "Future Programming Language",
		},
		{
			ID:      ai.ID(),
			Name:    "Supra",
			Address: "Madiun",
			Job:     "Entrepreneur",
			Reason:  "Make Website for Company",
		},
		{
			ID:      ai.ID(),
			Name:    "Sukiyem",
			Address: "Bantul",
			Job:     "Merchant",
			Reason:  "Want to Changing Job",
		},
		{
			ID:      ai.ID(),
			Name:    "Cimol",
			Address: "Bandung",
			Job:     "FrontEnd Developer",
			Reason:  "Want to be Fullstack",
		},
	}

	if absent < 1 || absent > len(biodataList) {
		fmt.Println("Wrong Input ID, Biodata are 1 to", len(biodataList))
		os.Exit(1)
	}

	//how to get absent == ID
	return biodataList[absent-1] //0 are null ???
}

func main() {
	var absent int
	if len(os.Args) > 0 {
		fmt.Sscanf(os.Args[1], "%d", &absent)
	}

	data := GetBiodata(absent)
	fmt.Println("ID :", absent)
	fmt.Println("nama :", data.Name)
	fmt.Println("alamat :", data.Address)
	fmt.Println("pekerjaan :", data.Job)
	fmt.Println("alasan :", data.Reason)
}

/*
https://stackoverflow.com/questions/64631848/how-to-create-an-autoincrement-id-field
https://stackoverflow.com/questions/38511526/check-empty-float-or-integer-value-in-golang
https://pkg.go.dev/os
*/
