package domain

import "testing"

func TestCacheFileWriter_Write(t *testing.T) {
	dir := "cache"
	filename := "iris.json"
	writer := NewCacheFileWriter(dir, filename)
	err := writer.Write(2)
	if err != nil {
		t.Fatal(err)
	}

}
