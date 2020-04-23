package prstat

// Aggregate ...
type Aggregate struct {
	PRCount               int `json:"prCount"`
	AvgReviewCommentCount int `json:"avgReviewCommentCount"`
	AvgMinutesOpen        int `json:"avgDescription"`
	AvgReviewCount        int `json:"avgReviewCount"`
	AvgApproverCount      int `json:"avgApproverCount"`
}

// PRStats ...
type PRStats struct {
	PRStats []PRStat `json:"prstats"`
}

// PRStat ...
type PRStat struct {
	ReviewCommentCount int `json:"reviewCommentCount"`
	MinutesOpen        int `json:"description"`
	ReviewCount        int `json:"reviewCount"`
	ApproverCount      int `json:"approverCount"`
}
