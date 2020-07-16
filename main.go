// main.go
package main

import (
	"dogecrypt/src"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dogecrypt.Dogecrypt_init()
	
	action := flag.Bool("")
	flag.Parse()
	
	if == "aesenc" {
		flag.Parse()
		raw := flag.String("filename", "", "the name of the file to be encrypted")
		key := flag.String("key", "ERROR", "key to encrypt the string")
		if *raw != "" || *key != "ERROR" {
			filename := *raw + ".enc"
			fmt.Printf("writing to %s\n", filename)
			content, _ := ioutil.ReadFile(*raw)
			str := dogecrypt.AesEnc(string(content), *key)
			ioutil.WriteFile(filename, []byte(str), 0644)
			fmt.Printf("done")
		} else {
			fmt.Println("ERROR")
			fmt.Printf("raw: %s\nkey: %s", *raw, *key)
			os.Exit(1)
		}
	} else if os.Args[1] == "aesdec" {
		raw := flag.String("raw", "Hello, World!", "the message to be decrypted")
		key := flag.String("key", "7E892875A52C59A3B588306B13C31FBD", "key to encrypt the string")
		flag.Parse()
		fmt.Println(dogecrypt.AesDec(*raw, *key))
	} else if os.Args[1] == "bamboozle" {
		//raw := flag.String("raw", "Hello, World!", "the message to be decrypted")
		//bits := flag.Int("borks", 256, "") //TODO
	} else {
		fmt.Println("ERROR: no operation specified")
	}
}

func fuck() {
	fmt.Println("ERROR")
	return
}
