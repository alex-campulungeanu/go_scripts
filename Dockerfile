FROM golang:1.21.4-bullseye

ARG http_proxy
ARG https_proxy

ENV http_proxy ${http_proxy}
ENV https_proxy ${https_proxy}
ENV GOINSECURE=*
ENV GONOSUMDB=*
ENV GOPRIVATE=*

RUN apt-get update
RUN apt install -y openssl vim
RUN git config --global http.sslverify false
RUN git config --global http.postBuffer 524288000
        # Or set double value  1048576000
# RUN git config --global https.postBuffer
RUN git config --global core.compression -1

RUN \
    echo 'alias ls="ls --color=auto"' >> /root/.bashrc && \
    echo 'PS1="\[\033[1;33m\][\u@\h \W >>>] \$ \[\033[0m\]"' >> /root/.bashrc && \
#fix vim on Debian
    echo 'if filereadable("/usr/share/vim/vim81/defaults.vim") \n \
source /usr/share/vim/vim81/defaults.vim \n \
endif \n \
" now set the line that the defaults file is not reloaded afterwards! \n \
let g:skip_defaults_vim = 1 \n \
" turn of mouse \n \
set mouse= \n \
" other override settings go here ' >> /etc/vim/vimrc.local

# RUN go get -u -v golang.org/x/tools/cmd/gopls 2>&1

