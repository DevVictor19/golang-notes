# Concorrência vs Paralelismo

Go é uma linguagem concorrente

Paralelismo - executar código simultaneamente em processadores físicos diferentes.

Concorrência - intercalar vários processos ao mesmo tempo (pode ocorrer no mesmo processador físico)

Processo - execução de um programa, um processo pode ter várias Threads

Thread - sequência de instruções que a CPU deve executar

Cada processador pode executar 1 thread por vez (ao menos que possua uma tecnologia que permite a execução de mais de uma thread por vez - SMT)

Concorrência viabiliza paralelismo

É possível que a concorrência seja melhor que o paralelismo!
Paralelismo exige muito mais do SO e do hardware

Concorrência - é uma forma também de estruturar seu programa.
Composição de processos (tipicamente funções) que executam de forma idependente, assim não é preciso fazer de forma sequêncial
