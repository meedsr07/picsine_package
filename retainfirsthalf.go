package piscine

func RetainFirstHalf(str string) string {
	if len(str)==1||str==""{
		return str
	}
	return str [0:len(str)/2]

}