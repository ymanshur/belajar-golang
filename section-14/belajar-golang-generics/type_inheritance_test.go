package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetSalary() int
}

type ManagerA struct {
	Name   string
	Salary int
}

func (m *ManagerA) GetName() string {
	return m.Name
}

func (m *ManagerA) GetSalary() int {
	return m.Salary
}

type VicePresident interface {
	GetName() string
	GetPeriod() time.Duration
}

type VicePresidentA struct {
	Name   string
	Period time.Duration
}

func (v *VicePresidentA) GetName() string {
	return v.Name
}

func (v *VicePresidentA) GetPeriod() time.Duration {
	return v.Period
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Yusuf", GetName[Manager](&ManagerA{Name: "Yusuf"}))
	assert.Equal(t, "Manshur", GetName[VicePresident](&VicePresidentA{Name: "Manshur"}))
}
