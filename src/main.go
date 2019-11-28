package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	//"encoding/hex"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	r "math/rand"
	"time"
)

type Entry struct {
	Key   string
	Value string
}

/** GLOBAL VARIABLES **/
var client *redis.Client // Redis Client
var cli bool             // Boolean flag to determine if the program should run as gui or terminal cli
var help bool            // Help flag
var add string           // Add new password to the keychain in CLI mode
var del string           //	Delete password from the keychain in CLI mode
var list bool            // List all passwords in the keychain in CLI mode

/**
 * Initialization Function
 */
func init() {
	// Initialize the redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	flag.BoolVar(&cli, "c", false, "Use terminal cli mode")
	flag.BoolVar(&help, "h", false, "Display flag options and usage")
	flag.StringVar(&add, "a", "", "Add a new `password` to the kaychain in CLI mode")
	flag.StringVar(&del, "d", "", "Delete a `password` from the keychain in CLI mode")
	flag.BoolVar(&list, "l", false, "List all passwords in the keychain in CLI mode")
	flag.Parse()
}

/**
 * Generate Key
 */
func GenerateKey() string {
	r.Seed(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$%^&*()")
	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

/**
 * Add
 */
func Add(key string, pass string) error {
	err := client.Set(key, Encrypt([]byte("try this out"), pass), 0).Err()
	if err != nil {
		return err
	}

	return nil
}

/**
 * Delete
 */
func Delete(key string) error {
	err := client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil
}

/**
 * List
 */
func List() ([]Entry, error) {
	entries := []Entry{}
	list, err := client.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	for i := range list {
		val, err := Get(list[i])
		if err != nil {

		} else {
			entries = append(entries, Entry{
				Key:   list[i],
				Value: val,
			})
		}
	}

	return entries, nil

}

func Get(key string) (string, error) {
	val, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return Decrypt([]byte("try this out"), val), nil
}

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//return hex.EncodeToString(hasher.Sum(nil))
	return "qwertyuiopasdfghjklzxcvbnm123456"
}

func Encrypt(data []byte, pass string) string {
	block, _ := aes.NewCipher([]byte("qwertyuiopasdfghjklzxcvbnm123456"))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error())
	}
	encrypted := gcm.Seal(nonce, nonce, data, nil)

	return string(encrypted)
}

func Decrypt(data []byte, pass string) string {
	key := []byte("qwertyuiopasdfghjklzxcvbnm123456")
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, encrypted := data[:nonceSize], data[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(decrypted)
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
				for i := range entries {
					fmt.Printf("Key: %s, Value: %s\n", entries[i].Key, entries[i].Value)
				}
			}

		} else {
			fmt.Println("A flag is required")
			flag.PrintDefaults()
		}

	}
}
