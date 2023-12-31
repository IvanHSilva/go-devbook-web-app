package requests

import (
	"io"
	"net/http"
	"web/src/cookies"
)

func DoRequestWithAuth(r *http.Request, method, url string,
	data io.Reader) (*http.Response, error) {

	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.ReadCookie(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
