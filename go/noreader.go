// noreader
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//input := flag.String("input", "/home/gfeng/Desktop/History/Original", "input file or folder")
	//output := flag.String("output", "/home/gfeng/Desktop/History/Readable1", "save converted file to")
	input := flag.String("input", "", "input file or folder, ~/Desktop/History/Original")
	output := flag.String("output", "", "save converted file to, ~/Desktop/History/Readable")

	flag.Parse()

	fmt.Println("input = " + *input)
	fmt.Println("output = " + *output)

	fi, err := os.Stat(*input)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(*output); os.IsNotExist(err) {
		os.Mkdir(*output, os.ModePerm)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("directory: " + *input)
		files, err := ioutil.ReadDir(*input)
		check(err)
		for _, file := range files {
			filename := filepath.Join(*input, file.Name())
			fmt.Println("file: " + filename)
			convert(filename, *output)
		}

	case mode.IsRegular():
		fmt.Println("file: " + *input)
		convert(*input, *output)
	}

	fmt.Println("Done!")
}

func convert(input string, output string) bool {
	data, err := ioutil.ReadFile(input)
	check(err)

	var outstr = ""

	for i := 0; i < len(data); i++ {
		var b = int(data[i])
		if data[i] < 0 {
			b = 255 + int(data[i])
		}

		if b > 31 && b < 127 || b == 10 || b == 13 {
			outstr += string(data[i])
		}
	}
	//fmt.Println(string(outstr))

	_, filename := filepath.Split(input)

	ioutil.WriteFile(filepath.Join(output, filename), []byte(outstr), os.ModePerm)

	return true
}
