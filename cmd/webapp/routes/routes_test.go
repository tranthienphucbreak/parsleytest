package routes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tranthienphucbreak/parsleytest/cmd/webapp/routes"
)

func TestPatientService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Patient Service Suite")
}

var _ = Describe("Patient Service", func() {

	defer GinkgoRecover()

	//var patientHandler routes.PatientHandler

	BeforeEach(func() {
		//patientHandler = routes.NewServicesHandler(&patient.PatientService{})
	})

	AfterEach(func() {

	})

	Context("Routes", func() {
		It("DobValidate wrong format", func() {
			isOk := routes.DobValidateInternal("2020/10/20")
			Expect(isOk).To(Equal(false))
		})

		It("DobValidate wrong date/month position", func() {
			isOk := routes.DobValidateInternal("2020-31-11")
			Expect(isOk).To(Equal(false))
		})
		It("DobValidate good", func() {
			isOk := routes.DobValidateInternal("2020-11-30")
			Expect(isOk).To(Equal(true))
		})

		It("PhoneValidateInternal wrong format", func() {
			isOk := routes.PhoneValidateInternal("a 812912823")
			Expect(isOk).To(Equal(false))

			isOk = routes.PhoneValidateInternal("(1223) | 812912823")
			Expect(isOk).To(Equal(false))
		})

		It("PhoneValidateInternal good", func() {
			isOk := routes.PhoneValidateInternal("844-517-4868")
			Expect(isOk).To(Equal(true))
			isOk = routes.PhoneValidateInternal("+84 812912823")
			Expect(isOk).To(Equal(true))
		})
	})
})
