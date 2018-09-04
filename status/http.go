package status

import "net/http"

// HTTPStatusFromCode takes an protocol code and maps it to the corresponding HTTP status.
func HTTPStatusFromCode(code Code) int {
	switch code {
	case OK:
		return http.StatusOK
	case Canceled:
		return http.StatusRequestTimeout
	case Unknown:
		return http.StatusInternalServerError
	case InvalidArgument:
		return http.StatusBadRequest
	case DeadlineExceeded:
		return http.StatusRequestTimeout
	case NotFound:
		return http.StatusNotFound
	case AlreadyExists:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case Unauthenticated:
		return http.StatusUnauthorized
	case ResourceExhausted:
		return http.StatusForbidden
	case FailedPrecondition:
		return http.StatusPreconditionFailed
	case Aborted:
		return http.StatusConflict
	case OutOfRange:
		return http.StatusBadRequest
	case Unimplemented:
		return http.StatusNotImplemented
	case Internal:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	case DataLoss:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}

// HTTPStatusFromCode takes an protocol code and maps it to the corresponding HTTP status.
func CodeFromHTTPStatus(code int) Code {
	switch code {
	case http.StatusOK:
		return OK
	case http.StatusRequestTimeout:
		return Canceled
	case http.StatusBadRequest:
		return InvalidArgument
	// case http.StatusRequestTimeout:
	// return DeadlineExceeded
	case http.StatusNotFound:
		return NotFound
	case http.StatusConflict:
		return AlreadyExists
	case http.StatusForbidden:
		return PermissionDenied
	case http.StatusUnauthorized:
		return Unauthenticated
		//	case http.StatusForbidden:
		//		return ResourceExhausted
	case http.StatusPreconditionFailed:
		return FailedPrecondition
		//	case http.StatusConflict:
		//		return Aborted
		//	case http.StatusBadRequest:
		// return OutOfRange
	case http.StatusNotImplemented:
		return Unimplemented
	case http.StatusInternalServerError:
		return Internal
	case http.StatusServiceUnavailable:
		return Unavailable
		//	case http.StatusInternalServerError:
		//		return DataLoss
	}

	return Unknown
}
