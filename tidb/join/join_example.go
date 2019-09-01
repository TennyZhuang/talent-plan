package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// JoinExample performs a simple hash join algorithm.
func JoinExample(f0, f1 string, offset0, offset1 []int) (sum uint64) {
	tbl0, tbl1 := readCSVFileIntoTbl(f0), readCSVFileIntoTbl(f1)
	hashtable := buildHashTable(tbl0, offset0)
	for _, row := range tbl1 {
		rowIDs := probe(hashtable, row, offset1)
		for _, id := range rowIDs {
			v, err := strconv.ParseUint(tbl0[id][0], 10, 64)
			if err != nil {
				panic("JoinExample panic\n" + err.Error())
			}
			sum += v
		}
	}
	return sum
}

func readCSVFileIntoTbl(f string) (tbl [][]string) {
	csvFile, err := os.Open(f)
	if err != nil {
		panic("ReadFileIntoTbl " + f + " fail\n" + err.Error())
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic("ReadFileIntoTbl " + f + " fail\n" + err.Error())
		}
		tbl = append(tbl, row)
	}
	return tbl
}

type htable map[string][]int

func (m *htable) Put(k []byte, v int) {
	sk := string(k)
	(*m)[sk] = append((*m)[sk], v)
}

func (m *htable) Get(k []byte) []int {
	return (*m)[string(k)]
}

func buildHashTable(data [][]string, offset []int) (hashtable *htable) {
	var keyBuffer []byte
	hashtable = &htable{}
	for i, row := range data {
		for j, off := range offset {
			if j > 0 {
				keyBuffer = append(keyBuffer, '_')
			}
			keyBuffer = append(keyBuffer, []byte(row[off])...)
		}
		hashtable.Put(keyBuffer, i)
		keyBuffer = keyBuffer[:0]
	}
	return
}

func probe(hashtable *htable, row []string, offset []int) (rowIDs []int64) {
	var keyHash []byte
	var vals []int
	for i, off := range offset {
		if i > 0 {
			keyHash = append(keyHash, '_')
		}
		keyHash = append(keyHash, []byte(row[off])...)
	}
	vals = hashtable.Get(keyHash)
	for _, val := range vals {
		rowIDs = append(rowIDs, int64(val))
	}
	return rowIDs
}
