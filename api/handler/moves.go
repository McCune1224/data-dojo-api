package handler

type MoveResponse struct {
    ID        uint   `json:"id"`
	Name      string `json:"name"`
	Input     string `json:"input"`
	Startup   string `json:"startup"`
	Active    string `json:"active"`
	Recovery  string `json:"recovery"`
	OnBlock   string `json:"on_block"`
	OnHit     string `json:"on_hit"`
	OnCounter string `json:"on_counter"`
}
