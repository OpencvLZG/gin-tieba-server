/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/19
  @desc: //TODO
**/

package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

// 生成 ECC 私钥
func generatePrivateKey() (key *ecdsa.PrivateKey) {
	key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return key
}

// 保存证书
func saveCert(cert *x509.Certificate, fileName string) {
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}
	pemData := pem.EncodeToMemory(certBlock)
	if err := os.WriteFile(fileName, pemData, 0644); err != nil {
		panic(err)
	}
}
func savePrivateKey(key *ecdsa.PrivateKey, fileName string) {
	keyDer, err := x509.MarshalECPrivateKey(key)

	keyBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyDer,
	}

	keyData := pem.EncodeToMemory(keyBlock)

	if err = os.WriteFile(fileName, keyData, 0644); err != nil {
		panic(err)
	}

}

// GenerateCert 生成根证书，中级证书， 终端证书
func GenerateCert(fileType string, organization string, country string, province string, locality string, organizationalUnit string, commonName string, dnsDomain string) {

	filePath := "./cert/" + organization + "/"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			panic(err)
		}
	}
	// 生成根证书
	rootCsr := &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{country},
			Province:           []string{province},
			Locality:           []string{locality},
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationalUnit},
			CommonName:         commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
		MaxPathLenZero:        false,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	rootKey := generatePrivateKey()
	rootDer, err := x509.CreateCertificate(rand.Reader, rootCsr, rootCsr, rootKey.Public(), rootKey)
	if err != nil {
		panic(err)
	}
	rootCert, err := x509.ParseCertificate(rootDer)
	if err != nil {
		panic(err)
	}
	saveCert(rootCert, filePath+"root_pem.cert")
	savePrivateKey(rootKey, filePath+"root_pem.key")
	// 根证书签证中级证书
	interCsr := &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{country},
			Province:           []string{province},
			Locality:           []string{locality},
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationalUnit},
			CommonName:         commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            0,
		MaxPathLenZero:        true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}
	interKey := generatePrivateKey()
	interDer, err := x509.CreateCertificate(rand.Reader, interCsr, rootCert, interKey.Public(), rootKey)
	if err != nil {
		panic(err)
	}
	interCert, err := x509.ParseCertificate(interDer)
	if err != nil {
		panic(err)
	}
	saveCert(interCert, filePath+"inter_pem.cert")
	savePrivateKey(interKey, filePath+"inter_pem.key")
	// 中级证书签证终端证书
	csr := &x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{country},
			Province:           []string{province},
			Locality:           []string{locality},
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationalUnit},
			CommonName:         commonName,
		},
		DNSNames:              []string{dnsDomain},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  false,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	key := generatePrivateKey()
	der, err := x509.CreateCertificate(rand.Reader, csr, interCert, key.Public(), interKey)
	if err != nil {
		panic(err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		panic(err)
	}
	saveCert(cert, filePath+"server.cert")
	savePrivateKey(key, filePath+"server.key")
}
