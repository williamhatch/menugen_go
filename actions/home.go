package actions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// Add uploadPath helper for the template
	c.Set("uploadPath", func() string {
		return "/upload"
	})
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

// UploadHandler handles the file upload and processing
func UploadHandler(c buffalo.Context) error {
	// Get the file from the request
	file, header, err := c.Request().FormFile("menu_image")
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{
			"error": "No file uploaded",
		}))
	}
	defer file.Close()

	// Validate file type
	contentType := header.Header.Get("Content-Type")
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif"}
	isValidType := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{
			"error": "Invalid file type. Please upload PNG, JPG, or GIF files only.",
		}))
	}

	// Validate file size (10MB limit)
	if header.Size > 10*1024*1024 {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{
			"error": "File size too large. Please upload files smaller than 10MB.",
		}))
	}

	// Create uploads directory if it doesn't exist
	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{
			"error": "Failed to create upload directory",
		}))
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), 
		strings.ReplaceAll(header.Filename, ext, ""), ext)
	filePath := filepath.Join(uploadsDir, filename)

	// Save the file
	dst, err := os.Create(filePath)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{
			"error": "Failed to save file",
		}))
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{
			"error": "Failed to save file",
		}))
	}

	// For now, simulate processing and return a success response
	// In a real implementation, you would:
	// 1. Use OCR to extract menu items from the image
	// 2. Generate images for each menu item using AI
	// 3. Store results in database
	// 4. Return results to user

	// Set data for the results template
	c.Set("filename", header.Filename)
	c.Set("filesize", formatFileSize(header.Size))
	c.Set("uploadPath", filePath)
	
	// Simulate some menu items extracted from OCR
	menuItems := []map[string]string{
		{"name": "Caesar Salad", "description": "Fresh romaine lettuce with parmesan cheese", "image": "/assets/images/caesar-salad.jpg"},
		{"name": "Grilled Salmon", "description": "Atlantic salmon with lemon butter sauce", "image": "/assets/images/grilled-salmon.jpg"},
		{"name": "Beef Tenderloin", "description": "Premium cut with roasted vegetables", "image": "/assets/images/beef-tenderloin.jpg"},
		{"name": "Chocolate Cake", "description": "Rich chocolate cake with vanilla ice cream", "image": "/assets/images/chocolate-cake.jpg"},
	}
	
	c.Set("menuItems", menuItems)
	c.Set("processingTime", "2.3 seconds")

	return c.Render(http.StatusOK, r.HTML("upload/results.plush.html"))
}

// formatFileSize converts bytes to human readable format
func formatFileSize(bytes int64) string {
	if bytes == 0 {
		return "0 Bytes"
	}
	
	const unit = 1024
	sizes := []string{"Bytes", "KB", "MB", "GB"}
	
	if bytes < unit {
		return fmt.Sprintf("%d %s", bytes, sizes[0])
	}
	
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), sizes[exp+1])
}
