package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"inventaris/model"
	"inventaris/service"
	"inventaris/utils"
	"inventaris/validation"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthHandler struct {
	service service.UserService
	tmpl    *template.Template
}

func NewAuthHandler(service service.UserService) *AuthHandler {
	tmpl, _ := template.ParseFiles("view/login.html", "view/dashboard.html")

	return &AuthHandler{service: service, tmpl: tmpl}
}

type LoginResponse struct {
	Message string `json:"message"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := validation.ValidateUser(&user, true); err != nil {
			utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		userFromDb, err := h.service.LoginService(user)

		if err != nil || userFromDb == nil {
			utils.SendJSONResponse(w, false, http.StatusBadRequest, "Invalid username or password", nil)
			return
		}

		sessionID := uuid.NewString()
		_, err = h.service.CreateSession(userFromDb.ID)

		if err != nil {
			http.Error(w, "Error creating session", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		w.Header().Set("Content-Type", "application/json")
		utils.SendJSONResponse(w, true, http.StatusOK, "login success", user)

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := validation.ValidateUser(&user, false); err != nil {
			fmt.Println(err, "err validation <<<")
			utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		valid, err := h.service.RepoUser.UsernameExists(user.Username)
		fmt.Println(err, "err username exist <<<", valid)

		if err != nil || valid {
			utils.SendJSONResponse(w, false, http.StatusConflict, "Username is already taken", nil)
			return
		}

		err = h.service.RegisterUser(user)
		if err != nil {
			fmt.Println(err, "register user <<<")
			utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed to register user", nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		utils.SendJSONResponse(w, true, http.StatusOK, "registration success", user)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteSession(cookie.Value)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})

	w.Write([]byte("Logged out successfully"))
}
