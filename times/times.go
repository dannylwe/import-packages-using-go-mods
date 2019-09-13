package times

import (
	"fmt"
	"time"
)

func TimeNow() {
	dt := time.Now()
	fmt.Println("the time now is: ", dt.String())
}
