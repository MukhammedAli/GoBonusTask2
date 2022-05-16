package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// "go run "c:\Users\muham\Desktop\Main\golang\tcp\simple-tcp-service\cmd\tcp-client.go" C:/Users/muham/ Desktop/txtfile.txt,Desktop/txtfile2.txt,Desktop/txtfile3.txt"

	if len(os.Args) < 2 {
		fmt.Println("Where is file name?")
		return
	}

	name := strings.Split(os.Args[2], ",")

	data, err := ioutil.ReadFile(os.Args[1] + name[0])

	something, good := ioutil.ReadFile(os.Args[1] + name[1])

	everything, so := ioutil.ReadFile(os.Args[1] + name[2])

	if err != nil || good != nil || so != nil {
		fmt.Println("You have a problem")
		fmt.Println(err)
		fmt.Println(good)
		fmt.Println(so)
	}

	fmt.Print(string(data))
	fmt.Print(string(something))
	fmt.Print(string(everything))

	fmt.Println("Zip archive creation")
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("First file checking")
	f1, err := os.Open(os.Args[1] + name[0])
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("First file writing into archive")
	w1, err := zipWriter.Create("txt/test1.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	fmt.Println("Second file checking")
	f2, err := os.Open(os.Args[1] + name[1])
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	fmt.Println("Second file writing into archive")
	w2, err := zipWriter.Create("txt/test2.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w2, f2); err != nil {
		panic(err)
	}

	fmt.Println("Third file checking")
	f3, err := os.Open(os.Args[1] + name[2])
	if err != nil {
		panic(err)
	}

	defer f3.Close()

	fmt.Println("Third file writing into archive")
	w3, err := zipWriter.Create("txt/test.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w3, f3); err != nil {
		panic(err)
	}
	fmt.Println("closing zip archive...")
	zipWriter.Close()
}
