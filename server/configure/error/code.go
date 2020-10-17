package error

type StatusCode int

const (
	BadRequest          StatusCode = 400001
	Forbidden           StatusCode = 403001
	NotFound            StatusCode = 404001
	InternalServerError StatusCode = 500001
)

var statusMap = map[StatusCode]map[string]interface{}{
	BadRequest:          {"httpCode": 400, "message": "허용하지 않는 파라미터입니다."},
	Forbidden:           {"httpCode": 403, "message": "권한이 없습니다."},
	NotFound:            {"httpCode": 404, "message": "리소스를 찾을 수 없습니다."},
	InternalServerError: {"httpCode": 500, "message": "알 수 없는 에러가 발생하였습니다."},
}

func castInt(s StatusCode) int {
	return int(s)
}

func httpCode(s StatusCode) int {
	return statusMap[s]["httpCode"].(int)
}

func statusText(s StatusCode) string {
	return statusMap[s]["message"].(string)
}
