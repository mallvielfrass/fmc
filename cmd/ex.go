package main

import (
	"errors"
	"os"

	"github.com/mallvielfrass/fmc"
)

type DV struct {
	Z, F string
}
type Date struct {
	Key   string
	Value string
	GR    int
	S     DV
}
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"database"`
}

func beta() {
	fmc.ErrorHandle(errors.New("Err"), "file")
}
func betaR() {
	fmc.ErrorHandleFatal(errors.New("Err"))
}
func main() {
	// fmc.Printf("#yst I#ybt  love#bbt  Go#gbt!\n")
	// fmc.Printf("#ybt I#rbt  love#bbt  Go#gbt!\n")
	// fmc.Printf("#ybt %s#rbt  %s#bbt  %s#gbt%s\n", "I", "love", "Go", "!")
	// fmc.Println("#ybt lol")
	// fmc.Printfln("#ybt %s", "I")
	// fmc.Println("#bbt Println")
	// fmc.Printfln(`
	// #RRR1) #bst  !#bst aqua 		#bbt !#bbt bold aqua
	// #RRR2) #0st  !#0st black 		#0bt !#0bt bold black
	// #RRR3) #Bst  !#Bst blue  		#Bbt !#Bbt bold blue
	// #RRR4) #fst  !#fst fuchsia     	#fbt !#fbt bold fuchsia
	// #RRR5) #1st  !#1st gray 		#1bt !#1bt bold gray
	// #RRR6) #Gst  !#Gst green 		#Gbt !#Gbt bold green
	// #RRR7) #gst  !#gst lime 		#gbt !#gbt bold lime
	// #RRR8) #mst  !#mst maroon		#mbt !#mbt bold maroon
	// #RRR9) #nst  !#nst navy 		#nbt !#nbt bold navy
	// #RRR10) #ost !#ost olive		#obt !#obt bold olive
	// #RRR11) #pst !#pst purple		#pbt !#pbt bold purple
	// #RRR12) #rst !#rst red   		#rbt !#rbt bold red
	// #RRR13) #sst !#sst silver		#sbt !#sbt bold silver
	// #RRR14) #tst !#tst teal        	#tbt !#tbt bold teal
	// #RRR15) #wst !#wst white		#wbt !#wbt bold white
	// #RRR16) #yst !#yst yellow		#ybt !#ybt bold yellow
	// #RRR17) #RRR !#RRR reset to default terminal font color
	// `)
	// jsonConfig := []byte(`{
	//     "server":{
	//         "host":"localhost",
	//         "port":"8080"},
	//     "database":{
	//         "host":"localhost",
	//         "user":"db_user",
	//         "password":"supersecret",
	//         "db":"my_db"}}`)
	// var config Config
	// err := json.Unmarshal(jsonConfig, &config)
	// if err != nil {
	// 	panic(err)
	// }
	// fmc.PrintStruct(config)
	//	beta()
	//	fmc.ErrorHandle(errors.New("Err"), "file")
	fmc.ErrorHandle(errors.New("Err"))
	//	betaR()
	//fmc.ErrorHandleFatal(errors.New("Err"))
	token := os.Getenv("token")
	if token == "" {
		fmc.ErrorHandleFatal(errors.New("os variable 'token' not found"))
	}
}
