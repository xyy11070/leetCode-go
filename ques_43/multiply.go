package ques_43

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	sumArr := make([]int, len(num2)+len(num1))
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n1*n2 + sumArr[i+j+1]
			sumArr[i+j+1] = sum % 10
			sumArr[i+j] += sum / 10
		}
	}
	sumStr := ""
	for k, v := range sumArr {
		if k == 0 && v == 0 {
			continue
		}
		sumStr += string(v + '0')
	}
	return sumStr
}

func multiply1(num1 string, num2 string) string {
	num1Arr, num2Arr, k := []string{}, []string{}, 2
	//num1Len, num2Len := len(num1), len(num2)
	for len(num1) > k {
		num1Arr = append(num1Arr, num1[len(num1)-k:])
		num1 = num1[:len(num1)-k]
	}
	if len(num1) > 0 {
		num1Arr = append(num1Arr, num1)
	}
	for len(num2) > k {
		num2Arr = append(num2Arr, num2[len(num2)-k:])
		num2 = num2[:len(num2)-k]
	}
	if len(num2) > 0 {
		num2Arr = append(num2Arr, num2)
	}
	return ""
}
