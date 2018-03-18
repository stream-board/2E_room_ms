FROM microsoft/dotnet:2.0-sdk
WORKDIR /rooms_ms

# Copy csproj and restore as distinct layers
COPY *.csproj /rooms_ms
RUN dotnet restore

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o out

#esto ya estaba
ENTRYPOINT ["dotnet", "out/rooms_ms.dll"]
