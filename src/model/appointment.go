package model

type Result struct {
	Msg string `json:"msg"`
}

type ReturnAppoint struct {
	Msg         string        `json:"msg"`
	AppointList []Appointment `json:"appoint_list"`
}
type Appointment struct {
	ID                string `json:"id"`
	Date              string `json:"date"`
	HospitalName      string `json:"hospital_name"`
	Name              string `json:"name"`
	Department        string `json:"department"`
	TimeSlot          int    `json:"time_slot"`
	ProfessionalTitle string `json:"professional_title"`
	Price             int    `json:"price"`
	Details           string `json:"details"`
	Status            string `json:"status"`
}
