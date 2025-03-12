package domain

// User Entity
type User struct {
	id        string
	firstName string
	lastName  string
	email     string
	address   Address
}

// NewUser creates a new User entity
func NewUser(id, firstName, lastName, email string, address Address) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		address:   address,
	}
}

// ID returns the user ID
func (u *User) ID() string { return u.id }

// FirstName returns the user first name
func (u *User) FirstName() string { return u.firstName }

// LastName returns the user last name
func (u *User) LastName() string { return u.lastName }

// Email returns the user email
func (u *User) Email() string { return u.email }

// Address returns the user address
func (u *User) Address() Address { return u.address }

// Address Value Object
type Address struct {
	street  string
	city    string
	state   string
	zipCode string
}

// NewAddress creates a new Address value object
func NewAddress(street, city, state, zipCode string) Address {
	return Address{
		street:  street,
		city:    city,
		state:   state,
		zipCode: zipCode,
	}
}

// Street returns the address street
func (a Address) Street() string { return a.street }

// City returns the address city
func (a Address) City() string { return a.city }

// State returns the address state
func (a Address) State() string { return a.state }

// ZipCode returns the address zip code
func (a Address) ZipCode() string { return a.zipCode }

// MaritalStatus enum
type MaritalStatus string

const (
	// Single ...
	Single MaritalStatus = "single"
	// Married ...
	Married MaritalStatus = "married"
)

// Gender ENUM
type Gender string

const (
	// Male gender
	Male Gender = "male"
	//Female gender
	Female Gender = "female"
)

// BloodGroup : Blood group
type BloodGroup string

const (
	// OPositive : O+
	OPositive BloodGroup = "O+"
	// ONegative : O-
	ONegative BloodGroup = "O-"
	// APositive : A+
	APositive BloodGroup = "A+"
	// ANegative : A-
	ANegative BloodGroup = "A-"
	// BPositive : B+
	BPositive BloodGroup = "B+"
	// BNegative : B-
	BNegative BloodGroup = "B-"
	// ABPositive : AB+
	ABPositive BloodGroup = "AB+"
	// ABNegative : AB-
	ABNegative BloodGroup = "AB-"
)

// Genotype ...
type Genotype string

const (
	// AS Genotype
	AS Genotype = "AS"
	// AA Genotype
	AA Genotype = "AA"
	// SS Genotype
	SS Genotype = "SS"
	// Others : other genotypes
	Others Genotype = "Others"
)
