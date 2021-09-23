package staff

import "fmt"

func TotalSalary() {
	permanent := ListPermanent{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	contract := ListContract{
		{1, 2},
		{1, 3},
	}
	clonePermanent := permanent.Clone()
	cloneContract := contract.Clone()
	fmt.Println((clonePermanent.SalaryStaff() + cloneContract.SalaryStaff()))
}
