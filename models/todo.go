package models

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	//Due       time.Time `json:"due"`
}


