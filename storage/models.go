package storage

/*
	Struct to get all information from the bearer token returned by MySuperMon upon login
	It must match the json object that we get from MySuperMon
*/
type Token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

/*
	Struct for the calls that the user makes with a Usecase and a ApplicationIdentifier.
	The information is gathered from the URL and added to the struct.
	Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/
type StartAndStopCommand struct {
	Usecase               string `uri:"Usecase" binding:"required"`
	ApplicationIdentifier string `uri:"Appiden" binding:"required"`
}

/*
	Struct for the calls that the user makes with a Username and a Password.
	The information is gathered from the URL and added to the struct.
	Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/
type LoginCommand struct {
	Username string `uri:"Username" binding:"required"`
	Password string `uri:"Password" binding:"required"`
}
