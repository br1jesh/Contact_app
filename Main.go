package main

import (
	"Contact_App/user"
)

func main() {
	admin := user.NewAdmin("Brijesh", "Mavani")
	staff := admin.NewStaff("Jay", "Shah")
	contact := staff.CreateContact("Samir", "Patel")

	staff.CreateContactDetails(contact.ContactID, "phone", "1234567890")
	staff.CreateContactDetails(contact.ContactID, "email", "samir@gmail.com")
	staff.CreateContactDetails(contact.ContactID, "address", "Ahmedabad")

	contact.ReadAllContactDetails()
	contact.ReadContactDetailByID(2)
	contact.UpdateContactDetail(2, "Value", "samir.new@gmail.com")
	contact.DeleteContactDetail(1)

	contact.ReadAllContactDetails()


	contact.UpdateContact("FName", "Sameer")
	contact.UpdateContact("LName", "PatelUpdated")
	contact.DeleteContact()

	
	admin.UpdateUser(staff, "FName", "Jaydeep")
	admin.UpdateUser(staff, "LName", "Kapoor")
	admin.UpdateUser(staff, "IsAdmin", true)

	
	admin.DeleteUser(staff)
}
