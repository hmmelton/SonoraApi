package main

type (
	Ingredient struct {
		Name		string	`json:"name"`
		Count		float32	`json:"count"`
		Count_Type	string	`json:"count_type"`
	}

	Recipe struct {
		Id				string			`json:"id,omitempty"`
		User_Id			string			`json:"user_id"`
		First_Name		string			`json:"first_name"`
		Last_Name		string			`json:"last_name"`
		Name			string			`json:"name"`
		Image			string			`json:"image"`
		Instructions	[]string		`json:"instructions"`
		Ingredients		[]Ingredient	`json:"ingredients"`
		Tags			[]string		`json:"tags"`
		Featured		bool			`json:"featured"`
	}

	Menu struct {
		Id			string		`json:"id,omitempty"`
		UserId		string		`json:"user_id"`
		Image		string		`json:"image"`
		ProfImage	string		`json:"prof_image"`
		Name		string		`json:"name"`
		Appetizers	*[]Recipe	`json:"appetizers"`
		Entrees		*[]Recipe	`json:"entrees"`
		Sides		*[]Recipe	`json:"sides"`
		Desserts	*[]Recipe	`json:"desserts"`
		Beverages	*[]Recipe	`json:"beverages"`
		Tags		[]string	`json:"tags"`
		Featured	bool		`json:"featured"`
	}
)