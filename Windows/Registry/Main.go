package main

import (
	"log"
	"fmt"
	
	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	fmt.Printf("Displaying programs that run on startup . . .\n")
	stat, err := k.Stat()
	if err != nil {
		log.Fatal(err)
	}
	valueCount := stat.ValueCount
	values, err := k.ReadValueNames(int(valueCount))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)

	fmt.Println("Setting test value . . .")

	err = k.SetStringValue("notepad", "C:\\windows\\notepad.exe")
	if err != nil {
		log.Fatal(err)
	}
	
}

