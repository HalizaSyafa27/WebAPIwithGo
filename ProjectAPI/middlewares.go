package middlewares

import (
    "your-app/helpers"
    "github.com/gin-gonic/gin"
)

// Middleware untuk otentikasi JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Dapatkan token JWT dari header atau parameter
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "Token JWT diperlukan",
            })
            c.Abort()
            return
        }

        // Verifikasi token JWT
        user, err := helpers.VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "Token JWT tidak valid",
            })
            c.Abort()
            return
        }

        // Simpan informasi pengguna dalam konteks Gin
        c.Set("user", user)

        c.Next()
    }
}
