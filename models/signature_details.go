package models

// SignatureDetails structure for the response from a submitted form
type SignatureDetails struct {
	Name			string `form:"name" validate:"required"`
	JobTitle		string `form:"jobTitle" validate:"required"`
	OfficeNumber	string `form:"officeNo"`
	MobileNumber	string `form:"mobileNo"`
	Email   		string `form:"email" validate:"required,email"`
}
