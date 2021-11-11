package Tool

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func Base64KeysGenerator(size int) string{
	bytes := make([]byte,size)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

