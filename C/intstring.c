#include <stdio.h>
#include <ctype.h>
#include <string.h>

int main()
{
    char inp[100];
    printf("Enter something: ");
    scanf("%s", &inp);
    for (int i = 0; inp[i] !=0; i++) 
	{
   		if (isdigit(inp[i]))
		{
   		printf("%c is an int\n", inp[i]);
		}
		else
		{
   		printf("%c is a char\n", inp[i]);
   		}
	}
}
