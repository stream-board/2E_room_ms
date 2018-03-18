FROM microsoft/aspnetcore-build:2.0
WORKDIR /rooms_ms

# Copy csproj and restore as distinct layers
COPY *.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o out

# Build runtime image
FROM microsoft/aspnetcore:2.0
WORKDIR /rooms_ms
#COPY ./rooms_ms/out ./
COPY ./rooms_ms/out .
RUN ls
ENTRYPOINT ["dotnet", "rooms_ms.dll"]


