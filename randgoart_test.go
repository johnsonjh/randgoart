package randgoart_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	u "github.com/johnsonjh/leaktestfe"
	"github.com/johnsonjh/randgoart"
)

func TestGenerate(
	t *testing.T,
) {
	u.Leakplug(t)
	tests := []struct {
		input  string
		output string
	}{
		{
			"f5f8a4b7ac83b0b9aa9ad4306269e738",
			"" +
				"+-----------------+\n" +
				"|                 |\n" +
				"|                 |\n" +
				"|          .      |\n" +
				"|  .      . o     |\n" +
				"|.* .    S . o    |\n" +
				"|+ B   .    +     |\n" +
				"| E o   + .. o    |\n" +
				"|...   o . .o .   |\n" +
				"|o.......  .oo    |\n" +
				"+-----------------+",
		},
		{
			"5a46a29f198bdb5a988ccef224acafbf",
			"" +
				"+-----------------+\n" +
				"|                 |\n" +
				"|                 |\n" +
				"|      . .        |\n" +
				"|     . o         |\n" +
				"|    . . S        |\n" +
				"|.  o = O         |\n" +
				"|..o = B          |\n" +
				"|o=   +           |\n" +
				"|o*E.o..          |\n" +
				"+-----------------+",
		},
		{
			"fc94b0c1e5b0987c5843997697ee9fb7",
			"" +
				"+-----------------+\n" +
				"|       .=o.  .   |\n" +
				"|     . *+*. o    |\n" +
				"|      =.*..o     |\n" +
				"|       o + ..    |\n" +
				"|        S o.     |\n" +
				"|         o  .    |\n" +
				"|          .  . . |\n" +
				"|              o .|\n" +
				"|               E.|\n" +
				"+-----------------+",
		},
	}

	for _, tt := range tests {
		data, _ := hex.DecodeString(
			tt.input,
		)
		var err error
		sshhash := randgoart.NewSSH()
		_, err = sshhash.Write(
			data,
		)
		if err != nil {
			t.Fatal(
				fmt.Sprintf(
					"\nrandgoart.sshhash.Write failed: %v",
					err,
				),
			)
		}
		s := sshhash.String()
		if s != tt.output {
			t.Errorf(
				"ssh visual hash test failed: got:\n%v\nwanted:\n%v",
				s,
				tt.output,
			)
		}
	}
}
