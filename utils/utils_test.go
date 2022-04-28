package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain_InputScan(t *testing.T) {
	str, err := InputScan()
	assert.Error(t, err, "Harus Error")
	assert.EqualValues(t, str, "", "kosong")
}

func TestMain_TimeDateNow_BedaFormat(t *testing.T) {
	assert.NotEqual(t, TimeDateNow(), time.Now().String(), "Cek tanggal")
}

func TestMain_TimeDateNow_Ok(t *testing.T) {
	date := TimeDateNow()
	strDate := time.Now().Format(formatLayout)
	assert.Equal(t, date, strDate, "Harus sesuai")
}
