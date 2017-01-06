#include <unistd.h>
#include <stdio.h>

/*
 *  ssize_t read(int fd, void *buf, size_t count);
 *  ssize_t write(int fd, const void *buf, size_t count);
*/

int
main(int argc, char* argv[])
{
    char write_buf[32000];
    char* write_ptr = write_buf;
    int hit_nl = 0;

    for (;;) {
        char c;
        const int r = read(0, &c, 1);
        if (r > 0) {
            *write_ptr++ = c;
            if ('\n' == c) { /* finished line */
                hit_nl = 1;
                break;
            }
        } else {
            break;
        }
    }

    if (write_ptr > write_buf) {
        *write_ptr = '\0';
        write(1, write_buf, write_ptr - write_buf);
        return 0;
    } else {
        if (hit_nl) {
            return 0;
        } else {
            return 1;
        }
    }
}


