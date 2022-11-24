package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	//"regexp"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/tranthienphucbreak/parsleytest/internal/patient"
)

const DEFAULT_PAGE_LIMIT = 10
const PAGE_LIMIT_MAX = 30

var validate *validator.Validate

type ServiceError struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}

type PatientProvider interface {
	GetPatients(ctx context.Context, req *patient.GetPatientsRequest) (*patient.GetPatientsResponse, error)
	GetPatientByID(ctx context.Context, req *patient.GetPatientByIDRequest) (*patient.GetPatientByIDResponse, error)
	DeletePatientByID(ctx context.Context, req *patient.GetPatientByIDRequest) (*patient.DeletePatientByIDResponse, error)
	UpdatePatientByID(ctx context.Context, req *patient.UpdatePatientRequest) (*patient.GetPatientByIDResponse, error)
	CreatePatient(ctx context.Context, req *patient.CreatePatientRequest) (*patient.GetPatientByIDResponse, error)
}

type PatientHandler struct {
	PatientService PatientProvider
}

func NewServicesHandler(s PatientProvider) PatientHandler {
	validate = validator.New()
	_ = validate.RegisterValidation("dob", DobValidate)
	_ = validate.RegisterValidation("phone", PhoneValidate)
	return PatientHandler{s}
}

func (sh *PatientHandler) Set(r *chi.Mux) {
	r.Get("/patients", sh.GetPatients)
	r.Get("/patients/{id}", sh.GetPatientByID)
	r.Delete("/patients/{id}", sh.DeletePatientByID)
	r.Put("/patients/{id}", sh.UpdatePatientByID)
	r.Post("/patient", sh.CreatePatient)
}

func (sh *PatientHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	request, err := createGetPatientsRequest(r)
	if err != nil {
		serviceCallErrorHandler(400, err, w, r)
		return
	}
	rs, err := sh.PatientService.GetPatients(r.Context(), request)
	if err != nil {
		fmt.Println("[GetPatients] ", err)
		serviceCallErrorHandler(500, err, w, r)
		return
	}
	writeJsonResponse(w, rs)
}

func (sh *PatientHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	request, err := createGetPatientByIDRequest(r)
	if err != nil {
		serviceCallErrorHandler(400, err, w, r)
		return
	}
	rs, err := sh.PatientService.GetPatientByID(r.Context(), request)
	if err != nil {
		fmt.Println("[GetPatientByID] ", err)
		if err.Error() == patient.ERROR_ITEM_NOT_FOUND {
			serviceCallErrorHandler(404, err, w, r)
		} else {
			serviceCallErrorHandler(500, err, w, r)
		}
		return
	}
	writeJsonResponse(w, rs)
}

func (sh *PatientHandler) DeletePatientByID(w http.ResponseWriter, r *http.Request) {
	request, err := createGetPatientByIDRequest(r)
	if err != nil {
		serviceCallErrorHandler(400, err, w, r)
		return
	}
	rs, err := sh.PatientService.DeletePatientByID(r.Context(), request)
	if err != nil {
		fmt.Println("[DeletePatientByID] ", err)
		if err.Error() == patient.ERROR_ITEM_NOT_FOUND {
			serviceCallErrorHandler(404, err, w, r)
		} else {
			serviceCallErrorHandler(500, err, w, r)
		}
		return
	}
	writeJsonResponse(w, rs)
}

func (sh *PatientHandler) UpdatePatientByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	request, err := createUpdatePatientByIDRequest(r)
	if err != nil {
		serviceCallErrorHandler(400, err, w, r)
		return
	}
	rs, err := sh.PatientService.UpdatePatientByID(r.Context(), request)
	if err != nil {
		fmt.Println("[UpdatePatientByID] ", err)
		if err.Error() == patient.ERROR_ITEM_NOT_FOUND {
			serviceCallErrorHandler(404, err, w, r)
		} else {
			serviceCallErrorHandler(500, err, w, r)
		}
		return
	}
	writeJsonResponse(w, rs)
}

func (sh *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	request, err := createCreatePatientRequest(r)
	if err != nil {
		serviceCallErrorHandler(400, err, w, r)
		return
	}
	rs, err := sh.PatientService.CreatePatient(r.Context(), request)
	if err != nil {
		fmt.Println("[CreatePatient] ", err)
		if err.Error() == patient.ERROR_ITEM_NOT_FOUND {
			serviceCallErrorHandler(404, err, w, r)
		} else {
			serviceCallErrorHandler(500, err, w, r)
		}
		return
	}
	writeJsonResponse(w, rs)
}

func createGetPatientsRequest(r *http.Request) (*patient.GetPatientsRequest, error) {
	request := patient.GetPatientsRequest{
		Limit: DEFAULT_PAGE_LIMIT,
		Page:  1,
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}
		request.Page = pageInt
	}
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		request.Limit = limitInt
	}
	if request.Limit > PAGE_LIMIT_MAX {
		request.Limit = PAGE_LIMIT_MAX
	}
	return &request, nil
}

func createGetPatientByIDRequest(r *http.Request) (*patient.GetPatientByIDRequest, error) {
	request := patient.GetPatientByIDRequest{
		ID: chi.URLParam(r, "id"),
	}
	return &request, nil
}

func createUpdatePatientByIDRequest(r *http.Request) (*patient.UpdatePatientRequest, error) {
	var request patient.UpdatePatientRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, err
	}
	request.ID = chi.URLParam(r, "id")
	err = validate.Struct(request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func createCreatePatientRequest(r *http.Request) (*patient.CreatePatientRequest, error) {
	var request patient.CreatePatientRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func serviceCallErrorHandler(errorCode int, err error, w http.ResponseWriter, r *http.Request) {
	error := ServiceError{
		Timestamp: time.Now().UTC(),
		Status:    errorCode,
		Error:     http.StatusText(errorCode),
		Message:   fmt.Sprint(err),
		Path:      r.URL.Path,
	}
	errorResponse, _ := json.Marshal(error)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	w.Write(errorResponse)
}

func writeJsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func DobValidate(fl validator.FieldLevel) bool {
	return DobValidateInternal(fl.Field().String())
}

func DobValidateInternal(dob string) bool {
	_, err := time.Parse("2006-01-02", dob)
	return err == nil
}

func PhoneValidate(fl validator.FieldLevel) bool {
	return PhoneValidateInternal(fl.Field().String())
}

func PhoneValidateInternal(phoneNumber string) bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(phoneNumber)
}
