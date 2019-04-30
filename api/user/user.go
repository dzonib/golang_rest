package user

import (
	"fmt"
	"net/http"
)

type UserAPI struct{}

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got call from %v", r.URL)
}
