package main

//
//import (
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/x509"
//	"crypto/x509/pkix"
//	"encoding/pem"
//	"log"
//	"math/big"
//	"net/http"
//	"os"
//	"time"
//
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	// 生成RSA密钥对
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		log.Fatalf("Failed to generate RSA key pair: %v", err)
//	}
//
//	// 生成证书原始内容
//	template := x509.Certificate{
//		SerialNumber: big.NewInt(1),
//		Subject: pkix.Name{
//			Organization: []string{"Gin Example"},
//		},
//		NotBefore: time.Now(),
//		NotAfter:  time.Now().Add(time.Hour * 24 * 365),
//		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
//		ExtKeyUsage: []ExtKeyUsage{
//			ExtKeyUsageServerAuth,
//		},
//		DNSNames: []string{
//			"localhost",
//		},
//	}
//	certDERBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
//	if err != nil {
//		log.Fatalf("Failed to create certificate: %v", err)
//	}
//
//	// 将证书保存到文件
//	certOut, err := os.Create("cert.pem")
//	if err != nil {
//		log.Fatalf("Failed to open cert.pem for writing: %v", err)
//	}
//	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certDERBytes}); err != nil {
//		log.Fatalf("Failed to write data to cert.pem: %v", err)
//	}
//	if err := certOut.Close(); err != nil {
//		log.Fatalf("Error closing cert.pem: %v", err)
//	}
//
//	// 将私钥保存到文件
//	keyOut, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
//	if err != nil {
//		log.Fatalf("Failed to open key.pem for writing: %v", err)
//	}
//	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
//	if err != nil {
//		log.Fatalf("Unable to marshal private key: %v", err)
//	}
//	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
//		log.Fatalf("Failed to write data to key.pem: %v", err)
//	}
//	if err := keyOut.Close(); err != nil {
//		log.Fatalf("Error closing key.pem: %v", err)
//	}
//
//	// 使用服务器证书启动Gin应用程序
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "Hello, World!")
//	})
//	log.Fatal(router.RunTLS(":8080", "cert.pem", "key.pem"))
//}
