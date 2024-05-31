# introducing-go 

Repositório dedicado a leitura do Livro [_Introducing Go_](https://www.amazon.com.br/Introducing-Go-Caleb-Doxsey/dp/1491941952) do autor Caleb Doxsey.
O autor aborda os principais recursos da linguagem e também tem os exercícios de cada capítulo para ajuda-lo a praticar o que aprendeu.

## Capítulo 1 - Introdução

### Package main
> É uma declaração de pacote e todo programa Go deve começar com ela.
> Os pacotes são a maneira do Go organizar e reutilizar os códigos.
> Existem dois tipos de programas em Go: Executáveis e bibliotecas. Executáveis são is tips de programas que podemos executar diretamente do terminal e Bibliotecas são coleções de código que empacotamos para que possamos usá-los em outros programas.

### Import "fmt"
> O import é como incluímos código de outros pacotes para usar com o nosso programa.
> O "fmt" (abreviação de formato) implementa a formatação de entrada e saída.
> Para comentários -> // e /**/.
> Funções são os blocos de construção de um programa Go. Eles têm entrada, saídas e uma série de etapas chamadas de instruções que são executadas em ordem. Todas as instruções devem começar com "func" e em seguida com o nome da função (principal, neste caso), uma lista de zero ou mais parâmetros entre parênteses, um tipo de retorno opcional e um corpo cercado por chaves.

## Capítulo 2 - Tipos
### Números
> - Inteiros: São números sem componentes decimais. (...,0,1, 2, ...). Existem também três tipos inteiros dependentes da 
máquina: uint, int e uintptr. Eles dependem da máquina porque seu tamanho depende 
do tipo de arquitetura que você está usando.
>
> - Floats: Contém um componente decimal.
> 
> - String: Sequência de caracteres com comprimeito definido usadas para representar texto. As strings Go são compostas de bytes individuais, geralmente um para cada caractere.
>
> - Um espaço também é considerado um caractere.
>
> - Strings são indexadas começando em 0, não em 1.
>
> - A concatenação utiliza o mesmo símbolo da adição. Ex: Se ambos os lados forem strings, então o compiçador assume que você quer dizer concatenação e não adição.
>
> - O comprimento de uma string pode ser encontrado usando a função len (len("a string")).
>
> - Booleanos: Um valor booleano (em homenagem a George Boole) é um tipo inteiro especial de 1 bit usado para representar verdadeiro e falso. Existem três operadores lógicos principais: AND (&&), OR (||) e NOT (!).
>
> - Frequentemente os booleanos são usados para tomada de decisão.

## Capítulo 3 - Variáveis
>
> - Uma variável é um local de armazenamento, como um tipo específico e um nome associado.
>
> - O compilador não se importa com o nome que você da a uma variável,mas é importante escolher um nome que descreve claramente o propósito da variável.
>
> - camelCase é um estilo para escrever palavras compostas em que a primeira letra de cada nova palavra ou frase é maiúscula.
>
> - Duas formas de criar variáveis: var x = 5 ou x := 5.
>
> - Escopo é o intervalo de lugares em que uma variável pode ser usada. Ele é determinado lexicamente usando blocos.
>
> Constantes:
> - São variáveis cujo os valores não podem ser alterados posteriormente. Elas criadas da mesma forma que as variáveis, mas em vez de usar a palavra-chave var, usamos a palavra-chave const.

## Capítulo 4 - Estrutura de Controle
>
> - A instrução for nos permite repetir uma lista de instruções (um bloco) várias vezes.
>
> - Go só tem uma maneira de escrever loop, mas essa pode ser usada de várias maneiras.
>
> - Uma instrução if é semelhante a uma instrução for porque possui uma condição seguida por um bloco. As 
instruções if também têm uma parte else opcional. Se a condição for avaliada como verdadeira, o bloco após 
a condição será executado; caso contrário, o bloco será ignorado ou, se o bloco else estiver presente, esse 
bloco será executado.
>
> - 

