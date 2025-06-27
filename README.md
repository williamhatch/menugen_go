# MenuGen - Turn Menus into Magic ✨🍽️

A beautiful web application built with Go Buffalo that transforms menu images into stunning visual representations using AI-powered OCR and image generation.

Inspired by [Karpathy's MenuGen project](https://karpathy.bearblog.dev/vibe-coding-menugen/), this Go implementation provides a modern, responsive interface for uploading menu images and visualizing each dish.

## ✨ Features

- **Beautiful UI**: Modern gradient design with smooth animations
- **Drag & Drop Upload**: Intuitive file upload with drag and drop support
- **File Validation**: Supports PNG, JPG, GIF files up to 10MB
- **Real-time Processing**: Visual feedback during upload and processing
- **Responsive Design**: Works perfectly on desktop and mobile devices
- **Results Visualization**: Beautiful grid layout showing extracted menu items

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or later
- Node.js and Yarn
- Buffalo CLI

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd menugen
   ```

2. **Install Buffalo CLI (if not already installed)**
   ```bash
   go install github.com/gobuffalo/cli/cmd/buffalo@latest
   ```

3. **Install dependencies**
   ```bash
   make deps
   # or manually:
   go mod tidy
   yarn install
   ```

4. **Start the development server**
   ```bash
   make dev
   # or manually:
   buffalo dev
   ```

5. **Open your browser**
   Visit `http://localhost:3000` to see the application

## 🛠️ Development

### Available Make Commands

```bash
make dev        # Start development server with hot reload
make build      # Build the application
make test       # Run tests
make clean      # Clean generated files
make deps       # Install dependencies
make help       # Show help message
```

### Project Structure

```
menugen/
├── actions/           # Controllers and route handlers
│   ├── app.go        # Main application and routing
│   └── home.go       # Home and upload handlers
├── assets/           # Frontend assets (CSS, JS, images)
├── templates/        # HTML templates
│   ├── home/         # Home page templates
│   └── upload/       # Upload result templates
├── uploads/          # Directory for uploaded files
├── models/           # Database models
├── public/           # Static assets
└── Makefile         # Development commands
```

## 🎨 Features Walkthrough

### 1. Beautiful Landing Page
- Gradient background with animated title
- Clear call-to-action for uploading menus
- Feature highlights with icons

### 2. File Upload Interface
- Drag and drop support
- File type and size validation
- Real-time file information display
- Smooth animations and hover effects

### 3. Processing Feedback
- Loading spinner during processing
- Visual indication of processing state
- Automatic form submission

### 4. Results Display
- Grid layout for menu items
- Upload details summary
- Interactive hover effects
- Action buttons for additional operations

## 🔧 Technical Implementation

### Backend (Go Buffalo)
- **File Upload Handling**: Secure file upload with validation
- **File Processing**: Simulated OCR and AI processing
- **Error Handling**: Comprehensive error responses
- **Routing**: RESTful API design

### Frontend
- **Modern CSS**: Gradients, backdrop filters, animations
- **Responsive Design**: Mobile-first approach
- **JavaScript**: Interactive elements and form handling
- **Accessibility**: Semantic HTML and ARIA labels

### Current Implementation Status

✅ **Completed Features**:
- Beautiful responsive UI matching MenuGen design
- File upload with drag & drop support
- File validation (type, size)
- Upload processing simulation
- Results display with mock data
- Error handling and user feedback

🔄 **Future Enhancements** (not yet implemented):
- Real OCR integration (OpenAI API)
- AI image generation (Replicate API)
- User authentication
- Database storage
- Payment processing
- Background job processing

## 🌐 Deployment

For production deployment:

1. **Build the application**
   ```bash
   make build-prod
   ```

2. **Set environment variables**
   ```bash
   export GO_ENV=production
   export PORT=8080
   ```

3. **Run the binary**
   ```bash
   ./bin/menugen
   ```

## 🎯 API Endpoints

- `GET /` - Home page with upload interface
- `POST /upload` - Handle file upload and processing

## 📝 Development Notes

This implementation focuses on the user interface and file handling aspects of MenuGen. The core AI functionality (OCR and image generation) is simulated with mock data to demonstrate the complete user experience.

To implement the full AI functionality, you would need to:

1. **Add OCR Service**: Integrate OpenAI Vision API for menu text extraction
2. **Add Image Generation**: Integrate Replicate or similar for dish image generation
3. **Add Database**: Store processing results and user data
4. **Add Authentication**: User accounts and session management
5. **Add Payment**: Credit system for processing

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- Inspired by [Andrej Karpathy's MenuGen](https://karpathy.bearblog.dev/vibe-coding-menugen/)
- Built with [Buffalo Framework](https://gobuffalo.io/)
- UI design inspired by modern web aesthetics

---

**Note**: This is a demonstration implementation focusing on UI/UX. For production use, implement proper AI services, authentication, and payment systems as needed.
