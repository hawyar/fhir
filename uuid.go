package main

import (
	"log"
	"os/exec"
)

// regex fromat: [A-Za-z0-9\-\.]{1,64}
func RandoID() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(uuid)
}
