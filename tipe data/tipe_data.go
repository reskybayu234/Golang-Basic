package main

import "fmt"

// =============================
// tipe data numerik Non-Desimal
// =============================

// func main(){
// 	var bilanganPertama uint8 = 89
// 	var bilanganKedua = -1243423644

// 	fmt.Printf("bilangan positif : %d\n", bilanganPertama)
// 	fmt.Printf("bilangan negatif : %d\n", bilanganKedua)
// }

// =========================
// tipe data numerik Desimal
// =========================

// func main(){
// 	var bilanganDesimal = 2.63

// 	fmt.Printf("bilangan desimal : %f\n", bilanganDesimal)
// 	fmt.Printf("bilangan desimal : %.3f\n", bilanganDesimal)
// }

// =========================
// tipe data bool (Boolean)
// =========================

// func main(){
// 	var bole bool = true
// 	fmt.Printf("hasil %t\n", bole)
// }

// =========================
// tipe data string
// =========================

func main(){
	var pesan string = "bayu"
	fmt.Printf("hello %s",pesan)

	var backticks string = `Saya "bayu", bekerja sebagai junior progammer di salah satu perusahaan IT di jakarta`
	fmt.Printf(backticks)
}