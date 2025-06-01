package cmd

import "os"

var rootVersion = os.Getenv("APP_VERSION")

const (
	rootCmdUse   = "apg"
	rootCmdShort = "CLI para o curso Aprenda Go"
	rootCmdLong  = "Este projeto tem como objetivo criar um CLI para facilitar a navegação e o acesso ao conteúdo do curso Aprenda Go. O CLI foi desenvolvido em Go e utiliza as bibliotecas padrão da linguagem para criar um menu interativo que permite ao usuário acessar os tópicos do curso."
)
