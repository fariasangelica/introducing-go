# introducing-go 

Repositório dedicado a leitura do Livro [_Introducing Go_](https://www.amazon.com.br/Introducing-Go-Caleb-Doxsey/dp/1491941952) do autor Caleb Doxsey.
O autor aborda os principais recursos da linguagem e também tem os exercícios de cada capítulo para ajuda-lo a praticar o que aprendeu.

## Capítulo 1 - Introdução

### Package main
- É uma declaração de pacote e todo programa Go deve começar com ela.
- Os pacotes são a maneira do Go organizar e reutilizar os códigos.
- Existem dois tipos de programas em Go: Executáveis e bibliotecas. Executáveis são is tips de programas que podemos executar diretamente do terminal e Bibliotecas são coleções de código que empacotamos para que possamos usá-los em outros programas.

### Import "fmt"
- O import é como incluímos código de outros pacotes para usar com o nosso programa.
- O "fmt" (abreviação de formato) implementa a formatação de entrada e saída.
- Para comentários -> // e /**/.
- Funções são os blocos de construção de um programa Go. Eles têm entrada, saídas e uma série de etapas chamadas de instruções que são executadas em ordem. Todas as instruções devem começar com "func" e em seguida com o nome da função (principal, neste caso), uma lista de zero ou mais parâmetros entre parênteses, um tipo de retorno opcional e um corpo cercado por chaves.

## Capítulo 2 - Tipos
### Números
- Inteiros: São números sem componentes decimais. (...,0,1, 2, ...). Existem também três tipos inteiros dependentes da 
máquina: uint, int e uintptr. Eles dependem da máquina porque seu tamanho depende 
do tipo de arquitetura que você está usando.
-
- Floats: Contém um componente decimal.
- 
- String: Sequência de caracteres com comprimeito definido usadas para representar texto. As strings Go são compostas de bytes individuais, geralmente um para cada caractere.
- Um espaço também é considerado um caractere.
- Strings são indexadas começando em 0, não em 1.
- 