package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// HashPassword hashes the password using SHA-256
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// CheckPasswordHash compares a hashed password with its plaintext equivalent
func CheckPasswordHash(password, hashedPassword string) bool {
	// Hash the incoming plain password
	log.Println(HashPassword("1234"))
	passwordHash := HashPassword(password)
	log.Println("Entered password", passwordHash)
	log.Println("stored password", hashedPassword)
	// Compare the hashed input with the stored hashed password
	return passwordHash == hashedPassword
}
