#include <stdio.h>

#define nullptr 0

typedef struct List {
    int m_Val;
    struct List* m_Next;
} List;


List* sortList(List* head, int length, int dep)
{
    if (length > 1) {
        int i;
        List *secondHead, *f, *s, *newHead, *newTail;
        int firstHalfLen = (length + 1) / 2;

        /* Split lists in half */
        f = head;
        for (i = 1; i < firstHalfLen; ++i) {
            f = f->m_Next;
        }
        secondHead = f->m_Next;
        printf("%d: [%d,%d]  %d\n", dep, head->m_Val, f->m_Val, secondHead->m_Val);
        f->m_Next = nullptr;

        /* Sort both lists */
        f = sortList(head, firstHalfLen, dep+1);
        printf("%d: f %d\n", dep, f->m_Val);
        s = sortList(secondHead, length - firstHalfLen, dep+1);
        printf("%d: s %d\n", dep, s->m_Val);


        /* Go through both lists and merge */
        if (f->m_Val <= s->m_Val) {
            newHead = newTail = f;
            f = f->m_Next;
        } else {
            newHead = newTail = s;
            s = s->m_Next;
        }

        while (f && s) {
            if (f->m_Val <= s->m_Val) {
                newTail->m_Next = f;
                newTail = f;
                f = f->m_Next;
            } else {
                newTail->m_Next = s;
                newTail = s;
                s = s->m_Next;
            }
        }

        if (f) {
            newTail->m_Next = f;
        } else if (s) {
            newTail->m_Next = s;
        } else {
            newTail->m_Next = nullptr;
        }
        return newHead;
    } else {
        return head;
    }
}

int main()
{
    enum { SIZE = 7 };
    List *newHead, *p;
    List listArr[SIZE];

    for (int i = 0; i < SIZE-1; ++i) {
        listArr[i].m_Next = &listArr[i+1];
        listArr[i].m_Val  = 2*SIZE - i;
    }
    listArr[SIZE-1].m_Next = nullptr;
    listArr[SIZE-1].m_Val  = 2*SIZE - (SIZE-1);
    for (p = &listArr[0]; p; p = p->m_Next) {
        printf("%d  ", p->m_Val);
    }
    printf("\n");

    newHead = sortList(&listArr[0], SIZE, 0);

    for (p = newHead; p; p = p->m_Next) {
        printf("%d  ", p->m_Val);
    }
    printf("\n");
    return 0;
}

