package endpoints

type Status struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
