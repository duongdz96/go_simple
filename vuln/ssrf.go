package vulnerabilities

import (
	"fmt"
	"net/http"
)

// G107: Url provided to HTTP request as taint input (Potential SSRF)
func DemonstrateG107(untrustedURL string) {
	// Giả sử untrustedURL đến từ người dùng hoặc nguồn không đáng tin cậy
	resp, err := http.Get(untrustedURL) // G107
	if err != nil {
		fmt.Printf("G107 Demo: Error making HTTP GET to %s: %v\n", untrustedURL, err)
		return // Xử lý lỗi
	}
	defer resp.Body.Close()
	fmt.Printf("G107 Demo: HTTP GET request to %s successful (status: %s).\n", untrustedURL, resp.Status)
}
