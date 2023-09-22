package router;

import (
    "your-app/controllers"
    "your-app/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Middleware untuk otentikasi JWT
    r.Use(middlewares.AuthMiddleware())

    // Endpoint untuk mengunggah gambar profil
    r.POST("/profile/upload", controllers.UploadProfileImage)

    // Endpoint untuk menghapus gambar profil
    r.DELETE("/profile/delete", controllers.DeleteProfileImage)

    return r
}
