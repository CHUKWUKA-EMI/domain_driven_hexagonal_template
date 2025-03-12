package domain

// Role represents the user roles
type Role string

const (
	// AdminRole represents the admin role
	AdminRole Role = "admin"
	// DoctorRole represents the doctor role
	DoctorRole Role = "doctor"
	// PatientRole represents the patient role
	PatientRole Role = "patient"
	// SuperAdminRole represents the super admin role
	SuperAdminRole Role = "superAdmin"
)
