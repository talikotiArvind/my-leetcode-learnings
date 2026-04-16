func convertToTitle(columnNumber int) string {
    result := []byte{}
    for columnNumber > 0 {
        columnNumber--
        result = append(result, byte(columnNumber%26)+'A')
        columnNumber /= 26
    }
    
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result)
}

func main() {
    fmt.Println(convertToTitle(1))    // "A"
    fmt.Println(convertToTitle(28))   // "AB"
    fmt.Println(convertToTitle(701))  // "ZY"
}
