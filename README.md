Nama: Marshely Ayu Iswanto 
NIM: 2311102073
Kelas: IF-11-03
“MODUL 2 RIVIEW STRUKTUR KONTROL”
Penjelasan Program
Soal 2a
No.1 
Program tersebut merupakan program dari go yaitu yang menerima tiga input berupa string dari yang kita masukkan lalu menukarnya secara melingkar dan menampilkan output hasil sebelum dan sesudah pertukaran. Nilai awal dari satu disimpan di variabel sementara “temp” lalu nilai dua dipindahkan ke satu, nilai tiga dipindahkan ke dua dan untuk nilai yang tersimpan di temp (nilai awal dari satu) dipindahkan ke tiga. 
No. 2
Program tersebut merupakan program dari go yaitu yang memeriksa apakah suatu tahun yang kita inputkan apakah tahun kabisat atau tidak. Program ini menggunakan “if” untuk menentukan apakah tahun tersebut termasuk kabisat atau tidak ( Tahun harus habis dibagi 4 (tahun%4 == 0)). Jika tahun yang di inputkan pengguna merupakan kabisat program akan mencetak “Kabisat: True”, namun jika tahun yang diinputkan pengguna bukan kabisat program akan mencetak “Kabisat: false”
No. 3
Program tersebut merupakan program dari  go yaitu menghitung volume dan luas permukaan bola berdasarkan input jejari yang kita inputkan. Volume bola dihitung dengan menggunakan  rumus 4/3 x π x r^3 , di mana math.Pi merepresentasikan konstanta π, dan math.Pow(jejari, 3) digunakan untuk menghitung pangkat tiga dari jejari. Luas permukaan bola dihitung dengan menggunakan  rumus 4 x π x r^2, di mana math.Pow(jejari, 2) digunakan untuk menghitung pangkat dua dari jejari. Output dari kode program tersebut adalah hasil perhitungan volume dan luas permukaan yang dicetak dengan menggunakan fmt.Printf(), menampilkan volume dan luas permukaan hingga 4 angka di belakang koma (%.4f).
No. 4
Program tersebut merupakan program dari go  yaitu mengonversi suhu dari satuan Celcius ke Reamur, Fahrenheit, dan Kelvin. Pengambilan input yaitu meminta pengguna untuk memasukkan suhu dalam derajat Celcius menggunakan fmt.Scan(&celcius),  untuk output dari kode program tersebut yaitu menampilkan hasil konversi ke Reamur, Fahrenheit, dan Kelvin. Untuk Kelvin, program menggunakan int(kelvin) yang mengubah hasil menjadi integer (angka bulat).
No. 5
Program tersebut merupakan program dari go yaitu mengambil input dari pengguna berupa lima angka integer dan tiga karakter. Setelah itu, program mencetak karakter yang sesuai dengan kode ASCII dari angka-angka yang dimasukkan serta karakter yang telah dimodifikasi dengan menambahkan nilai 1 pada setiap karakter (untuk beralih ke karakter berikutnya dalam urutan ASCII). Dalam inputnya fmt.Scanln(&angka1, &angka2, &angka3, &angka4, &angka5) digunakan untuk membaca lima angka integer dari input sedangkan fmt.Scanf("%c%c%c", &char1, &char2, &char3) digunakan untuk mengambil tiga karakter secara berurutan dari input. Output dari kode program tersebut adalah fmt.Printf("%c%c%c%c%c\n", angka1, angka2, angka3, angka4, angka5) mencetak lima angka dalam bentuk karakter berdasarkan nilai ASCII dari angka yang dimasukkan. Nilai angka yang dimasukkan akan dikonversi ke karakter ASCII sesuai dengan kode yang diwakili angka tersebut.sedangkan fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1) mencetak tiga karakter dengan menambahkan 1 ke setiap karakter ASCII, sehingga menghasilkan karakter berikutnya dalam urutan ASCII.

Soal 2b
No.1 
Program tersebut merupakan program dari go yaitu melakukan lima kali percobaan dengan menerima input berupa empat nama bunga (atau warna bunga) dari pengguna pada setiap percobaan. Program kemudian memeriksa apakah ada perbedaan antara input yang diberikan dengan set warna tertentu (merah, kuning, hijau, ungu) selama kelima percobaan. Program ini menggunakan perulangan (loop) for i := 1; i <= 5; i++ untuk melakukan lima percobaan. Output dari kode program ini  adalah ketika semua percobaan sesuai dengan urutan warna yang diharapkan (merah, kuning, hijau, ungu), maka BERHASIL : true, tetapi  jika ada satu atau lebih percobaan yang berbeda, maka BERHASIL : false.
No.2
Program tersebut merupakan program dari go  yaitu kita dapat memasukkan nama bunga hingga pengguna memasukkan kata "SELESAI". Setiap bunga yang dimasukkan akan ditambahkan ke sebuah string pita dengan format "bunga - ". Setelah pengguna memasukkan "SELESAI", program menghentikan loop dan mencetak daftar bunga yang telah dimasukkan, serta jumlah total bunga yang dimasukkan. Program ini memanfaatkan perulangan loop for {} yang berjalan terus-menerus hingga kondisi break dipenuhi, ketika input bunga adalah "SELESAI", program menghentikan loop dengan break.Input dari kode program ini adalah mencetak variabel pita yang berisi daftar bunga yang dimasukkan dalam format "bunga1 - bunga2 - bunga3 - ". Lalu program juga mencetak total bunga yang dimasukkan berdasarkan nilai i
No.3
Program tersebut merupakan program dari go yaitu yang meminta kita untuk memasukkan berat belanjaan di dua kantong secara berulang hingga salah satu atau kedua berat kantong mencapai atau melebihi 9 kg. Jika salah satu kantong beratnya 9 kg atau lebih, program menghentikan proses dan menampilkan pesan "Proses selesai", tetapi jika berat dikedua kantong masih kurang dari 9 kg maka program akan mengulang dan meminta input berat belanjaan kembali. Program tersebut memanfaatkan perulangan yaitu loop for {} yang berjalan terus-menerus hingga mencapai kondisi break.
No.4
Program tersebut merupakan program dari go  yaitu menghitung nilai fungsi matematis f(k) berdasarkan input nilai k dari yang kita masukkan. Program menghitung nilai f(k) menggunakan rumus yang telah diberikan. Rumus tersebut melibatkan operasi aritmetika dan fungsi pangkat. math.Pow(4*k + 2, 2) digunakan untuk menghitung kuadrat dari 4k+24k + 24k+2. Output dari kode program ini adalah mencetak nilai f(k) dengan format 10 angka di belakang koma menggunakan fmt.Printf("nilai f(k) = %.10f", fk).
No.5
Program tersebut merupakan program dari go yaitu menghitung nilai fungsi f(k) berdasarkan input k dari pengguna. Fungsi ini dihitung menggunakan perulangan yang mengalikan nilai setiap iterasi dengan rumus yang telah ditentukan. Program ini memanfaatkan perulangan loop for i := 0; i <= k; i++ untuk melakukan iterasi dari 0 hingga k. Output dari kode program tersebut adalah setelah loop selesai, program mencetak nilai f(k) dengan format 10 angka di belakang koma menggunakan fmt.Printf("nilai f(k) = %.10f", fk).

Soal 2c
No.1
Program tersebut merupakan program dari go  yaitu menghitung biaya pengiriman parsel berdasarkan berat dalam gram yang diinput oleh pengguna. Program mengonversi berat tersebut ke dalam kilogram dan gram, menghitung biaya pengiriman berdasarkan berat, dan mencetak detail berat dan biaya. Untuk system perhitungannya adalah beratkg dihitung dengan membagi gram dengan 1000 untuk mendapatkan berat dalam kilogram. beratGram dihitung dengan menggunakan (%) untuk mendapatkan sisa berat dalam gram (yaitu berat setelah dikonversi ke kilogram). Output dari  kode program ini adalah program mencetak detail berat parsel dengan menggunakan format kilogram dan gram menggunakan fmt.Printf("Detail berat: %v kg + %v gr\n", beratkg, beratGram), program juga mencetak detail biaya pengiriman untuk kilogram dan gram, dan total biaya pengiriman dihitung dengan menjumlahkan biayakg dan biayagram, kemudian dicetak.
No.2
Program tersebut merupakan program dari go  yaitu menentukan nilai huruf (grade) berdasarkan nilai akhir mata kuliah yang kita inputkan. Program ini mengkategorikan nilai akhir menjadi beberapa grade sesuai dengan rentang nilai yang telah ditentukan. Program menggunakan serangkaian pernyataan if-else untuk menentukan grade berdasarkan nilai. Output dari program ini adalah mencetak grade (nilai huruf) yang telah ditentukan berdasarkan nilai akhir menggunakan fmt.Println("Nilai mata kuliah: ", nmk).
No.3
Program tersebut merupakan program dari go  yaitu menemukan dan mencetak semua faktor dari bilangan bulat yang diinput oleh pengguna. Dengan menggunakan loop, program ini memeriksa setiap angka dari 1 hingga bilangan tersebut untuk menentukan apakah angka tersebut merupakan faktor. Proses mencetak faktornya adalah program menggunakan loop for i := 1; i < bil+1; i++ untuk mengiterasi dari 1 hingga bilangan yang diinput (inklusif). Dalam setiap iterasi, program memeriksa apakah bil dapat dibagi oleh i tanpa sisa menggunakan kondisi if bil%i == 0. Jika kondisi ini terpenuhi, maka i adalah faktor dari bil, dan program mencetak i dengan fmt.Print(i, " "). Output dari program tersebut adalah Setelah loop selesai, program mencetak baris baru dengan fmt.Println() untuk memberikan jarak setelah daftar faktor yang dicetak. 








