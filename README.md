# 🔒 password-generator
Um gerador de senhas seguro e personalizável, criado com a linguagem Go. Este projeto permite criar senhas fortes baseadas em diferentes critérios e salvar os resultados em arquivos para uso posterior. É ideal para quem busca uma solução prática e eficiente para proteger suas contas ou sistemas.

## Project Structure

- **`main.go`**: Arquivo principal que inicializa o servidor e configura o roteamento da API.
- **`app.go`**: Arquivo que contém a lógica principal do programa
- **`generator.go`**: Arquivo que gera as senhas 
- **`writefile`**: Arquivo que gera os arquivosv .txt

## Technologies used:
- Go 
    
## 🛠️ Funcionalidades
Comprimento é personalizável: Apenas escolha o número de caracteres para a sua senha.
Tipos de senha variam entre:
Apenas letras.
Apenas números.
Apenas símbolos especiais.
Combinação de letras, números e símbolos especiais.

Geração de arquivos: Salve as senhas geradas em um arquivo de texto.
Interface interativa: O programa guia o usuário com prompts no terminal.

Como Usar: 

### Installing dependencies:
Após a clonagem, entre no diretório do projeto e instale as dependências necessárias.
Instale as dependências do Go: Se você não tiver o Go instalado
Baixe as dependências do projeto: Dentro do diretório do projeto, execute:

- [Go](https://go.dev/doc/install) 1.20

Rodando o projeto:

Clone este repositório:
````
git clone https://github.com/seu-usuario/password-generator.git
````
Navegue até o diretório do projeto:
````
cd password-generator
````

Execute o programa:
````
go run cmd/main.go
````

Funcionalidade Interativa
- Informe o comprimento desejado para a senha.<br>
- Escolha entre os tipos de caracteres para compor a senha.<br>
- Receba a senha gerada no terminal.<br>
- O arquivo com a senha será salvo automaticamente em output/fileData.txt.<br>
<br>
<br>
🛡️ Segurança
Este projeto é voltado para fins educacionais. Apesar de criar senhas fortes, não é recomendado para uso em aplicações críticas sem melhorias adicionais de segurança.
