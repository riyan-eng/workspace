package util

import "sort"

var jabatanRole = map[string]string{
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

var jenisKelaminOrder = map[int]string{
	1: "L",
	2: "P",
}

var jenisKelamin = map[string]string{
	"L": "Laki-laki",
	"P": "Perempuan",
}

var agamaOrder = map[int]string{
	1: "I",
	2: "KR",
	3: "KA",
	4: "H",
	5: "B",
	6: "K",
	7: "L",
}

var agama = map[string]string{
	"I":  "Islam",
	"KR": "Kristen",
	"KA": "Katolik",
	"H":  "Hindu",
	"B":  "Budha",
	"K":  "Konghucu",
	"L":  "Lainnya",
}

var pendidikanOrder = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
}

var pendidikan = map[string]string{
	"A": "Tidak/Belum Sekolah",
	"B": "Belum Tamat SD/Sederajat",
	"C": "Tamat SD/Sederajat",
	"D": "SLTP/Sederajat",
	"E": "SLTA/Sederajat",
	"F": "Diploma I/II",
	"G": "Akademi/Diploma III/S. Muda",
	"H": "Diploma IV/Strata I",
	"I": "Strata II",
	"J": "Strata III",
}

var statusKeluargaOrder = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
}

var statusKeluarga = map[string]string{
	"A": "Kepala Keluarga",
	"B": "Suami",
	"C": "Istri",
	"D": "Anak",
	"E": "Menantu",
	"F": "Cucu",
	"G": "Orang Tua",
	"H": "Mertua",
	"I": "Family Lain",
	"J": "Lainnya",
}

var golonganDarahOrder = map[int]string{
	1: "A",
	2: "B",
	3: "AB",
	4: "O",
	5: "T",
}

var golonganDarah = map[string]string{
	"A":  "A",
	"B":  "B",
	"AB": "AB",
	"O":  "O",
	"T":  "Tidak Tahu",
}

var statusPerkawinanOrder = map[int]string{
	1: "K",
	2: "BK",
	3: "CH",
	4: "CM",
}

var statusPerkawinan = map[string]string{
	"K":  "Kawin",
	"BK": "Belum Kawin",
	"CH": "Cerai Hidup",
	"CM": "Cerai Mati",
}

var kelainanFisikMentalOrder = map[int]string{
	1: "A",
	2: "TA",
}

var kelainanFisikMental = map[string]string{
	"A":  "Ada",
	"TA": "Tidak Ada",
}

var bulan = map[int]string{
	1:  "Januari",
	2:  "Februari",
	3:  "Maret",
	4:  "April",
	5:  "Mei",
	6:  "Juni",
	7:  "Juli",
	8:  "Agustus",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Desember",
}

var pertanyaanSurveiPilihanOrder = map[int]string{
	1: "SB",
	2: "B",
	3: "C",
	4: "K",
	5: "SK",
}

var pertanyaanSurveiPilihan = map[string]string{
	"SB": "Sangat Baik",
	"B":  "Baik",
	"C":  "Cukup",
	"K":  "Kurang",
	"SK": "SangatKurang",
}

var kunjunganTamuKategori = map[string]string{
	"P": "Personal",
	"K": "Kelompok",
}

type enum struct{}

func NewEnum() *enum {
	return &enum{}
}

func (m *enum) JabatanRole(jabatanCode *string) string {
	return jabatanRole[*jabatanCode]
}

func (m *enum) JenisKelaminListName() []string {
	var key []int
	for k := range jenisKelaminOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, jenisKelamin[jenisKelaminOrder[k]])
	}
	return name
}

func (m *enum) JenisKelaminCodeByName(name string) string {
	code := new(string)
	for k, v := range jenisKelamin {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) AgamaListName() []string {
	var key []int
	for k := range agamaOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, agama[agamaOrder[k]])
	}
	return name
}

func (m *enum) AgamaCodeByName(name string) string {
	code := new(string)
	for k, v := range agama {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) PendidikanListName() []string {
	var key []int
	for k := range pendidikanOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, pendidikan[pendidikanOrder[k]])
	}
	return name
}

func (m *enum) PendidikanCodeByName(name string) string {
	code := new(string)
	for k, v := range pendidikan {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) StatusKeluargaListName() []string {
	var key []int
	for k := range statusKeluargaOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, statusKeluarga[statusKeluargaOrder[k]])
	}
	return name
}

func (m *enum) StatusKeluargaCodeByName(name string) string {
	code := new(string)
	for k, v := range statusKeluarga {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) GolonganDarahListName() []string {
	var key []int
	for k := range golonganDarahOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, golonganDarah[golonganDarahOrder[k]])
	}
	return name
}

func (m *enum) GolonganDarahCodeByName(name string) string {
	code := new(string)
	for k, v := range golonganDarah {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) StatusPerkawinanListName() []string {
	var key []int
	for k := range statusPerkawinanOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, statusPerkawinan[statusPerkawinanOrder[k]])
	}
	return name
}

func (m *enum) StatusPerkawinanCodeByName(name string) string {
	code := new(string)
	for k, v := range statusPerkawinan {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) KelainanFisikMentalListName() []string {
	var key []int
	for k := range kelainanFisikMentalOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, kelainanFisikMental[kelainanFisikMentalOrder[k]])
	}
	return name
}

func (m *enum) KelainanFisikMentalCodeByName(name string) string {
	code := new(string)
	for k, v := range kelainanFisikMental {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

func (m *enum) Bulan(bulanInt *int) string {
	return bulan[*bulanInt]
}

func (m *enum) PertanyaanSurveiPilihanOrder() map[int]string {
	return pertanyaanSurveiPilihanOrder
}

func (m *enum) PertanyaanSurveiPilihan() map[string]string {
	return pertanyaanSurveiPilihan
}

func (m *enum) KunjunganTamuKategori(kategori *string) string {
	return kunjunganTamuKategori[*kategori]
}
