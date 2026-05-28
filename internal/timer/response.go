package timer

type Response struct {
	Success          bool   `json:"success"`
	Message          string `json:"msg"`
	Type             string `json:"type"`
	QueryTime        int64  `json:"queryTime"`
	Estimate         int64  `json:"estimate"`
	EstimateRelative string `json:"estimateRelative"`
	Num              int    `json:"num"`
}
