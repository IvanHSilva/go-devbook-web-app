package cookies

import (
	"net/http"
	"time"
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

func ReadCookie(r *http.Request) (map[string]string, error) {

	cookie, err := r.Cookie("devbook")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("devbook", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func CleanCookie(w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:     "devbook",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
