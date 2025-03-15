package service

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/database"
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/model"
)

func GetAllInvoiceData() (*[]model.Invoice, error) {
	var invoices []model.Invoice
	res := database.DB.Find(&invoices)
	return &invoices, res.Error
}

func GetInvoiceData(id uint) (*model.Invoice, error) {
	var invoice model.Invoice
	res := database.DB.First(&invoice, id)
	return &invoice, res.Error
}

func PostInvoiceData(invoice *model.Invoice) (*model.Invoice, error) {
	res := database.DB.Create(&invoice)
	return invoice, res.Error
}
