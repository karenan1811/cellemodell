package cell

import "testing"

func TestInitCell(t *testing.T) {
wanted := false; //Bruker semantikken i Golang konseptuelle modell
state := InitCell();
if state != wanted {
t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)}
}

