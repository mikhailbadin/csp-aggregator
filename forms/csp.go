package forms

type CSPReport struct {
	/* Basic directives */
	BlockedURI         string `json:"blocked-uri"`
	Disposition        string `json:"disposition"`
	DocumentURI        string `json:"document-uri"`
	Referrer           string `json:"referrer"`
	ViolatedDirective  string `json:"violated-directive"`
	EffectiveDirective string `json:"effective-directive"`
	OriginalPolicy     string `json:"original-policy"`
	StatusCode         int    `json:"status-code"`
	/* Sctipt directives */
	SourceFile   string `json:"source-file"`
	LineNumber   int    `json:"line-number"`
	ColumnNumber int    `json:"column-number"`
}

type CSPReportMeta struct {
	Report CSPReport `json:"csp-report" binding:"required"`
}

type Headers struct {
	Referer string
}
