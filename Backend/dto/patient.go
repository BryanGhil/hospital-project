package dto

type Patient struct {
	ID        int    `json:"patient_id"`
	FullName  string `json:"full_name"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	CreatedBy int    `json:"created_by,omitempty"`
}

type ReqAddPatient struct {
	FullName string `json:"full_name" binding:"required"`
	DOB      string `json:"dob" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
