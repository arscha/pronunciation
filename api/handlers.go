package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const URL = "https://dictionary.cambridge.org/pronunciation/english/"

func handleGetPron(ctx *gin.Context) {
	query := ctx.Params.ByName("query")
	if query == "" {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.Writer.WriteString("Missing query")
		return
	}

	pronResp := PronResponse{}

	err := pronResp.getPron(URL + query)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.Writer.WriteString(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"word pron": pronResp.Pronunciation, "letters": pronResp.LettersPron})
}
