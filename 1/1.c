#include <stdio.h>
#include <stdlib.h>

struct elf
{
    int elfId;
    int totalCals[100];
};

int main() {
    FILE    *textfile;
    char    *text;
    long    numbytes;
     
    textfile = fopen("input.txt", "r");
    if(textfile == NULL)
        return 1;
     
    fseek(textfile, 0L, SEEK_END);
    numbytes = ftell(textfile);
    fseek(textfile, 0L, SEEK_SET);  
 
    text = (char*)calloc(numbytes, sizeof(char));   
    if(text == NULL)
        return 1;
 
    fread(text, sizeof(char), numbytes, textfile);
    fclose(textfile);
 
    // assess total amount of elves 
    for (size_t i = 0; i < numbytes; i++)
    {
        /* code */
    }
    
    printf(text);
 
    return 0;
}