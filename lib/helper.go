package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Trim(value string) string {
	return strings.TrimSpace(value)
}

func ParseBody[T any](r *http.Request) T {
	var model T
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(body, &model)
	}
	return model
}

func SendResponse(w http.ResponseWriter, statusCode int, payload map[string]any) {
	w.WriteHeader(statusCode)
	payload["status_code"] = statusCode
	b, _ := json.Marshal(payload)
	w.Write(b)
}

func SendData(w http.ResponseWriter, statusCode int, data any) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"data":    data,
	})
}

func SendDataMessage(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendDataMessageFailed(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
		"data":    data,
	})
}

func SendMessage(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
	})
}

func SendMessageFail(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
	})
}

func GetLimitOffset(r *http.Request) (limit int, offset int) {
	limit = ParseToInt(r.URL.Query().Get("limit"))
	offset = ParseToInt(r.URL.Query().Get("offset"))
	if limit == 0 {
		limit = 15
	}
	return
}

func ParseToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func ToSliceAny[T any](data []T) []any {
	result := make([]any, len(data))
	for i, val := range data {
		result[i] = val
	}
	return result
}

func DiffDays(start, end time.Time) int {
	diff := end.Sub(start)
	days := diff.Hours() / 24
	return ParseToInt(fmt.Sprintf("%.0f", days))
}

func ParseTime(timeString string) (time.Time, error) {
	layouts := []string{
		// Full date-time layouts
		time.RFC3339,           // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,       // "2006-01-02T15:04:05.999999999Z07:00"
		"2006-01-02 15:04:05",  // "2024-12-08 11:45:00"
		"2006/01/02 15:04:05",  // "2024/12/08 11:45:00"
		"02-01-2006 15:04:05",  // "08-12-2024 11:45:00"
		"02 Jan 2006 15:04:05", // "08 Dec 2024 11:45:00"
		"2006-01-02 15:04",     // "2024-12-08 11:45"
		"02 Jan 2006",          // "08 Dec 2024"
		"January 2, 2006",      // "December 8, 2024"

		// Date-only layouts
		"2006-01-02",      // "2024-12-08"
		"2006/01/02",      // "2024/12/08"
		"02-01-2006",      // "08-12-2024"
		"02 Jan 2006",     // "08 Dec 2024"
		"January 2, 2006", // "December 8, 2024"

		// Time-only layouts
		"15:04:05", // "11:45:00"
		"15:04",    // "11:45"

		// AM/PM formats
		"2006-01-02 03:04:05 PM",  // "2024-12-08 11:45:00 PM"
		"02 Jan 2006 03:04:05 PM", // "08 Dec 2024 11:45:00 PM"
		"03:04:05 PM",             // "11:45:00 PM"
		"03:04 PM",                // "11:45 PM"
	}

	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, timeString)
		if err == nil {
			return parsedTime, nil // Successfully parsed
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", timeString)
}
