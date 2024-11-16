package main

import (
	"PencraftB/utils"	
	db "PencraftB/repository"
	"time"
)

func main(){
	utils.Logger("Started main driver function")

	client := db.NewDBClient()
	defer client.Close();
	time.Sleep(3 * time.Second)
}