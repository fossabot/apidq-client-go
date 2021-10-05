package apidq

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikitaksv/apidq-client-go/dto/address"
)

const TestAPIKey = "testApiKey123"

func NewTestClient(h http.HandlerFunc) (*Client, *httptest.Server) {
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				authKey := r.Header.Get(authorization)
				if authKey == "" {
					w.WriteHeader(http.StatusUnauthorized)
					if _, err := w.Write([]byte(`{"code": 16, "message": "Ключ API обязателен"}`)); err != nil {
						panic(err)
					}
					return
				} else if authKey != TestAPIKey {
					if _, err := w.Write([]byte(`{"code":16,"message":"Ошибка авторизации"}`)); err != nil {
						panic(err)
					}
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				h(w, r)
			},
		),
	)
	c, err := NewClient(http.DefaultClient, s.URL)
	if err != nil {
		panic(err)
	}
	return c.WithAuth(TestAPIKey), s
}

func TestAuth(t *testing.T) {
	client, tS := NewTestClient(func(w http.ResponseWriter, r *http.Request) {})
	defer tS.Close()

	_, _, err := client.WithReqOptions(func(r *http.Request) error {
		r.Header.Del(authorization)
		return nil
	}).Address.Clean(context.Background(), &address.CleanRequest{})
	if err == nil {
		t.Fatal(errors.New("need ErrorResponse"))
	}
	if err.Error() != "[16] Ключ API обязателен" {
		t.Fatal(err)
	}

	_, _, err = client.WithAuthService(TestAPIKey, "address").Address.Clean(context.Background(), &address.CleanRequest{})
	if err != nil {
		t.Fatal(err)
	}

	client, tS = NewTestClient(func(w http.ResponseWriter, r *http.Request) {})
	defer tS.Close()

	_, _, err = client.Address.Clean(context.Background(), &address.CleanRequest{})
	if err != nil {
		t.Fatal(err)
	}
}
