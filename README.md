# Cellemodell (fra konseptuel modell på norsk til konseptuell på Golang)
Simulering av en digital celle.

Funksjoner:
-Lage en celle (og eventuelt initalisere den)
-sette en verdi i celle (tillate verdier 0 eller 1)
- avlese den verdien som cellenn har (tilstand)


#lage en celle

Hvordan allokerer vi minne for nye typer i vårt programmingsspråk?
Kan vurdere å bruke bool for representere cellen.
https://go.dev/ref/spec#Types
TestInitCell

#Stte en verdi
TestSetCellValue

#Avlese verdien i cellen
TestViewCellStatus
TestGetStatus