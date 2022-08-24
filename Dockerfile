FROM ubuntu
COPY ./case ./case
ENTRYPOINT [ "./case" ]