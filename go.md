Aprender GO
* Nasceu na Google;
* Linguagem pensada para a era moderna (Multi core)
* Sintaxe simples e limpa;
* Retrocompativel entre suas versões;
* Linguagem compilada;
* Linguagem estaticamente tipada (Variavel inicia int, morre em int);
* Todo o compilado em GO gera um unico arquivo compialdo;


PQ Go?
* Python -> Lento demais
* Java -> Quanto maior o projeto, pior é mante-lo;
* C++ -> Complexo demais;
* GO implementa o conceito de multicore, diferente das linguagens que implementaram durante o tempo;
* Frameworks nativos;
* Simplicidae na escrita;
* Ele é limpo, ele força a somente usar o necessário, se não ele nem compila;
* Ela é retro compativel entre todas as versões, atual está na 1.22 e é possível executar a do modelo 1;

Executa o programa sem compilar:
```
go run <arquivo>
```

Buildar o programa:
```
go build <arquivo>
```

Obter pacotes
```
go get <pacote>
go install <pacote@version>
```

Estrutura de diretorio do GO
bin -> Binarios compilados que a linguagem vai usar ou que o proprio go usa
pwd -> Arquivos pré-compilados para acelerar o processo de compilação
src -> Source code (Pacotes que podem ser utilizados dentro do projeto GO)

Mostra a versão do GO
```
go version
```

Variaveis de ambiente do GO
```
go env
```

GOPATH -> Diretorio principal do GO
GOOS -> Sistema operacioonal que o GO identificou

Build para outra versão de OS:
```
GOOS=windows go build t1.go
```

Go Routines e channels

Concorrencia e paralelismo
* C -> Alterar entre as tarefas
* P -> Conseguir lidar com várias ações ao mesmo tempo sem denegrir o resultado

Uma thread normal:
* Consumo normal de 1MB por thread
* Tem que fazer um system call para requisitar a criação da thread
* Se tiver uma area comum, tem que dar lock na variavel de momento para evitar a interferencia dos processos
* Para alterar entre as thread, existe o scheduler, que dá um tempo de execução para o processo e independe de ele terminar ou não, ele alterar o processo conforme combinado (Custoso)

Go thread
* Runtime proprio (User land)
	* Não tem chamada para SO
	* Scheduler proprio
	* Threads menores que 1MB (2K)
	* Ele possue suas proprias regras de alterar processos, onde tudo que ocupa mais tempo que um processo local fica por ultimo e aqueles processos locais são os primeiros a serem trabalhados (Não exatamente isso, más é de forma proxima)


Iniciar um projeto GO
* go mod init <namespace -> Por padrão coloque o repo do git que vai ter o projeto>
	* Ex: ```go mod init github.com/guilhermeG23/t10```
* Para os MAINs, crie um diretorio CMD e coloque o main lá, se tiver vários endpoints, crie um diretorio para cada endpoint e coloque o main dentro do diretorio
* Crie o diretorio internal na raiz do seu projeto, esse é referente a somente "coisas locais" do seu projeto
	* Crie os subdiretorios necessários sobre o internal;
	* Sempre no singular;
	* O nome do diretorio tem que ser o mesmo do arquivo go
	* Letras minusculas e uso de underline
	* Se for testar as coisas de um arquivo, gere o test com o mesmo nome + _test antes da extensão e torne os valores publicos do arquivo que ele quer usar de base para os testes
* pkg é o diretorio para globais do projeto
* Crie o diretorio build para coisas necessárias a buildar projeto
* configs é o diretorio que tem as configurações globais que são usadas no projeto
* docs é o diretorio para documentações
* api é o diretorio para a geração de documentações de api
* scripts é o diretorio de scripts de automatização
* test é o diretorio de testes gerais
	* Teste gerais é nesse diretorio
	* Testes isolados é dentro do proprio arquivo
* examples é o diretorio de exemplos para amostrar o funcionamento de algo
* web é o diretorio para assets
* website é o diretorio para necessidade de web
* arquivo global de Makefile para ações globais no projeto