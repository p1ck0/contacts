package contact

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/p1ck0/contacts/database"
	"gorm.io/gorm"
)

var db = database.Connector()

func GetContacts(c *fiber.Ctx) error {
	var contacts []database.Contact
	db.Find(&contacts)
	return c.JSON(contacts)
}

func NewContact(c *fiber.Ctx) error {
	contact := new(database.Contact)
	if err := c.BodyParser(contact); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	if !checkNumber(contact.Number) {
		return c.Status(500).SendString("Contact already exist")
	}
	db.Create(&contact)
	return c.JSON(contact)
}

func EditContact(c *fiber.Ctx) error {
	id := c.Params("id")
	contact := new(database.Contact)
	editedContact := new(database.Contact)
	if err := c.BodyParser(editedContact); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	if err := db.First(&contact, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).SendString("No Contact Found with ID")
	}
	if editedContact.Name != "" {
		contact.Name = editedContact.Name
	}
	if editedContact.Number != "" {
		if !checkNumber(editedContact.Number) {
			return c.Status(500).SendString("Contact already exist")
		}
		contact.Number = editedContact.Number
	}
	db.Save(&contact)
	return c.JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var contact database.Contact
	if err := db.First(&contact, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).SendString("No Contact Found with ID")
	}
	db.Delete(&contact)
	return c.JSON(contact)
}

func checkNumber(number string) bool {
	var c database.Contact
	if err := db.Select("number").Where("number = ?", number).First(&c).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
