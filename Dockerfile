FROM microsoft/aspnetcore-build:2.0
RUN mkdir /rooms_ms
WORKDIR /rooms_ms
COPY . /rooms_ms
RUN dotnet publish -c Release -o /rooms_ms/out

FROM microsoft/aspnetcore:2.0
WORKDIR /rooms_ms
###COPY /rooms_ms/out ./
RUN ls
#COPY /rooms_ms/out .
RUN dotnet run 

