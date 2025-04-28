package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IsiahMc/HTML2PDF/internal/convert"
	"github.com/IsiahMc/HTML2PDF/internal/data"
)

// POST : /convert
func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	data := data.Conversion{
		PageWidth:  8.5,
		PageHeight: 11,
	}
	// 1. Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	// 2. Validate the input (either HTML or URL must be present)
	var source string
	if data.URL != "" {
		source = data.URL
	} else if data.HTML != "" {
		source = data.HTML
	} else {
		http.Error(w, "Please submit either html or a url", http.StatusBadRequest)
	}
	// 3. Create a new job
	pdf, err := convert.Convert(ctx, source, data)
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// 4. Return job ID/status to client
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=document.pdf")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdf)))

	// Write PDF to response
	w.Write(pdf)
}

// GET : /job
