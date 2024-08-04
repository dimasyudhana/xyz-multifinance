package response

import (
	"net/http"
	"strings"
)

func CheckHandlerSuccessCode(msg string) int {
	switch true {
	case strings.Contains(msg, "insert") || strings.Contains(msg, "create"):
		return http.StatusCreated
	case strings.Contains(msg, "read") || strings.Contains(msg, "get") || strings.Contains(msg, "update") || strings.Contains(msg, "delete"):
		return http.StatusOK
	default:
		return http.StatusOK
	}
}

func CheckHandlerErrorCode(err error) (responseCode int, layerCode string, errConst error) {
	switch err.Error() {
	case ERR_AuthWrongCredentials:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case JWT_InvalidJwtToken:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case JWT_FailedCastingJwtToken:
		return http.StatusInternalServerError, LAYER_HANDLER_CODE, err

	case JWT_FailedCreateToken:
		return http.StatusInternalServerError, LAYER_SERVICE_CODE, err

	case REQ_ErrorBindData:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case REQ_InvalidParam:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case REQ_InvalidIdParam:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case REQ_InvalidPageParam:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case REQ_InvalidLimitParam:
		return http.StatusBadRequest, LAYER_HANDLER_CODE, err

	case VAL_InvalidImageFileType:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case VAL_InvalidFileSize:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case VAL_InvalidValidation:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case VAL_IncompleteAnswer:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case VAL_PasswordNotSet:
		return http.StatusBadRequest, LAYER_SERVICE_CODE, err

	case VAL_Unauthorized:
		return http.StatusUnauthorized, LAYER_HANDLER_CODE, err

	case DB_ERR_RECORD_NOT_FOUND:
		return http.StatusNotFound, LAYER_DATA_CODE, err

	case DB_ERR_MISSING_WHERE_CLAUSE:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_UNSUPPORTED_RELATION:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_INVALID_DATA:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_INVALID_FIELD:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_PRELOAD_NOT_ALLOWED:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_INVALID_DB:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_PRIMARY_KEY_REQUIRED:
		return http.StatusInternalServerError, LAYER_DATA_CODE, err

	case DB_ERR_DUPLICATE_KEY:
		return http.StatusBadRequest, LAYER_DATA_CODE, err

	default:
		return http.StatusInternalServerError, LAYER_DEFAULT_CODE, err
	}
}
