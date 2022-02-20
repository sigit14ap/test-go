package domain

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID           string `json:"id"`
	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Name         string `json:"name"`
}
