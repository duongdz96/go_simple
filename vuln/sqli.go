package vulnerabilities

import (
	"database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3" // Import driver nếu muốn chạy thực tế
)

// G201/G202: SQL query constructed from user input
func DemonstrateG201(db *sql.DB, userInput string) {
	// Lỗ hổng SQL Injection do nối chuỗi trực tiếp
	query := "SELECT data FROM products WHERE name = '" + userInput + "'"
	fmt.Printf("G201/G202 Demo: Vulnerable SQL Query: %s\n", query)

	// Để gosec phân tích, không nhất thiết phải có DB thực sự chạy.
	// Nếu bạn muốn chạy thử, bỏ comment các dòng sau và import driver:
	// if db != nil {
	//     rows, err := db.Query(query) // gosec sẽ cảnh báo dòng này
	//     if err != nil {
	//         fmt.Printf("G201 Demo: SQL query execution error: %v\n", err)
	//         return
	//     }
	//     defer rows.Close()
	//     fmt.Println("G201 Demo: Query executed (conceptually).")
	// }
}
