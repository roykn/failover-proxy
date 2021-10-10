# Failover Proxy

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

## About the project
This project is the playground for a real world use case i had lately. On the client (browser) i need a response even if the server (hosted web server) is down or unavailable for some reason.

So I developed a failover proxy that can be installed on the client. The proxy handles server unavailablity and send back a copy response of the last successfull client request.

## Usage
In order o try it out yourself you have to follow the following steps:
1. Start the server
```bash
go run ./cmd/server/main.go
Starting server on :8080
```

2. Start the proxy
```bash
go run ./cmd/proxy/main.go
Starting proxy on :8081
```

3. Open your browser and navigate to [http://localhost:8081](http://localhost:8081)
you should see the servers response.

4. Stop the server with Ctrl+C

5. Again, navigate to [http://localhost:8081](http://localhost:8081) in your browser

Even though the server is down you will get a response.

## Contributing
Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.