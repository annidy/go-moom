package internal

import (
	"log"
	"os/exec"
)

func Restart()  {
	log.Println("kill app")
	exec.Command("/usr/bin/killall", "Moom").Run()
	log.Println("start app")
	exec.Command("/usr/bin/open", "-a", "Moom").Run()
}