# Atualização de Canonical SKU


O script irá atualizar o campo `seller_custom_field` no item e em todas as variações do mesmo.


### Parâmetros

- `-f` => Caminho do arquivo que contenha a lista de sku e IDs que serão atualizados _(string)_
- `-c` => Canonical SKU _(string)_
- `-i` => ID do produto no Meli _(string)_

> Você deve escolher em atualizar por uma lista ou um item 

- `-t` => Access Token _(string)_

#### Arquivo de dados 

O arquivo de dados deve estar no formato

```console
SKU0000001,MLB00001
SKU0000002,MLB00002
SKU0000003,MLB00003
```

## Rodando o programa compilado

Rode o comando passando as flags.

#### Atualizando uma lista  
```sh
./putcanonical -f <file_path> -t <access_token> 
```

#### Atualizando um item 
```sh
./putcanonical -c <canonical_sku> -i <meli_id> -t <access_token>  
```

## Rodando o programa com o `go run`

#### Atualizando uma lista  
```sh
go run cmd/main.go -f <file_path> -t <access_token> 
```

#### Atualizando um item 
```sh
go run cmd/main.go -c <canonical_sku> -i <meli_id> -t <access_token>  
```

#### Saida

```console
****************************************** *******
*** Start update Meli Items with Canonical SKU ***
****************************************** *******

Items to process: 1

--------------------------------------------------------
Processing: sku: SKU0001
ID: MLB0001 Payload: {"seller_custom_field":"SKU0001","variations":[]}

--------------------------------------------------------
[SUCCESS]
sku: SKU0001
ID: MLB0001


****** Finished *******
```
