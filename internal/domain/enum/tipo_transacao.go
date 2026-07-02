package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// TipoTransacao define o tipo da transação: se é receita ou despesa.
type TipoTransacao int

const (
	TipoTransacaoReceita TipoTransacao = 1
	TipoTransacaoDespesa TipoTransacao = 2
)

var tipoTransacaoNames = map[TipoTransacao]string{
	TipoTransacaoReceita: "Receita",
	TipoTransacaoDespesa: "Despesa",
}

var tipoTransacaoValues = map[string]TipoTransacao{
	"Receita": TipoTransacaoReceita,
	"Despesa": TipoTransacaoDespesa,
}

// String retorna o nome legível do enum.
func (tt TipoTransacao) String() string {
	if name, ok := tipoTransacaoNames[tt]; ok {
		return name
	}
	return fmt.Sprintf("TipoTransacao(%d)", int(tt))
}

// IsValid verifica se o valor do enum é válido.
func (tt TipoTransacao) IsValid() bool {
	_, ok := tipoTransacaoNames[tt]
	return ok
}

// MarshalJSON serializa o enum como string no JSON.
func (tt TipoTransacao) MarshalJSON() ([]byte, error) {
	if name, ok := tipoTransacaoNames[tt]; ok {
		return json.Marshal(name)
	}
	return nil, fmt.Errorf("valor inválido de TipoTransacao: %d", int(tt))
}

// UnmarshalJSON desserializa o enum a partir de string ou número.
func (tt *TipoTransacao) UnmarshalJSON(data []byte) error {
	// Tenta como string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if val, ok := tipoTransacaoValues[str]; ok {
			*tt = val
			return nil
		}
		return fmt.Errorf("valor inválido de TipoTransacao: %q", str)
	}

	// Tenta como número
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		val := TipoTransacao(num)
		if val.IsValid() {
			*tt = val
			return nil
		}
		return fmt.Errorf("valor inválido de TipoTransacao: %d", num)
	}

	return fmt.Errorf("formato inválido para TipoTransacao")
}

// Value implementa driver.Valuer para salvar como int no banco.
func (tt TipoTransacao) Value() (driver.Value, error) {
	return int64(tt), nil
}

// Scan implementa sql.Scanner para ler do banco.
func (tt *TipoTransacao) Scan(value interface{}) error {
	if value == nil {
		*tt = 0
		return nil
	}
	switch v := value.(type) {
	case int64:
		*tt = TipoTransacao(v)
	case int:
		*tt = TipoTransacao(v)
	default:
		return fmt.Errorf("não é possível converter %T para TipoTransacao", value)
	}
	return nil
}
