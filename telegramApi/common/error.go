package common

func CheckError(err error, hType int) {
	if err == nil {
		return
	}
	panic(err)
	//switch hType {
	//case 1:
	//	log.Println(err)
	//case 2:
	//	log.Fatalln(err)
	//}
}
