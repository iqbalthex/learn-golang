package main

import "fmt"

var running bool = true
var todolist = []string { "Ngoding", "Belajar", "Makan", "Minum" }
var err error

func main() {
  welcome()

  for running {
    listMenu()
  }
}

func welcome() {
  fmt.Println(`Selamat datang di aplikasi "to-do-list" berbasis CLI`)
}

func listMenu() {
  fmt.Println("[INFO] Silahkan pilih salah satu menu berikut.")
  fmt.Println("1. Tampilkan semua item")
  fmt.Println("2. Tambah item")
  fmt.Println("3. Ubah item")
  fmt.Println("4. Hapus item")
  fmt.Println("5. Keluar aplikasi")

  selectMenu()
}

func selectMenu() {
  var selection int

  fmt.Print(">> ")
  fmt.Scanln(&selection)
  fmt.Println()

  switch selection {
    case 1: show()
    case 2: store()
    case 3: update()
    // case 4: destroy()
    // case 5: exitApp()
    default: invalidSelection()
  }

  var lanjut string

  fmt.Println("Lanjut? (Kosongi dan tekan Enter bila tidak)")
  fmt.Print(">> ")
  fmt.Scanln(&lanjut)

  if lanjut == "" {
    running = false

    fmt.Println("Terima kasih telah menggunakan aplikasi kami")
  }

  fmt.Println()
}

func show() {
  fmt.Println("[INFO] Daftar item")

  for index, item := range todolist {
    fmt.Printf("(%d) %s\n", (index + 1), item)
  }

  fmt.Println()
}

func store() {
  fmt.Println("[MENU] Tambah item")

  var item string

  fmt.Print("Masukkan nama item >> ")
  fmt.Scanln(&item)

  todolist = append(todolist, item)

  fmt.Println("[INFO] Item berhasil ditambahkan")
  fmt.Println()

  show()
}

func update() {
  fmt.Println("[MENU] Ubah item")
  fmt.Println()
  show()

  var selection int

  fmt.Println("Masukkan nomor item yang ingin diubah")
  fmt.Print(">> ")
  fmt.Scanln(&selection)
  fmt.Println()

  var item string = todolist[selection - 1]
  var newItem string

  fmt.Printf("Item asal: %s\n", item)
  fmt.Print("Ubah menjadi >> ")
  fmt.Scanln(&newItem)
  fmt.Println()

  todolist[selection - 1] = newItem
  fmt.Printf(`Berhasil mengubah "%s" menjadi "%s"\n`, item, newItem)
  fmt.Println()

  show()
}

func invalidSelection() {
  fmt.Println("[INFO] Menu yang dipilih tidak tersedia.")
  fmt.Println()
}
