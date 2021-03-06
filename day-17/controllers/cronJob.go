package controllers

import (
	"time"
	"xapiens-bootcamp-golang/day-17/helpers"
	"xapiens-bootcamp-golang/day-17/models"
)

func (StrDB *StrDB) Job() {
	var (
		mail []models.RegistMailer
	)
	// dbPG := config.Connection()

	StrDB.DB.Where("status = ?", false).Find(&mail)

	// fmt.Println(len(mail))

	for i := 0; i < len(mail); i++ {
		helpers.RegisterMailer(mail[i].Email, mail[i].Message)
		mail[i].Status = true
		mail[i].DeliveredAt = time.Now()
		StrDB.DB.Save(&mail[i])
	}
}
