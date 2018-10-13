package sample

import "testing"

func TestNewStore(t *testing.T) {

	tests := []struct {
		name string
		want int
	}{
		{name: "OK", want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			store, err := initStore()
			if err != nil {
				t.Fatal(err)
			}

			want := tt.want
			got := store.Sell()

			if got != want {
				t.Errorf("store.Sell() = %v, want %v", got, want)
			}
		})
	}
}
