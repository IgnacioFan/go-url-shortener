package handler

type ShortUrlRequest struct {
	Url                 string `json:"url"`
	ExpirationlenInMins int    `json:"expiration_len_in_mins"`
}
