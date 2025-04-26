package data

// what does the converter need to be abler to convert?
// html/url
// dimensions
// page count
//
// job struct
type Conversion struct {
	HTML string `json:"html,omitempty"`
	URL  string `json:"url,omitempty"`

	PageWidth   float64 `json:"page_width,omitempty"`
	PageHeight  float64 `json:"page_height,omitempty"`
	Orientation string  `json:"orientation,omitempty"`
}

type Response struct {
}
