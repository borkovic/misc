#include <unistd.h>
#include <stdio.h>

/*
 *  ssize_t read(int fd, void *buf, size_t count);
 *  ssize_t write(int fd, const void *buf, size_t count);
*/

#if 1

int
main()
{
    unsigned char write_buf[32000];
    unsigned char* write_ptr = write_buf;
    int hit_nl = 0;

    for (;;) {
        unsigned char c;
        const int r = read (0, &c, 1);
        if (r > 0) {
            if ('\n' != c) {
                *write_ptr++ = c;
            } else {
                hit_nl = 1;
                break;
            }
        } else {
            break;
        }
    }

    if (write_ptr > write_buf) {
        *write_ptr++ = '\n';
        *write_ptr = '\0';
        write(1, write_buf, write_ptr - write_buf);
        return ! hit_nl;
    } else {
        return 1;
    }
}

#else

/* From rc shell FAQ */
int
main ()
{
    unsigned char c;
    while (read (0, &c, 1) == 1 && c != '\n') {
        write (1, &c, 1);
    }
    write (1, "\n", 1);
    return c != '\n';
}

#endif

