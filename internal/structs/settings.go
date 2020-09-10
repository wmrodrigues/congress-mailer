package structs

type Settings struct {
	Smtp struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		From     string `json:"from"`
	} `json:"smtp"`
	Mail struct {
		Subject string `json:"subject"`
	} `json:"mail"`
}