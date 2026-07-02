package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// CategoriaFinalidade define o contexto da categoria: se é receita, despesa ou ambas.
type CategoriaFinalidade int

const (
	CategoriaFinalidadeReceita CategoriaFinalidade = 1
	CategoriaFinalidadeDespesa CategoriaFinalidade = 2
	CategoriaFinalidadeAmbas   CategoriaFinalidade = 3
)

var categoriaFinalidadeNames = map[CategoriaFinalidade]string{
	CategoriaFinalidadeReceita: "Receita",
	CategoriaFinalidadeDespesa: "Despesa",
	CategoriaFinalidadeAmbas:   "Ambas",
}

var categoriaFinalidadeValues = map[string]CategoriaFinalidade{
	"Receita": CategoriaFinalidadeReceita,
	"Despesa": CategoriaFinalidadeDespesa,
	"Ambas":   CategoriaFinalidadeAmbas,
}

// String retorna o nome legível do enum.
func (cf CategoriaFinalidade) String() string {
	if name, ok := categoriaFinalidadeNames[cf]; ok {
		return name
	}
	return fmt.Sprintf("CategoriaFinalidade(%d)", int(cf))
}

// IsValid verifica se o valor do enum é válido.
func (cf CategoriaFinalidade) IsValid() bool {
	_, ok := categoriaFinalidadeNames[cf]
	return ok
}

// MarshalJSON serializa o enum como string no JSON (equivalente ao JsonStringEnumConverter do C#).
func (cf CategoriaFinalidade) MarshalJSON() ([]byte, error) {
	if name, ok := categoriaFinalidadeNames[cf]; ok {
		return json.Marshal(name)
	}
	return nil, fmt.Errorf("valor inválido de CategoriaFinalidade: %d", int(cf))
}

// UnmarshalJSON desserializa o enum a partir de string ou número.
func (cf *CategoriaFinalidade) UnmarshalJSON(data []byte) error {
	// Tenta como string primeiro
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if val, ok := categoriaFinalidadeValues[str]; ok {
			*cf = val
			return nil
		}
		return fmt.Errorf("valor inválido de CategoriaFinalidade: %q", str)
	}

	// Tenta como número
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		val := CategoriaFinalidade(num)
		if val.IsValid() {
			*cf = val
			return nil
		}
		return fmt.Errorf("valor inválido de CategoriaFinalidade: %d", num)
	}

	return fmt.Errorf("formato inválido para CategoriaFinalidade")
}

// Value implementa driver.Valuer para salvar como int no banco.
func (cf CategoriaFinalidade) Value() (driver.Value, error) {
	return int64(cf), nil
}

// Scan implementa sql.Scanner para ler do banco.
func (cf *CategoriaFinalidade) Scan(value interface{}) error {
	if value == nil {
		*cf = 0
		return nil
	}
	switch v := value.(type) {
	case int64:
		*cf = CategoriaFinalidade(v)
	case int:
		*cf = CategoriaFinalidade(v)
	default:
		return fmt.Errorf("não é possível converter %T para CategoriaFinalidade", value)
	}
	return nil
}

// ParseCategoriaFinalidade converte uma string para CategoriaFinalidade.
func ParseCategoriaFinalidade(s string) (CategoriaFinalidade, error) {
	if val, ok := categoriaFinalidadeValues[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("valor inválido de CategoriaFinalidade: %q", s)
}
