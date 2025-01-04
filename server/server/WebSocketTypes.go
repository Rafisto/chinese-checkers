package server

type WSRequest struct {
	Type     string `json:"type"`
	Action   string `json:"action"`
	PlayerID int    `json:"player_id,omitempty"`
	Start    struct {
		Row int `json:"row"`
		Col int `json:"col"`
	} `json:"start,omitempty"`
	End struct {
		Row int `json:"row"`
		Col int `json:"col"`
	} `json:"end,omitempty"`
}
