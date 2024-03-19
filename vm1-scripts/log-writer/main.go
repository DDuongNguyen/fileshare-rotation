package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeChunk(filePath string, data []byte) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

func writeNoCloseChunk(filePath string, data []byte) (*os.File, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(data)
	return f, err
}

func main() {
	var f *os.File
	var err error

	logFilePath := "/logs/writer.log"
	logFilePath2 := "/logs/writer-not-closed.log"

	chunkSize := 1024 * 10 // 10 KB chunks
	data := make([]byte, 0, chunkSize*2)
	data2 := make([]byte, 0, chunkSize*2)

	for i := 0; i < chunkSize; i++ {
		data = append(data, 'A')
		data = append(data, []byte(strconv.Itoa(i+1))...) // Convert integer to string, then to byte slice
		data = append(data, '\n')
	}

	for i := 0; i < chunkSize; i++ {
		data2 = append(data2, 'B')
		data2 = append(data2, []byte(strconv.Itoa(i+1))...) // Convert integer to string, then to byte slice
		data2 = append(data2, '\n')
	}

	fmt.Printf("chunk size: %d\n", len(data))
	for i := 0; i < 500; i++ {
		err := writeChunk(logFilePath, data)
		if err != nil {
			fmt.Println("Error writing chunk:", err)
		} else {
			// fmt.Printf("Iteration %d: Wrote chunk, pausing for 1 microseconds\n", i+1)
			fmt.Printf("Closed Iteration %d\n", i+1)
			// time.Sleep(1 * time.Microsecond)
		}
	}

	for i := 0; i < 500; i++ {
		f, err = writeNoCloseChunk(logFilePath2, data2)
		if err != nil {
			fmt.Println("Error writing chunk:", err)
		} else {
			fmt.Printf("No Closed Iteration %d\n", i+1)
		}
	}

	if f != nil {
		err = f.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}
}
