package modulelooping

import "fmt"		

func GetLooping(inputLooping []int)  {
	for _, item := range inputLooping {
		fmt.Print(item);
	};
}