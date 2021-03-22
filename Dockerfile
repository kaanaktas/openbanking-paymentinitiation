FROM scratch

COPY ./certs ./certs
COPY ./.env ./

# Copy the binary file
COPY ./payment ./
COPY ./callback ./

EXPOSE 8080 8081

#services can be run with following commands
#docker run -p 8080:8080 kaktas/openbanking-paymentinitiation ./payment
#docker run -p 8081:8081 kaktas/openbanking-paymentinitiation ./callback