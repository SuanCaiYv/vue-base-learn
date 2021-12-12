package db

import (
	"fmt"
	"testing"
)

func TestGridFSDaoService_UploadFile(t *testing.T) {
	instance := NewGridFSDaoService()
	err := instance.DeleteFile("IMG_8559.JPG")
	fmt.Println(err)
}
