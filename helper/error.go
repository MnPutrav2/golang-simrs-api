package helper

func Error(err error) {
	if err != nil {
		panic(err.Error())
	}
}
