package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type SignatureDetails struct {
	Name			string
	JobTitle		string
	OfficeNumber	string
	MobileNumber	string
	Email   		string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := SignatureDetails{
			Name:			strings.ToUpper(r.FormValue("Name")),
			JobTitle:		strings.ToUpper(r.FormValue("JobTitle")),
			OfficeNumber:	r.FormValue("OfficeNumber"),
			MobileNumber:	r.FormValue("MobileNumber"),
			Email:			r.FormValue("Email"),
		}

		if details.Name == "" || details.JobTitle == "" || details.Email == "" {
			tmpl.Execute(w, struct { 
				Success 	bool
				Error 	bool
				Message	string
			}{false, true, "Must have Name, Job Title, and Email provided"})
			return
		}
		
		fmt.Printf("Name => %s\n", details.Name)
		fmt.Printf("Job Title => %s\n", details.JobTitle)
		fmt.Printf("Office Number => %s\n", details.OfficeNumber)
		fmt.Printf("Mobile Number => %s\n", details.MobileNumber)
		fmt.Printf("EMAIL => %s\n", details.Email)

		data := struct { 
			Success 	bool
			Signature 	SignatureDetails
		}{true, details}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
