package main

import (
	"encoding/hex"
)

func findAccount(secretKey, accountkey string) bool {
	secretKeyBytes, _ := hex.DecodeString(secretKey)
	accountKeyBytes, _ := hex.DecodeString(accountkey)
	hash := hash(secretKeyBytes, accountKeyBytes)

	row := db.QueryRow("SELECT id FROM key_rotation.accounts WHERE hash1 = $1 OR hash2 = $1", hash)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}
