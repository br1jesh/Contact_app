package user

import (
	"Contact_App/contact"
)

type Role int

const (
	Staff Role = iota
	Admin
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Staff:
		return "Staff"
	default:
		return "Unknown"
	}
}

type User struct {
	UserID    int
	FirstName string
	LastName  string
	Role      Role
	IsActive  bool
	Contacts  []*contact.Contact
}

var (
	users  []*User
	userId = 1
)

func newUser(firstName, lastName string, role Role) *User {
	u := &User{
		UserID:    userId,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		IsActive:  true,
		Contacts:  []*contact.Contact{},
	}
	users = append(users, u)
	userId++
	println("Created User:", firstName, lastName, "Role:", role.String(), "ID:", u.UserID)
	return u
}

func NewAdmin(firstName, lastName string) *User {
	return newUser(firstName, lastName, Admin)
}

func (u *User) NewStaff(firstName, lastName string) *User {
	if !u.checkIsAdminAndIsActive() {
		return nil
	}
	return newUser(firstName, lastName, Staff)
}

func (u *User) checkIsAdminAndIsActive() bool {
	if u.Role != Admin {
		println("Only Admin can perform this operation.")
		return false
	}
	if !u.IsActive {
		println("Inactive Admin cannot perform this operation.")
		return false
	}
	return true
}

func (u *User) checkIsStaffAndIsActive() bool {
	if u.Role != Staff {
		println("Only Staff can perform this operation.")
		return false
	}
	if !u.IsActive {
		println("Inactive Staff cannot perform this operation.")
		return false
	}
	return true
}

func (u *User) GetUserById(id int) *User {
	if !u.checkIsAdminAndIsActive() {
		return nil
	}
	for _, usr := range users {
		if usr.UserID == id && usr.IsActive {
			println("Fetched User ID:", id)
			return usr
		}
	}
	println("User not found with ID:", id)
	return nil
}

func (u *User) GetAllUsers() []*User {
	if !u.checkIsAdminAndIsActive() {
		return nil
	}
	return users
}


func (u *User) UpdateFirstName(target *User, firstName string) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.FirstName = firstName
	println("Updated FirstName for User ID:", target.UserID)
}

func (u *User) UpdateLastName(target *User, lastName string) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.LastName = lastName
	println("Updated LastName for User ID:", target.UserID)
}

func (u *User) UpdateRole(target *User, role Role) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.Role = role
	println("Updated Role for User ID:", target.UserID, "to", role.String())
}

func (u *User) UpdateIsActive(target *User, isActive bool) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.IsActive = isActive
	println("Updated IsActive for User ID:", target.UserID, "to", isActive)
}

func (u *User) UpdateUser(target *User, firstName *string, lastName *string, role *Role, isActive *bool) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	if firstName != nil {
		u.UpdateFirstName(target, *firstName)
	}
	if lastName != nil {
		u.UpdateLastName(target, *lastName)
	}
	if role != nil {
		u.UpdateRole(target, *role)
	}
	if isActive != nil {
		u.UpdateIsActive(target, *isActive)
	}
}

func (u *User) DeleteUser(target *User) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.IsActive = false
	println(" deleted user ID:", target.UserID)
}


func (u *User) CreateContact(fName, lName string) *contact.Contact {
	if !u.checkIsStaffAndIsActive() {
		return nil
	}
	cID := len(u.Contacts) + 1
	newContact := contact.NewContact(fName, lName, cID)
	u.Contacts = append(u.Contacts, newContact)
	println("Created contact", fName, lName, "ID:", cID)
	return newContact
}

func (u *User) GetContactById(contactId int) *contact.Contact {
	if !u.checkIsStaffAndIsActive() {
		return nil
	}
	for _, c := range u.Contacts {
		if c.ContactID == contactId && c.IsActive {
			println(" contact ID:", contactId)
			return c
		}
	}
	println("Contact not found  ID:", contactId)
	return nil
}

func (u *User) GetAllContacts() []*contact.Contact {
	if !u.checkIsStaffAndIsActive() {
		return nil
	}
	return u.Contacts
}

func (u *User) UpdateContact(contactId int, param string, value interface{}) {
	c := u.GetContactById(contactId)
	if c != nil {
		c.UpdateContact(param, value)
	}
}

func (u *User) DeleteContact(contactId int) {
	c := u.GetContactById(contactId)
	if c != nil {
		c.DeleteContact()
	}
}

func (u *User) CreateContactDetail(contactId int, ctype, cvalue string) {
	c := u.GetContactById(contactId)
	if c != nil {
		c.CreateContactDetail(ctype, cvalue)
	}
}

func (u *User) UpdateContactDetail(contactId, detailId int, param string, value interface{}) {
	c := u.GetContactById(contactId)
	if c != nil {
		c.UpdateContactDetail(detailId, param, value)
	}
}

func (u *User) DeleteContactDetail(contactId, detailId int) {
	c := u.GetContactById(contactId)
	if c != nil {
		c.DeleteContactDetail(detailId)
	}
}
