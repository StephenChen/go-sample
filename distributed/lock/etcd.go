package lock

import (
	"github.com/zieckey/etcdsync"
	"log"
)

func Etcd() {
	m, err := etcdsync.New("/lock", 10, []string{"http://localhost:2379"})
	if m == nil || err != nil {
		log.Println("etcdsync.New failed")
		return
	}
	err = m.Lock()
	if err != nil {
		log.Println("etcdsync.Lock falied")
		return
	}

	log.Println("etcdsync.Lock OK")
	log.Println("Get the lock. Do something.")

	err = m.Unlock()
	if err != nil {
		log.Println("etcdsync.Unlock failed")
	} else {
		log.Println("etcdsync.Unlock OK")
	}
}
