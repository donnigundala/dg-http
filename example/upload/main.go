package main

import (
	"fmt"
	"log"
	"os"

	"github.com/donnigundala/dg-http/response"
	"github.com/donnigundala/dg-http/upload"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Single file upload example
	router.POST("/upload", func(c *gin.Context) {
		// Handle upload with validation
		file, err := upload.HandleUpload(c.Request, upload.Config{
			MaxSize:      5 * 1024 * 1024, // 5MB
			AllowedTypes: []string{"image/jpeg", "image/png", "image/gif"},
			FieldName:    "avatar",
		})
		if err != nil {
			response.BadRequest(c, err.Error())
			return
		}

		// Save to local storage
		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, 0755)

		err = os.WriteFile(uploadDir+"/"+file.Filename, file.Data(), 0644)
		if err != nil {
			response.InternalServerError(c, "Failed to save file")
			return
		}

		response.Success(c, gin.H{
			"filename":     file.Filename,
			"size":         file.Size,
			"content_type": file.ContentType,
			"extension":    file.Extension(),
			"is_image":     file.IsImage(),
		}, "File uploaded successfully")
	})

	// Multiple files upload example
	router.POST("/upload-multiple", func(c *gin.Context) {
		// Handle multiple uploads
		files, err := upload.HandleMultipleUploads(c.Request, upload.Config{
			MaxSize:   10 * 1024 * 1024, // 10MB per file
			MaxFiles:  5,
			FieldName: "documents",
		})
		if err != nil {
			response.BadRequest(c, err.Error())
			return
		}

		// Save all files
		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, 0755)

		uploaded := []gin.H{}
		for _, file := range files {
			err = os.WriteFile(uploadDir+"/"+file.Filename, file.Data(), 0644)
			if err != nil {
				response.InternalServerError(c, "Failed to save file: "+file.Filename)
				return
			}

			uploaded = append(uploaded, gin.H{
				"filename":     file.Filename,
				"size":         file.Size,
				"content_type": file.ContentType,
			})
		}

		response.Success(c, gin.H{
			"files": uploaded,
			"count": len(uploaded),
		}, "Files uploaded successfully")
	})

	// Custom validation example
	router.POST("/upload-image", func(c *gin.Context) {
		// Handle upload with custom validators
		file, err := upload.HandleUploadWithValidators(c.Request,
			upload.DefaultConfig(),
			upload.SizeValidator(2*1024*1024), // 2MB
			upload.ImageValidator(),
			upload.ExtensionValidator([]string{".jpg", ".jpeg", ".png"}),
		)
		if err != nil {
			response.BadRequest(c, err.Error())
			return
		}

		// Save file
		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, 0755)
		err = os.WriteFile(uploadDir+"/"+file.Filename, file.Data(), 0644)
		if err != nil {
			response.InternalServerError(c, "Failed to save file")
			return
		}

		response.Success(c, gin.H{
			"filename": file.Filename,
			"size":     file.Size,
		}, "Image uploaded successfully")
	})

	fmt.Println("Upload example server running on :8082")
	fmt.Println("Try:")
	fmt.Println("  curl -F 'avatar=@image.jpg' http://localhost:8082/upload")
	fmt.Println("  curl -F 'documents=@file1.pdf' -F 'documents=@file2.pdf' http://localhost:8082/upload-multiple")
	log.Fatal(router.Run(":8082"))
}
