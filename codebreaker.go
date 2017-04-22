package codebreaker

import "fmt"

func codebreaker(s string) string{
  var equis string = ""
  var rayita string = ""
  secret := "2584"
  for i, r := range s {
        c := string(r)
        for j, k := range secret {
              d := string(k)
          if i == j && c == d {
            fmt.Println("i= %d j= %d",i,j)
            equis = equis + "X"
          }
          if i != j && c == d {
            rayita = rayita + "_"
          }

          }
    }
  return(equis + rayita)
}
