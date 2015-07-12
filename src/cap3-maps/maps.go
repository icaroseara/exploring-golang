package main

import (
    "os"
    "fmt"
    "strings"
)

func colherEstatisticas(palavras []string) map[string]int {
    estatisticas := make(map[string]int)
    
    for _, p := range palavras {
        inicial := strings.ToUpper(string(p[0]))
        contador, encontrado := estatisticas[inicial]
        
        if encontrado {
            estatisticas[inicial] = contador + 1
        } else {
             estatisticas[inicial] = 1
        }
    }
    
    return estatisticas
}

func imprimir(estatisticas map[string]int) {
    fmt.Printf("Contagem de palavras por letra")
    
    for inicial, contador := range estatisticas {
        fmt.Printf("%s = %d\n", inicial, contador)
    }
    
}

func main() {
    palavras := os.Args[1:]
    
    estatisticas := colherEstatisticas(palavras)
    
    imprimir(estatisticas)
}