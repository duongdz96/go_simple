package vulnerabilities

import (
	"fmt"
	"log"
	"net/http"
)

// G102: Bind to all network interfaces
func DemonstrateG102() {
	// gosec sẽ quét dòng http.ListenAndServe bên dưới và cảnh báo G102.
	// Khối `if false` này để ngăn server thực sự khởi chạy khi chạy demo,
	// nhưng gosec vẫn phân tích mã tĩnh.
	if false {
		http.HandleFunc("/g102test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "G102 Test Endpoint")
		})
		log.Println("G102 Demo: Server would start on 0.0.0.0:8082 if condition was true.")
		http.ListenAndServe("0.0.0.0:8082", nil) // G102 ở đây
	}
	fmt.Println("G102 Demo: Code for binding to all interfaces is present for gosec to scan.")
}
