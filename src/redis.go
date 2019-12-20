package main

import (
	"github.com/go-redis/redis"
)

/*GLOBAL VARIABLES*/
var client *redis.Client // Redis Client

type Entry struct {
	Index int
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
	encrypted, err := Encrypt([]byte(config.Hash), []byte(pass))
	if err != nil {
		return err
	}
	err = client.LPush(config.Key, encrypted).Err()
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
	list, err := client.LRange(config.Key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for i := range list {
		val, err := Decrypt([]byte(config.Hash), []byte(list[i]))
		if err != nil {
			entries = append(entries, Entry{
				Index: i,
				Value: "Could not be decrypted",
			})
		} else {
			entries = append(entries, Entry{
				Index: i,
				Value: string(val),
			})
		}
	}

	return entries, nil

}
