# teste_zarv

## Estrutura do Projeto

O projeto está organizado da seguinte forma:

```
teste_zarv/
├── main.go                # Ponto de entrada da aplicação
├── main_test.go           # Testes do package main
├── manhattan/
│   ├── manhattan.go       # Lógica de cálculo da distância de Manhattan
│   └── manhattan_test.go  # Testes do package manhattan
```

## Lógica do Projeto

O objetivo do projeto é calcular a **distância de Manhattan** entre dois pontos em uma matriz.  
A matriz deve conter exatamente dois pontos com valor `1`, e os demais valores devem ser `0`.

A distância de Manhattan entre dois pontos `(x1, y1)` e `(x2, y2)` é calculada por:

```
|x1 - x2| + |y1 - y2|
```

No projeto, a matriz é percorrida para encontrar as coordenadas dos dois pontos com valor `1`.  
Depois, a distância de Manhattan é calculada usando a fórmula acima.

## Como Executar

Para rodar a aplicação principal:

```bash
go run main.go
```

## Como Rodar os Testes

Para rodar todos os testes do projeto (incluindo os dos subdiretórios):

```bash
go test ./...
```

Para rodar os testes com mais detalhes (verbose):

```bash
go test -v ./...
```

Para ver a cobertura dos testes:

```bash
go test -cover ./...
```
