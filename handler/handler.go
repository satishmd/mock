package handler

import (
	"github/satishmd/mock/utils"
	"net/http"
)

func DlWidget(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "dlWidget.html")
}
