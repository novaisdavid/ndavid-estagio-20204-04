package encomenda_test

import(
	encomenda "Stock_Acme/encomenda"
	lote      "Stock_Acme/lote"
	"testing"
)

func Verifica(t *testing.T, valorEsperado, valorAtual int){

	if valorEsperado != valorAtual {
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func TestEncomenda(t *testing.T){
	// cliente---Nome
	// o tipo de produto
	// a quantidade que ele quer
	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
				Cliente: "Zafir",
				IdentificadorProduto: "001",
				Quantidade: 200,
	}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

	}

	v := encomend.RetiraEncomenda(encomendar, lotes)

	Verifica(t, v, 2)
}