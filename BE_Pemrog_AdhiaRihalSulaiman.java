import java.util.Scanner;

public class BE_Pemrog_AdhiaRihalSulaiman{

    // Fungsi untuk menghitung jumlah bit yang diperlukan
    static int alokasikanMemori(int n) {
        // Cek kombinasi bit [3, 5, 7] untuk mencapai n
        for (int i = 0; i <= n / 3; i++) {
            for (int j = 0; j <= n / 5; j++) {
                for (int k = 0; k <= n / 7; k++) {
                    if (3 * i + 5 * j + 7 * k == n) {
                        return i + j + k; // Kembalikan jumlah bit yang digunakan
                    }
                }
            }
        }
        return -1; // Jika tidak ada solusi
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        // Input jumlah tes alokasi
        int t = sc.nextInt();
        
        // Untuk setiap tes alokasi
        for (int i = 0; i < t; i++) {
            // Input nilai n
            int n = sc.nextInt();
            
            // Alokasikan memori dan cetak hasilnya
            int result = alokasikanMemori(n);
            if (result == -1) {
                System.out.println("TIDAK"); // Tidak bisa dialokasikan
            } else {
                System.out.println(result); // Cetak banyak bit yang digunakan
            }
        }
        
        sc.close();
    }
}
