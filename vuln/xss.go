package vulnerabilities

import (
	"fmt"
	"net/http"
)

// G203: Use of unescaped data in HTML response
func DemonstrateG203(w http.ResponseWriter, userInput string) {
	// Giả sử hàm này là một phần của HTTP handler.
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Xin chào, %s!</h1>", userInput) // G203: userInput được in trực tiếp
}
