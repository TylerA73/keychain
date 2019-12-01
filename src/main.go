package main

import (
	"flag"
	"fmt"
)

/** GLOBAL VARIABLES **/
var cli bool   // Boolean flag to determine if the program should run as gui or terminal cli
var help bool  // Help flag
var add string // Add new password to the keychain in CLI mode
var del string //	Delete password from the keychain in CLI mode
var list bool  // List all passwords in the keychain in CLI mode

/**
 * Initialization Function
 */
func init() {

	flag.BoolVar(&cli, "c", false, "Use terminal cli mode")
	flag.BoolVar(&help, "h", false, "Display flag options and usage")
	flag.StringVar(&add, "a", "", "Add a new `password` to the kaychain in CLI mode")
	flag.StringVar(&del, "d", "", "Delete a `password` from the keychain in CLI mode")
	flag.BoolVar(&list, "l", false, "List all passwords in the keychain in CLI mode")
	flag.Parse()
}

/**
 * Main Function
 */
func main() {
	/*s := "key"
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(s, sha1_hash, GenerateKey())*/

	var err error

	if help {
		flag.PrintDefaults()
	} else if !cli {
		fmt.Println("GUI MODE")
	} else {
		fmt.Println("CLI MODE")

		if add != "" {
			err = Add(GenerateKey(), add)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Key added to keychain")
			}
		} else if del != "" {
			err = Delete(del)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Key deleted from keychain")
			}
		} else if list {
			entries, err := List()
			if err != nil {
				fmt.Println(err)
			} else {
				if len(entries) < 1 {
					fmt.Println("You have no keys stored")
				} else {
					for i := range entries {
						fmt.Printf("Key: %s, Value: %s\n", entries[i].Key, entries[i].Value)
					}
				}
			}

		} else {
			//fmt.Println("A flag is required")
			//flag.PrintDefaults()

			s := "password"
			enc, _ := Encrypt(CreateHash("string"), []byte(s))
			fmt.Println("Encrypted: ", string(enc))
			dec, _ := Decrypt(CreateHash("string"), []byte(enc))
			fmt.Println("Decrypted: ", string(dec))
		}

	}
}
