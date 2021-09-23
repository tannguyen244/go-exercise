package staff

type iStaff interface {
	Clone() iStaff
	SalaryStaff() int
}
