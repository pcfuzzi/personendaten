// Deklaration des Hauptpakets
package main2

//Importfunktion
import (
	// Importieren der fmt-Bibliothek für Ein- und Ausgaben
	"fmt"
    // Importieren des Testpakets
	"testing"
)
// Funktion, die die beiden Argumente addiert und das Ergebnis zurückgibt
func add(a, b int) int { 
	return a + b
}
// Benchmark-Testfunktion für die add-Funktion
func BenchmarkAdd(b *testing.B) { 
	    /* Schleife, die "b.N" Mal ausgeführt wird, wobei "b.N" eine vom
		Testpaket bereitgestellte Variable ist, die angibt, wie oft die Funktion
		ausgeführt werden soll, um eine aussagekräftige Leistungsmessung zu erhalten.
        Schleife, die "b.N" Mal ausgeführt wird, wobei "b.N" eine vom
		Testpaket bereitgestellte Variable ist, die angibt, wie oft die Funktion
		ausgeführt werden soll, um eine aussagekräftige Leistungsmessung zu erhalten.*/
	for i := 0; i < b.N; i++ { 
		// Aufruf der add-Funktion mit dem aktuellen Schleifenindex und dem nächsten Index als Argumenten
		result := add(i, i+1) 
        // Ausgabe des Ergebnisses in der Konsole
		fmt.Println(result)   
	}
}
