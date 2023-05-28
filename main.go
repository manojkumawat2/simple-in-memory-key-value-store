package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type Entry struct {
	Key    string
	Offset int64
}

type KeyValue map[string]Entry

func (e KeyValue) put(key string, value string, file *os.File) error {
	// Write the value to the file
	offset, err := file.Seek(0, io.SeekEnd)

	if err != nil {
		return err
	}

	_, err = file.WriteString(value)
	if err != nil {
		return err
	}

	// Store the entry in the hash map
	entry := Entry{
		Key:    key,
		Offset: offset,
	}
	e[key] = entry

	return nil
}

func (e KeyValue) get(key string, file *os.File) (string, error) {
	entry, ok := e[key]
	if !ok {
		return "", fmt.Errorf("Key not found")
	}

	// Read the value from the file
	_, err := file.Seek(entry.Offset, io.SeekStart)
	if err != nil {
		return "", err
	}

	value, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func main() {
	filePath := "data.txt"

	// Open the file for read and write operations
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	kv := KeyValue{}

	i := 0
	for i < 5000 {
		err = kv.put(strconv.Itoa(i), "manoj", file)
		if err != nil {
			fmt.Println(err)
			return
		}
		i++
	}

	val, err := kv.get("499", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
