1. Get the repository as a url request (needs api)
2. Clone the latest code source | version tag (needs api)
3. Build it and upload the binary (needs container)
4. Curl to get the binary (from the container)
5. install on the system (proxy request to the container)

## POC

Container the entire thing,
api runs in the container
creates the binary locally
uploads to am exposed minio server
curl script gets the file
installs on whatever
