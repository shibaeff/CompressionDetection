package dataload

import (
	"fmt"
	"github.com/go-gota/gota/series"
	"github.com/kniren/gota/dataframe"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
)

// DataReader reads the corpus from files
type DataReader struct {
	df         dataframe.DataFrame
	Categories series.Series
	Names      []string
}

// Init provides intial workaround
func (dr *DataReader) Init(path string) {
	xl, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Print(err)
		log.Panic("Failed to read the xlsx file")
	}
	slice, err := xl.ToSlice()
	if err != nil {
		log.Panic(err)
	}
	df := dataframe.LoadRecords(slice[1])
	dr.df = df
	dr.Categories = df.Col("Category")
	dr.Names = df.Col("File").Records()
}

// FormCorpus forms a corpus from the provided files
func (dr DataReader) FormCorpus() [][]byte {
	const bound = 1
	var corpus [][]byte
	count := int(bound*float64(len(dr.Names)) - 5)
	for _, name := range dr.Names {
		if count < 0 {
			break
		}
		count--
		data, err := ioutil.ReadFile("./cor/" + name)
		if err != nil {
			log.Panic("File read problem")
		}

		corpus = append(corpus, data)

	}
	for _, let := range []string{"a", "b", "c", "d", "e"} {
		data, err := ioutil.ReadFile("./cor/" + "orig_task" + let + ".txt")
		if err != nil {
			log.Panic("File read problem")
		}
		corpus = append(corpus, data)
	}
	return corpus
}
