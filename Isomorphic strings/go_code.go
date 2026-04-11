// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func isIsomorphic(s string, t string) bool {
    
    dataS := make(map[byte]int)
    dataT := make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        _, elementS := dataS[s[i]]
        _, elementT := dataT[t[i]]
        
        if !elementS {
            dataS[s[i]] = i
        }
        
        if !elementT{
            dataT[t[i]] = i
        }
        
        if dataS[s[i]] != dataT[t[i]] {
            return false
        }
    }
    return true
}


func main() {
  s,t := "egl", "add"
  output := isIsomorphic(s,t)
  fmt.Println(output)
}
