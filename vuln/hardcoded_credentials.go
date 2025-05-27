package vulnerabilities

import "fmt"

// G101: Potential hardcoded credentials
const SuperSecretApiKey = "aks_thisIsMySuperSecretKey12345!"

func DemonstrateG101() {
	fmt.Printf("G101 Demo: Using API Key: %s\n", SuperSecretApiKey)
}
