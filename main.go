package deanio

func Init() error {
	go startUI()
	initBrain()
	return initUtil()
}
