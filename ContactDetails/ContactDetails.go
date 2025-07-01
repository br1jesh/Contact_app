package contactDetails

import "fmt"

var contactDetailId = 1

type ContactDetail struct {
	ContactDetailsID int
	Type             string
	Value            string
}

func NewContactDetail(typ, value string) *ContactDetail {
	cd := &ContactDetail{
		ContactDetailsID: contactDetailId,
		Type:             typ,
		Value:            value,
	}
	fmt.Println("Created ContactDetail:", typ, "=", value, "(ID:", cd.ContactDetailsID, ")")
	contactDetailId++
	return cd
}

func (cd *ContactDetail) UpdateContact(param string, value interface{}) {
	switch param {
	case "Type":
		cd.updateType(value)
	case "Value":
		cd.updateValue(value)
	default:
		fmt.Println("Unknown parameter:", param)
	}
}

func (cd *ContactDetail) updateType(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		cd.Type = str
		fmt.Println("Type updated successfully.")
	} else {
		fmt.Println("updateType: invalid string")
	}
}

func (cd *ContactDetail) updateValue(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		cd.Value = str
		fmt.Println("Value updated successfully.")
	} else {
		fmt.Println("updateValue: invalid string")
	}
}
