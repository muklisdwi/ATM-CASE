package utils

import (
	"bufio"
	"os"
	"strings"
	"time"
)

const formatLayout = "2006-01-02 15:04:05"

// fungsi untuk menerima input keyboard
func InputScan() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, " ", "")
	str = strings.Replace(str, "\n", "", 1)
	return str, nil
}

// fungsi convert time.Time
func TimeDateNow() string {
	now := time.Now()
	dateNow := now.Format(formatLayout)
	return dateNow
}
