Timeline do Go
2007 -> Inicio do processo
	* Viu que o C++ não dava conta, partiu nova linguagem
2009 -> Open-source 
	* Liberou para primeiro uso
2012 -> V1.0
	* Liberou a primeira versão estavel da linguagem
	* Atualmente na 1.22 (Tudo que funcioava lá atrás roda aqui)
2015 -> Melhorias V1.5
	* Aqui o compilado do Go foi feito em Go
2018 -> Consolidação do sistema
	* V1.11
	* Gerenciamento de pacotes
2022 -> Generics
	* V1.18
	* Manter tudo retrocompativel com a V1.0
2024 -> HTTP server
	* Melhorias gerais de HTTP
	* Correção de loop var
	* V1.22

O GO é totalmente retroativo, isso é, GO v1 roda no v1.22 -> Nunca pretende ignorar o legado

Go se tornou uma plataforma
* CLi, web e IDE
* Go security (SBOM, fuzzing, vulnerability management)
	* SBOM -> procura se as libs usadas tem propriedade de terceiros (Importante em projetos grandes)
	* Fuzzy -> Estresse de projeto
* Go modeules (Module proxy, workspaces, checksum DB)
* Go core (Language, compiler, runtime, standard library)

* Linguagem de programação
* testes
* Geração de documentação
* Multiplataforma
* Profilling
* Performance "out-of-box" (Se pode fazer muita cagada, a performance anda é boa)
* Build e deploy
* Go actions

Caracteristicas
* Foco na simplicidade
* Extrema performance
* Sples acesso a recursos e protocolos de rede
* Nasceu otimizada para multi-core
* Baixo uso de memoria
* Compilado para unico binario
* Multiplataforma

Go tem uma base de dados de segurança propria que vare a net e confere se tal lib tem um chabu caso baixado

```
govulcheck <diretorio>
```

Arquitetura runtime
* GO não roda em VM, ele pega o code, joga em runtime e num compilado temporaria

Pq GO é rapido?
Quando uma thread sobe, ele liga uma scheduler M:P:G que realiza algumas rotinas
* M -> Machine
	* Cria um thread de 2MB, que chama o processador lógico e gera uma gorotine de 2KB -> Essi são valores fixos
* P -> processes -> Instancia o M e usa o G
* G -> Global queue

O Processamento do Go funciona com 2 filas, a fila global onde fica todas as operações e a fila de processamento local que é separada por processador lógico, ele pega um processo global e traz para o processamento local para realizar a ação.

O Go tem o proprio gerenciador de memoria -> Mallocgc
tiny -> Objetos menores que 16 bytes
Small -> Objetos entre 16 bytes a 32 KB
Large -> Objetos maiores que 32 KB

Baseado no projeto do google chamado mc malloc

Garbade collector
* Não geracional, concorrente, tri-collor mark e sweep collect
* Extremamente otimizado

Linguagem de uso geral


Futuro do Go
* Não havera 2.0 pq ele é retrocompativel, então nunca vai mudar sua estrutura

GO não tem classe, ele tem embading


```
package main

func main() {
	print("Hello")
}
```

```
GOOS=windows go build -o "hello.exe"
```

Ex de server HTTP:
```
package main

import "net/http"

func main() {
	http.ListenAndServe(":3333", nil)
}
```

Linguagem fortemente tipada, más tmb trabalha com inferencia

Simplificado:
```
a := 1
```

Completa
```
var a int
a = 1
```


Criar um tipo
```
package main

type PASTEL string

func main() {
	var teste PASTEL
	teste = "teste"
	print(teste)
}
```


Estrutura de dados

```
package main

type Teste struct {
	t1 string
	t2 string
	t3 string
}

func main() {
	pastel := Teste{t1:"t1", t2:"t2", t3:"t3"}
	fmt.Println(pastel.t1)
	fmt.Println(pastel.t2)
	fmt.Println(pastel.t3)
}
```

metodo/funcao
Faz referencia a metodo usado
```
package main

type Teste struct {
	t1 string
	t2 string
}

func (t Teste) EchoT1() {
	fmt.Println("T1:" + t.t1)
}

func main() {
	teste := {t1:"t1", t2:"t2"}
	tmt.Println(pastel.t1)

	teste.EchoT1()
}
```

Server HTTP JSON
```
package main

import "net/http"

// Oculta a saida do T2 no retorno do json
type Teste struct {
	t1 string `json:"t1"`
	t2 string `json:"-"`
}

func main() {
	teste := {t1:"t1", t2:"t2"}

	// Funcao anonima
	http.HandleFunc("/", func(w httpResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(teste)
	})

	// Usando funcao
	http.HandleFunc("/", home)

	http.ListenAndServe(":3333", nil)
}

func home(w httpResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(teste)
}

```

Scheduler -> gereciamento de processos multicore
	* Define quem vai ir executar em tal CPU
	* Tipos:
		* Colaborativo/Cooperativo
			* Inicia o processo e deixa ele terminar, fica no CPU até terminar (Sequestra o CPU)
		* Preemptivo
			* Altera entre processo com tempos definidos
			* Deixa mais lento no geral pq tem que ficar switando entre processos
	
Threads
* Instancia de execução de um processo
* Fila de execução das tarefas do processo
* O GO trabalha com light threads, são threads geradas pelo GO e não pelo OS

Go routines -> Gerenciado pelo Go Scheduler
* Gera uma lite thread de 2KB para executar
* Thread gerada pelo runtime do go, não afeta os processos do S.O

Problema comum no GO
* Racing condition -> usar o mesmo valor global de memoria
	* Se uma thread muda, o restante é afetado tmb, tem que tomar cuidado com isso
* Para evitar isso, vc usa chanells para fazer a comunicação entre as threads


Oq acontence se o ambiente tem 1vCPU e rodar o threads GO?
O GO tem uma Scheduler preempitivo, ele altera entre as threads para executar as ações, caso ele tenha um hardware melhor, ele pode separar a rotina por nucleo
