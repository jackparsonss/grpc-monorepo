variable "HOME" {
  default = "$HOME"
}

group "production" {
  targets = ["web-server-production", "inventory-production"]
}

target "web-server-production" {
  context    = "./web-server"
  dockerfile = "Dockerfile"
  tags       = ["jackpar/web-server:latest"]
  platforms  = ["linux/amd64"]
  secret     = ["type=file,id=token,src=${HOME}/.netrc"]
  output     = ["type=image,push=true"]
}

target "inventory-production" {
  context    = "./inventory"
  dockerfile = "Dockerfile"
  tags       = ["jackpar/inventory:latest"]
  platforms  = ["linux/amd64"]
  secret     = ["type=file,id=token,src=${HOME}/.netrc"]
  output     = ["type=image,push=true"]
}

group "dev" {
  targets = ["web-server-dev", "inventory-dev"]
}

target "web-server-dev" {
  context    = "./web-server"
  dockerfile = "Dockerfile"
  tags       = ["jackpar/web-server:dev"]
  platforms  = ["linux/arm64"]
  secret     = ["type=file,id=token,src=${HOME}/.netrc"]
}

target "inventory-dev" {
  context    = "./inventory"
  dockerfile = "Dockerfile"
  tags       = ["jackpar/inventory:dev"]
  platforms  = ["linux/arm64"]
  secret     = ["type=file,id=token,src=${HOME}/.netrc"]
}