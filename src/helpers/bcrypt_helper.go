package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func CompareHash(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)

}
