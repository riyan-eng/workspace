package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumJabatanRole(t *testing.T) {
	var data = map[string]string{
		"A": "KEPALA",
		"B": "PERANGKAT",
		"C": "PERANGKAT",
		"D": "PERANGKAT",
		"E": "PERANGKAT",
		"F": "PERANGKAT",
		"G": "PERANGKAT",
		"H": "PERANGKAT",
		"I": "PERANGKAT",
		"J": "PERANGKAT",
		"K": "PERANGKAT",
		"L": "PERANGKAT",
		"M": "PERANGKAT",
		"N": "PERANGKAT",
		"O": "PERANGKAT",
		"P": "PERANGKAT",
		"Q": "STAF",
	}

	for i, d := range data {
		expect := d
		actual := NewEnum().JabatanRole(&i)
		assert.Equal(t,
			expect,
			actual,
		)
	}
}

func TestEnumJenisKelaminListName(t *testing.T) {
	expect := []string{"Laki-laki", "Perempuan"}
	actual := NewEnum().JenisKelaminListName()
	assert.Equal(t,
		expect,
		actual,
	)
}

func TestEnumJenisKelaminCodeByName(t *testing.T) {
	var data = map[string]string{
		"Laki-laki": "L",
		"Perempuan": "P",
	}

	for i, d := range data {
		expect := d
		actual := NewEnum().JenisKelaminCodeByName(i)
		assert.Equal(t,
			expect,
			actual,
		)
	}
}

func TestEnumAgamaListName(t *testing.T) {
	expect := []string{"Islam", "Kristen", "Katolik", "Hindu", "Budha", "Konghucu", "Lainnya"}
	actual := NewEnum().AgamaListName()
	assert.Equal(t,
		expect,
		actual,
	)
}
