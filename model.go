package main

type PerguruanTinggi struct {
	IDSp   string `json:"id_sp"`
	KodePt string `json:"kode_pt"`
	NamaPt string `json:"nama_pt"`
}

type DataMengajar struct {
	IDSmt  string `json:"id_smt"`
	NmKls  string `json:"nm_kls"`
	KodeMk string `json:"kode_mk"`
	NmMk   string `json:"nm_mk"`
	Namapt string `json:"namapt"`
	Linkpt string `json:"linkpt"`
}

type DataPendidikan struct {
	ThnLulus     int    `json:"thn_lulus"`
	NmSpFormal   string `json:"nm_sp_formal"`
	Namajenjang  string `json:"namajenjang"`
	SingkatGelar string `json:"singkat_gelar"`
}

type DataUmum struct {
	IDSdm           string `json:"id_sdm"`
	NmSdm           string `json:"nm_sdm"`
	Jk              string `json:"jk"`
	TmptLahir       string `json:"tmpt_lahir"`
	Namapt          string `json:"namapt"`
	Linkpt          string `json:"linkpt"`
	Linkprodi       string `json:"linkprodi"`
	Namaprodi       string `json:"namaprodi"`
	Statuskeaktifan string `json:"statuskeaktifan"`
	PendTinggi      string `json:"pend_tinggi"`
	Fungsional      string `json:"fungsional"`
	Ikatankerja     string `json:"ikatankerja"`
}

type DosenDetail struct {
	Datamengajar   []DataMengajar   `json:"datamengajar"`
	Datapendidikan []DataPendidikan `json:"datapendidikan"`
	Dataumum       DataUmum         `json:"dataumum"`
}

type Dosen struct {
	ID      string `json:"id"`
	Nama    string `json:"nama"`
	Nip     string `json:"nip"`
	Pt      string `json:"pt"`
	Jenjang string `json:"jenjang"`
	Prodi   string `json:"prodi"`
}
