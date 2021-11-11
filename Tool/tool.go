package Tool

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
)

func Base64KeysGenerator(size int) string {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func PathExists(path string) (bool,error){
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}