FROM alpine
ADD dict-srv /dict-srv
ENTRYPOINT [ "/dict-srv" ]
