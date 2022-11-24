package suite

import (
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tranthienphucbreak/parsleytest/internal/patient"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

const HOST = "http://localhost:18443"

var patientID = ""

func TestPatientService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Patient Service Functional Test")
}

var _ = Describe("Patient Service Functional Test", func() {
	defer GinkgoRecover()
	var (
	//patientHandler routes.PatientHandler
	)

	BeforeEach(func() {
		//patientHandler = routes.NewServicesHandler(&patient.PatientService{})
	})

	AfterEach(func() {

	})

	It("GetPatients API OK", func() {
		req := newGetRequest("/patients")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		Expect(resp.StatusCode).To(Equal(200))
	})

	It("GetPatients API invalid page", func() {
		req := newGetRequest("/patients")
		q := url.Values{}
		q.Add("page", "pageInvalid")
		q.Add("limit", "-1")
		req.URL.RawQuery = encodeQuerySimple(q)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		Expect(resp.StatusCode).To(Equal(400))
	})

	It("GetPatients API FirstPage HasNextPage = true", func() {
		req := newGetRequest("/patients")
		q := url.Values{}
		q.Add("page", "1")
		q.Add("limit", "10")
		req.URL.RawQuery = encodeQuerySimple(q)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		patientsResponse := patient.GetPatientsResponse{}

		err = json.Unmarshal(result, &patientsResponse)
		if err != nil {
			Fail(err.Error())
		}
		Expect(patientsResponse.HasNextPage).To(Equal(true))
	})

	It("GetPatients API FirstPage HasNextPage = false", func() {
		req := newGetRequest("/patients")
		q := url.Values{}
		q.Add("page", "1000")
		q.Add("limit", "10")
		req.URL.RawQuery = encodeQuerySimple(q)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		patientsResponse := patient.GetPatientsResponse{}

		err = json.Unmarshal(result, &patientsResponse)
		if err != nil {
			Fail(err.Error())
		}
		Expect(patientsResponse.HasNextPage).To(Equal(false))
	})

	It("GetPatients API SecondPage FirstID", func() {
		req := newGetRequest("/patients")
		q := url.Values{}
		q.Add("page", "2")
		q.Add("limit", "10")
		req.URL.RawQuery = encodeQuerySimple(q)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		patientsResponse := patient.GetPatientsResponse{}

		err = json.Unmarshal(result, &patientsResponse)
		if err != nil {
			Fail(err.Error())
		}
		Expect(patientsResponse.Data[0].ID).To(Equal("1zsjhd0d"))
	})

	It("CreatePatient API", func() {
		req := newPostRequest("/patient", "../../testdata/patient.json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		response := patient.GetPatientByIDResponse{}

		err = json.Unmarshal(result, &response)
		if err != nil {
			Fail(err.Error())
		}
		Expect(response.ID).ToNot(Equal(""))
		patientID = response.ID
	})

	It("GetPatient API", func() {
		req := newGetRequest("/patients/" + patientID)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		response := patient.GetPatientByIDResponse{}

		err = json.Unmarshal(result, &response)
		if err != nil {
			Fail(err.Error())
		}
		Expect(response.ID).To(Equal(patientID))
	})

	It("Update Patient API", func() {
		req := newPutRequest("/patients/"+patientID, "../../testdata/update_patient.json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		Expect(resp.StatusCode).To(Equal(200))

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Fail(err.Error())
		}
		response := patient.GetPatientByIDResponse{}

		err = json.Unmarshal(result, &response)
		if err != nil {
			Fail(err.Error())
		}
		Expect(response.Email).To(Equal("Rowland.Jack@hotmail.com"))
		Expect(response.Phone).To(Equal("844-517-2343"))
		Expect(response.DOB).To(Equal("2006-04-22"))
	})

	It("Update Patient API | Wrong DOB", func() {
		req := newPutRequest("/patients/"+patientID, "../../testdata/update_patient_wrong_dob.json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		Expect(resp.StatusCode).To(Equal(400))
	})

	It("Delete Patient API | Wrong DOB", func() {
		req := newDeleteRequest("/patients/" + patientID)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Fail(fmt.Sprintf("Request failed with error: %s", err.Error()))
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		Expect(resp.StatusCode).To(Equal(200))
	})
})

func newGetRequest(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, HOST+path, nil)
	return req
}

func newDeleteRequest(path string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, HOST+path, nil)
	return req
}

func newPostRequest(path string, payloadPath string) *http.Request {
	payload, _ := os.Open(payloadPath)
	req, _ := http.NewRequest(http.MethodPost, HOST+path, payload)
	return req
}

func newPutRequest(path string, payloadPath string) *http.Request {
	payload, _ := os.Open(payloadPath)
	req, _ := http.NewRequest(http.MethodPut, HOST+path, payload)
	return req
}

func encodeQuerySimple(q url.Values) string {
	var queries []string
	for k, vs := range q {
		for _, v := range vs {
			queries = append(queries, fmt.Sprintf("%v=%v", k, v))
		}
	}
	return strings.ReplaceAll(strings.Join(queries, "&"), " ", "%20")
}
