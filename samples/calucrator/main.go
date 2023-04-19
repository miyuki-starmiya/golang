package main

import (
	"fmt"
)

func main() {
	// 変数の宣言
	var num1, num2, lastResult float64
	var operator, keepCalculating string

	// breakしない限りloop
	for {
		fmt.Printf("現在の計算結果: %f\n", lastResult)

		fmt.Print("数値1を入力してください (直前の結果を使用する場合は 'R' を入力): ")
		// num1のアドレスにある変数を格納
		// return 読み取った値, error
		_, err := fmt.Scanf("%f", &num1)
		if err != nil {
			var useResult string
			fmt.Scanf("%s", &useResult)
			if useResult == "R" || useResult == "r" {
				num1 = lastResult
			} else {
				fmt.Println("数値1の入力に誤りがあります。")
				continue
			}
		}

		fmt.Print("演算子を入力してください (+, -, *, /): ")
		fmt.Scanf("%s", &operator)

		fmt.Print("数値2を入力してください: ")
		_, err = fmt.Scanf("%f", &num2)
		if err != nil {
			fmt.Println("数値2の入力に誤りがあります。")
			continue
		}

		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%f %s %f = %f\n", num1, operator, num2, result)
		lastResult = result

		fmt.Print("続けますか？(Y/N/Reset): ")
		fmt.Scanf("%s", &keepCalculating)

		if keepCalculating == "N" || keepCalculating == "n" {
			break
		} else if keepCalculating == "Reset" || keepCalculating == "reset" {
			lastResult = 0
		}
	}
}

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("0で割ることはできません。")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("無効な演算子: %s", operator)
	}
}
