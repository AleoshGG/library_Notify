package domain

type Reader struct {
	IdReader      int
	First_name     string
	Last_name      string
	Email         string
	Account_status string
}

type Links struct {
	Self string
}

type Response struct {
	Data   []Reader
	Links  Links
	Status bool
}