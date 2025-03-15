package model

type Invoice struct {
	Abstract
	InvoiceType string `json:"invoice_type" gorm:"not null; size:20"` //Fee, Payslip
	StudentID   uint   `json:"student_id"`
	TeacherID   uint   `json:"teacher_id"`
}

func (Invoice) TableName() string {
	return "pp_invoice"
}

//invoice <--> student ---one2one
//invoice <--> teacher ---one2one
