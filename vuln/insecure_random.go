package vulnerabilities

import (
	"fmt"
	"math/rand" // G404
)

// G404: Use of weak random number generator (math/rand)
func DemonstrateG404() {
	// math/rand không an toàn cho các giá trị cần tính bảo mật (ví dụ: token)
	insecureSessionID := rand.Int63()
	fmt.Printf("G404 Demo: Insecurely generated session ID: %d\n", insecureSessionID)
}
