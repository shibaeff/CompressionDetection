package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"

	// "strings"
	simhash "minHashLzwAsync/src/simhash"
)

// Worker is a struct type in charge of concurrent hashing and compression
type Worker struct {
	c chan ChanType
}

var rw sync.RWMutex

func (w *Worker) acceptString() {
	log.Info("inside worker")
	for w.c != nil {
		if w.c == nil {
			log.Panic("Channel is nil!")
		}
		mes := <-w.c
		s := mes.payload
		counter := mes.i
		if len(s) == 0 || counter >= uint64(len(docs)) {
			return
		}
		hashvalue := simhash.Simhash(simhash.NewWordFeatureSet(s))
		rw.RLock()
		if !w.isStringHashed(hashvalue) {
			hashes[counter] = hashvalue
			// log.Info(fmt.Sprintf("Got string %v, calculated hash %v", s, hashes[counter]))
			log.Info("Adding new string yield!!!")
			cmp.M.Lock()
			cmp.Compress(string(s))
			cmp.M.Unlock()
		} else {
			log.Info(fmt.Sprintf("String ignored: %v", s))
		}
		rw.RUnlock()

	}

}

func (w *Worker) isStringHashed(hashvalue uint64) bool {
	for _, item := range hashes {
		val := simhash.Compare(hashvalue, item)
		if val != 0 && val < 9 {
			return true
		}

	}
	return false
}
