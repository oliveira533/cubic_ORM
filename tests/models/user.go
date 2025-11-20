package models_test

type User struct {
	ID    int    `cubic:"id,primary_key,auto_increment"`
	Name  string `cubic:"name,size=100,not_null"`
	Email string `cubic:"email,unique"`
}
