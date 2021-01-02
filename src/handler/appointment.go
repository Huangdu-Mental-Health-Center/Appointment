package handler

import (
	"Huangdu_HMC_Appointment/src/driver"
	"Huangdu_HMC_Appointment/src/logger"
	"Huangdu_HMC_Appointment/src/model"
	"database/sql"
)

func (h *Handler) deleteFromDB(pDB *sql.DB) bool {
	var result = false
	stmt, err := pDB.Prepare("delete from appointment where appointment_id=? AND user_id=?")
	if err != nil {
		logger.Error.Println(err)
		return result
	}
	res, err1 := stmt.Exec(h.orderID, h.userID)

	if err1 != nil {
		logger.Error.Println(err1)
		return result
	}
	affect, err2 := res.RowsAffected()
	if err2 != nil {
		logger.Error.Println(err2)
	}
	if affect == 0 {
		logger.Error.Println("Delete Nothing")
		return false
	}
	result = true
	return result
}
func (h *Handler) queryFromDB(pDB *sql.DB) []model.Appointment {
	var resList []model.Appointment
	//query the needed rows
	rows, err := pDB.Query(
		"SELECT appointment_id,hospital,doctor_name,department,appoint_date,time_slot,professional_title,price,status FROM appointment  WHERE user_id=?", h.userID)
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	// close the rows unitl everything is done
	defer rows.Close()
	//get info   row by row
	for rows.Next() {
		var tempAppoint model.Appointment
		err1 := rows.Scan(&tempAppoint.ID, &tempAppoint.HospitalName, &tempAppoint.Name, &tempAppoint.Department, &tempAppoint.Date, &tempAppoint.TimeSlot,
			&tempAppoint.ProfessionalTitle, &tempAppoint.Price, &tempAppoint.Status)
		// scan is failed ,continue to the next row
		if err1 != nil {
			logger.Error.Println(err1)
			continue
		}
		// insert the appoint item into the list
		resList = append(resList, tempAppoint)
	}
	//collect the rows' error except scan error
	err = rows.Err()
	if err != nil {
		logger.Error.Println(err)
	}
	return resList
}
func (h *Handler) createFromDB(pDB *sql.DB, order model.Appointment) bool {
	var result = false
	stmt, err :=
		pDB.Prepare("INSERT INTO appointment(user_id ,hospital,doctor_name,department,appoint_date,time_slot,professional_title,price,details,status)	VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Error.Println(err)
		return result
	}
	res, err1 := stmt.Exec(h.userID, order.HospitalName, order.Name, order.Department,
		order.Date, order.TimeSlot, order.ProfessionalTitle,
		order.Price, order.Details, "申请中")
	logger.Info.Println(res)
	if err1 != nil {
		logger.Error.Println(err1)
		return result
	}
	result = true
	return result
}

/*****/
func (h *Handler) queryAppoint() []model.Appointment {
	var result []model.Appointment
	isOK, pDB := driver.ConnectDB()
	if !isOK {
		logger.Error.Printf("Connect DB:%s failed", driver.DBNAME)
	}
	result = h.queryFromDB(pDB)
	err := pDB.Close()
	if err != nil {
		logger.Error.Println(err)
	}
	return result
}

func (h *Handler) deleteAppoint() bool {
	res := false
	isOK, pDB := driver.ConnectDB()
	if !isOK {
		logger.Error.Printf("Connect DB:%s failed", driver.DBNAME)
		return false
	}
	res = h.deleteFromDB(pDB)
	err := pDB.Close()
	if err != nil {
		logger.Error.Println(err)
	}
	return res
}
func (h *Handler) addAppoint(appoint model.Appointment) bool {
	res := false
	isOK, pDB := driver.ConnectDB()
	if !isOK {
		logger.Error.Printf("Connect DB:%s failed", driver.DBNAME)
		return false
	}

	res = h.createFromDB(pDB, appoint)
	err := pDB.Close()
	if err != nil {
		logger.Error.Println(err)
	}
	return res
}
