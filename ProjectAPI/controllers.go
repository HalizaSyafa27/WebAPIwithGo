package controllers;

import (
    "net/http"
    "your-app/helpers"
    "github.com/gin-gonic/gin"
)

// Handler untuk mengunggah gambar profil
func UploadProfileImage(c *gin.Context) {
    // Dapatkan informasi pengguna dari token JWT
    var user := helpers.GetUserFromToken(c)

    // Lakukan logika unggah gambar dan simpan ke database
    // Pastikan gambar diupload oleh pengguna yang sesuai

    c.JSON(http.StatusOK, gin.H{
        "message": "Gambar profil berhasil diunggah",
    })
}

// Handler untuk menghapus gambar profil
func DeleteProfileImage(c *gin.Context) {
    // Dapatkan informasi pengguna dari token JWT
    user := helpers.GetUserFromToken(c)

    // Lakukan logika penghapusan gambar profil
    // Pastikan hanya pengguna yang memiliki gambar profil yang dapat menghapusnya

    c.JSON(http.StatusOK, gin.H{
        "message": "Gambar profil berhasil dihapus",
    })
}
