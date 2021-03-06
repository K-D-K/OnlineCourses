package status

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Status enum
type Status string

const (
	STATUS_PUBLISHED Status = "published"
	STATUS_MERGED           = "merged"
	STATUS_SAVED            = "saved"
	STATUS_DRAFT            = "draft"
)

var statusArr = [...]Status{
	STATUS_PUBLISHED,
	STATUS_MERGED,
	STATUS_SAVED,
	STATUS_DRAFT,
}

var statusVsIndexMap map[Status]int

// GetIndexForStatus .
func GetIndexForStatus(status Status) int {
	if len(statusVsIndexMap) == 0 {
		statusVsIndexMap = make(map[Status]int)
		for index, statusData := range statusArr {
			statusVsIndexMap[statusData] = index
		}
	}
	return statusVsIndexMap[status]
}

var byNameMap = map[string]Status{
	"published": STATUS_PUBLISHED,
	"merged":    STATUS_MERGED,
	"saved":     STATUS_SAVED,
	"draft":     STATUS_DRAFT,
}

// Scan for status
func (status *Status) Scan(value interface{}) error {
	*status = statusArr[value.(int64)]
	return nil
}

// Value for status
func (status Status) Value() (driver.Value, error) {
	for index, statusEnum := range statusArr {
		if statusEnum == status {
			return int64(index), nil
		}
	}
	return int64(3), nil
}

// UnmarshalJSON for status
func (status *Status) UnmarshalJSON(b []byte) error {
	type TempType Status
	var a *TempType = (*TempType)(status)
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	isValid := false
	for _, st := range statusArr {
		if st == *status {
			isValid = true
		}
	}
	if !isValid {
		var errStatus string
		json.Unmarshal(b, &errStatus)
		err = errors.New("Invalid status " + errStatus)
	}
	return err
}

func (status Status) String() string {
	return string(status)
}

// StatusValueOf for Status equivalent to Java Enum.ValueOf
func StatusValueOf(status string) Status {
	return byNameMap[status]
}

// CompareStatus .
func CompareStatus(status1 Status, status2 Status) int {
	return GetIndexForStatus(status1) - GetIndexForStatus(status2)
}
