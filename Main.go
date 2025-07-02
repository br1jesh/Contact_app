package main

import (
	user "Contact_App/User"
	"fmt"
)

func main() {
	admin := user.NewAdmin("Vishav", "Pathania")
	fmt.Println("Admin created:", *admin)
	staff := admin.NewStaff("Brijesh", "Mavani")
	fmt.Println("Staff created:", *staff)
	contact := staff.CreateContact("Aniket", "Pardeshi")
	fmt.Println("Contact created:", *contact)

	staff.CreateContactDetail(contact.ContactID, "Email", "aniketpardeshi@gmail.com")
	staff.CreateContactDetail(contact.ContactID, "Phone", "9876543210")
	staff.CreateContactDetail(contact.ContactID, "email", "ABS@gmail.com")

	contact.ReadAllContactDetails()

	contact.UpdateContactDetail(2, "Value", "9876501234")
	contact.ReadAllContactDetails()

	contact.DeleteContactDetail(1)
	contact.ReadAllContactDetails()

	contact.UpdateContact("FName", "AniketUpdated")
	contact.UpdateContact("LName", "PardeshiUpdated")
	staff.DeleteContact(contact.ContactID)

	foundContact := staff.GetContactById(contact.ContactID)
	fmt.Println("Retrieved after delete:", foundContact)

	for _, u := range admin.GetAllUsers() {
		fmt.Println(u)
	}

	admin.DeleteUser(staff)

	for _, u := range admin.GetAllUsers() {
		fmt.Println(u)
	}
}
