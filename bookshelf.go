package main

import "fmt"
import "strings"

type Book struct {
  title  string
  author string
  year   int
  read   bool
}

var err     error
var running bool            = true
var books   map[string]Book = map[string]Book {
  "001": { "title 1", "author 1", 2023, false },
  "002": { "title 2", "author 2", 2021, true },
  "003": { "title 3", "author 1", 2022, false },
  "004": { "title 4", "author 1", 2022, true },
  "005": { "title 5", "author 2", 2020, false },
}

func main() {
  welcome()

  for running {
    listMenu()
  }
}

func welcome() {
  fmt.Println(`Selamat datang di aplikasi "bookshelf" berbasis CLI`)
}

func listMenu() {
  fmt.Println("[INFO] Silahkan pilih salah satu menu berikut.")
  fmt.Println("1. Tampilkan semua buku")
  fmt.Println("2. Tampilkan semua buku telah dibaca")
  fmt.Println("3. Tampilkan semua buku belum dibaca")
  fmt.Println("4. Tambah buku")
  fmt.Println("5. Ubah buku")
  fmt.Println("6. Hapus buku")
  fmt.Println("7. Keluar aplikasi")

  selectMenu()
}

func selectMenu() {
  var selection int

  fmt.Print(">> ")
  fmt.Scanln(&selection)
  fmt.Println()

  switch selection {
    case 1: show("")
    case 2: show("read")
    case 3: show("unread")
    case 4: store()
    // case 5: update()
    // case 6: destroy()
    case 7:
      exitApp()
      return
    default: invalidSelection()
  }

  // var next string

  // fmt.Println("Lanjut? (Kosongi dan tekan Enter bila tidak)")
  // fmt.Print(">> ")
  // fmt.Scanln(&next)

  // if next == "" {
    // exitApp()
  // }

  // fmt.Println()
}

func show(read string) {
  fmt.Println("[INFO] Daftar buku")
  fmt.Println()

  var readBooks, unreadBooks, bookStr string

  for code, book := range books {
    bookStr = strings.Join([]string {
      fmt.Sprintf("(%s) Judul: %s", code, book.title),
      fmt.Sprintf("Penulis: %s", book.author),
      fmt.Sprintf("Tahun terbit: %d\n", book.year),
    }, "\t")

    if book.read {
      readBooks += bookStr
    } else {
      unreadBooks += bookStr
    }
  }

  if read == "" || read == "read" {
    fmt.Println("[SECTION] Read book list")
    fmt.Println(readBooks)
  }

  if read == "" || read == "unread" {
    fmt.Println("[SECTION] Unread book list")
    fmt.Println(unreadBooks)
  }

  fmt.Println()
}

func store() {
  fmt.Println("[MENU] Tambah buku")

  var code string
  var book Book

  fmt.Print("Masukkan kode buku >> ")
  fmt.Scan(&code)

  fmt.Print("Masukkan judul buku >> ")
  fmt.Scan(&book.title)

  fmt.Print("Masukkan nama penulis buku >> ")
  fmt.Scan(&book.author)

  fmt.Print("Masukkan tahun terbit buku >> ")
  fmt.Scan(&book.year)

  books[code] = book

  fmt.Println("[INFO] Berhasil menambahkan buku")
  fmt.Println()

  show("")
}

// func update() {
  // fmt.Println("[MENU] Ubah item")
  // fmt.Println()
  // show()

  // var selection int

  // fmt.Println("Masukkan nomor item yang ingin diubah")
  // fmt.Print(">> ")
  // fmt.Scanln(&selection)
  // fmt.Println()

  // var item string = todolist[selection - 1]
  // var newItem string

  // fmt.Printf("Item asal: %s\n", item)
  // fmt.Print("Ubah menjadi >> ")
  // fmt.Scanln(&newItem)
  // fmt.Println()

  // todolist[selection - 1] = newItem
  // fmt.Printf("[INFO] Berhasil mengubah %q menjadi %q\n", item, newItem)
  // fmt.Println()

  // show()
// }

// func destroy() {
  // fmt.Println("[MENU] Hapus item")
  // fmt.Println()
  // show()

  // var selection int

  // fmt.Println("Masukkan kode buku yang ingin dihapus")
  // fmt.Print(">> ")
  // fmt.Scanln(&selection)
  // fmt.Println()

  // var index int = selection - 1
  // var item string = todolist[index]

  // todolist = append(todolist[:index], todolist[(index + 1):]...)
  // fmt.Printf("[INFO] Berhasil menghapus %q\n", item)
  // fmt.Println()

  // show()
// }

func exitApp() {
  running = false

  fmt.Println("[INFO] Terima kasih telah menggunakan aplikasi kami")
}

func invalidSelection() {
  fmt.Println("[INFO] Menu yang dipilih tidak tersedia.")
  fmt.Println()
}
