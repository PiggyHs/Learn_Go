package main

import (
	"../../pipeline"
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := 10000
	fileName := "small.in"
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	witer := bufio.NewWriter(file)
	p := pipeline.ReadomNumber(count)
	pipeline.WriterSink(witer, p)
	witer.Flush()

	file2, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	num := pipeline.ReaderSource(bufio.NewReader(file2))
	Numbercount := 0
	for v := range num {
		fmt.Println(v)
		Numbercount++
		if Numbercount > 100{
			break
		}
	}

}

func merageDemo()  {
	p := pipeline.Merage(pipeline.InMemSort(pipeline.ArraySource(3, 4, 5, 7, 1, 10)),
		pipeline.InMemSort(pipeline.ArraySource(1, 0, 9, 6, 7, 2, 4, 6,15)))
	for v := range p {
		fmt.Println(v)
	}
}
