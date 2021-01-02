package handler

import (
	"Huangdu_HMC_Appointment/src/logger"
	"Huangdu_HMC_Appointment/src/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userID  string
	orderID string
}

/******/
func parseToken(tokenStr string) string {
	var userID string
	//split "Bearer" and JWT Token
	tokenStr = strings.Fields(tokenStr)[1]
	//turn token string to token struct
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	claim := token.Claims.(jwt.MapClaims)
	userID = claim["http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier"].(string)
	logger.Info.Printf("User ID : %s", userID)
	//if return is "",parse fail
	return userID
}
func parseJSON(context *gin.Context) model.Appointment {
	var res model.Appointment
	var jsonMap map[string]interface{}
	err := context.ShouldBindJSON(&jsonMap)
	if err != nil {
		logger.Error.Println(err)
		return res
	}
	if len(jsonMap) != 8 {
		return res
	}
	res.Date = jsonMap["date"].(string)
	res.HospitalName = jsonMap["hospital_name"].(string)
	res.Name = jsonMap["doctor_name"].(string)
	res.Department = jsonMap["department"].(string)
	res.TimeSlot = int(jsonMap["time_slot"].(float64))
	res.ProfessionalTitle = jsonMap["professional_title"].(string)
	res.Price = int(jsonMap["price"].(float64))
	res.Details = jsonMap["details"].(string)
	return res
}

/******/
func (h *Handler) GetAppoint(context *gin.Context) {
	var res model.ReturnAppoint
	// get the token from the header
	token := context.Request.Header.Get("Authorization")
	//if token is empty
	if token == "" {
		logger.Error.Printf("No Authorization here, invalid request")
		res.Msg = "Invalid Authorization here,check it"
		//return 400 then terminate
		context.JSON(http.StatusBadRequest, res)
		return
	}
	h.userID = parseToken(token)
	//get result of the query
	res.AppointList = h.queryAppoint()
	if res.AppointList == nil {
		res.Msg = "No Data"
		context.JSON(http.StatusNotFound, res)
	}
	res.Msg = "OK"
	context.JSON(http.StatusOK, res)
}

/*****/
func (h *Handler) NewAppoint(context *gin.Context) {
	var res model.Result
	token := context.Request.Header.Get("Authorization")
	//if token is empty
	if token == "" {
		logger.Error.Printf("No Authorization here, invalid request")
		res.Msg = "Invalid Authorization here,check it"
		//return 400 then terminate
		context.JSON(http.StatusBadRequest, res)
		return
	}
	h.userID = parseToken(token)

	appoint := parseJSON(context)
	if appoint != (model.Appointment{}) {
		if h.addAppoint(appoint) {
			res.Msg = "Create Successfully"
			context.JSON(http.StatusCreated, res)
			return
		} else {
			res.Msg = "Create Failed"
		}
	} else {
		res.Msg = "Invalid JSON "
	}
	context.JSON(http.StatusBadRequest, res)
}

/*****/
func (h *Handler) DelAppoint(context *gin.Context) {
	var res model.Result
	h.orderID = context.Query("order_id")

	// get the token from the header
	token := context.Request.Header.Get("Authorization")
	//if token is empty
	if token == "" {
		logger.Error.Printf("No Authorization here, invalid request")
		res.Msg = "Invalid Authorization here,check it"
		//return 400 then terminate
		context.JSON(http.StatusBadRequest, res)
		return
	}
	h.userID = parseToken(token)
	if h.orderID != "" && h.deleteAppoint() {
		res.Msg = "Delete the order successfully"
		context.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "Bad Parameter"
	context.JSON(http.StatusBadRequest, res)
}
