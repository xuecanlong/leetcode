package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate("(1+2)*5*(2+3)*1"));
	//fmt.Println(calculate("1+2*3+(4*5+6)*7"));
}

func calculate(s string) int32 {
	//中缀表达式转化为后缀表达式
	numStack := make([]int32, 0)
	opStack := make([]rune, 0)
	for _, data := range s {
		switch data {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': //数字直接入栈numStack
			numStack = append(numStack, data-'0')
			//如果遇到任何其他的操作符，如（“+”， “*”，“（”）等，从栈中弹出元素直到遇到发现更低优先级的元素(或者栈为空)为止。弹出完这些元素后，才将遇到的操作符压入到栈中。有一点需要注意，只有在遇到" ) "的情况下我们才弹出" ( "，其他情况我们都不会弹出" ( "
		case '('://遇到'('直接入栈opStack
			opStack = append(opStack, data)
		case ')'://遇到')'把'('之前的栈顶移到numStack
			for i := len(opStack) - 1; i >= 0; i -- {
				op := opStack[i]
				opStack = opStack[:i]
				if op != '(' {
					numStack = append(numStack, op)
				} else {
					break
				}
			}
		case '*'://遇到'*'把'*'和'+'的栈顶移到numStack
			for i := len(opStack) - 1; i >= 0; i -- {
				op := opStack[i]
				if op == '*' {
					numStack = append(numStack, op)
					opStack = opStack[:i]
				} else {
					break
				}
			}
			opStack = append(opStack, data)
		case '+'://遇到'+'把'+'的栈顶移到numStack
			for i := len(opStack) - 1; i >= 0; i -- {
				op := opStack[i]
				if op == '+' || op == '*' {
					numStack = append(numStack, op)
					opStack = opStack[:i]
				} else {
					break
				}
			}
			opStack = append(opStack, data)
		}
	}

	for i := len(opStack) - 1; i >= 0; i -- { //把opStack放到numStack
		op := opStack[i]
		numStack = append(numStack, op)
		opStack = opStack[:i]
	}
	
	//计算
	//遇到数字，入栈
	//遇到运算符，弹出栈顶两个元素，做运算，并将结果入栈
	resStack := make([]int32, 0)
	for _, data := range numStack {
		switch data {
		case 0,1,2,3,4,5,6,7,8,9:
			resStack = append(resStack, data)
		case '*':
			num1 := resStack[len(resStack) - 1]
			num2 := resStack[len(resStack) - 2]
			resStack = resStack[:len(resStack) - 2]
			resStack = append(resStack, num1 * num2)
		case '+':
			num1 := resStack[len(resStack) - 1]
			num2 := resStack[len(resStack) - 2]
			resStack = resStack[:len(resStack) - 2]
			resStack = append(resStack, num1 + num2)
		}
	}
	return resStack[0]
}