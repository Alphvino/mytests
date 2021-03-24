#include <stdio.h>
#include <stdlib.h>

int main ()
{
    int i, x, max, nums;
    printf("How many numbers do you want to check? ");
    scanf("%d", &nums);
    for(i = 0; i < nums; i++)
    {
        printf("Enter the number: ");
        scanf("%d", &x);
 
        if (i == 0)
        {
	
            max = x;
        }
 
        if (x > max)
        {
            max = x;
        }
    }
 
    printf("The biggest number is: %d\n", max);
    getchar();
    return 0;
}
