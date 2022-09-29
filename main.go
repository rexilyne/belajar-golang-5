package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	// penggunaan dari channel
	// pembuatan worker pool

	// [1, 2, 3, 4, 5, 6, 7, 8, 9]

	chanInt := make(chan int)
	go func() {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < len(arr); i++ {
			chanInt <- arr[i]
		}
	}()

	for i := 0; i < 3; i++ {
		go func() {
			for val := range chanInt {
				fmt.Println(val * 10)
			}
		}()
	}
	time.Sleep(10 * time.Second)
}

func channeling() {
	// deferAndExit()

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Printf("end of loop%v\n", i)
	// }
	// fmt.Println("outside function deferAndExit")
	defer recoveryFunction()

	// panic vs exit
	// panic bisa ditangkap oleh recover function
	// terhadap error yang terjadi
	// exit akan langsung mengeluarkan program
	// os.Exit(1)

	// Channel
	// pointer yang digunakan untuk
	// komunikasi antar go routine
	// secara aman -> bisa menghindari data race

	chanInt := make(chan int, 3)

	chanInt2 := make(chan int)
	chanInt3 := make(chan int)
	// chanInt <- 10 // artinya kita memberikan value ke channel tsb
	// <- chanInt // artinya kita mengambil value dari channel tsb

	// deadlock -> tidak bisa memproses process selanjutnya, saling tunggu
	go func() {
		for i := 0; i < 10; i++ {
			chanInt <- i // berhenti tapi di background
		}

		// close -> mengindikasikan bahwa
		// channel sudah tidak dapat dimasukin nilai lagi
		close(chanInt)

		// "send on closed channel"
		// chanInt <- -
	}()

	go func() {
		for val := range chanInt3 {
			chanInt2 <- val
		}
		close(chanInt2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chanInt3 <- i
			// time.Sleep(time.Duration(int(time.Millisecond) * i))
		}
		close(chanInt3)
	}()

	// ketika channel di assign
	// proses assignment akan diproses lanjut setelah
	// channel diambil juga nilainya

	// jika hal seperti ini terjadi, akan terjadi deadlock
	// untuk menghindari deadlock
	// 1. bisa menggunakan close
	// 2. kita hanya menerima sebanyak channel capacity

	for val := range chanInt {
		// proses ini akan mendengarkan channel
		// sehingga ketika channel diassign suatu value
		// dia akan langsung menangkap value tsb
		fmt.Println(val)
	}

	// for {
	// 	select {
	// 	case int2 := <-chanInt2:
	// 		fmt.Println("got from chan2", int2)
	// 	case int3 := <-chanInt3:
	// 		fmt.Println("got from chan3", int3)
	// 	}
	// }

	for val := range chanInt2 {
		fmt.Println("test", val)
	}

	// Dengan menggunakan channel, kita bisa tau
	// kapan sebuah go routine selesai menjalankan program / tugasnya

	// apakah goroutine + channel
	// bisa mempercepat program kita?
	// tidak selalu -> semakin banyak go routine akan memakan resource
}

func panicExplain() {
	var err error
	email := "calmantarasp@gmail.com"
	if err = isEmailExist(email); err != nil {
		fmt.Printf("email not valid : %v", err)
		panic(err)
		return
	}
	fmt.Println("email is exist, you are ready to go!")
	fmt.Printf("got error: %v", err.Error())

	// Panic
	// untuk mengeluarkan program dengan suatu indikasi
	// misalnya, ketika program gagal connect to database
	// program tidak seharusnya berjalan
	// ketika kita mencoba memanggil suatu function dari interface
	// ketika interface itu masih nil

	// panic => dia akan langsung mengeluarkan program
	// err => data type yang akan menampung informasi kesalahan pada function
}

func recoveryFunction() {
	// recover akan menangkap error yang terjadi saat panic
	// atau saat program selesai dijalankan

	if err := recover(); err != nil {
		fmt.Printf("program exit caused by error:%v\n", err)
	}
	fmt.Printf("program exit normally\n")
}

func isEmailExist(email string) (err error) {
	emails := map[string]bool{
		"calman@gmail.com": true,
		"tara@gmail.com":   true,
		"abdi@gmail.com":   true,
		"gulam@gmail.com":  true,
	}

	if !emails[email] {
		// err = errors.New("email is not exist in our system")

		// tergantung linter => convertion yang digunakan saat ngoding
		// error tidak boleh berakhir dengan tanda seru / line baru
		// error tidak boleh mengandung huruf kapital
		err = fmt.Errorf("\n%v is not exist in our system", email)
		return err
	}

	return nil
}

func erorHandling() {
	// error, panic, recover
	// error => situasi yang tidak diinginkan
	// baik itu dari data yang tidak valid,
	//	- password salah
	//	- user tidak ditemukan
	// kondisi yang tidak normal
	//	- database tidak bisa connect
	//	- menghubungkan dengan server lain, tapi tidak bisa connect
	// kegunaan => mengindikasikan bahwa program kita atau data kita
	// tidak baik-baik saja

	// interface dengan function Error()
	// => mengubah error menjadi string
	var err error

	err = errors.New("this is custom error")
	fmt.Println(err.Error())
	// error biasanya digunakan untuk return/output dari suatu function
}

func deferAndExit() {
	// defer recoverFunction()
	var pwdEnv string
	defer func() {
		fmt.Printf("current dir is: %v\n", pwdEnv)
	}()

	name := "test1"
	if name == "test" {
		fmt.Println("name is test")
		return
	}
	fmt.Println("PAKSA KELUAR")
	// os package adalah package yang
	// mengakses system komputer kita secar langsung

	// 0 mengindikasikan success
	// != 0 mengindikasikan error
	// exit => akan memaksa program untuk keluar
	//exit ini berguna nantinya untuk:
	//	1. mengetahui program kita mati karena apa
	//	2. kita bisa memanfaatkan grace exit di program go
	//		- ini akan dibahas pada web application

	// di line ini dia akan langsung mematikan program
	// os.Exit(0)

	pwdEnv = os.Getenv("PWD")
	fmt.Println(pwdEnv)

	hostName, _ := os.Hostname()
	fmt.Println(hostName)

	fmt.Println("this is not test")
}
