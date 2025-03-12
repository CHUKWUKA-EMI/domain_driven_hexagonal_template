package domain

// OnboardingStep tells the user the next required
// action to take  when onboarding
type OnboardingStep string

const (
	// ConfirmIdentity [ Required Action ]
	ConfirmIdentity OnboardingStep = "confirmIdentity"
	// AddSubscriptionPlan [ Required Action ]
	AddSubscriptionPlan OnboardingStep = "addSubscriptionPlan"
	// ProvideMedicalLicense [ Required Action ]
	ProvideMedicalLicense OnboardingStep = "provideMedicalLicense"
	// MedicalLicensePendingVerification [ Required Action ]
	MedicalLicensePendingVerification OnboardingStep = "medicalLicensePendingVerification"
	// Onboarded [Completed onboarding]
	Onboarded OnboardingStep = "onboarded"
)
