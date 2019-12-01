package main

import (
	"github.com/go-redis/redis"
)

/*GLOBAL VARIABLES*/
var client *redis.Client // Redis Client

type Entry struct {
	Key   string
	Value string
}

func init() {
	// Initialize the redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

/**
 * Add
 */
func Add(key string, pass string) error {
	encrypted, err := Encrypt([]byte(CreateHash("string")), []byte(pass))
	if err != nil {
		return err
	}
	err = client.Set(key, encrypted, 0).Err()
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
	decrypted, err := Decrypt([]byte(CreateHash("string")), []byte(val))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
