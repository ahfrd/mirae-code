package response

type GenericResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type StatusData int

type StatusMessagePair struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const (
	Success                              StatusData = 2002
	SuccessCreating                      StatusData = 2010
	BadRequestMissingMandatory           StatusData = 4001
	BadRequestNoDataFound                StatusData = 4040
	BadRequestUnauthorized               StatusData = 4010
	BadRequestMissingAuth                StatusData = 4170
	FailedGettingDataErrorDB             StatusData = 5001
	FailedGettingDataErrorBindingRequest StatusData = 5002
	FailedStoringDataErrorDB             StatusData = 5003
	FailedGettingDataErrorInternal       StatusData = 5004
	FailedCreatingDataErrorInternal      StatusData = 5020
	GeneralError                         StatusData = 4002
	BadRequestMarshalError               StatusData = 4003
)

func (sD StatusData) Message() string {
	return statusMessage[sD].Message
}

func (sD StatusData) Code() int {
	return statusMessage[sD].Code
}

func (sD StatusData) Status() string {
	return statusMessage[sD].Status
}

func (gR GenericResponse) GenericSuccess() GenericResponse {
	return GenericResponse{
		Status: "success",
	}
}

func (gR GenericResponse) GenericError() GenericResponse {
	return GenericResponse{
		Status: "error",
	}
}

var statusMessage = map[StatusData]StatusMessagePair{
	Success: {
		Status:  "success",
		Message: "success %s",
		Code:    200,
	},
	SuccessCreating: {
		Status:  "success",
		Message: "success creating",
		Code:    201,
	},
	BadRequestMissingMandatory: {
		Status:  "error",
		Message: "missing required data of %s",
		Code:    400,
	},
	BadRequestNoDataFound: {
		Status:  "error",
		Message: "no data for %s could be found",
		Code:    404,
	},
	FailedGettingDataErrorDB: {
		Status:  "error",
		Message: "failed while getting data from database (%s)",
		Code:    500,
	},
	FailedGettingDataErrorBindingRequest: {
		Status:  "error",
		Message: "failed while binding data (%s)",
		Code:    500,
	},
	FailedStoringDataErrorDB: {
		Status:  "error",
		Message: "failed while storing data (%s)",
		Code:    500,
	},
	FailedGettingDataErrorInternal: {
		Status:  "error",
		Message: "failed while calling internal (%s)",
		Code:    500,
	},
	BadRequestMissingAuth: {
		Status:  "error",
		Message: "missing auth",
		Code:    417,
	},
	FailedCreatingDataErrorInternal: {
		Status:  "error internal",
		Message: "error creating data %s",
		Code:    502,
	},
	GeneralError: {
		Status:  "error",
		Message: "error %s",
		Code:    400,
	},
	BadRequestMarshalError: {
		Status:  "error",
		Message: "error marshaling JSON : %s",
		Code:    400,
	},
}
