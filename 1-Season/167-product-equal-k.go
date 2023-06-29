package main

/*
//enter n, k and m: 7
//2
//27
//
//enter 7 numbers: 9
//1
//1
//27
//3
//27
//3
//
//9 * 3 = 27

#include <stdio.h>
#include <stdlib.h>

int found(int n, const int set[], int flags[], int i, int k, int m) {
    if (k <= 0)
        return m == 1;
    if (i >= n)
        return 0;
    if (set[i] == 0) {  // must special case 0
        if (m == 0) {
            if (i + k <= n) {
                while (k-- > 0)
                    flags[i++] = 1;
                return 1;
            }
            return 0;
        }
    } else if (m % set[i] == 0) {
        if (found(n, set, flags, i + 1, k - 1, m / set[i])) {
            flags[i] = 1;
            return 1;
        }
    }
    return found(n, set, flags, i + 1, k, m);
}

// get a number from the command line or from stdin
int getnum(int *index, char *argv[]) {
    int value;
    if (argv[*index]) {
        value = strtol(argv[(*index)++], NULL, 0);
        printf("%d ", value);
        return value;
    }
    if (scanf("%d", &value) != 1) {
        printf("invalid input\n");
        exit(1);
    }
    return value;
}

int main(int argc, char *argv[]) {
    int n, m, k, arg_i = 1;

    printf("enter n, k and m: ");
    n = getnum(&arg_i, argv);
    k = getnum(&arg_i, argv);
    m = getnum(&arg_i, argv);
    printf("\n");
    int set[n];
    int flags[n];
    printf("enter %d numbers: ", n);
    for (int i = 0; i < n; i++) {
        set[i] = getnum(&arg_i, argv);
        flags[i] = 0;
    }
    printf("\n");
    if (found(n, set, flags, 0, k, m)) {
        for (int i = 0, j = k; i < n; i++) {
            if (flags[i])
                printf("%d %c ", set[i], --j > 0 ? '*' : '=');
        }
        printf("%d\n", m);
    } else {
        printf("no solution\n");
    }
    return 0;
}
*/
