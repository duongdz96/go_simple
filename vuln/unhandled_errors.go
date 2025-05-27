package vulnerabilities

import (
	"fmt"
	"os"
)

// G104: Errors unhandled
func DemonstrateG104() {
	file, err := os.Open("some_random_file_that_does_not_exist.txt")
	// Nếu không có kiểm tra lỗi 'if err != nil' ở đây, gosec sẽ cảnh báo.
	// Hiện tại, ta kiểm tra lỗi, nhưng lỗi từ file.Close() không được xử lý.

	if err != nil {
		fmt.Printf("G104 Demo: Error opening file (handled for now): %v\n", err)
		// Nếu không return ở đây, và file là nil, file.Close() sẽ panic.
		// return
	}

	// Giả sử file được mở thành công (dù trong thực tế là không)
	if file != nil {
		errClose := file.Close() // G104: Lỗi từ file.Close() không được kiểm tra và xử lý.
		if errClose != nil {
			// Đây là cách xử lý đúng, nếu bỏ đi sẽ là G104
			// fmt.Printf("G104 Demo: Error closing file: %v\n", errClose)
		}
	}
	fmt.Println("G104 Demo: An error from file.Close() might be unhandled.")
}
