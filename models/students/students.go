package students

import "learn-go/errors"

type StudentModel struct {
	RollNo string `json:"roll_no" bson:"Roll_No"`
	Name   string `json:"name" bson:"Student_Name"`
	Bday   string `json:"bday" bson:"Birth_Date"`
}

func (s *StudentModel) Validate() error {
	ve := errors.ValidationErrs()
	if s.RollNo == "" {
		ve.Add("rollNo", "cannot be empty")
	}
	if s.Name == "" {
		ve.Add("name", "cannot be empty")
	}
	if s.Bday == "" {
		ve.Add("bday", "cannot be empty")
	}
	return ve.Err()
}
