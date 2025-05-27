package vulnerabilities

import (
	"crypto/md5" // G501: Import thư viện MD5 nằm trong danh sách đen
	"fmt"
	"io"
)

// G501: Use of MD5 (flagged due to import)
func DemonstrateG501(dataToHash string) {
	hasher := md5.New()
	io.WriteString(hasher, dataToHash)
	hashedData := fmt.Sprintf("%x", hasher.Sum(nil))
	fmt.Printf("G501 Demo: MD5 Hash of '%s': %s\n", dataToHash, hashedData)
}
