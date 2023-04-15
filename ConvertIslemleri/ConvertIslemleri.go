package ConvertIslemleri

import "time"

type Convert[T any] struct {
}

func (c *Convert[T]) Chunk(slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

/*
t := time.Date(2023, time.April, 14, 10, 0, 0, 0, time.UTC)
millis := TimeToMillis(t)
fmt.Println(millis)
*/
func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

/*
timeStr := "2023-04-14T10:00:00Z"
format := "2006-01-02T15:04:05Z"
millis, err := StringToMillis(timeStr, format)

	if err != nil {
	    fmt.Println(err)
	} else {

	    fmt.Println(millis)
	}
*/
func StringToMillis(timeStr string, format string) (int64, error) {
	t, err := time.Parse(format, timeStr)
	if err != nil {
		return 0, err
	}
	return t.UnixNano() / int64(time.Millisecond), nil
}
