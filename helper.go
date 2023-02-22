package main

type Customer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getResponse(code string) Response {
	// Gets a code type string based on HTTP Status Codes
	// Returns a Struct containing the code and a message.

	switch code {
	case "400":
		return Response{code, "Bad Request"}
	case "401":
		return Response{code, "Unauthorized"}
	case "403":
		return Response{code, "Forbidden"}
	case "404":
		return Response{code, "Not Found"}
	case "409":
		return Response{code, "Conflict"}
	default:
		return Response{"400", "Bad Request"}
	}
}
