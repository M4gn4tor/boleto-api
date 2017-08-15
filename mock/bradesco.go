package mock

import gin "gopkg.in/gin-gonic/gin.v1"

func registerBoletoBradesco(c *gin.Context) {

const tok = `
{
    "merchant_id": "90000",
    "meio_pagamento": "800",
    "pedido": {
        "numero": "0-9_A-Z_.MAX-27-CH99",
        "valor": 15000,
        "descricao": "Descritivo do pedido"
    },
    "boleto": {
        "valor_titulo": 15000,
        "data_geracao": "2016-04-22T08:10:43",
        "data_atualizacao": null,
        "linha_digitavel": "23790000255123456789223000000002867240000015000",
        "linha_digitavel_formatada": "23790.00025 51234.567892 23000.000002 8 67240000015000",
        "token": "c3ZtRGVKRDFoUlRESmxRNnhKQnpJalFrb0VueXdVdUxnT2FVMG45cm1qMFMyRDcwRWZ0cFVBS0o0\nMFAxOHY0aTdJK3E1MXVjUVJjNEpBdUxvcE15T1E9PQ==",
        "url_acesso": "http://localhost:9080/boleto/titulo?token=c3ZtRGVKRDFoUlRESmxRNnhKQnpJalFrb0VueXdVdUxnT2FVMG45cm1qMFMyRDcwRWZ0cFVBS0o0\nMFAxOHY0aTdJK3E1MXVjUVJjNEpBdUxvcE15T1E9PQ=="
    },
    "status": {
        "codigo": 0,
        "mensagem": "OPERACAO REALIZADA COM SUCESSO",
        "detalhes": null
    }
}
`
	c.Data(200, "text/json", []byte(tok))
}
