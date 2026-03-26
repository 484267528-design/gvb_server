package pwd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPwd,hash密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CheckPwd(HashPwd string, pwd string) bool {
	byteHash := []byte(HashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
