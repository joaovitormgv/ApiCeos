package models

import "encoding/json"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// função de desserializar dados
func (u *User) UnmarshalData(data []byte) error {
	err := json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	return nil
}

// função de serializar dados
func (u *User) MarshalData() ([]byte, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return data, nil
}
