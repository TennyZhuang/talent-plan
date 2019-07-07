package main

import (
	"fmt"
	"strconv"
	"strings"
)

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs
	// round 1: do url count
	args = append(args, RoundArgs{
		MapFunc:    URLCountMap,
		ReduceFunc: URLCountReduce,
		NReduce:    nWorkers,
	})
	// round 2: sort and get the 10 most frequent URLs
	args = append(args, RoundArgs{
		MapFunc:    URLTop10Map,
		ReduceFunc: ExampleURLTop10Reduce,
		NReduce:    1,
	})
	return args
}

func URLCountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(string(contents), "\n")
	cntMap := make(map[string]int)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		cntMap[l] += 1
	}
	kvs := make([]KeyValue, 0, 10)
	for url, cnt := range cntMap {
		kvs = append(kvs, KeyValue{url, strconv.Itoa(cnt)})
	}
	return kvs
}

func URLCountReduce(key string, values []string) string {
	total := 0
	for _, v := range values {
		cnt, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		total += cnt
	}
	return fmt.Sprintf("%v %v\n", key, strconv.Itoa(total))
}

func URLTop10Map(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	kvs := make([]KeyValue, 0, len(lines))
	cntMap := make(map[string]int)
	for _, v := range lines {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cntMap[tmp[0]] = n
	}

	urls, cnts := TopN(cntMap, 10)
	for i, url := range urls {
		kvs = append(kvs, KeyValue{"", fmt.Sprintf("%v %v", url, cnts[i])})
	}
	return kvs
}
