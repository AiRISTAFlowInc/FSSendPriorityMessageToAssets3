package SendPriorityMessageToAssets

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IP string  `md:"IP,required"`
	CustomerId string`md:"CustomerId,required"`
	Username string `md:"Username,required"`
	Password string `md:"Password,required"`
	StaffIdList string `md:"StaffIdList,required"`
	Message string `md:"Message,required"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["IP"])
	i.IP = strVal

	strVal, _ = coerce.ToString(values["CustomerId"])
	i.CustomerId = strVal

	strVal, _ = coerce.ToString(values["Username"])
	i.Username = strVal

	strVal, _ = coerce.ToString(values["Password"])
	i.Password = strVal

	strVal, _ = coerce.ToString(values["StaffIdList"])
	i.StaffIdList = strVal

	strVal, _ = coerce.ToString(values["Message"])
	i.Message = strVal

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP": i.IP,
		"CustomerId": i.CustomerId,
		"Username": i.Username,
		"Password": i.Password,
		"StaffIdList": i.StaffIdList,
		"Message": i.Message,
	}
}

type Output struct {
	Status string `md:"Status"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	Status, ok:= (values["Status"]).(string) // type assertion
	if ok {
		o.Status = Status
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Status":    o.Status,
	}
}

type Asset struct {
	ID                           int    `json:"Id"`
} 