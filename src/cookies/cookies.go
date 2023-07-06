package cookies

import (
	"net/http"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func ConfigureCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, id, token string) error {

	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encryptData, err := s.Encode("devbook", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "devbook",
		Value:    encryptData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
