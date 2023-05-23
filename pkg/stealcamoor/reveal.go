package stealcamoor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"

	mail "github.com/0xmichalis/stealcamoor/pkg/client/email"
)

func (sc *Stealcamoor) reveal(creator common.Address, id uint64) {
	url, err := sc.apiClient.RevealMemoryWithRetries(id, sc.ourAddress, sc.ourSignature, 30)
	if err != nil {
		log.Printf("Cannot reveal memory %d: %v", id, err)
		return
	}

	// Download image
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot download image: %v", err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Cannot save image: %v", err)
		return
	}

	// Read the first 512 bytes of the file to try to detect the content type
	buffer := make([]byte, 512)
	if _, err = bytes.NewReader(content).Read(buffer); err != nil {
		log.Printf("Cannot read body: %v", err)
		return
	}
	contentType := http.DetectContentType(buffer)
	log.Println("Detected content type: " + contentType)
	fileExtension, err := mime.ExtensionsByType(contentType)
	if err != nil {
		log.Printf("Cannot get file extension: %v", err)
		return
	}

	// Pack as an email attachment
	filename, msg := sc.getFilenameAndMessage(creator, id, fileExtension[0])
	attachment := mail.Attachment{
		Name:        filename,
		ContentType: contentType,
		Content:     content,
	}

	// Send email
	log.Printf("Sending email for memory %d...", id)
	if err := sc.emailClient.Send(msg, attachment); err != nil {
		log.Printf("Cannot send email: %v", err)
	}

	// Email content may be corrupted, write content in a local file too
	if err := os.WriteFile(filepath.Join(sc.backupDir, filename), content, 0644); err != nil {
		log.Printf("Cannot write file: %v", err)
	}
}

func (sc *Stealcamoor) getFilenameAndMessage(creator common.Address, id uint64, fileExtension string) (string, string) {
	creatorString := creator.String()
	username := sc.addressToTwitter[creatorString]
	usernameHtml := creatorString

	if username != "" {
		usernameHtml = fmt.Sprintf(`<a href="https://twitter.com/%s">@%s</a>`, username, username)
	} else {
		// Creator has not set up a Twitter account
		// Fallback to their address
		username = creatorString
	}

	filename := username + "_" + fmt.Sprintf("%d", id) + fileExtension
	msg := fmt.Sprintf(`Just revealed memory <a href="https://www.stealcam.com/memories/%d">%d</a> for %s`, id, id, usernameHtml)

	return filename, msg
}
