package entity

import "encoding/json"

type CustomerOneMessage struct {
	UUID  string  `json:"uuid"`
	Value float64 `json:"value"`
}

func (s *CustomerOneMessage) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func UnMarshal(bytes []byte) (*CustomerOneMessage, error) {
	var message CustomerOneMessage
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
