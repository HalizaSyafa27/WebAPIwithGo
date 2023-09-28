package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

// User represents the user model
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

// Photo represents the photo model
type Photo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
	UserID      string `json:"user_id"`
	CreatedOn   string `json:"created_on"`
}

// In-memory database
var (
	users  []User
	photos []Photo
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for user registration
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	// Check if email already exists
	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Email already exists")
			return
		}
	}

	// Generate ID for new user
	user.ID = generateID()

	// Set created and updated timestamps
	user.CreatedOn = time.Now().Format(time.RFC3339)
	user.UpdatedOn = time.Now().Format(time.RFC3339)

	// Add user to the database
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User registered successfully")
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for user login
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	// Find user by email
	var user User
	for _, existingUser := range users {
		if existingUser.Email == credentials.Email {
			user = existingUser
			break
		}
	}

	// Check if user exists and password is correct
	if user.ID == "" || user.Password != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid email or password")
		return
	}

	// Generate JWT token
	token := generateToken(user.ID)

	// Return the token in the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, token)
}

// UpdateUser handles user update
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for user update
	params := mux.Vars(r)
	userID := params["userId"]

	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	// Find user by ID
	var user *User
	for i := range users {
		if users[i].ID == userID {
			user = &users[i]
			break
		}
	}

	// Check if user exists
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}

	// Update user details
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	user.UpdatedOn = time.Now().Format(time.RFC3339)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User updated successfully")
}

// DeleteUser handles user deletion
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for user deletion
	params := mux.Vars(r)
	userID := params["userId"]

	// Find user index by ID
	var userIndex int
	found := false
	for i := range users {
		if users[i].ID == userID {
			userIndex = i
			found = true
			break
		}
	}

	// Check if user exists
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}

	// Delete user from the database
	users = append(users[:userIndex], users[userIndex+1:]...)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User deleted successfully")
}

// CreatePhoto handles photo creation
func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	// Implementation for photo creation
	var photo Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	// Generate ID for new photo
	photo.ID = generateID()

	// Set created timestamp
	photo.CreatedOn = time.Now().Format(time.RFC3339)

	// Add photo to the database
	photos = append(photos, photo)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Photo created successfully")
}

// GetPhotos handles fetching all photos
func GetPhotos(w http.ResponseWriter, r *http.Request) {
	// Implementation for fetching all photos
	json.NewEncoder(w).Encode(photos)
}

// GetPhotoByID handles fetching a photo by ID
func GetPhotoByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for fetching a photo by ID
	params := mux.Vars(r)
	photoID := params["photoId"]

	// Find photo by ID
	var photo *Photo
	for i := range photos {
		if photos[i].ID == photoID {
			photo = &photos[i]
			break
		}
	}

	// Check if photo exists
	if photo == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Photo not found")
		return
	}

	json.NewEncoder(w).Encode(photo)
}

// DeletePhoto handles photo deletion
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	// Implementation for photo deletion
	params := mux.Vars(r)
	photoID := params["photoId"]

	// Find photo index by ID
	var photoIndex int
	found := false
	for i := range photos {
		if photos[i].ID == photoID {
			photoIndex = i
			found = true
			break
		}
	}

	// Check if photo exists
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Photo not found")
		return
	}

	// Delete photo from the database
	photos = append(photos[:photoIndex], photos[photoIndex+1:]...)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Photo deleted successfully")
}

// Middleware function for JWT authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check JWT token and authenticate user
		// Add your JWT authentication logic here

		// Example JWT authentication logic:
		tokenString := r.Header.Get("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Add your JWT verification logic here
			// Example JWT verification logic:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("your-secret-key"), nil
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println("User authenticated:", claims["username"])
			// Add user authentication logic here
			// Example: You can set the authenticated user's ID in the request context
			// to be used in subsequent handlers
			// r = r.WithContext(context.WithValue(r.Context(), "userID", claims["userID"]))
		} else {
			log.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Helper function to generate a unique ID
func generateID() string {
	// Implement your own logic to generate a unique ID
	// Example: You can use a UUID library or generate a random string
	return "unique-id"
}

// Helper function to generate a JWT token
func generateToken(userID string) string {
	// Create the token claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token with the claims and sign it using a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("your-secret-key"))

	return signedToken
}

func main() {
	r := mux.NewRouter()

	// User endpoints
	r.HandleFunc("/users/register", RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", LoginUser).Methods("POST")
	r.HandleFunc("/users/{userId}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", DeleteUser).Methods("DELETE")

	// Photo endpoints
	r.HandleFunc("/photo", CreatePhoto).Methods("POST")
	r.HandleFunc("/photo", GetPhotos).Methods("GET")
	r.HandleFunc("/photos/{photoId}", GetPhotoByID).Methods("GET")
	r.HandleFunc("/photos/{photoId}", DeletePhoto).Methods("DELETE")

	// Apply JWT authentication middleware to protected routes
	r.Use(authMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
