#! /bin/sh

## Variables set by crontab: SHELL, LOGNAME, HOME

PATH="$HOME/bin:$PATH"
## echo $PATH

LOG=$HOME/daily.log

#########################################################
rc  $HOME/bin/daily.rc  >$LOG  2>&1

#########################################################
if test -f $LOG
then
    if test -s $LOG
    then   ## non-zero log
        cat $LOG            ## send to cron and cron sends to mail
    else
        rm $LOG             ## Remove empty log.
    fi
fi

