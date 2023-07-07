package main

func generateJson() jsonResponse {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	return payload
}
