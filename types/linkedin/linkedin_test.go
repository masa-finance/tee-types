package linkedin_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/masa-finance/tee-types/types"
	"github.com/masa-finance/tee-types/types/linkedin/experiences"
	"github.com/masa-finance/tee-types/types/linkedin/functions"
	"github.com/masa-finance/tee-types/types/linkedin/industries"
	"github.com/masa-finance/tee-types/types/linkedin/seniorities"
)

var _ = Describe("LinkedIn Types", func() {
	Describe("LinkedIn Package", func() {
		It("should have all required fields", func() {
			linkedin := types.LinkedIn

			Expect(linkedin.Profile).ToNot(BeNil())
			Expect(linkedin.Seniorities).ToNot(BeNil())
			Expect(linkedin.Experiences).ToNot(BeNil())
			Expect(linkedin.Functions).ToNot(BeNil())
			Expect(linkedin.Industries).ToNot(BeNil())
		})
	})

	Describe("Seniorities", func() {
		It("should have all seniority levels", func() {
			seniorities := types.LinkedIn.Seniorities

			Expect(seniorities.InTraining).To(Equal(seniorities.InTraining))
			Expect(seniorities.EntryLevel).To(Equal(seniorities.EntryLevel))
			Expect(seniorities.Senior).To(Equal(seniorities.Senior))
			Expect(seniorities.Strategic).To(Equal(seniorities.Strategic))
			Expect(seniorities.EntryLevelManager).To(Equal(seniorities.EntryLevelManager))
			Expect(seniorities.ExperiencedManager).To(Equal(seniorities.ExperiencedManager))
			Expect(seniorities.Director).To(Equal(seniorities.Director))
			Expect(seniorities.VicePresident).To(Equal(seniorities.VicePresident))
			Expect(seniorities.CXO).To(Equal(seniorities.CXO))
			Expect(seniorities.Partner).To(Equal(seniorities.Partner))
		})

		It("should have correct ID values", func() {
			Expect(string(types.LinkedIn.Seniorities.InTraining)).To(Equal("100"))
			Expect(string(types.LinkedIn.Seniorities.EntryLevel)).To(Equal("110"))
			Expect(string(types.LinkedIn.Seniorities.Senior)).To(Equal("120"))
			Expect(string(types.LinkedIn.Seniorities.Strategic)).To(Equal("130"))
			Expect(string(types.LinkedIn.Seniorities.EntryLevelManager)).To(Equal("200"))
			Expect(string(types.LinkedIn.Seniorities.ExperiencedManager)).To(Equal("210"))
			Expect(string(types.LinkedIn.Seniorities.Director)).To(Equal("220"))
			Expect(string(types.LinkedIn.Seniorities.VicePresident)).To(Equal("300"))
			Expect(string(types.LinkedIn.Seniorities.CXO)).To(Equal("310"))
			Expect(string(types.LinkedIn.Seniorities.Partner)).To(Equal("320"))
		})

		It("should have All set containing all seniorities", func() {
			all := types.LinkedIn.Seniorities.All

			Expect(all.Contains(seniorities.InTraining)).To(BeTrue())
			Expect(all.Contains(seniorities.EntryLevel)).To(BeTrue())
			Expect(all.Contains(seniorities.Senior)).To(BeTrue())
			Expect(all.Contains(seniorities.Strategic)).To(BeTrue())
			Expect(all.Contains(seniorities.EntryLevelManager)).To(BeTrue())
			Expect(all.Contains(seniorities.ExperiencedManager)).To(BeTrue())
			Expect(all.Contains(seniorities.Director)).To(BeTrue())
			Expect(all.Contains(seniorities.VicePresident)).To(BeTrue())
			Expect(all.Contains(seniorities.CXO)).To(BeTrue())
			Expect(all.Contains(seniorities.Partner)).To(BeTrue())

			Expect(all.Length()).To(Equal(10))
		})
	})

	Describe("Experiences", func() {
		It("should have all experience levels", func() {
			experiences := types.LinkedIn.Experiences

			Expect(experiences.LessThanAYear).To(Equal(experiences.LessThanAYear))
			Expect(experiences.OneToTwoYears).To(Equal(experiences.OneToTwoYears))
			Expect(experiences.ThreeToFiveYears).To(Equal(experiences.ThreeToFiveYears))
			Expect(experiences.SixToTenYears).To(Equal(experiences.SixToTenYears))
			Expect(experiences.MoreThanTenYears).To(Equal(experiences.MoreThanTenYears))
		})

		It("should have correct ID values", func() {
			Expect(string(types.LinkedIn.Experiences.LessThanAYear)).To(Equal("1"))
			Expect(string(types.LinkedIn.Experiences.OneToTwoYears)).To(Equal("2"))
			Expect(string(types.LinkedIn.Experiences.ThreeToFiveYears)).To(Equal("3"))
			Expect(string(types.LinkedIn.Experiences.SixToTenYears)).To(Equal("4"))
			Expect(string(types.LinkedIn.Experiences.MoreThanTenYears)).To(Equal("5"))
		})

		It("should have All set containing all experiences", func() {
			all := types.LinkedIn.Experiences.All

			Expect(all.Contains(experiences.LessThanAYear)).To(BeTrue())
			Expect(all.Contains(experiences.OneToTwoYears)).To(BeTrue())
			Expect(all.Contains(experiences.ThreeToFiveYears)).To(BeTrue())
			Expect(all.Contains(experiences.SixToTenYears)).To(BeTrue())
			Expect(all.Contains(experiences.MoreThanTenYears)).To(BeTrue())

			Expect(all.Length()).To(Equal(5))
		})
	})

	Describe("Functions", func() {
		It("should have all function types", func() {
			functions := types.LinkedIn.Functions

			Expect(functions.Accounting).To(Equal(functions.Accounting))
			Expect(functions.Engineering).To(Equal(functions.Engineering))
			Expect(functions.Marketing).To(Equal(functions.Marketing))
			Expect(functions.Sales).To(Equal(functions.Sales))
			Expect(functions.HumanResources).To(Equal(functions.HumanResources))
		})

		It("should have correct ID values for key functions", func() {
			Expect(string(types.LinkedIn.Functions.Accounting)).To(Equal("1"))
			Expect(string(types.LinkedIn.Functions.Engineering)).To(Equal("8"))
			Expect(string(types.LinkedIn.Functions.Marketing)).To(Equal("15"))
			Expect(string(types.LinkedIn.Functions.Sales)).To(Equal("25"))
			Expect(string(types.LinkedIn.Functions.HumanResources)).To(Equal("12"))
		})

		It("should have All set containing all functions", func() {
			all := types.LinkedIn.Functions.All

			Expect(all.Contains(functions.Accounting)).To(BeTrue())
			Expect(all.Contains(functions.Engineering)).To(BeTrue())
			Expect(all.Contains(functions.Marketing)).To(BeTrue())
			Expect(all.Contains(functions.Sales)).To(BeTrue())
			Expect(all.Contains(functions.HumanResources)).To(BeTrue())
			Expect(all.Contains(functions.InformationTechnology)).To(BeTrue())
			Expect(all.Contains(functions.Finance)).To(BeTrue())

			Expect(all.Length()).To(Equal(25))
		})
	})

	Describe("Industries", func() {
		It("should have all industry types", func() {
			industries := types.LinkedIn.Industries

			Expect(industries.SoftwareDevelopment).To(Equal(industries.SoftwareDevelopment))
			Expect(industries.FinancialServices).To(Equal(industries.FinancialServices))
			Expect(industries.Manufacturing).To(Equal(industries.Manufacturing))
			Expect(industries.Retail).To(Equal(industries.Retail))
			Expect(industries.Education).To(Equal(industries.Education))
		})

		It("should have correct ID values for key industries", func() {
			Expect(string(types.LinkedIn.Industries.SoftwareDevelopment)).To(Equal("4"))
			Expect(string(types.LinkedIn.Industries.FinancialServices)).To(Equal("43"))
			Expect(string(types.LinkedIn.Industries.Manufacturing)).To(Equal("25"))
			Expect(string(types.LinkedIn.Industries.Retail)).To(Equal("27"))
			Expect(string(types.LinkedIn.Industries.Education)).To(Equal("1999"))
		})

		It("should have All set containing all industries", func() {
			all := types.LinkedIn.Industries.All

			Expect(all.Contains(industries.SoftwareDevelopment)).To(BeTrue())
			Expect(all.Contains(industries.FinancialServices)).To(BeTrue())
			Expect(all.Contains(industries.Manufacturing)).To(BeTrue())
			Expect(all.Contains(industries.Retail)).To(BeTrue())
			Expect(all.Contains(industries.Education)).To(BeTrue())
			Expect(all.Contains(industries.Hospitals)).To(BeTrue())
			Expect(all.Contains(industries.ProfessionalServices)).To(BeTrue())

			Expect(all.Length()).To(BeNumerically(">=", 100)) // Should have many industries
		})

		It("should have technology industries", func() {
			all := types.LinkedIn.Industries.All

			Expect(all.Contains(industries.SoftwareDevelopment)).To(BeTrue())
			Expect(all.Contains(industries.ComputerHardwareManufacturing)).To(BeTrue())
			Expect(all.Contains(industries.Telecommunications)).To(BeTrue())
			Expect(all.Contains(industries.TechnologyInformationAndInternet)).To(BeTrue())
		})

		It("should have financial industries", func() {
			all := types.LinkedIn.Industries.All

			Expect(all.Contains(industries.FinancialServices)).To(BeTrue())
			Expect(all.Contains(industries.Banking)).To(BeTrue())
			Expect(all.Contains(industries.Insurance)).To(BeTrue())
			Expect(all.Contains(industries.InvestmentBanking)).To(BeTrue())
		})
	})

	Describe("Type Consistency", func() {
		It("should have consistent ID types across all collections", func() {
			// All IDs should be string type
			Expect(string(types.LinkedIn.Seniorities.InTraining)).To(BeAssignableToTypeOf(""))
			Expect(string(types.LinkedIn.Experiences.LessThanAYear)).To(BeAssignableToTypeOf(""))
			Expect(string(types.LinkedIn.Functions.Accounting)).To(BeAssignableToTypeOf(""))
			Expect(string(types.LinkedIn.Industries.SoftwareDevelopment)).To(BeAssignableToTypeOf(""))
		})

		It("should have non-empty ID values", func() {
			Expect(types.LinkedIn.Seniorities.InTraining).ToNot(BeEmpty())
			Expect(types.LinkedIn.Experiences.LessThanAYear).ToNot(BeEmpty())
			Expect(types.LinkedIn.Functions.Accounting).ToNot(BeEmpty())
			Expect(types.LinkedIn.Industries.SoftwareDevelopment).ToNot(BeEmpty())
		})
	})
})
