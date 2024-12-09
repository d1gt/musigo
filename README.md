# **Musigo - CLI Music Streaming App (Under Development)**  

Musigo is a command-line interface (CLI) application written in **Go**, designed for music lovers who prefer a terminal-based experience. While still in its early development stage, Musigo currently supports **search suggestions** and **YouTube search**.

---

## **Features (Current & Planned)**  

### ✅ **Currently Working:**  
- 🔍 **Search Suggestions** - Get autocomplete suggestions for songs and artists.  
- 📺 **YouTube Search** - Search for music on YouTube (streaming unavailable).  

### 🚧 **Planned Features:**  
- 🎵 **Music Streaming** - Stream tracks from YouTube, Spotify, and other services.  
- 📂 **Offline Caching** - Save songs locally for offline playback.  
- 🎶 **Playlist Management** - Create and manage playlists.  
- 🌐 **Multi-Platform Support** - Available on Linux, macOS, and Windows.

---

## **Current Limitations**  
- **YouTube API Issues:** Due to recent API changes, **music streaming is unavailable**.  
- **Spotify Integration:** Integration with Spotify’s API is **pending**.  
- **Offline Caching:** Not implemented yet.  

---

## **Installation**  

**Prerequisites:**  
- Go 1.22+ installed  
- Make installed  

### **Steps:**  
```bash
# Clone the repository
git clone https://github.com/d1gt/musigo.git

# Navigate to the project directory
cd musigo

# Build the app
make build

# Run the app
./musigo
```

---

## **Usage**  
Simply run the Musigo binary with no arguments:

```bash
./musigo
```

This will start the app and you can interact with it directly in the terminal.

---

## **Roadmap**  
- [ ] Fix YouTube streaming issues  
- [ ] Add Spotify API integration  
- [ ] Implement offline caching  
- [ ] Add full playlist management  
- [ ] Enhance CLI interface with better output formatting  

---

## **Contributing**  
Contributions are welcome! Feel free to open issues or submit pull requests.

---

**Stay Tuned!** 🚀  
