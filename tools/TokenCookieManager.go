package tools

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Stocker le token API dans un JWT et mettre dans un cookie
func StoreAPITokenInJWT(w http.ResponseWriter, tokenAPI string) {
	claims := jwt.MapClaims{
		"token_api": tokenAPI,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("HASH_KEY")))
	if err != nil {
		http.Error(w, "Erreur lors de la création du token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,  // Activer en production
		MaxAge:   86400, // 24 heures
	})
}

// Récupérer le token API depuis le JWT dans le cookie
func GetAPITokenFromJWT(r *http.Request) (string, error) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature inattendue: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("HASH_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("token invalide")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("impossible d'extraire les claims")
	}

	tokenAPI, ok := claims["token_api"].(string)
	if !ok {
		return "", fmt.Errorf("token API non trouvé")
	}

	return tokenAPI, nil
}
