package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func generateData(total int, start int) {
	idChan := make(chan int, 100)

	for w := 1; w <= 8; w++ {
		wg.Add(1)
		go dataWorker(idChan)
	}

	for id := start; id <= total+start; id++ {
		idChan <- id
	}
	close(idChan)

	wg.Wait()
}

func dataWorker(idChan <-chan int) {
	for id := range idChan {
		secret, _ := hex.DecodeString("988d441fa6c532c583102360a7629691")
		hash1 := hash(secret, generateRandomKey())
		hash2 := hash(secret, generateRandomKey())
		_, err = db.Exec("INSERT INTO key_rotation.accounts VALUES ($1, $2, $3)", id, hash1, hash2)
		if err != nil {
			log.Fatalln(err)
		}
		if id%10000 == 0 {
			log.Printf("inserted %d rows", id)
		}
	}
	defer wg.Done()
}

func generateRandomKey() []byte {
	b := make([]byte, 16)
	rand.Read(b)
	// log.Printf("key: %s\n", hex.EncodeToString(b))
	return b
}
