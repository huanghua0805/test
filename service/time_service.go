package service

import (
	"time"
)

func CircleTime(TimeStart int64,CircleLimitHour int64)(CircleState bool){
	TimeStop := time.Now().Unix()
	RunTime :=TimeStop - TimeStart
	if RunTime < CircleLimitHour*60*60{
		return  true
	}else{
		return  false
	}

}

