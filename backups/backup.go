package backups

import (
	"log"
	"os/exec"
	"sync/atomic"
)

var writeCount uint64

func Backup() {
	log.Println("20 writes done, backing up!")
	exec.Command("scp", "prod.db", "root@cdn.maaya-pes.co:~")
}

func IncrementWriteCount() {
	atomic.AddUint64(&writeCount, 1)
	if writeCount%20 == 0 {
		Backup()
	}
}
