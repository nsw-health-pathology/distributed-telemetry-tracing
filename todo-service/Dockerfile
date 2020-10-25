##################################################################

FROM mcr.microsoft.com/dotnet/core/sdk:2.1 AS builder
WORKDIR /app

# Copy csproj and restore as distinct layers
COPY *.csproj .
# RUN dotnet restore

# Copy everything else and build
COPY . .
RUN dotnet publish -c Release -o out

##################################################################

# Build runtime image
FROM mcr.microsoft.com/dotnet/core/aspnet:2.1

ENV PORT 5000
WORKDIR /app

COPY --from=builder /app/out .

ENTRYPOINT ["dotnet", "todo-service.dll"]
