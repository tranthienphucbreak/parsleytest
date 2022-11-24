package patient

//import "time"

type GetPatientsRequest struct {
	Page  int
	Limit int
}

type GetPatientByIDRequest struct {
	ID string
}

type DeletePatientByIDRequest struct {
	ID string
}

type UpdatePatientRequest struct {
	ID            string `json:"id" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	MiddleName    string `json:"middle_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	DOB           string `json:"dob" validate:"required,dob"`
	Status        string `json:"status" validate:"required,oneof=active inactive"`
	Gender        string `json:"gender" validate:"required,oneof=male female"`
	TermsAccepted int32  `json:"terms_accepted" validate:"oneof=0 1"`
	AddressStreet string `json:"address_street" validate:"required"`
	AddressCity   string `json:"address_city" validate:"required"`
	AddressState  string `json:"address_state" validate:"required"`
	AddressZip    int32  `json:"address_zip" validate:"required"`
	Phone         string `json:"phone" validate:"required,phone"`
}

type CreatePatientRequest struct {
	FirstName     string `json:"first_name" validate:"required"`
	MiddleName    string `json:"middle_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	DOB           string `json:"dob" validate:"required,dob"`
	Gender        string `json:"gender" validate:"required,oneof=male female"`
	TermsAccepted int32  `json:"terms_accepted" validate:"oneof=0 1"`
	AddressStreet string `json:"address_street" validate:"required"`
	AddressCity   string `json:"address_city" validate:"required"`
	AddressState  string `json:"address_state" validate:"required"`
	AddressZip    int32  `json:"address_zip" validate:"required"`
	Phone         string `json:"phone" validate:"required,phone"`
}

type ExecResult struct {
}

type Person struct {
	ID              string
	FirstName       string
	MiddleName      string
	LastName        string
	Email           string
	DOB             string
	Gender          string
	Status          string
	TermsAccepted   int32
	TermsAcceptedAt string
	AddressStreet   string
	AddressCity     string
	AddressState    string
	AddressZip      int32
	Phone           string
}

type Patient struct {
	ID              string `json:"id,omitempty"`
	FullName        string `json:"fullname,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	DOB             string `json:"dob,omitempty"`
	Age             int    `json:"age,omitempty"`
	Gender          string `json:"gender,omitempty"`
	Status          string `json:"status,omitempty"`
	TermsAccepted   int32  `json:"terms_accepted"`
	TermsAcceptedAt int64  `json:"terms_accepted_at,omitempty"`
	Address         string `json:"address,omitempty"`
	AddressStreet   string `json:"address_street,omitempty"`
	AddressCity     string `json:"address_city,omitempty"`
	AddressState    string `json:"address_state,omitempty"`
	AddressZip      int32  `json:"address_zip,omitempty"`
	Phone           string `json:"phone,omitempty"`
}

type GetPatientsResponse struct {
	Data        []Patient `json:"data,omitempty"`
	HasNextPage bool      `json:"has_next_page,omitempty"`
}

type GetPatientByIDResponse struct {
	Patient
}

type DeletePatientByIDResponse struct {
	PatientID string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
}

const (
	ERROR_ITEM_NOT_FOUND = "Item Not Found"
)
