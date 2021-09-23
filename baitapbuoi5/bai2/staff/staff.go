package staff

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

type ListPermanent []Permanent

type ListContract []Contract

func (s ListPermanent) Clone() iStaff {
	return &s
}

func (s ListContract) Clone() iStaff {
	return &s
}

func (s ListPermanent) SalaryStaff() int {
	totalsalary := 0
	for _, item := range s {
		totalsalary += (item.basicpay + item.pf)
	}
	return totalsalary
}

func (s ListContract) SalaryStaff() int {
	totalsalary := 0
	for _, item := range s {
		totalsalary += item.basicpay
	}
	return totalsalary
}
