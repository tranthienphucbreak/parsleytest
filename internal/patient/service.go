package patient

import (
	"context"
	//"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	//"errors"
	age "github.com/bearbin/go-age"
	"github.com/google/uuid"
)

const (
	DATE_TIME_FORMAT = "2006-01-02T15:04:05.000Z"
)

type PatientService struct {
	DBService DatabaseProvider
}

type DatabaseProvider interface {
	Exec(query string, params []interface{}) (int64, error)
	Query(query string, params []interface{}) ([]Person, error)
	QueryRow(query string, params []interface{}) (*Person, error)
}

func (ss *PatientService) GetPatients(ctx context.Context, req *GetPatientsRequest) (*GetPatientsResponse, error) {
	patients, err := ss.DBService.Query("SELECT * FROM person ORDER BY last_name ASC LIMIT ? OFFSET ?", []interface{}{
		req.Limit + 1,
		req.Limit * (req.Page - 1),
	})
	if err != nil {
		return nil, err
	}

	patientsRs := []Patient{}

	limit := req.Limit
	hasNextPage := len(patients) > req.Limit
	if !hasNextPage {
		limit = len(patients)
	}

	for _, person := range patients[:limit] {
		patientsRs = append(patientsRs, ss.formatPatient(person))
	}

	return &GetPatientsResponse{
		Data:        patientsRs,
		HasNextPage: hasNextPage,
	}, nil
}

func (ss *PatientService) GetPatientByID(ctx context.Context, req *GetPatientByIDRequest) (*GetPatientByIDResponse, error) {
	patient, err := ss.DBService.QueryRow("SELECT * FROM person WHERE id = ?", []interface{}{
		req.ID,
	})
	if err != nil {
		return nil, err
	}
	return &GetPatientByIDResponse{
		ss.formatPatient(*patient),
	}, nil
}

func (ss *PatientService) DeletePatientByID(ctx context.Context, req *GetPatientByIDRequest) (*DeletePatientByIDResponse, error) {
	_, err := ss.DBService.Exec("DELETE FROM person WHERE id = ?", []interface{}{
		req.ID,
	})
	if err != nil {
		return nil, err
	}
	return &DeletePatientByIDResponse{
		PatientID: req.ID,
		Status:    "DELETED",
	}, nil
}

func (ss *PatientService) UpdatePatientByID(ctx context.Context, req *UpdatePatientRequest) (*GetPatientByIDResponse, error) {
	updateFields := []string{}
	updateParams := []interface{}{}
	updateFields = append(updateFields, "first_name=?")
	updateParams = append(updateParams, req.FirstName)

	updateFields = append(updateFields, "last_name=?")
	updateParams = append(updateParams, req.LastName)

	updateFields = append(updateFields, "middle_name=?")
	updateParams = append(updateParams, req.MiddleName)

	updateFields = append(updateFields, "dob=?")
	updateParams = append(updateParams, req.DOB)

	updateFields = append(updateFields, "email=?")
	updateParams = append(updateParams, req.Email)

	updateFields = append(updateFields, "gender=?")
	updateParams = append(updateParams, req.Gender)

	updateFields = append(updateFields, "status=?")
	updateParams = append(updateParams, req.Status)

	updateFields = append(updateFields, "terms_accepted=?")
	updateParams = append(updateParams, req.TermsAccepted)
	if req.TermsAccepted == 1 {
		updateFields = append(updateFields, "terms_accepted_at=?")
		updateParams = append(updateParams, time.Now().UTC().Format(DATE_TIME_FORMAT))
	} else {
		updateFields = append(updateFields, "terms_accepted_at=?")
		updateParams = append(updateParams, "")
	}

	updateFields = append(updateFields, "address_street=?")
	updateParams = append(updateParams, req.AddressStreet)

	updateFields = append(updateFields, "address_city=?")
	updateParams = append(updateParams, req.AddressCity)

	updateFields = append(updateFields, "address_state=?")
	updateParams = append(updateParams, req.AddressState)

	updateFields = append(updateFields, "address_zip=?")
	updateParams = append(updateParams, req.AddressZip)

	updateFields = append(updateFields, "phone=?")
	updateParams = append(updateParams, req.Phone)

	updateParams = append(updateParams, req.ID)

	_, err := ss.DBService.Exec("UPDATE person SET "+strings.Join(updateFields, ",")+" WHERE id = ?", updateParams)
	if err != nil {
		return nil, err
	}
	return ss.GetPatientByID(ctx, &GetPatientByIDRequest{
		ID: req.ID,
	})
}

func (ss *PatientService) CreatePatient(ctx context.Context, req *CreatePatientRequest) (*GetPatientByIDResponse, error) {
	fields := []string{}
	params := []interface{}{}

	id := uuid.New().String()

	fields = append(fields, "id")
	params = append(params, id)

	fields = append(fields, "first_name")
	params = append(params, req.FirstName)

	fields = append(fields, "last_name")
	params = append(params, req.LastName)

	fields = append(fields, "middle_name")
	params = append(params, req.MiddleName)

	fields = append(fields, "dob")
	params = append(params, req.DOB)

	fields = append(fields, "email")
	params = append(params, req.Email)

	fields = append(fields, "gender")
	params = append(params, req.Gender)

	fields = append(fields, "status")
	params = append(params, "active")

	fields = append(fields, "terms_accepted")
	params = append(params, req.TermsAccepted)
	if req.TermsAccepted == 1 {
		fields = append(fields, "terms_accepted_at")
		params = append(params, time.Now().UTC().Format(DATE_TIME_FORMAT))
	} else {
		fields = append(fields, "terms_accepted_at")
		params = append(params, "")
	}

	fields = append(fields, "address_street")
	params = append(params, req.AddressStreet)

	fields = append(fields, "address_city")
	params = append(params, req.AddressCity)

	fields = append(fields, "address_state")
	params = append(params, req.AddressState)

	fields = append(fields, "address_zip")
	params = append(params, req.AddressZip)

	fields = append(fields, "phone")
	params = append(params, req.Phone)

	fieldValues := []string{}

	for range fields {
		fieldValues = append(fieldValues, "?")
	}

	_, err := ss.DBService.Exec("INSERT INTO person("+strings.Join(fields, ",")+") values("+strings.Join(fieldValues, ",")+")", params)
	if err != nil {
		return nil, err
	}
	return ss.GetPatientByID(ctx, &GetPatientByIDRequest{
		ID: id,
	})
}

func (ss *PatientService) formatPatient(person Person) (patient Patient) {
	patient.ID = person.ID
	patient.FirstName = person.FirstName
	patient.MiddleName = person.MiddleName
	patient.LastName = person.LastName

	patient.Gender = person.Gender
	patient.Email = person.Email
	patient.Status = person.Status
	patient.DOB = person.DOB

	patient.TermsAccepted = person.TermsAccepted

	patient.AddressStreet = person.AddressStreet
	patient.AddressCity = person.AddressCity
	patient.AddressState = person.AddressState
	patient.AddressZip = person.AddressZip
	patient.Phone = person.Phone

	patient.Age = ss.GetAgeFromDOB(person.DOB)
	patient.TermsAcceptedAt = ss.StringToDate(person.TermsAcceptedAt)
	space := regexp.MustCompile(`\s+`)
	patient.FullName = space.ReplaceAllString(person.FirstName+" "+person.MiddleName+" "+person.LastName, " ")
	patient.Address = space.ReplaceAllString(person.AddressStreet+", "+person.AddressCity+", "+person.AddressState, " ")
	return
}

func (ss *PatientService) GetAgeFromDOB(dob string) int {
	dobYearMonthDate := strings.Split(dob, "-")
	if len(dobYearMonthDate) != 3 {
		return 0
	}
	year, err := strconv.Atoi(dobYearMonthDate[0])
	if err != nil {
		return 0
	}
	month, err := strconv.Atoi(dobYearMonthDate[1])
	if err != nil {
		return 0
	}
	date, err := strconv.Atoi(dobYearMonthDate[2])
	if err != nil {
		return 0
	}
	dobDate := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
	return age.Age(dobDate)
}

func (ss *PatientService) StringToDate(dateTimeString string) int64 {
	dateTime, err := time.Parse(DATE_TIME_FORMAT, dateTimeString)
	if err != nil {
		return 0
	}
	return dateTime.Unix()
}
