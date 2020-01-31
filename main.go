package main

import (
	"datasetload/src/dataload"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"minHashLzwAsync/src/compressor"
	"runtime"
	"sync"
	"time"
	// simhash "minHashLzwAsync/src/simhash"
)

// Message is a struct type containing information transmitted over channels
type Message struct {
	payload []byte
	i       uint64
}

func staticCounter() (f func() int) {
	var i int
	f = func() int {
		i++
		//  fmt.Println(i)
		return i
	}
	return
}

var workers [maxWorkers]Worker

//var docs = [][]byte{
//	[]byte("this is a test phrase"),
//	[]byte("this is a test phrass"),
//	[]byte("foo bar"),
//}

var docs [][]byte
var dr dataload.DataReader
var hashes []uint64

var cmp compressor.Compressor

var counter uint64

func setWorkers(channel chan ChanType) {
	counter = 0
	for i := 0; i < maxWorkers; i++ {
		workers[i].c = channel
	}
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetLevel(log.FatalLevel)
	//path, err := os.Getwd()
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(path)
	rand.Seed(time.Now().Unix())

	cmp.Compress("")
	dr.Init("corpus-final09.xlsx")
	docs = dr.FormCorpus()

	avg := 0.0
	for i := 0; i < nreps; i++ {
		t1 := time.Now()
		hashes = make([]uint64, len(docs))
		channel := make(chan ChanType)
		var w sync.WaitGroup
		w.Add(maxWorkers)

		setWorkers(channel)
		log.Info(channel)
		log.Info("Initialization is done")

		for i := 0; i < maxWorkers; i++ {
			go func(i int) {
				log.Info("Waiting to accept string")
				workers[i].acceptString()
				w.Done()
			}(i)
		}
		go generate(channel)
		w.Wait()
		delta := time.Since(t1)
		avg += float64(delta.Seconds())
	}

	fmt.Printf("Avg time elapsed: %v seconds, compression achieved %v", avg/nreps)
	// defer log.Fatal(fmt.Sprintln(len(hashes)))
	// defer fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	//fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	// w.Wait()
}
