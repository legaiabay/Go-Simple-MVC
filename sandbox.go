package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kataras/iris"
)

var d, e, f int = 4, 5, 6

func ngomong(kata string) string {
	return kata
}

func hitung(a int, b int, c int, exVar string) int {
	if c == 0 {
		return a * b
	} else if exVar == "y" {
		return a * b * c * d * e * f
	}

	return a * b * c
}

func testingSwitch(option int) string {
	switch option {
	case 0:
		return "ini satu!"
	case 1:
		return "ini dua"
	}

	return ""
}

func testingWhile(i int) int {
	sum := 0
	for sum < i {
		sum++
	}
	return sum
}

func testingFor(loopTime int) int {
	sum := 0
	for i := 0; i < loopTime; i++ {
		sum++
	}

	return sum
}

func testingArray(a []string) {
	fmt.Println("start")
	for value := range a {
		fmt.Println(a[value])
	}
	fmt.Println("end")
}

func testingArraySlice(a []string, n int) {
	fmt.Println(a[0:n])
}

type unkVariable struct {
	g, h, i int
	j       float32
}

type testingJeson struct {
	Nama       string `json:"nama"`
	Nilai      int    `json:"nilai"`
	Keterangan string `json:"keterangan"`
}

func startServer(port int) {
	// Kill server : ctrl+c
	// source ~/.profile
	// sudo netstat -lpn |grep :8080
	// sudo kill -9 PID
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("hello"))
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func toJson(data interface{}) string {
	var response []byte
	response, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Print(err)
	}

	return string(response)
}

func startIrisServer(port int) {
	app := iris.Default()

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Get("/test/{nilai:int}", func(ctx iris.Context) {
		nilai := ctx.Params().GetIntDefault("nilai", 0)

		response := testingJeson{
			Nama:       "baru lagi",
			Nilai:      nilai,
			Keterangan: "Mantappp",
		}

		ctx.JSON(response)
	})
	app.Run(iris.Addr(":" + strconv.Itoa(port)))
}

func mainn() {
	// Testing defer (diexecute di akhir)
	defer fmt.Println("ini defer")

	// Testing function with parameter, if else
	fmt.Println(ngomong("anjing "), hitung(1, 2, 3, "y"))

	// Testing struct
	unk := unkVariable{7, 8, 9, 4.0}
	fmt.Println(hitung(unk.g, unk.h, unk.i, "y"))

	// Testing for loop
	fmt.Println(testingFor(10))

	// Testing switch case
	fmt.Println(testingSwitch(1))

	// Testing while
	fmt.Println(testingWhile(100))

	// Testing array + foreach
	alay := []string{"aku", "tampan"}
	testingArray(alay)

	// Testing array slice
	alay2 := []string{"aku", "tampan", "tiada", "duanya", "boong deh"}
	testingArraySlice(alay2, 4)

	// Testing Map
	var newMap = map[string]int{"abay": 1, "abay2": 2, "abay3": 3}
	var newMap2 = map[string]int{"popp": 3, "popp1": 2, "popp2": 1}
	var newMapArray []map[string]int
	newMapArray = append(newMapArray, newMap)
	newMapArray = append(newMapArray, newMap2)
	fmt.Println(newMap)

	// Testing JSON
	//response, _ := json.Marshal(newMapArray, "", "  ") // One line
	//response, _ := json.MarshalIndent(newMapArray, "", "  ") // Pretty JSON
	aaa := testingJeson{
		Nama:       "baru lagi",
		Nilai:      200,
		Keterangan: "Mantappp",
	}

	var response []byte
	response, err := json.MarshalIndent(aaa, "", "  ")
	if err != nil {
		log.Print(err)
	}

	fmt.Println(string(response))

	// Testing webserver (native)
	// note : build dulu sebelum start
	// startServer(8787)

	// Testing webserver (Iris)
	//startIrisServer(8787)
}
