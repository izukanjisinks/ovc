package handlers

import (
	"net/http"

	"github.com/izukanji/ovc/pkg/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.OK(w, map[string]string{"status": "ok"})
}
