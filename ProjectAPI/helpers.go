package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

// Fungsi untuk mendapatkan informasi pengguna dari token JWT
func GetUserFromToken(c *gin.Context) *User {
    user, _ := c.Get("user")
    return user.(*User)
}

// Fungsi untuk memverifikasi token JWT
func VerifyToken(tokenString string) (*User, error) {
    // Lakukan verifikasi token JWT menggunakan library JWT Go
    // Kembalikan informasi pengguna jika token valid, jika tidak kembalikan error
}
