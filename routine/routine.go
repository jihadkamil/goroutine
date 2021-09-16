package routine

import (
	"fmt"
	"time"
)



func ResetAttemptRoutine(timer int64){
	firstLoop := true
	var success bool
	var err error
	fmt.Println("resert auto",time.Now())


	for {
		// loop delay control
		fmt.Println("firstLoop",firstLoop, "succes",success, "err",err)
		if !firstLoop{
			time.Sleep(
				time.Duration(
					// per 15 minutes
					15*60*time.Second,
				),
			)
			// dataFetched := fetchDataFromAPI()
			// myData,success, err = myUseCase.GetData(dataFetched)
		if !success && err != nil {
			fmt.Println("return data from usecase")
			/*
			ctx.JSON(http.StatusOK, gin.H{
				"data":      myData,
				"error":     "",
			})
			// */
			}
		}
		firstLoop = false
	}
}