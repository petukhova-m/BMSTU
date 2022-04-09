#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Elem { 
        struct Elem *next; 
        char *word; 
}; 

void swap(struct Elem* a, struct Elem* b){
	char* tmp;
        tmp = a->word; 
        a->word = b->word;
        b->word = tmp;
}

void createlist(char* str, struct Elem* tmp) {
        tmp = calloc(1, sizeof(struct Elem));
        tmp->word = calloc(100, sizeof(char));
        strcpy(tmp->word, str);
        tmp->next = NULL;
        //return tmp;
        free(tmp);
        //free(tmp->word);
}

void insertafter(struct Elem* list, char* str) {
        struct Elem* z;
        z = calloc(1, sizeof(struct Elem));
        z->word = calloc(100, sizeof(char));
        strcpy(z->word, str);
        z->next = NULL;
        //createlist(str, &z);
        struct Elem *y;
        y = list;
        while (y->next != NULL) {
                //y = list;
                y = y->next;
        }
        y->next = z;
        free(z);
}

void to_struct(struct Elem* list, char* str){
	int i = 0, j = 0;
	char *buf = calloc(1000, sizeof(char));
	//for (int k = 0; k < 1000; k++)
	//	buf[k] = 0;
	while(str[i] != '\0'){
		//char buf[1000];
		
                if (str[i] != ' ') {
                        buf[j] = str[i];
                        j++;
                }
                if (str[i] == ' ' && j > 0) {
                        insertafter(&list, buf);
                        for (int k = 0; k < j; k++) {
                                buf[k] = 0;  
                        }
                        j = 0;
                }
                i++;
        }
        if (j != 0) {
                insertafter(list, buf);
                for (int k = 0; k < j; k++) {
                        buf[k] = 0;  
                }
                j = 0;                
        }
        free(buf);
}

int compare(struct Elem *a, struct Elem *b) {
        if (strlen(a->word) > strlen(b->word))
                return 1;
        else return 0;
}

void bsort(struct Elem *start) {
        struct Elem* i= start->next; 
        struct Elem* j;
        while(i != NULL){
        	j = start->next;
        	while(j->next != NULL){
                        if (strlen(j->word) > strlen(j->next->word))
                                swap(j, j->next);
                        j = j->next;
                }
                i = i->next;
	}
}

int main() {
        long int i, j, k, n;
        char *str = calloc(10000, sizeof(char));
        //char *buf = calloc(1000, sizeof(char));
        struct Elem* list;
        createlist("", &list);
        fgets(str, 101, stdin);
        to_struct(&list, str);
        bsort(list);
	struct Elem *p = list->next;
        while (p != NULL) {
		printf("%s ", p->word);
		free(p->word);
		//free(list->word);
		//list = list->next;
		p = p->next;
	} 
	free(p);
	free(str);
        free(list->word);
        //free(list);
        return 0;
}
