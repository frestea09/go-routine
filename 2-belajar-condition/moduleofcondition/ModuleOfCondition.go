package moduleofcondition

func GetValidation(inputInt int) string {
	if inputInt > 3 {
		return "benar"
	} else {
		return "salah"
	}
}