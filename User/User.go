package user

import (
	contact "Contact_App/Contact"
	"fmt"
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

var users []*User
var userId = 1

func newUser(firstName, lastName string, role Role) *User {
	u := &User{
		UserID:    userId,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		IsActive:  true,
		Contacts:  []*contact.Contact{},
	}
	fmt.Println("Created User:", firstName, lastName, "(Role:", role.String(), ", ID:", u.UserID, ")")
	userId++
	users = append(users, u)
	return u
}

func NewAdmin(firstName, lastName string) *User {
	return newUser(firstName, lastName, Admin)
}

func (u *User) NewStaff(firstName, lastName string) *User {
	if u.Role != Admin {
		panic("Only Admin can create Staff. Current user role: " + u.Role.String())
	}
	return newUser(firstName, lastName, Staff)
}

func (u *User) checkIsAdminAndIsActive() bool {
	if u.Role != Admin {
		fmt.Println("Only Admin can perform this operation.")
		return false
	}
	if !u.IsActive {
		fmt.Println("Inactive Admin cannot perform this operation.")
		return false
	}
	return true
}

func (u *User) checkIsStaffAndIsActive() bool {
	if u.Role != Staff {
		fmt.Println("Only Staff can manage contacts.")
		return false
	}
	if !u.IsActive {
		fmt.Println("Inactive Staff cannot manage contacts.")
		return false
	}
	return true
}

func (u *User) GetAllUsers() []*User {
	if !u.checkIsAdminAndIsActive() {
		return nil
	}
	return users
}

func (u *User) UpdateUser(target *User, param string, value interface{}) {
	switch param {
	case "FName":
		u.UpdateUserFirstName(target, value)
	case "LName":
		u.UpdateUserLastName(target, value)
	case "IsAdmin":
		u.UpdateIsAdminStatus(target, value)
	case "IsActive":
		u.UpdateUserIsActiveStatus(target, value)
	default:
		fmt.Println("Unknown parameter:", param)
	}
}

func (u *User) UpdateUserFirstName(target *User, value interface{}) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("UpdateUserFirstName: invalid string")
		return
	}
	target.FirstName = strVal
	fmt.Println("First name updated successfully.")
}

func (u *User) UpdateUserLastName(target *User, value interface{}) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	strVal, ok := value.(string)
	if !ok || strVal == "" {
		fmt.Println("UpdateUserLastName: invalid string")
		return
	}
	target.LastName = strVal
	fmt.Println("Last name updated successfully.")
}

func (u *User) UpdateIsAdminStatus(target *User, value interface{}) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	isAdmin, ok := value.(bool)
	if !ok {
		fmt.Println("UpdateIsAdminStatus: invalid bool")
		return
	}
	if isAdmin {
		target.Role = Admin
	} else {
		target.Role = Staff
	}
	fmt.Println("Role updated successfully.")
}

func (u *User) UpdateUserIsActiveStatus(target *User, value interface{}) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	status, ok := value.(bool)
	if !ok {
		fmt.Println("UpdateUserIsActiveStatus: invalid bool")
		return
	}
	target.IsActive = status
	fmt.Println("IsActive status changed to:", status)
}

func (u *User) DeleteUser(target *User) {
	if !u.checkIsAdminAndIsActive() {
		return
	}
	target.IsActive = false
	fmt.Println("User", target.FirstName, target.LastName, "(ID:", target.UserID, ") has been soft deleted (IsActive=false).")
}

func (u *User) CreateContact(fName, lName string) *contact.Contact {
	if !u.checkIsStaffAndIsActive() {
		return nil
	}
	cID := len(u.Contacts) + 1
	newContact := contact.NewContact(fName, lName, cID)
	u.Contacts = append(u.Contacts, newContact)
	return newContact
}

func (u *User) GetContactById(contactId int) *contact.Contact {
	if !u.checkIsStaffAndIsActive() {
		return nil
	}
	for _, c := range u.Contacts {
		if c.ContactID == contactId && c.IsActive {
			fmt.Println("Contact data fetched with ContactId:", contactId)
			return c
		}
	}
	fmt.Println("Contact Not Present with ContactId:", contactId)
	return nil
}

func (u *User) UpdateContact(contactId int, param string, value interface{}) {
	if !u.checkIsStaffAndIsActive() {
		return
	}
	contactObj := u.GetContactById(contactId)
	if contactObj != nil {
		contactObj.UpdateContact(param, value)
	}
}

func (u *User) CreateContactDetails(contactId int, ctype, cvalue string) {
	if !u.checkIsStaffAndIsActive() {
		return
	}
	contactObj := u.GetContactById(contactId)
	if contactObj != nil {
		contactObj = u.GetContactById(contactId)
	}
}

func (u *User) DeleteContact(contactId int) {
	if !u.checkIsStaffAndIsActive() {
		return
	}
	contactObj := u.GetContactById(contactId)
	if contactObj != nil {
		contactObj.IsActive = false
		fmt.Println("Contact", contactObj.FName, contactObj.LName, "(ID:", contactObj.ContactID, ") has been soft deleted (IsActive=false).")
	}
}
