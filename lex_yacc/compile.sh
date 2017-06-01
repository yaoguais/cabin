lex example.l
yacc -d example.y
cc lex.yy.c y.tab.c -o example
