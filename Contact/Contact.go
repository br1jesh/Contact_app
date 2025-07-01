package contact

import (
	"Contact_App/contactDetails"
	"fmt"
)

type Contact struct {
	ContactID      int
	FName          string
	LName          string
	IsActive       bool
	ContactDetails []*contactDetails.ContactDetail
}

func NewContact(fName, lName string, id int) *Contact {
	c := &Contact{
		ContactID:      id,
		FName:          fName,
		LName:          lName,
		IsActive:       true,
		ContactDetails: []*contactDetails.ContactDetail{},
	}
	fmt.Println("Created Contact:", fName, lName, "(ID:", c.ContactID, ")")
	return c
}

func (c *Contact) UpdateContact(param string, value interface{}) {
	if !c.IsActive {
		fmt.Println("Contact is inactive; update skipped.")
		return
	}
	switch param {
	case "FName":
		c.updateFirstName(value)
	case "LName":
		c.updateLastName(value)
	case "IsActive":
		c.updateIsActiveStatus(value)
	default:
		fmt.Println("Unknown parameter:", param)
	}
}

func (c *Contact) updateFirstName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.FName = str
		fmt.Println("First name updated successfully.")
	} else {
		fmt.Println("updateFirstName: invalid string")
	}
}

func (c *Contact) updateLastName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.LName = str
		fmt.Println("Last name updated successfully.")
	} else {
		fmt.Println("updateLastName: invalid string")
	}
}

func (c *Contact) updateIsActiveStatus(value interface{}) {
	if status, ok := value.(bool); ok {
		c.IsActive = status
		fmt.Println("IsActive status changed to:", status)
	} else {
		fmt.Println("updateIsActiveStatus: invalid bool")
	}
}


func (c *Contact) DeleteContact() {
	c.IsActive = false
	fmt.Println("Contact", c.FName, c.LName, "(ID:", c.ContactID, ") has been soft deleted (IsActive=false).")
}

//ContactDetails------->>>>>>>

func (c *Contact) CreateContactDetail(typeName, value string) {
	if !c.IsActive {
		fmt.Println("Cannot add contact details to inactive contact.")
		return
	}
	detail := contactDetails.NewContactDetail(typeName, value)
	c.ContactDetails = append(c.ContactDetails, detail)
	fmt.Println("Contact detail added successfully:", typeName, "-", value)
}

func (c *Contact) ReadAllContactDetails() {
	if len(c.ContactDetails) == 0 {
		fmt.Println("No contact details found.")
		return
	}
	for _, detail := range c.ContactDetails {
		fmt.Printf("ID: %d, Type: %s, Value: %s\n", detail.ContactDetailsID, detail.Type, detail.Value)
	}
}


func (c *Contact) ReadContactDetailByID(id int) *contactDetails.ContactDetail {
	for _, detail := range c.ContactDetails {
		if detail.ContactDetailsID == id {
			fmt.Printf("Found ContactDetail: ID=%d, Type=%s, Value=%s\n", detail.ContactDetailsID, detail.Type, detail.Value)
			return detail
		}
	}
	fmt.Println("Contact detail not found for ID:", id)
	return nil
}

func (c *Contact) UpdateContactDetail(id int, param string, value interface{}) {
	detail := c.ReadContactDetailByID(id)
	if detail != nil {
		detail.UpdateContact(param, value)
	}
}


func (c *Contact) DeleteContactDetail(id int) {
	index := -1
	for i, detail := range c.ContactDetails {
		if detail.ContactDetailsID == id {
			index = i
			break
		}
	}
	if index != -1 {
		c.ContactDetails = append(c.ContactDetails[:index], c.ContactDetails[index+1:]...)
		fmt.Println("ContactDetail with ID", id, "has been deleted permanently.")
	} else {
		fmt.Println("Contact detail not found for deletion with ID:", id)
	}
}


