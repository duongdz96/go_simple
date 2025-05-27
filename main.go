package main

import (
	"database/sql" // Cần cho G201
	"fmt"
	"log"
	"net/http" // Cần cho G107, G203
	"os"       // Cần cho G304

	// Sử dụng đường dẫn module cho package vulnerabilities
	vulnerabilities "go_simple/vuln"
)

// DummyResponseWriter để cung cấp cho hàm DemonstrateG203
type DummyResponseWriter struct{}

func (drw *DummyResponseWriter) Header() http.Header { return http.Header{} }
func (drw *DummyResponseWriter) Write(p []byte) (int, error) {
	// In ra console thay vì HTTP response thực sự cho demo
	return fmt.Print(string(p))
}
func (drw *DummyResponseWriter) WriteHeader(statusCode int) {
	fmt.Printf("DummyResponseWriter: Status Code %d\n", statusCode)
}

func main() {
	log.Println("Bắt đầu demo các lỗ hổng cho gosec...")
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG101()
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG102()
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG104()
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG107("http://malicious-example.com/data.txt") // SSRF
	vulnerabilities.DemonstrateG107("file:///etc/hosts")                     // Local file access
	log.Println("--------------------------------------")

	var dummyDB *sql.DB // Không cần DB thật, gosec quét mã tĩnh
	vulnerabilities.DemonstrateG201(dummyDB, "admin' OR '1'='1")
	log.Println("--------------------------------------")

	var dummyWriter = &DummyResponseWriter{}
	vulnerabilities.DemonstrateG203(dummyWriter, "<script>alert('XSS Attack!');</script>")
	fmt.Println() // Thêm dòng mới sau output của G203
	log.Println("--------------------------------------")

	// Tạo file tạm để demo G304
	tempFilePath := "temp_g304_test_file.txt"
	err := os.WriteFile(tempFilePath, []byte("some data for G304"), 0644)
	if err != nil {
		log.Fatalf("Không thể tạo file tạm: %v", err)
	}
	vulnerabilities.DemonstrateG304(tempFilePath)
	vulnerabilities.DemonstrateG304("/etc/passwd") // Ví dụ đường dẫn nguy hiểm
	os.Remove(tempFilePath)                        // Dọn dẹp
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG404()
	log.Println("--------------------------------------")

	vulnerabilities.DemonstrateG501("my secret data")
	log.Println("--------------------------------------")

	log.Println("Hoàn thành demo. Hãy chạy 'gosec ./...' trong thư mục dự án")
}
