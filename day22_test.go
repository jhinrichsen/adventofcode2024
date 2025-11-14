package adventofcode2024

import "testing"

func TestDay22SecretEvolution(t *testing.T) {
	// Test the evolution of secret 123
	expected := []uint{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}

	secret := uint(123)
	for i, want := range expected {
		secret = nextSecret(secret)
		if secret != want {
			t.Errorf("iteration %d: want %d, got %d", i+1, want, secret)
		}
	}
}

func TestDay22Part1Example(t *testing.T) {
	tests := []struct {
		initial uint
		want    uint
	}{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := evolveSecret(tt.initial, 2000)
			if got != tt.want {
				t.Errorf("evolveSecret(%d, 2000) = %d, want %d", tt.initial, got, tt.want)
			}
		})
	}

	// Test the full example
	testWithParserNoErr(t, 22, exampleFilename, true, NewDay22, Day22, 37327623)
}

func TestDay22Part1(t *testing.T) {
	testWithParserNoErr(t, 22, filename, true, NewDay22, Day22, 13004408787)
}

func TestDay22Part2Example(t *testing.T) {
	testWithParserNoErr(t, 22, example2Filename, false, NewDay22, Day22, 23)
}

func TestDay22Part2(t *testing.T) {
	testWithParserNoErr(t, 22, filename, false, NewDay22, Day22, 1455)
}

func BenchmarkDay22Part1(b *testing.B) {
	benchWithParserNoErr(b, 22, true, NewDay22, Day22)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchWithParserNoErr(b, 22, false, NewDay22, Day22)
}
