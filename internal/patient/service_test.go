package patient_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tranthienphucbreak/parsleytest/internal/patient"
	"github.com/tranthienphucbreak/parsleytest/mocks"
	"testing"
)

func TestPatientService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Patient Service Suite")
}

var _ = Describe("Patient Service", func() {
	defer GinkgoRecover()
	var (
		mockCtrl            *gomock.Controller
		patientService      *patient.PatientService
		mockDatabaseService *mocks.MockDatabaseProvider
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockDatabaseService = mocks.NewMockDatabaseProvider(mockCtrl)
		patientService = &patient.PatientService{
			DBService: mockDatabaseService,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("GetPatients", func() {
		It("GetPatients error", func() {
			ctx := context.Background()
			mockDatabaseService.EXPECT().Query(gomock.Any(), gomock.Any()).Return(nil, errors.New("Database is locked"))
			_, err := patientService.GetPatients(ctx, &patient.GetPatientsRequest{
				Page:  1,
				Limit: 10,
			})
			Expect(err).To(HaveOccurred())
		})

		It("GetPatients HasNextPage = false", func() {
			ctx := context.Background()
			personList := []patient.Person{}

			for i := 0; i < 1; i++ {
				personList = append(personList, patient.Person{})
			}
			mockDatabaseService.EXPECT().Query(gomock.Any(), gomock.Any()).Return(personList, nil)
			rs, err := patientService.GetPatients(ctx, &patient.GetPatientsRequest{
				Page:  1,
				Limit: 10,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(rs.HasNextPage).To(Equal(false))
		})

		It("GetPatients HasNextPage = false", func() {
			ctx := context.Background()
			personList := []patient.Person{}

			for i := 0; i < 11; i++ {
				personList = append(personList, patient.Person{})
			}
			mockDatabaseService.EXPECT().Query(gomock.Any(), gomock.Any()).Return(personList, nil)
			rs, err := patientService.GetPatients(ctx, &patient.GetPatientsRequest{
				Page:  1,
				Limit: 10,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(rs.HasNextPage).To(Equal(true))
		})
	})

	Context("GetAgeFromDOB", func() {
		It("GetAgeFromDOB error", func() {
			age := patientService.GetAgeFromDOB("2012/10/20")
			Expect(age).To(Equal(0))
		})

		It("GetAgeFromDOB good", func() {
			age := patientService.GetAgeFromDOB("2002-10-20")
			Expect(age).To(Equal(20))
		})
	})

	Context("GetAgeFromDOB", func() {
		It("GetAgeFromDOB error", func() {
			age := patientService.GetAgeFromDOB("2012/10/20")
			Expect(age).To(Equal(0))
		})

		It("GetAgeFromDOB good", func() {
			age := patientService.GetAgeFromDOB("2002-10-20")
			Expect(age).To(Equal(20))
		})
	})

	Context("StringToDate", func() {
		It("StringToDate error", func() {
			timestamp := patientService.StringToDate("2012/10/20 02:20:20")
			Expect(timestamp).To(Equal(int64(0)))
		})

		It("StringToDate good", func() {
			timestamp := patientService.StringToDate("2006-01-02T15:04:05.000Z")
			Expect(timestamp).To(Equal(int64(1136214245)))
		})
	})
})
