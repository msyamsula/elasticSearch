package utils

type JsonBody struct {
	Size  int64 `json:"size"`
	Query `json:"query"`
}

type Boost struct {
	MP struct {
		Title struct {
			Q string  `json:"query"`
			B float64 `json:"boost,omitempty"`
		} `json:"title"`
	} `json:"match_phrase"`
}

type BoostingQuery struct {
	Q struct {
		B struct {
			S []interface {
			} `json:"should"`
		} `json:"bool"`
	} `json:"query"`

	S int64 `json:"size"`
}

type ExplainBody struct {
	Query Query `json:"query"`
}

type Query struct {
	MultiMatch `json:"multi_match"`
}

type MultiMatch struct {
	QueryString string   `json:"query"`
	Fields      []string `json:"fields"`
	Type        string   `json:"type,omitempty"`
	TieBreaker  string   `json:"tie_breaker,omitempty"`
}

type Response struct {
	BigHit `json:"hits"`
	Error  string `json:"error"`
}

type BigHit struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Source map[string]interface{} `json:"_source"`
	Score  float64                `json:"_score"`
	ID     string                 `json:"_id"`
}

type ExplainResponse struct {
	Explanations []Explanation `json:"explanations"`
}

type Explanation struct {
	Index string `json:"index"`
	Valid bool   `json:"valid"`
	Exp   string `json:"explanation"`
}
