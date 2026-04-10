package types

type City struct {
	ID            int    `json:"id"`
	ProvinsiID    int    `json:"provinsi_id"`
	KabupatenName string `json:"kabupaten_name"`
	Type          string `json:"type"`
	PostalCode    string `json:"postal_code"`
}

type District struct {
	ID            int    `json:"id"`
	KecamatanName string `json:"kecamatan_name"`
	KabupatenID   int    `json:"kabupaten_id"`
}

type SubDistrict struct {
	ID            int    `json:"id"`
	KelurahanName string `json:"kelurahan_name"`
	KecamatanID   int    `json:"kecamatan_id"`
}

type AddressByNameResult struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
