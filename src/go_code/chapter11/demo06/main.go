package main

import (
	"fmt"
)

func main() {
	var score [3][5]float64
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}
	fmt.Println(score)
	totalSum := 0.0
	for i := 0; i < len(score); i++ {
		sum := 0.0
		for j := 0; j < len(score[i]); j++ {
			sum += score[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分为%v，平均分%v\n", i+1, sum, sum/float64(len(score[i])))
	}

	fmt.Printf("所有班级的总分为%v，所有班级平均分为%v\n", totalSum, totalSum/15)

}
