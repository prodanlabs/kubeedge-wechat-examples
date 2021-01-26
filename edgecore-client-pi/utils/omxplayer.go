package utils

import (
	"log"
	"os/exec"
)

func StopOmxplayer() {
	log.Println("Stopping omxplayer...")
	cmd := exec.Command("pkill", "-9", "omxplayer")
	// 执行命令，并返回结果
	_, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
}

func StartOmxplayer(index int) {

	StopOmxplayer()
	f := resFileName(index)
	log.Printf("The song playing is %s\n", f)
	cmd := exec.Command("omxplayer", "-o", "local", f)
	//cmd := exec.Command("tail", "-f", "/dev/null")
	go cmd.Run()
	/*	if err != nil {
		log.Printf("Error while playing track: %v\n", err)
	}*/
}
