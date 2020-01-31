package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync/atomic"
)

//
//func generate(channel chan string) {
//	for ;; {
//		var l []string
//		for i := 0; i < maxWords; i++ {
//			l = append(l, String(rand.Int() % maxStringLength))
//		}
//		channel <- strings.Join(l, " ")
//	}
//}

func generate(channel chan ChanType) {
	log.Info("Inside generator")
	log.Info(fmt.Sprintf("Current counter is %d", counter))
	for counter < uint64(len(docs)) {
		var mes Message
		mes.i = counter
		mes.payload = docs[counter]
		channel <- ChanType(mes)
		atomic.AddUint64(&counter, 1)
	}
	close(channel)
}
