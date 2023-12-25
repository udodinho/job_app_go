package utils

import "golang.org/x/crypto/bcrypt"

// NormalizePassword func for returning the users input as a byte slice.
func NormalizePassword(p string) []byte {
	return []byte(p)
}

// GeneratePassword func for making hash & salt with user password.
func GeneratePassword(p string) string {
	bytePwd := NormalizePassword(p)

	hash, err := bcrypt.GenerateFromPassword(bytePwd, 14)

	if err != nil {
		return err.Error()
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it.
	return string(hash)
}

// ComparePasswords func for comparing password.
func ComparePasswords(hashedPwd, inputPwd string) bool {
	byteHash := NormalizePassword(hashedPwd)
	
	byteInput := NormalizePassword(inputPwd)

	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}