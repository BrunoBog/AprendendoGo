func Test_divideInteiros(t *testing.T) {
    
	type args struct {
		dividendo int
		divisor   int
	}
	
	type expected struct {
		quociente int
		resto     int
		err       error
	}
	
	tests := []struct {
		name string
		args args
		want expected
	}{
		// Casos de teste para a função
		{
			name: "divide por zero",
			want: expected{
				err: errDivisaoInvalida,
			},
			args: args{
				dividendo: 10,
				divisor:   0,
			},
		},
		{
			name: "divisão sem resto",
			want: expected{
				err:       nil,
				resto:     0,
				quociente: 5,
			},
			args: args{
				dividendo: 10,
				divisor:   2,
			},
		},
		{
			name: "divisão com resto",
			want: expected{
				err:       nil,
				resto:     1,
				quociente: 3,
			},
			args: args{
				dividendo: 7,
				divisor:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuociente, gotResto, err := divideInteiros(tt.args.dividendo, tt.args.divisor)
			if err != tt.want.err {
				t.Errorf("divideInteiros() error = %v, wantErr %v", err, tt.want.err)
				return
			}
			if gotQuociente != tt.want.quociente {
				t.Errorf("divideInteiros() gotQuociente = %v, want %v", gotQuociente, tt.want.quociente)
			}
			if gotResto != tt.want.resto {
				t.Errorf("divideInteiros() gotResto = %v, want %v", gotResto, tt.want.resto)
			}
		})
	}
}