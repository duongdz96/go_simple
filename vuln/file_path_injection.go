package vulnerabilities

import (
	"fmt"
	"io/ioutil"
)

// G304: File path provided as taint input
func DemonstrateG304(untrustedFilePath string) {
	// Giả sử untrustedFilePath đến từ nguồn không đáng tin cậy
	data, err := ioutil.ReadFile(untrustedFilePath)
	if err != nil {
		fmt.Printf("G304 Demo: Error reading file '%s': %v\n", untrustedFilePath, err)
		return // Xử lý lỗi
	}
	fmt.Printf("G304 Demo: Read %d bytes from '%s'.\n", len(data), untrustedFilePath)
}
