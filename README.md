# URL Shortener
A simple and efficient URL shortener built with Go. 🔗✨

---

## 💡 Why URL Shortener?

- Shorten long URLs into compact, shareable links. 🔗
- Fast and lightweight — built with Go for performance. ⚡
- Easy to deploy and use. 🚀

---

## ✨ Features

- Create short URLs with a simple API.
- Redirect to original URLs seamlessly.
- Persistent storage for shortened URLs.
- Scalable and extensible architecture.

---

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/clementmns/go-url-shortener.git  
   cd go-url-shortener  
   ```  

2. **Install dependencies**
   ```bash
   go mod tidy  
   ```  

3. **Run the application**
   ```bash
   go run main.go  
   ```  

4. **Access the app**
    - Open your browser and visit: `http://localhost:9808`

---

## Development Tips 🛠

- Use `go mod tidy` to manage dependencies.
- Add environment variables in a `.env` file (e.g., `PORT`) — document them if needed.
- Use `godotenv` to load `.env` files during development.

---

## Usage Examples 📚

- **Create a short URL**
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:9808/url  
  ```  

- **Redirect using a short URL**
    - Visit: `http://localhost:9808/<shortUrl>`

---

## License & Contact

- Licensed under MIT — see `LICENSE`.

Thanks for using this Go URL Shortener! 🙏  