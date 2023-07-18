package main

import "fmt"

// program ini digunakan untuk konversi dari reamur ke celcius dan sebaliknya
func celcToFahr() {
	const step uint = 20
	const max uint = 100
	const min uint = 0
	var celcius, fahrenheit float64

	// celcius ke fahrenheit
	celcius = float64(min)
	fmt.Println("celcius", "fahrenheit")
	for celcius <= float64(max) {
		fahrenheit = (celcius * (float64(9) / 5)) + 32

		fmt.Println(celcius, fahrenheit)
		celcius = celcius + float64(step)
	}
}

func fahrToCelc() {
	const step uint = 6
	const max uint = 212
	const min uint = 0
	var fahrenheit, celcius int

	fahrenheit = int(min)
	fmt.Println("fahrenheit", "celcius")
	for fahrenheit <= int(max) {
		celcius = (fahrenheit - 32) * 5 / 9
		fmt.Println(fahrenheit, celcius)
		fahrenheit = fahrenheit + int(step)
	}
}

func searchGenap() {
	const max uint8 = 255
	var x uint16 = 0 // menggunakan uint16 karena, jika menggunakan uint8 ketika variabel x nilainya sudah sampai 255, di statement for loop akan menambahkan nilai variabel x, dan ketika tipe uint8 mencapai nilai maksimalnya, maka secara otomatis akan kembali lagi menjadi 0

	// selain menggunakan tipe uint16, juga bisa mengganti operator perbandingan dari <= menjadi <
	for x <= uint16(max) {
		if x%2 == 0 {
			fmt.Println(x)
		}
		x++
	}
}

func main() {
	celcToFahr()
	fmt.Println("=================")
	fahrToCelc()
	fmt.Println("=================")
	searchGenap()
}
