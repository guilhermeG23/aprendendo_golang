//Executar: go run teste.go

//Pacotes
package main

//Libs necessarias
//Ele da erro se tiver uma lib instanciada que não esta em uso

//Importar uma lib
//import "fmt"

//Importar varias libs
import (
	"fmt"
	"strings"
)

//Funcoes
//Funcao principal
func main() {
	
	//Nao precisa de ; no fim das linhas, é mais por custome

	//Print sem quebra de linha
	fmt.Print("Hello World.\n");
	fmt.Print("Teste 1 function main.\n");

	//Print com quebra de linha
	fmt.Println("Quebra de linha")

	//Use variable
	var teste = "t1";
	fmt.Print(teste, "\n");
	fmt.Println(teste);
	//Ele da espaço por automatico
	fmt.Println("Com texto:", teste);

	//Constantes
	const teste1 = "t1";
	fmt.Println("Teste constante:", teste1);

	//Teste
	var t1 = 1;
	var t2 = "teste";
	var t3 = 1.1;

	fmt.Println(t1);
	fmt.Println(t2);
	fmt.Println(t3);

	//Saida dentro de uma string
	fmt.Printf("teste: %v\n",t1);
	fmt.Printf("teste: %v\n",t2);
	fmt.Printf("teste: %v\n",t3);

	//Definindo o tipo das variaveis
	var t10 int = 10;
	var t11 string = "teste";
	fmt.Printf("teste: %v\n",t10);
	fmt.Printf("teste: %v\n",t11);
	//Ou
	var t12 int;
	t12  = 10;
	var t13 string;
	t13 = "teste";
	fmt.Printf("teste: %v\n",t12);
	fmt.Printf("teste: %v\n",t13);

	//Outra forma de definir uma variavel é assim
	t4 := "teste";
	fmt.Printf("t4: %v\n",t4);

	//Entrada
	var testeEntrada string;
	//Pointer -> Endereço da memoria que a variavel esta armazenada
	fmt.Print("Entrada: ");
	fmt.Scan(&testeEntrada);
	//Output do preenchimento do endereço de memoria
	fmt.Printf("Saida: %v", testeEntrada);

	//Array
	//Eles tem que ter tamanho definido
	//E tem que ter o tipo de dado definido
	//Se houver um valor maior que o array, esse valor é dado como erro
	var t15 [50]string;
	t15[0] = "t1";
	t15[1] = "t2";
	t15[2] = "t3";
	fmt.Print("\nArray: ",t15);
	fmt.Print("\nOne element Array: ",t15[0]);
	fmt.Println("");

	//Slice são abstrações de array e são dinamicos
	var t16[] string;
	t16 = append(t16, "teste");
	t16 = append(t16, "teste1");
	fmt.Println(t16);
	fmt.Println(t16[0]);
	fmt.Println(len(t16));
	
	//Loops
	//Só existe o for no GO
	//For true eterno
	for {
		var testeEntrada1 string;
		fmt.Print("Entrada: ");
		fmt.Scan(&testeEntrada1);
		fmt.Printf("Saida: %v", testeEntrada1);
		//Para o for de forma forçada
		break
	}

	fmt.Println("Array dinamico em FOR:");
	for index, tlist := range t16 {
		//strings.Fields -> Faz um  split de um valor que esteja com espaçamento por padrão
		//Ex: "teste1 teste2" -> ['teste1', 'teste2']
		fmt.Println(index)
		fmt.Println(strings.Fields(tlist));
	}

	//Limitar um loop
	var t22 int = 10;
	for t22 > 0 {
		fmt.Println("t22:", t22);
		t22=t22-1;
	}

	//Estrutura de Decisao
	if len(t16) == 0 {
		fmt.Println("Array t16 está vazio");
	} else {
		fmt.Println("Array t16 não está vazio");
	}

	if true {
		fmt.Println("True");
	}

	if false == false {
		fmt.Println("False");
	}

	if false {
		fmt.Println("True");
	} else {
		fmt.Println("False");
	}


	var t20 string = "t1";
	if t20 == "t0" {
		fmt.Println("t0");
	} else if t20 == "t1" {
		fmt.Println("t1");
	} else {
		fmt.Println("t*");
	}

}